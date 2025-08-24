package config

type GlobalConf struct {
	Server *ServerConf `yaml:"server"`
}

type ServerConf struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

func NewConfiguration(configPath string) (*GlobalConf, error) {
	return nil, nil
}
