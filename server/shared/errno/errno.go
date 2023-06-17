package errno

import "summer/server/shared/kitex_gen/errno"

type ErrNo struct {
	ErrCode int64
	ErrMsg  string
}

// NewErrNo return ErrNo
func NewErrNo(code int64, msg string) ErrNo {
	return ErrNo{
		ErrCode: code,
		ErrMsg:  msg,
	}
}

var (
	Success    = NewErrNo(int64(errno.Err_Success), "success")
	ParamsEr   = NewErrNo(int64(errno.Err_ParamsErr), "params err")
	ServiceErr = NewErrNo(int64(errno.Err_ServiceErr), "service err")

	RPCUserErr          = NewErrNo(int64(errno.Err_RPCUserErr), "rpc call user server error")
	UserServerErr       = NewErrNo(int64(errno.Err_UserServerErr), "user server error")
	UserAlreadyExistErr = NewErrNo(int64(errno.Err_UserAlreadyExistErr), "user already exist")
	UserNotFoundErr     = NewErrNo(int64(errno.Err_UserNotFoundErr), "user not found")
	AuthorizeFailErr    = NewErrNo(int64(errno.Err_AuthorizeFailErr), "authorize failed")

	RPCFileErr    = NewErrNo(int64(errno.Err_RPCFileErr), "rpc call file server error")
	FileServerErr = NewErrNo(int64(errno.Err_FileServerErr), "file server error")
)