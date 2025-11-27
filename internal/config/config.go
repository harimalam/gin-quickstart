package config

import "github.com/spf13/viper"

type Config struct {
	Port      string `mapstructure:"APP_PORT"`
	JWTSecret string `mapstructure:"JWT_SECRET"`

	DBHost     string `mapstructure:"DB_HOST"`
	DBPort     string `mapstructure:"DB_PORT"`
	DBUser     string `mapstructure:"DB_USER"`
	DBPassword string `mapstructure:"DB_PASSWORD"`
	DBName     string `mapstructure:"DB_NAME"`
	SSLMode    string `mapstructure:"SSL_MODE"`
}

// LoadConfig reads configuration from file or environment variables.
func LoadConfig() (config Config, err error) {
	// 1. Tell Viper to look for an .env file
	viper.AddConfigPath(".")    // Look in current directory
	viper.SetConfigName(".env") // Look for file named .env
	viper.SetConfigType("env")  // Use environment file type

	// 2. Instruct Viper to read from environment variables (important for deployment)
	viper.AutomaticEnv()

	// 3. Read the configuration file (optional, file might not exist)
	err = viper.ReadInConfig()
	if err != nil {
		// Log the error but proceed, as variables might come from the OS environment
	}

	// 4. Unmarshal the values into the Config struct
	err = viper.Unmarshal(&config)
	return
}
