package mysql

type Config struct {
	Endpoint string `yaml:"endpoint"`
	Name     string `yaml:"name"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}
