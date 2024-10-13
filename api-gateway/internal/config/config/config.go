package config

type Config struct {
	Listen    ListenConfig  `yaml:"listen"`
	Storage   StorageConfig `yaml:"storage"`
	SecretKey JwtConfig     `yaml:"jwt"`
	Redis     RedisConfig   `yaml:"redis"`
}

type JwtConfig struct {
	Secret string `yaml:"secret"`
}

type StorageConfig struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Database string `yaml:"database"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

type ListenConfig struct {
	BindIP string `yaml:"bind_ip" env-default:"127.0.0.1"`
	Port   string `yaml:"port" env-default:"8080"`
}

type RedisConfig struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}
