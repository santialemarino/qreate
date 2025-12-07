package config

type AppConfig struct {
	Name string `yaml:"name"`
	Env  string `yaml:"env"`
	Port int    `yaml:"port"`
}

type Environment string

type DatabaseConfig struct {
	Driver   string `yaml:"driver"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Name     string `yaml:"name"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	SSLMode  string `yaml:"sslmode"`
}

type RedisConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Password string `yaml:"password"`
	DB       int    `yaml:"db"`
}

type JWTConfig struct {
	Secret                 string `yaml:"secret"`
	AccessTokenTTLSeconds  int    `yaml:"accessTokenTTLSeconds"`
	RefreshTokenTTLSeconds int    `yaml:"refreshTokenTTLSeconds"`
}

type QRConfig struct {
	BaseURL         string `yaml:"base_url"`
	ShortCodeLength int    `yaml:"short_code_length"`
}

type Settings struct {
	App         AppConfig      `yaml:"app"`
	Environment Environment    `yaml:"Environment"`
	Database    DatabaseConfig `yaml:"Database"`
	Redis       RedisConfig    `yaml:"Redis"`
	JWT         JWTConfig      `yaml:"JWT"`
	QR          QRConfig       `yaml:"QR"`
}
