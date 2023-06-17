package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/cloudwego/kitex/pkg/klog"
	"summer/server/cmd/analyze/pkg"
	"summer/server/cmd/analyze/pkg/mapreduce"
	"summer/server/shared/consts"
	"summer/server/shared/errno"
	analyze "summer/server/shared/kitex_gen/analyze"
)

// AnalyzeServiceImpl implements the last service interface defined in the IDL.
type AnalyzeServiceImpl struct {
	MinioManger
}

type MinioManger interface {
	UploadFinalData(ctx context.Context, filepath string) error
}

// Analyze implements the AnalyzeServiceImpl interface.
func (s *AnalyzeServiceImpl) Analyze(ctx context.Context, req *analyze.AnalyzeRequest) (resp *analyze.AnalyzeResponse, err error) {
	resp = new(analyze.AnalyzeResponse)

	data1, err := pkg.DownloadFile(req.Url)
	if err != nil {
		klog.Error("download err,error", err)
		resp.StatusCode = int32(errno.FileServerErr.ErrCode)
		resp.StatusMsg = errno.FileServerErr.ErrMsg
		return resp, nil
	}

	data2, err := mapreduce.StartWordCount(data1, req.Filed)
	if err != nil {
		klog.Error("count err,error", err)
		resp.StatusCode = int32(errno.FileServerErr.ErrCode)
		resp.StatusMsg = errno.FileServerErr.ErrMsg
		return resp, nil
	}

	data3, err := json.Marshal(data2)
	if err != nil {
		klog.Error("count err,error", err)
		resp.StatusCode = int32(errno.FileServerErr.ErrCode)
		resp.StatusMsg = errno.FileServerErr.ErrMsg
		return resp, nil
	}

	pathToPut := fmt.Sprintf("%dput.json", req.UserId)
	err = pkg.NewFile(pathToPut, data3)
	if err != nil {
		klog.Error("count err,error", err)
		resp.StatusCode = int32(errno.FileServerErr.ErrCode)
		resp.StatusMsg = errno.FileServerErr.ErrMsg
		return resp, nil
	}

	filepath := fmt.Sprintf(consts.NewLocalFilePath, req.UserId)
	err = s.MinioManger.UploadFinalData(ctx, filepath)
	if err != nil {
		klog.Error("count err,error", err)
		resp.StatusCode = int32(errno.FileServerErr.ErrCode)
		resp.StatusMsg = errno.FileServerErr.ErrMsg
		return resp, nil
	}

	resp.StatusCode = int32(errno.Success.ErrCode)
	resp.StatusMsg = errno.Success.ErrMsg
	return resp, nil
}
