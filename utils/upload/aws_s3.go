package upload

import (
	"errors"
	"fmt"
	"mime/multipart"
	"time"

	"github.com/MjSteed/vue3-element-admin-go/common"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"go.uber.org/zap"
)

var sess = newSession()

func UploadFile(file *multipart.FileHeader) (string, string, error) {
	uploader := s3manager.NewUploader(sess)

	fileKey := fmt.Sprintf("%d%s", time.Now().Unix(), file.Filename)
	filename := common.Config.AwsS3.PathPrefix + "/" + fileKey
	f, openError := file.Open()
	if openError != nil {
		common.LOG.Error("function file.Open() Filed", zap.Any("err", openError.Error()))
		return "", "", errors.New("function file.Open() Filed, err:" + openError.Error())
	}
	defer f.Close() // 创建文件 defer 关闭

	_, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(common.Config.AwsS3.Bucket),
		Key:    aws.String(filename),
		Body:   f,
	})
	if err != nil {
		common.LOG.Error("function uploader.Upload() Filed", zap.Any("err", err.Error()))
		return "", "", err
	}

	return common.Config.AwsS3.BaseURL + "/" + filename, fileKey, nil
}

// newSession Create S3 session
func newSession() *session.Session {
	sess, _ := session.NewSession(&aws.Config{
		Region:           aws.String(common.Config.AwsS3.Region),
		Endpoint:         aws.String(common.Config.AwsS3.Endpoint),
		S3ForcePathStyle: aws.Bool(common.Config.AwsS3.S3ForcePathStyle),
		DisableSSL:       aws.Bool(common.Config.AwsS3.DisableSSL),
		Credentials: credentials.NewStaticCredentials(
			common.Config.AwsS3.SecretID,
			common.Config.AwsS3.SecretKey,
			"",
		),
	})
	return sess
}
