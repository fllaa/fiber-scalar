package scalar

type Config struct {
	// This parameter can be used to name different instances of the Scalar API Reference
	// default: ""
	InstanceName string `json:"-"` // Name of the instance

	// Title of the API
	// default: "API Reference"
	Title string `json:"-"` // Title of the API

	// OpenAPI spec content, can be a json or yaml string
	// default: ""
	Content string `json:"-"` // OpenAPI spec content

	// Content URL, can be a URL to the OpenAPI spec
	// default: ""
	ContentURL string `json:"-"` // OpenAPI spec URL

	// Proxy URL to avoid CORS issues
	// default: "https://proxy.scalar.com"
	ProxyURL string `json:"-"` // Proxy URL

	// Layout, can be "modern" or "classic"
	// default: "modern"
	Layout string `json:"-"` // Layout

	// Theme, can be "alternate", "default", "moon", "purple", "solarized", "bluePlanet", "saturn", "kepler", "mars", "deepSpace" or "none"
	// default: "default"
	Theme string `json:"-"` // Theme

	// Dark mode, enable or disable dark mode
	// default: false
	DarkMode bool `json:"-"` // Dark mode
}

var (
	ConfigDefault = Config{
		Title:    "API Reference",
		ProxyURL: "https://proxy.scalar.com",
		Layout:   "modern",
		Theme:    "default",
		DarkMode: false,
	}
)

func configDefault(config ...Config) Config {
	if len(config) < 1 {
		return ConfigDefault
	}

	// Override default config
	cfg := config[0]

	if cfg.Title == "" {
		cfg.Title = ConfigDefault.Title
	}

	if cfg.ProxyURL == "" {
		cfg.ProxyURL = ConfigDefault.ProxyURL
	}

	if cfg.Layout == "" {
		cfg.Layout = ConfigDefault.Layout
	}

	if cfg.Theme == "" {
		cfg.Theme = ConfigDefault.Theme
	}

	return cfg
}
