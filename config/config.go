package config

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

type Config struct {
	Environment string // develop, staging, production
	LogLevel    string

	RPCPort string

	TempBucketName string
	BucketName     string
	Region         string

	GpuTranscode bool

	AwsRegion        string
	AwsID            string
	AwsSecret        string
	AwsEC2InstanceID string

	FFmpegGPU  string
	FFmpegCPU  string
	FFprobeCPU string
	FFprobeGPU string

	DefaultWorkers int

	OutputDir string
	InputDir  string

	MongoDBHost     string
	MongoDBPassword string
	MongoDBDatabase string
	MongoDBUser     string
	MongoDBPort     int

	ScriptsFolder string
}

func Load() Config { //nolint
	config := Config{}

	absPath := getAbsPath()

	config.Environment = cast.ToString(env("ENVIRONMENT", EnvDevelop))
	config.LogLevel = cast.ToString(env("LOG_LEVEL", "debug"))
	config.GpuTranscode = cast.ToBool(env("GPU_TRANSCODE", true))

	config.RPCPort = cast.ToString(env("RPC_PORT", ":9112"))

	config.AwsID = cast.ToString(env("AWS_ID", "KEY"))
	config.AwsSecret = cast.ToString(env("AWS_SECRET", "SECRET"))

	config.MongoDBHost = cast.ToString(env("MONGO_DB_HOST", "localhost"))
	config.MongoDBPort = cast.ToInt(env("MONGO_DB_PORT", "27017"))
	config.MongoDBDatabase = cast.ToString(env("MONGO_DB_DATABASE", "content_service"))
	config.MongoDBUser = cast.ToString(env("MONGO_DB_USER", "mongo"))
	config.MongoDBPassword = cast.ToString(env("MONGO_DB_PASSWORD", "mongo"))

	config.Region = cast.ToString(env("BUCKET_REGION", "region"))
	config.TempBucketName = cast.ToString(env("TEMP_BUCKET_NAME", "temp-bucket"))
	config.BucketName = cast.ToString(env("BUCKET_NAME", "bucket"))

	config.FFmpegCPU = cast.ToString(env("FFMPEG_CPU", "ffmpeg_binary"))
	config.FFprobeCPU = cast.ToString(env("FFPROBE_CPU", "ffprobe_binary"))
	config.FFmpegGPU = cast.ToString(env("FFMPEG_GPU", "/usr/local/bin/ffmpeg"))
	config.FFprobeGPU = cast.ToString(env("FFMPROBE_GPU", "/usr/local/bin/ffprobe"))

	config.ScriptsFolder = cast.ToString(env("SCRIPTS_FOLDER", fmt.Sprintf(
		"%s%s", absPath, "/scripts",
	)))

	// config.OutputDir = cast.ToString(env("OUTPUT_DIRECTORY",
	// 	"/media/voxe/Expansion/download_transcode"))
	config.OutputDir = cast.ToString(env("OUTPUT_DIRECTORY",
		"/media/voxe/H1/downloaded_transcode"))
	config.InputDir = cast.ToString(env("INPUT_DIRECTORY", "/media/voxe/H1/gotta_transco	de"))

	return config
}

func env(key string, defaultValue interface{}) interface{} {
	_, exists := os.LookupEnv(key)
	if exists {
		return os.Getenv(key)
	}

	return defaultValue
}

func getAbsPath() string {
	dir, err := filepath.Abs(filepath.Dir("."))
	if err != nil {
		log.Println("failed to retrieve absolute path for .env")
		panic(err)
	}

	if err := godotenv.Load(filepath.Join(dir, ".env")); err != nil {
		log.Print("No .env file found")
	}

	return dir
}
