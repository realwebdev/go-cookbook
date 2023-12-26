/**
A practical use case for the singleton design pattern in Go: a configuration manager
that loads and provides application configurations from afile or external source.
*/

package main

import (
	"fmt"
	"sync"
)

type Config struct {
	ServerAddress string
	Port          int
	// other configurations fields...
}

// Config Manager manages application configurations as a singleton
type ConfigManager struct {
	ConfigData *Config
	// other ConfigManager fields...
}

var instance *ConfigManager
var once sync.Once

func GetConfigMAnagerInstance() *ConfigManager {
	once.Do(func() {
		instance = &ConfigManager{
			ConfigData: loadConfiguration(),
		}
	})
	return instance
}

// loadConfiguration simulates the loading configuraiton data
func loadConfiguration() *Config {
	// Simulated configuration loading
	return &Config{
		ServerAddress: "localhost",
		Port:          8080,
		// set other configuraitons
	}
}

// GetConfig returns the applicaiton configuration from ConfigManager
func (cm *ConfigManager) GetConfig() *Config {
	return cm.ConfigData
}

// SetConfig sets the application configuration in ConfigManager

func (cm *ConfigManager) SetConfig(config *Config) {
	cm.ConfigData = config
	fmt.Println("Configuraiton updated")
}

func main() {
	// Getting the config manager instance
	configManger := GetConfigMAnagerInstance()

	// retrieving and pring the configuration
	appConfig := configManger.GetConfig()
	fmt.Println("Server Address:", appConfig.ServerAddress)
	fmt.Println("Port:", appConfig.Port)

	// Modifying configuration

	newConfig := &Config{
		ServerAddress: "newhub.com",
		Port:          9090,
	}
	configManger.SetConfig(newConfig)

	// rertrieving and printing the updated configuraiton
	fmt.Println("Updated Server Address:", appConfig.ServerAddress)
	fmt.Println("Updated Port:", appConfig.Port)
}
