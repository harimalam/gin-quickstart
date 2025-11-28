package config

import (
	"errors" // Added for explicit validation errors
	"log"
	"time" // Added for time.Duration

	"github.com/spf13/viper"
)

// Config holds all application configuration values.
// The flat structure is used for guaranteed compatibility with your environment variables.
type Config struct {
	// Application Settings
	Port         string        `mapstructure:"APP_PORT"`
	JWTSecret    string        `mapstructure:"JWT_SECRET"`
	ReadTimeout  time.Duration `mapstructure:"READ_TIMEOUT"`  // Standard practice for server configuration
	WriteTimeout time.Duration `mapstructure:"WRITE_TIMEOUT"` // Standard practice for server configuration

	// Database Settings
	DBHost     string `mapstructure:"DB_HOST"`
	DBPort     string `mapstructure:"DB_PORT"`
	DBUser     string `mapstructure:"DB_USER"`
	DBPassword string `mapstructure:"DB_PASSWORD"`
	DBName     string `mapstructure:"DB_NAME"`
	SSLMode    string `mapstructure:"SSL_MODE"`
}

// LoadConfig reads configuration from file or environment variables.
// It applies defaults, reads from the environment, and performs critical validation.
func LoadConfig() (config Config, err error) {
	// 1. Set Defaults (Robustness and Predictability)
	viper.SetDefault("APP_PORT", "8080")
	viper.SetDefault("READ_TIMEOUT", 5*time.Second)
	viper.SetDefault("WRITE_TIMEOUT", 10*time.Second)
	viper.SetDefault("SSL_MODE", "disable")
	// Set an insecure default and rely on validation/warning if it's not changed
	viper.SetDefault("JWT_SECRET", "change-me-in-production-1234567890123456")

	// 2. Configure file search (for local .env)
	viper.AddConfigPath(".")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	// 3. Instruct Viper to read from OS environment (Critical for Docker)
	viper.AutomaticEnv()

	// 4. CRITICAL: Explicitly bind ALL environment variables.
	// This is the most reliable method for Docker to ensure environment variables are picked up.
	_ = viper.BindEnv("DB_HOST")
	_ = viper.BindEnv("DB_PORT")
	_ = viper.BindEnv("DB_USER")
	_ = viper.BindEnv("DB_PASSWORD")
	_ = viper.BindEnv("DB_NAME")
	_ = viper.BindEnv("SSL_MODE")
	_ = viper.BindEnv("APP_PORT", "PORT")
	_ = viper.BindEnv("JWT_SECRET")
	_ = viper.BindEnv("READ_TIMEOUT")
	_ = viper.BindEnv("WRITE_TIMEOUT")

	// 5. Read the configuration file (optional)
	err = viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// File not found is OK, rely on environment/defaults
			log.Println("Configuration file not found, relying on environment variables and defaults.")
		} else {
			// Malformed file is critical
			return Config{}, err
		}
	}

	// 6. Unmarshal values into the struct
	if err = viper.Unmarshal(&config); err != nil {
		return Config{}, err
	}

	// 7. 🛑 MANDATORY VALIDATION (Fail Fast Principle)
	// Check for critical database settings. The application cannot run without these.
	if config.DBUser == "" || config.DBPassword == "" || config.DBName == "" || config.DBHost == "" {
		return Config{}, errors.New("missing critical database credentials (DB_USER, DB_PASSWORD, DB_NAME, or DB_HOST is empty)")
	}

	// Warn if using the insecure default JWT secret
	if config.JWTSecret == "change-me-in-production-1234567890123456" {
		log.Println("⚠️ WARNING: JWT_SECRET is using the insecure default. Please set a unique JWT_SECRET environment variable.")
	}

	return config, nil
}
