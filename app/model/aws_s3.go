package model


import (
    "errors"
    "fmt"
    "mime/multipart"

    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/credentials"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/s3/s3manager"
)

// AwsS3 ...
type AwsS3 struct {
    Config *Config
    Keys AwsS3URLs
    Uploader *s3manager.Uploader
}

// AwsS3URLs ...
type AwsS3URLs struct {
    Test string
}

// NewAwsS3 ...
func NewAwsS3() *AwsS3 {

    config := NewConfig()

    // The session the S3 Uploader will use
    sess := session.Must(session.NewSessionWithOptions(session.Options{
        Config: aws.Config{
            // Access key ID と Secret Access Key は IAM から作成する。
            // 一度作成した Access key ID の Secret Access Key は csv ファイルでダウンロードしていない場合は、
            // 確認する手段がないので新しくアクセスキーを作成する必要がある。
            Credentials: credentials.NewStaticCredentialsFromCreds(credentials.Value{
                AccessKeyID: config.Aws.S3.AccessKeyID,
                SecretAccessKey: config.Aws.S3.SecretAccessKey,
            }),
            Region: aws.String(config.Aws.S3.Region),
        },
    }))

    return &AwsS3{
        Config: config,
        Keys: AwsS3URLs{
            Test: "userIcon/",
        },
        // Create an uploader with the session and default options
        Uploader: s3manager.NewUploader(sess),
    }
}


// UploadTest ...
func (a *AwsS3) UploadTest(file multipart.File,fileName string, extension string) (url string, err error) {

    var contentType string

    switch extension {
    case ".jpg":
        contentType = "image/jpeg"
    case ".jpeg":
        contentType = "image/jpeg"
    case ".gif":
        contentType = "image/gif"
    case ".png":
        contentType = "image/png"
    default:
				return "拡張子エラー", errors.New("this extension is invalid")
    }

    // Upload the file to S3.
    result, err := a.Uploader.Upload(&s3manager.UploadInput{
        // ACL の設定は重要
        ACL: aws.String("public-read"),
        Body: file,
        Bucket: aws.String(a.Config.Aws.S3.Bucket),
        // content-type の設定も忘れずに
        ContentType: aws.String(contentType),
        Key: aws.String(a.Keys.Test + "/" + fileName + extension),
    })

    if err != nil {
        return "アップロードエラー", fmt.Errorf("failed to upload file, %v", err)
    }

    return result.Location, nil
}
