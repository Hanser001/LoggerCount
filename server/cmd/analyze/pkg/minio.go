package pkg

import (
	"context"
	"fmt"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/minio/minio-go/v7"
	"strconv"
	"summer/server/cmd/analyze/config"
	"time"
)

type MinioManger struct {
	mc *minio.Client
}

func NewMinioManger(client *minio.Client) *MinioManger {
	return &MinioManger{mc: client}
}

func (m *MinioManger) UploadFinalData(ctx context.Context, filepath string) error {
	buckname := config.GlobalServerConfig.MinioInfo.Bucket

	newobjname := strconv.FormatInt(time.Now().Unix(), 10)
	objname := fmt.Sprintf("%s.json", newobjname)

	_, err := m.mc.FPutObject(ctx, buckname, objname, filepath, minio.PutObjectOptions{ContentType: "application/json"})
	if err != nil {
		klog.Fatal("upload data err,", err)
	}
	return nil
}
