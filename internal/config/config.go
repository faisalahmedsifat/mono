package config

type Config struct {
	Services map[string]Service `yaml:"services"`
}

type Service struct {
	Dir     string            `yaml:"dir"`
	Command string            `yaml:"command"`
	Watch   []string          `yaml:"watch"`
	Env     map[string]string `yaml:"env"`
	Port    int               `yaml:"port"`
}

func Load(path string) (*Config, error) {
	//TODO
	return nil, nil
}

func FilterServices(cfg *Config, onlyTasks []string) *Config {
	//TODO
	return cfg
}
