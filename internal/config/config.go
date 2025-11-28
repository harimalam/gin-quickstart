package config

import (
	"errors"
	"log"
	"strings"
	"time"

	"github.com/spf13/viper"
)

// Mappers for ENV to nested config keys
var mappers = map[string]string{
	// App Configs
	"APP_PORT":      "app.port",
	"JWT_SECRET":    "app.jwt_secret",
	"GIN_MODE":      "app.gin_mode",
	"READ_TIMEOUT":  "app.read_timeout",
	"WRITE_TIMEOUT": "app.write_timeout",

	// DB Configs
	"DB_HOST":     "db.host",
	"DB_PORT":     "db.port",
	"DB_USER":     "db.user",
	"DB_PASSWORD": "db.password",
	"DB_NAME":     "db.name",
	"SSL_MODE":    "db.ssl_mode",
}

type AppConfig struct {
	Port         string        `mapstructure:"port"`
	JWTSecret    string        `mapstructure:"jwt_secret"`
	GinMode      string        `mapstructure:"gin_mode"`
	ReadTimeout  time.Duration `mapstructure:"read_timeout"`
	WriteTimeout time.Duration `mapstructure:"write_timeout"`
}

type DBConfig struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Name     string `mapstructure:"name"`
	SSLMode  string `mapstructure:"ssl_mode"`
}

type Config struct {
	App AppConfig `mapstructure:"app"`
	DB  DBConfig  `mapstructure:"db"`
}

func LoadConfig() (cfg Config, err error) {
	v := viper.New()

	// ---- 1. System ENV must be readable ----
	v.AutomaticEnv()

	// ---- 2. Force ENV format: APP_PORT, DB_USER etc ----
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	// ---- 3. Read .env file if available (optional) ----
	v.AddConfigPath(".")
	v.SetConfigName(".env")
	v.SetConfigType("env")

	if err = v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Println("⚠️ .env file not found, using system ENV only")
		} else {
			return cfg, err
		}
	} else {
		log.Println("📌 Loaded .env file:", v.ConfigFileUsed())
	}

	// ---- 4. Map flat ENV → Nested keys manually (Reliable Fix) ----

	for envKey, viperKey := range mappers {
		if val := v.GetString(envKey); val != "" {
			v.Set(viperKey, val)
		}
	}

	// ---- 5. Set duration defaults properly for nested keys ----
	if !v.IsSet("app.read_timeout") {
		v.Set("app.read_timeout", 5*time.Second)
	}
	if !v.IsSet("app.write_timeout") {
		v.Set("app.write_timeout", 10*time.Second)
	}
	if !v.IsSet("db.port") {
		v.Set("db.port", "5432")
	}
	if !v.IsSet("db.ssl_mode") {
		v.Set("db.ssl_mode", "disable")
	}
	if !v.IsSet("app.port") {
		v.Set("app.port", "8080")
	}
	if !v.IsSet("app.gin_mode") {
		v.Set("app.gin_mode", "debug")
	}

	// ---- 6. Unmarshal into nested struct ----
	if err = v.Unmarshal(&cfg); err != nil {
		return cfg, err
	}

	// ---- 7. Validate critical DB creds ----
	if cfg.DB.User == "" || cfg.DB.Password == "" || cfg.DB.Name == "" {
		return cfg, errors.New("missing DB_USER, DB_PASSWORD or DB_NAME")
	}
	log.Println("⚙️ Config loaded successfully (Nested struct + ENV mode)")
	return cfg, nil
}
