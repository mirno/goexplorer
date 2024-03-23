package config

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/spf13/viper"
)

func LoadDefaultConfig() {
    viper.SetConfigName("default") // Name of the default config file
    viper.AddConfigPath("config")  // Path to look for the config file in
    viper.SetConfigType("yaml")    // Set the type of the configuration file

    if err := viper.ReadInConfig(); err != nil {
        log.Fatalf("Error loading default config: %v", err)
    }
}

// InitializeConfig initializes and merges configuration from default and environment-specific files.
func InitializeConfig(envConfigFile string) (*Config, error) {
    var conf Config

    viper.SetConfigType("yaml")

    // Load the default configuration
    viper.SetConfigName("default")
    viper.AddConfigPath("config") // Adjust the path to where your config files are located
    if err := viper.ReadInConfig(); err != nil {
        return nil, fmt.Errorf("error reading default config file: %w", err)
    }

    // If an environment-specific configuration file is provided, merge it
    if envConfigFile != "" {
        viper.SetConfigFile(filepath.Join("config", envConfigFile))
        if err := viper.MergeInConfig(); err != nil {
            return nil, fmt.Errorf("error merging environment-specific config file: %w", err)
        }
    }

    // Unmarshal the combined configuration into the Config struct
    if err := viper.Unmarshal(&conf); err != nil {
        return nil, fmt.Errorf("error unmarshalling config: %w", err)
    }

    return &conf, nil
}

func InitializeAppConfig(envConfigFile string) (*AppConfig, error) {
    var conf AppConfig

    viper.SetConfigType("yaml")

    // Load the default configuration
    viper.SetConfigName("default")
    viper.AddConfigPath("config") // Adjust the path to where your config files are located
    if err := viper.ReadInConfig(); err != nil {
        return nil, fmt.Errorf("error reading default config file: %w", err)
    }

    // If an environment-specific configuration file is provided, merge it
    if envConfigFile != "" {
        viper.SetConfigFile(filepath.Join("config", envConfigFile))
        if err := viper.MergeInConfig(); err != nil {
            return nil, fmt.Errorf("error merging environment-specific config file: %w", err)
        }
    }

    // Unmarshal the combined configuration into the Config struct
    if err := viper.Unmarshal(&conf); err != nil {
        return nil, fmt.Errorf("error unmarshalling config: %w", err)
    }

    return &conf, nil
}

func LoadConfig(configPath, configName string) (*AppConfig, error){
	viper.AddConfigPath(configPath)
	viper.SetConfigName(configName)
	viper.SetConfigType("yaml")

	var config AppConfig

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}
	return &config, nil
}