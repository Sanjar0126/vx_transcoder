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

	RPCPort string

	TempBucketName string
	BucketName     string
	Region         string

	RabbitMQUser     string
	RabbitMQPassword string
	RabbitMQHost     string
	RabbitMQPort     int

	ContentServiceHost string
	ContentServicePort int

	GpuTranscode bool

	AwsRegion        string
	AwsID            string
	AwsSecret        string
	AwsEC2InstanceID string

	// Constants
	VideoExtension string
	FFmpegGPU      string
	FFmpegCPU      string
	FFprobeCPU     string
	FFprobeGPU     string

	DefaultWorkers int

	OutputDirectory string
	InputDirectory  string

	MongoDBHost     string
	MongoDBPassword string
	MongoDBDatabase string
	MongoDBUser     string
	MongoDBPort     int
}

func Load() Config { //nolint
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

	config.RPCPort = cast.ToString(getOrReturnDefault("RPC_PORT", ":9112"))

	// GPU transcoding option
	config.GpuTranscode = cast.ToBool(getOrReturnDefault("GPU_TRANSCODE", true))

	config.Region = cast.ToString(getOrReturnDefault("BUCKET_REGION", "eu-north-1"))
	config.TempBucketName = cast.ToString(getOrReturnDefault("TEMP_BUCKET_NAME", "voxe-temp-1"))
	config.BucketName = cast.ToString(getOrReturnDefault("BUCKET_NAME", "voxe-cdn"))

	config.VideoExtension = cast.ToString(getOrReturnDefault("VIDEO_EXTENSION", "m3u8"))
	config.FFmpegCPU = cast.ToString(getOrReturnDefault("FFMPEG_CPU", "ffmpeg_binary"))
	config.FFprobeCPU = cast.ToString(getOrReturnDefault("FFPROBE_CPU", "ffprobe_binary"))
	config.FFmpegGPU = cast.ToString(getOrReturnDefault("FFMPEG_GPU", "/usr/local/bin/ffmpeg"))
	config.FFprobeGPU = cast.ToString(getOrReturnDefault("FFMPROBE_GPU", "/usr/local/bin/ffprobe"))

	// aws credentials
	config.AwsRegion = cast.ToString(getOrReturnDefault("AWS_REGION", "eu-north-1"))
	config.AwsID = cast.ToString(getOrReturnDefault("AWS_ID", "sample_aws_id"))
	config.AwsSecret = cast.ToString(getOrReturnDefault("AWS_SECRET", "sample_aws_secret"))
	config.AwsEC2InstanceID = cast.ToString(getOrReturnDefault("AWS_EC2_INSTANCE_ID",
		"aws_ec2_instance_id"))

	// rabbit-mq credentials
	config.RabbitMQHost = cast.ToString(getOrReturnDefault("RABBIT_MQ_HOST", "localhost"))
	config.RabbitMQPassword = cast.ToString(getOrReturnDefault("RABBIT_MQ_PASSWORD", "guest"))
	config.RabbitMQUser = cast.ToString(getOrReturnDefault("RABBIT_MQ_USER", "guest"))
	config.RabbitMQPort = cast.ToInt(getOrReturnDefault("RABBIT_MQ_PORT", RabbitMQPort))

	// content-service credentials
	config.ContentServiceHost = cast.ToString(getOrReturnDefault("CONTENT_SERVICE_HOST",
		"localhost"))
	config.ContentServicePort = cast.ToInt(getOrReturnDefault("CONTENT_SERVICE_PORT",
		ContentServicePort))

	config.OutputDirectory = cast.ToString(getOrReturnDefault("OUTPUT_DIRECTORY",
		"/home/samandar/Downloads/transcoded"))
	config.InputDirectory = cast.ToString(getOrReturnDefault("INPUT_DICTORY",
		"/home/samandar/Downloads/videos"))

	config.MongoDBHost = cast.ToString(getOrReturnDefault("MONGO_DB_HOST", "localhost"))
	config.MongoDBPort = cast.ToInt(getOrReturnDefault("MONGO_DB_PORT", "27017"))

	config.MongoDBDatabase = cast.ToString(getOrReturnDefault("MONGO_DB_DATABASE",
		"content_service"))
	config.MongoDBUser = cast.ToString(getOrReturnDefault("MONGO_DB_USER", "mongo"))
	config.MongoDBPassword = cast.ToString(getOrReturnDefault("MONGO_DB_PASSWORD", "mongo"))

	return config
}

func getOrReturnDefault(key string, defaultValue interface{}) interface{} {
	_, exists := os.LookupEnv(key)
	if exists {
		return os.Getenv(key)
	}

	return defaultValue
}
