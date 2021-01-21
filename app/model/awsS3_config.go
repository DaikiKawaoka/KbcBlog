package model


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

	c := new(Config)

	// ex) アジアパシフィック (東京): ap-northeast-1
	c.Aws.S3.Region = "ap-northeast-1"
	c.Aws.S3.Bucket = "www.kbcblog-s3.com"
	c.Aws.S3.AccessKeyID = "AKIARMBCWCNCA3BQRPZJ"
	c.Aws.S3.SecretAccessKey = "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"

	return c
}