package generator

type AppConfig struct {
	Version     string    `yaml:"version"`
	Name        string    `yaml:"name"`
	PackagePath string    `yaml:"package-path"`
	AppVersion  string    `yaml:"app-version"`
	Schemas     []*Schema `yaml:"schemas"`
	Flows       []*Flow   `yaml:"flows"`
}
type Schema struct {
	Name    string      `yaml:"name"`
	Payload interface{} `yaml:"payload"`
}
type Flow struct {
	Name         string      `yaml:"name"`
	Upstream     string      `yaml:"upstream"`
	Downstream   string      `yaml:"downstream"`
	Dependencies []string    `yaml:"dependencies"`
	Mapping      interface{} `yaml:"mapping"`
}
