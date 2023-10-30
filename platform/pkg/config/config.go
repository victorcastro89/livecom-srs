package config

import (
	"github.com/spf13/viper"
)

// Config represents your application configuration.
type Config struct {
    EncryptionKey string
    // Add other configuration fields here
}


// Global variable to hold the configuration.
var Cfg *Config // Exported variable

// Load loads the configuration.
func Load() (*Config, error) {
    viper.SetConfigName("config") // Name of your config file without extension
	viper.SetConfigType("yaml") 
    viper.AddConfigPath(".") // Path to your config file
 

    // Automatically bind environment variables
    viper.AutomaticEnv()

    if err := viper.ReadInConfig(); err != nil {
        return nil, err
    }

    // Unmarshal the configuration into the Config struct
    Cfg = new(Config) // Assign to the exported variable
    if err := viper.Unmarshal(Cfg); err != nil {
        return nil, err
    }

    return Cfg, nil
}
