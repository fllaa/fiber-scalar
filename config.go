package scalar

type Config struct {
	InstanceName string `json:"-"` // Name of the instance
	Title        string `json:"-"` // Title of the API
	Spec         string `json:"-"` // OpenAPI spec
	URL          string `json:"-"` // URL of the OpenAPI spec
}

var (
	ConfigDefault = Config{
		Title: "API Reference",
	}
)

func configDefault(config ...Config) Config {
	if len(config) < 1 {
		return ConfigDefault
	}
	return config[0]
}
