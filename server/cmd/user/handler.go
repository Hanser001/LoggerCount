package main

import (
	"context"
	"github.com/bwmarrin/snowflake"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/golang-jwt/jwt"
	models "summer/server/cmd/api/model"
	"summer/server/cmd/user/config"
	"summer/server/cmd/user/dao"
	"summer/server/cmd/user/model"
	"summer/server/cmd/user/pkg/md5"
	"summer/server/shared/consts"
	"summer/server/shared/errno"
	"summer/server/shared/kitex_gen/user"
	"summer/server/shared/middwares"
	"summer/server/shared/middwares/casbin"
	"time"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct {
	Jwt *middwares.JWT
	Dao *dao.UserManger
}

// Register implements the UserServiceImpl interface.
func (s *UserServiceImpl) Register(_ context.Context, req *user.RegisterRequest) (resp *user.RegisterResponse, err error) {
	resp = new(user.RegisterResponse)

	sf, err := snowflake.NewNode(consts.UserSnowflakeNode)
	if err != nil {
		klog.Errorf("generate user id failed: %s", err.Error())
		resp.StatusCode = int32(errno.UserServerErr.ErrCode)
		resp.StatusMsg = errno.UserServerErr.ErrMsg
		return resp, nil
	}

	usr := &model.User{
		ID:       sf.Generate().Int64(),
		Username: req.Username,
		Password: md5.Md5Crypt(req.Password, config.GlobalServerConfig.MysqlInfo.Salt),
	}

	if err = s.Dao.CreateUser(usr); err != nil {
		if err == dao.ErrUserExist {
			resp.StatusCode = int32(errno.UserAlreadyExistErr.ErrCode)
			resp.StatusMsg = errno.UserAlreadyExistErr.ErrMsg
		} else {
			klog.Error("create user error", err)
			resp.StatusCode = int32(errno.UserServerErr.ErrCode)
			resp.StatusMsg = "create user error"
		}
		return resp, nil
	}

	//When creates user,adding "user" grouping policy to new user
	e := casbin.GetEnforcer()
	_, err = e.AddGroupingPolicy(usr.ID, "user")
	if err != nil {
		klog.Error("add user policy", err)
		resp.StatusCode = int32(errno.UserServerErr.ErrCode)
		resp.StatusMsg = "add user policy err"
		return resp, nil
	}

	resp.StatusCode = int32(errno.Success.ErrCode)
	resp.StatusMsg = errno.Success.ErrMsg

	return resp, nil
}

// Login implements the UserServiceImpl interface.
func (s *UserServiceImpl) Login(_ context.Context, req *user.LoginRequest) (resp *user.LoginResponse, err error) {
	resp = new(user.LoginResponse)
	usr, err := s.Dao.GetUserByUsername(req.Username)
	if err != nil {
		if err == dao.ErrNoSuchUser {
			resp.StatusCode = int32(errno.UserNotFoundErr.ErrCode)
			resp.StatusMsg = errno.UserNotFoundErr.ErrMsg
		} else {
			klog.Errorf("get user by name err", err)
			resp.StatusCode = int32(errno.UserServerErr.ErrCode)
			resp.StatusMsg = "get user by name err"
		}
		return resp, nil
	}

	if usr.Password != md5.Md5Crypt(req.Password, config.GlobalServerConfig.MysqlInfo.Salt) {
		resp.StatusCode = int32(errno.UserNotFoundErr.ErrCode)
		resp.StatusMsg = "wrong password"
		return resp, nil
	}

	resp.Token, err = s.Jwt.CreateToken(models.CustomClaims{
		ID: usr.ID,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix(),
			ExpiresAt: time.Now().Unix() + consts.ThirtyDays,
			Issuer:    consts.JWTIssuer,
		},
	})
	if err != nil {
		klog.Error("create token err", err)
		resp.StatusCode = int32(errno.UserServerErr.ErrCode)
		resp.StatusMsg = errno.UserServerErr.ErrMsg
		return resp, nil
	}

	resp.StatusCode = int32(errno.Success.ErrCode)
	resp.StatusMsg = errno.Success.ErrMsg
	return resp, nil
}
