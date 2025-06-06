package database

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"strings"
	"testing"
)

type DBConfig struct {
	Database struct {
		Username string `destructure:"username"`
		Password string `destructure:"password"`
		Host     string `destructure:"host"`
		Port     string `destructure:"port"`
	} `destructure:"database"`
}

func TestConfig(t *testing.T) {
	viper.SetConfigName("local") // Config file name without extension
	viper.SetConfigType("yaml")  // Config file type
	viper.AddConfigPath(".")     // Look for the configs file in the current directory

	viper.AutomaticEnv()
	viper.SetEnvPrefix("env")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	// Read the configs file
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error reading configs file, %s", err)
	}

	viper.BindEnv("database.username", "DB_USERNAME")

	// Get the values, using env variables if present
	username := viper.GetString("database.username")
	password := viper.GetString("database.password")
	host := viper.GetString("database.host")
	port := viper.GetString("database.port")

	// Output the configuration values
	fmt.Printf("DB_USERNAME: %s\n", username)
	fmt.Printf("Password: %s\n", password)
	fmt.Printf("Host: %s\n", host)
	fmt.Printf("Port: %s\n", port)

	// Create an instance of Config
	var config DBConfig
	// Unmarshal the configs file into the Config struct
	err = viper.Unmarshal(&config)
	if err != nil {
		log.Fatalf("Unable to decode into struct, %v", err)
	}

	// Output the configuration values
	fmt.Printf("Config: %v\n", config)
}
