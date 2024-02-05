package config

import (
	"context"
	"github.com/intyouss/Traceability/utils"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/spf13/viper"
)

func InitOSS() (*utils.MinioClient, error) {
	extraConn, err := NewMinioExtraConn()
	if err != nil {
		return nil, err
	}
	intraConn, err := NewMinioIntraConn()
	if err != nil {
		return nil, err
	}
	client, err := NewMinioConn(extraConn, intraConn)
	if err != nil {
		return nil, err
	}
	return client, err
}

func NewMinioConn(extraConn utils.ExtraConn, intraConn utils.IntraConn) (*utils.MinioClient, error) {
	client := utils.NewMinioClient(extraConn, intraConn)
	err := client.CreateBucket(context.Background(), viper.GetString("minio.bucketName"))
	return client, err
}

func NewMinioExtraConn() (utils.ExtraConn, error) {
	extraConn, err := minio.New(viper.GetString("minio.endpointExtra"), &minio.Options{
		Creds: credentials.NewStaticV4(
			viper.GetString("minio.accessKeyId"), viper.GetString("minio.accessSecret"), ""),
		Secure: viper.GetBool("minio.UseSsl"),
	})
	return utils.NewMinioExtraConn(extraConn), err
}

func NewMinioIntraConn() (utils.IntraConn, error) {
	extraConn, err := minio.New(viper.GetString("minio.endpointIntra"), &minio.Options{
		Creds: credentials.NewStaticV4(
			viper.GetString("minio.accessKeyId"), viper.GetString("minio.accessSecret"), ""),
		Secure: viper.GetBool("minio.UseSsl"),
	})
	return utils.NewMinioIntraConn(extraConn), err
}
