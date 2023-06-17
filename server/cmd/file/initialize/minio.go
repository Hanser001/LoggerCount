package initialize

import (
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/minio/minio-go/v7/pkg/credentials"
	config2 "summer/server/cmd/file/config"

	"github.com/minio/minio-go/v7"
)

func InitMinio() *minio.Client {
	config := config2.GlobalServerConfig.MinioInfo
	// Initialize minio client object.
	mc, err := minio.New(config.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(config.AccessKeyID, config.SecretAccessKey, ""),
		Secure: false,
	})
	if err != nil {
		klog.Fatalf("create minio client err: %s", err.Error())
	}
	exists, err := mc.BucketExists(context.Background(), config.Bucket)
	if err != nil {
		klog.Fatal(err)
	}
	if !exists {
		err = mc.MakeBucket(context.Background(), config.Bucket, minio.MakeBucketOptions{Region: "cn-north-1"})
		if err != nil {
			klog.Fatalf("make bucket err: %s", err.Error())
		}
	}

	return mc
}
