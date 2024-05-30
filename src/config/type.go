package config

import "time"

type GlobalConfig struct {
	Server   ServerConfig       `mapstructure:"server"`
	Database DatabaseConfig     `mapstructure:"database"`
	Casbin   CasbinConfig       `mapstructure:"casbin"`
	JWT      JWTConfig          `mapstructure:"jwt"`
	GraphQL  GraphqlConfig      `mapstructure:"graphql"`
	Argon2   PasswordHashParams `mapstructure:"argon2"`
}

type ServerConfig struct {
	Address string     `mapstructure:"address"`
	CORS    CORSConfig `mapstructure:"cors"`
}

type DatabaseConfig struct {
	Driver string `mapstructure:"driver"`
	Url    string `mapstructure:"url"`
}

type JWTConfig struct {
	PrivateKeyPath string        `mapstructure:"private_key_path"`
	Duration       time.Duration `mapstructure:"duration"`
	RenewDuration  time.Duration `mapstructure:"renew_duration"`
}

type GraphqlConfig struct {
	EndPoint           string `mapstructure:"endpoint"`
	Introspection      bool   `mapstructure:"introspection"`
	Playground         bool   `mapstructure:"playground"`
	PlaygroundEndpoint string `mapstructure:"playground_endpoint"`
}

type PasswordHashParams struct {
	Memory      uint32 `mapstructure:"memory"`
	Iterations  uint32 `mapstructure:"iterations"`
	Parallelism uint8  `mapstructure:"parallelism"`
	SaltLength  uint32 `mapstructure:"salt_length"`
	KeyLength   uint32 `mapstructure:"key_length"`
}

type CasbinConfig struct {
	Driver string `mapstructure:"driver"`
	Url    string `mapstructure:"url"`
}

type CORSConfig struct {
	AllowedOrigins   []string `mapstructure:"allowed_origins"`
	AllowedMethods   []string `mapstructure:"allowed_methods"`
	AllowedHeaders   []string `mapstructure:"allowed_headers"`
	ExposedHeaders   []string `mapstructure:"exposed_headers"`
	AllowCredentials bool     `mapstructure:"allow_credentials"`
	MaxAge           int      `mapstructure:"max_age"`
}
