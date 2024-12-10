package utils

import (
	"context"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"net/url"
	"time"
)

// MinioEditProvider 是一个使用 minio 作为存储的编辑提供者
type MinioEditProvider struct {
	Client *minio.Client
	Bucket string
}

// NewMinioEditProvider 创建一个新的 MinioEditProvider
func NewMinioEditProvider(endpoint, accessKeyID, secretAccessKey, bucket string) (*MinioEditProvider, error) {
	// 使用指定的密钥创建一个客户端
	client, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: false, // 根据你的 minio 服务是否使用 HTTPS 设置
	})
	if err != nil {
		return nil, err
	}

	return &MinioEditProvider{
		Client: client,
		Bucket: bucket,
	}, nil
}

// UploadAddress 返回上传文件的地址
func (m *MinioEditProvider) UploadAddress(ctx context.Context, objectName string) (*url.URL, error) {
	// 获取上传地址
	presignedURL, err := m.Client.PresignedPutObject(ctx, m.Bucket, objectName, 24*time.Hour)
	if err != nil {
		return nil, err
	}

	return presignedURL, err
}
