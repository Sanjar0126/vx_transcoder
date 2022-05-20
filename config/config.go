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

	config.Environment = cast.ToString(env("ENVIRONMENT", EnvDevelop))
	config.LogLevel = cast.ToString(env("LOG_LEVEL", "debug"))

	config.RPCPort = cast.ToString(env("RPC_PORT", ":9112"))

	// GPU transcoding option
	config.GpuTranscode = cast.ToBool(env("GPU_TRANSCODE", true))

	config.Region = cast.ToString(env("BUCKET_REGION", "eu-north-1"))
	config.TempBucketName = cast.ToString(env("TEMP_BUCKET_NAME", "voxe-temp-1"))
	config.BucketName = cast.ToString(env("BUCKET_NAME", "voxe-cdn"))

	config.VideoExtension = cast.ToString(env("VIDEO_EXTENSION", "m3u8"))
	config.FFmpegCPU = cast.ToString(env("FFMPEG_CPU", "ffmpeg_binary"))
	config.FFprobeCPU = cast.ToString(env("FFPROBE_CPU", "ffprobe_binary"))
	config.FFmpegGPU = cast.ToString(env("FFMPEG_GPU", "/usr/local/bin/ffmpeg"))
	config.FFprobeGPU = cast.ToString(env("FFMPROBE_GPU", "/usr/local/bin/ffprobe"))

	// aws credentials
	config.AwsRegion = cast.ToString(env("AWS_REGION", "eu-north-1"))
	config.AwsID = cast.ToString(env("AWS_ID", "sample_aws_id"))
	config.AwsSecret = cast.ToString(env("AWS_SECRET", "sample_aws_secret"))
	config.AwsEC2InstanceID = cast.ToString(env("AWS_EC2_INSTANCE_ID", "aws_ec2_instance_id"))

	config.OutputDirectory = cast.ToString(env("OUTPUT_DIRECTORY",
		"/home/samandar/Downloads/transcoded"))
	config.InputDirectory = cast.ToString(env("INPUT_DICTORY",
		"/home/samandar/Downloads/videos"))

	config.MongoDBHost = cast.ToString(env("MONGO_DB_HOST", "localhost"))
	config.MongoDBPort = cast.ToInt(env("MONGO_DB_PORT", "27017"))

	config.MongoDBDatabase = cast.ToString(env("MONGO_DB_DATABASE",
		"content_service"))
	config.MongoDBUser = cast.ToString(env("MONGO_DB_USER", "mongo"))
	config.MongoDBPassword = cast.ToString(env("MONGO_DB_PASSWORD", "mongo"))

	return config
}

func env(key string, defaultValue interface{}) interface{} {
	_, exists := os.LookupEnv(key)
	if exists {
		return os.Getenv(key)
	}

	return defaultValue
}
