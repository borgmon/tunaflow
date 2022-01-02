package generator

type AppConfig struct {
	Version     string    `yaml:"version"`
	Name        string    `yaml:"name"`
	PackagePath string    `yaml:"package-path"`
	AppVersion  string    `yaml:"app-version"`
	Schemas     []Schemas `yaml:"schemas"`
	Flows       []Flows   `yaml:"flows"`
}
type Schemas struct {
	Name    string      `yaml:"name"`
	Payload interface{} `yaml:"payload"`
}
type Flows struct {
	Name         string      `yaml:"name"`
	Upstream     string      `yaml:"upstream"`
	Downstream   string      `yaml:"downstream"`
	Dependencies []string    `yaml:"dependencies"`
	Mapping      interface{} `yaml:"mapping"`
}
