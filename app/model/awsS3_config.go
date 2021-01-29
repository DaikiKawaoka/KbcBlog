package model

import (
	"os"
	"fmt"
	"github.com/joho/godotenv" //環境変数
)

// Config ...
type Config struct {
	Aws struct {
			S3 struct {
					Region string
					Bucket string
					AccessKeyID string
					SecretAccessKey string
			}
	}
}

// NewConfig ...
func NewConfig() *Config {
	_ = godotenv.Load(fmt.Sprintf("../../%s.env", os.Getenv("GO_ENV")))

	var (
		AwsS3Region = os.Getenv("AWS_S3_REGION")
		AwsS3Bucket = os.Getenv("AWS_S3_BUCKET")
		AwsS3AccessKeyID = os.Getenv("AWS_S3_ACCESS_KEY_ID")
		AwsS3SecretAccessKey = os.Getenv("AWS_S3_SECRET_ACCESS_KEY")
	)

	c := new(Config)

	// 環境変数から代入
	c.Aws.S3.Region = AwsS3Region
	c.Aws.S3.Bucket = AwsS3Bucket
	c.Aws.S3.AccessKeyID = AwsS3AccessKeyID
	c.Aws.S3.SecretAccessKey = AwsS3SecretAccessKey

	return c
}