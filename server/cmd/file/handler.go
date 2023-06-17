package main

import (
	"context"
	"strconv"
	"summer/server/cmd/file/config"
	"summer/server/cmd/file/dao"
	"summer/server/shared/errno"
	"summer/server/shared/kitex_gen/file"
	"time"
)

// FileServiceImpl implements the last service interface defined in the IDL.
type FileServiceImpl struct {
	Dao *dao.FileManger
	RedisManger
}

type RedisManger interface {
	NewUpload(ctx context.Context, userId int64, filed string, value int64) error
}

// UploadFile implements the FileServiceImpl interface.
func (s *FileServiceImpl) UploadFile(ctx context.Context, req *file.UploadRequest) (resp *file.UploadResponse, err error) {
	resp = new(file.UploadResponse)

	bucketName := config.GlobalServerConfig.MinioInfo.Bucket
	objectName := strconv.FormatInt(time.Now().Unix(), 10)
	score := time.Now().Unix()

	err = s.Dao.Upload(ctx, bucketName, objectName, req.FilePath)
	if err != nil {
		resp.StatusCode = int32(errno.FileServerErr.ErrCode)
		resp.StatusMsg = errno.FileServerErr.ErrMsg
		resp.Url = ""
		return resp, nil
	}

	//call redis client
	s.RedisManger.NewUpload(ctx, req.UserId, objectName, score)

	resp.StatusCode = int32(errno.Success.ErrCode)
	resp.StatusMsg = errno.Success.ErrMsg

	return resp, nil
}

// DownloadFile implements the FileServiceImpl interface.
func (s *FileServiceImpl) DownloadFile(ctx context.Context, req *file.DownloadRequest) (resp *file.DownloadResponse, err error) {
	resp = new(file.DownloadResponse)

	bucketName := config.GlobalServerConfig.MinioInfo.Bucket
	url, err := s.Dao.Download(ctx, bucketName, req.ObjectName)
	if err != nil {
		resp.StatusCode = int32(errno.FileServerErr.ErrCode)
		resp.StatusMsg = errno.FileServerErr.ErrMsg
		resp.Url = ""
		return resp, nil
	}

	resp.StatusCode = int32(errno.Success.ErrCode)
	resp.StatusMsg = errno.Success.ErrMsg
	resp.Url = url.String()

	return resp, nil
}
