package dao

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/minio/minio-go/v7"
	"net/url"
	"summer/server/shared/consts"
	"time"
)

type FileManger struct {
	mc *minio.Client
}

// NewFileManger create a user dao.
func NewFileManger(client *minio.Client) *FileManger {
	return &FileManger{mc: client}
}

func (m *FileManger) Upload(ctx context.Context, bucketName, objectName, filePath string) error {
	_, err := m.mc.FPutObject(ctx, bucketName, objectName, filePath, minio.PutObjectOptions{
		ContentType: "application/log",
	})

	if err != nil {
		klog.Fatal("upload file err", err)
	}

	return nil
}

func (m *FileManger) Download(ctx context.Context, bucketName, objectName string) (*url.URL, error) {
	reqParams := make(url.Values)
	reqParams.Set("response-content-disposition", "attachment; filename=\""+objectName+"\"")

	downloadUrl, err := m.mc.PresignedGetObject(ctx, bucketName, objectName, time.Second*consts.UrlExpiredTime, reqParams)
	if err != nil {
		klog.Fatal("get download url err", err)
	}

	return downloadUrl, nil
}
