package cloud

import (
	"context"

	"github.com/Darktrace1/Playble-PopupStore-Backend/pkg/common/utils"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func InitS3(AwsRegion string, AwsAccessKey string, AwsSecretKey string)  *manager.Uploader{
	creds := credentials.NewStaticCredentialsProvider(
		AwsAccessKey,
		AwsSecretKey,
		"",
	)

	cfg, err := config.LoadDefaultConfig(
		context.Background(),
		config.WithRegion(AwsRegion),
		config.WithCredentialsProvider(creds),
	)
	utils.CheckErr(err)

	client := s3.NewFromConfig(cfg)
	uploader := manager.NewUploader(client)

	return uploader
}