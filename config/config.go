package config

import (
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

type Config struct {
	Environment string // develop, staging, production
	LogLevel    string
	BaseCDN     string

	MongoDBHost     string
	MongoDBPassword string
	MongoDBDatabase string
	MongoDBUser     string
	MongoDBPort     int

	RPCPort string
}

func Load() Config {
	dir, err := filepath.Abs(filepath.Dir("."))
	if err != nil {
		log.Println("failed to retrieve absolute path for .env")
		panic(err)
	}

	if err := godotenv.Load(filepath.Join(dir, ".env")); err != nil {
		log.Print("No .env file found")
	}

	config := Config{}

	config.Environment = cast.ToString(getOrReturnDefault("ENVIRONMENT", EnvDevelop))
	config.LogLevel = cast.ToString(getOrReturnDefault("LOG_LEVEL", "debug"))
	config.BaseCDN = cast.ToString(
		getOrReturnDefault("BASE_CDN", "https://voxe-cdn.s3.eu-north-1.amazonaws.com"))

	config.MongoDBHost = cast.ToString(getOrReturnDefault("MONGO_DB_HOST", "localhost"))
	config.MongoDBPort = cast.ToInt(getOrReturnDefault("MONGO_DB_PORT", "27017"))

	config.MongoDBDatabase = cast.ToString(
		getOrReturnDefault("MONGO_DB_DATABASE", "content_service_new"))
	config.MongoDBUser = cast.ToString(getOrReturnDefault("MONGODB_USER", "mongo"))
	config.MongoDBPassword = cast.ToString(getOrReturnDefault("MONGODB_PASSWORD", "mongo"))

	config.RPCPort = cast.ToString(getOrReturnDefault("RPC_PORT", ":9100"))

	return config
}

func getOrReturnDefault(key string, defaultValue interface{}) interface{} {
	_, exists := os.LookupEnv(key)
	if exists {
		return os.Getenv(key)
	}

	return defaultValue
}
