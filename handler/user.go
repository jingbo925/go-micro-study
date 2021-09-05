package handler

import (
	"context"
	"go_micro_service/domain/model"
	"go_micro_service/domain/service"
	"go_micro_service/proto/user"
)

type User struct {
	UserDateService service.IUserDataService
}

// 注册

func (u *User) Register(ctx context.Context, userRegisterReq *user.UserRegisterReq, userRegisterResp *user.UserRegisterResp) error {
	userRegister := &model.User{
		UserName:     userRegisterReq.UserName,
		FirstName:    userRegisterReq.FirstName,
		HashPassword: userRegisterReq.Pwd,
	}
	_, err := u.UserDateService.AddUser(userRegister)

	if err != nil {
		return err
	}
	userRegisterResp.Message = "添加成功"
	return nil
}

func (u *User) Login(ctx context.Context, userLoginReq *user.UserLoginReq, userLoginResp *user.UserLoginResp) error {
	isOk, err := u.UserDateService.CheckPwd(userLoginReq.UserName, userLoginReq.Pwd)

	if err != nil {
		return err
	}
	userLoginResp.IsSuccess = isOk
	return nil
}

// 查询用户信息
func (u *User) GetUserInfo(ctx context.Context, userInfoReq *user.UserInfoReq, userInfoResp *user.UserInfoResp) error {
	userInfo, err := u.UserDateService.FindUserByName(userInfoReq.UserName)

	if err != nil {
		return err
	}

	userInfoResp = UserForResp(userInfo)
	return nil
}

func UserForResp(userModel *model.User) *user.UserInfoResp {
	resp := &user.UserInfoResp{}
	resp.UserName = userModel.UserName
	resp.FirstName = userModel.FirstName
	resp.UserId = userModel.ID
	return resp

}
