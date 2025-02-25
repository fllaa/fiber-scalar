package scalar

import (
	"fmt"
	"html/template"
	"path"
	"strings"
	"sync"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
	"github.com/swaggo/swag/v2"
)

const (
	defaultDocURL = "doc.json"
	defaultIndex  = "index.html"
)

var HandlerDefault = New()

func New(config ...Config) fiber.Handler {
	cfg := configDefault(config...)

	index, err := template.New("scalar_index.html").Parse(indexTmpl)
	if err != nil {
		panic(fmt.Errorf("fiber: scalar middleware error -> %w", err))
	}

	var (
		prefix string
		once   sync.Once
	)

	return func(c *fiber.Ctx) error {
		// Set prefix
		once.Do(
			func() {
				prefix = strings.ReplaceAll(c.Route().Path, "*", "")

				forwardedPrefix := getForwardedPrefix(c)
				if forwardedPrefix != "" {
					prefix = forwardedPrefix + prefix
				}

				// Set doc url
				if len(cfg.URL) == 0 {
					cfg.URL = path.Join(prefix, defaultDocURL)
				}

				// Set Spec
				if len(cfg.Spec) == 0 {
					cfg.Spec, err = swag.ReadDoc(cfg.InstanceName)
					if err != nil {
						return
					}
				}
			},
		)

		p := c.Path(utils.CopyString(c.Params("*")))

		switch p {
		case defaultIndex:
			c.Type("html")
			return index.Execute(c, cfg)
		case defaultDocURL:
			var doc string
			if doc, err = swag.ReadDoc(cfg.InstanceName); err != nil {
				fmt.Println(err)
				return err
			}
			return c.Type("json").SendString(doc)
		default:
			return c.Redirect(path.Join(prefix, defaultIndex), fiber.StatusMovedPermanently)
		}
	}
}

func getForwardedPrefix(c *fiber.Ctx) string {
	header := c.GetReqHeaders()["X-Forwarded-Prefix"]

	if len(header) == 0 {
		return ""
	}

	prefix := ""

	for _, rawPrefix := range header {
		endIndex := len(rawPrefix)
		for endIndex > 1 && rawPrefix[endIndex-1] == '/' {
			endIndex--
		}

		if endIndex != len(rawPrefix) {
			prefix += rawPrefix[:endIndex]
		} else {
			prefix += rawPrefix
		}
	}

	return prefix
}
