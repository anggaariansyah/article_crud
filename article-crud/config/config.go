package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

type Config struct {
	Database struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Name     string `yaml:"name"`
	} `yaml:"database"`
	Server struct {
		Port int `yaml:"port"`
	} `yaml:"server"`
	Auth struct {
		Username string `yaml:"username"`
		Password string `yaml:"password"`
	} `yaml:"auth"`
}

var AppConfig Config

func LoadConfig() {
	f, err := os.ReadFile("config.yaml")
	if err != nil {
		log.Fatal("Failed to read config.yaml:", err)
	}
	if err := yaml.Unmarshal(f, &AppConfig); err != nil {
		log.Fatal("Failed to parse config.yaml:", err)
	}
}

func ConnectDB() *gorm.DB {
	LoadConfig()
	d := AppConfig.Database
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		d.Host, d.User, d.Password, d.Name, d.Port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect database:", err)
	}
	return db
}
