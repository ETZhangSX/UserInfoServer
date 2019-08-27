package main

import (
	"errors"
	"time"

	"LifeService"
)

//UserInfoServiceImp Interface implement
type UserInfoServiceImp struct {
	app     *LifeService.DataService
	objName string
}

//init initialize struct
func (imp *UserInfoServiceImp) init() {
	imp.app = new(LifeService.DataService)
	imp.objName = "LifeService.DataServer.DataServiceObj"
	comm.StringToProxy(imp.objName, imp.app);
}

//SignUp Create a new account
func (imp *UserInfoServiceImp) SignUp(Wx_id string, UserInfo *LifeService.UserInfo, RetCode *int32) (int32, error) {
	currentTime := time.Now().Format("2006-01-02 15:04:05")

	UserInfo.Group = 0
	UserInfo.Registration_time = currentTime

	iRet, err := imp.app.CreateUser(Wx_id, UserInfo)
	if err != nil {
		SLOG.Error("Create user error with error message: ", err)
		*RetCode = 404
	} else {
		SLOG.Debug("Create success")
		*RetCode = 200
	}
	return iRet, nil
}

//SignIn Judge if the user exist. If exist, return user info, otherwise, return error message
func (imp *UserInfoServiceImp) SignIn(Wx_id string, SRsp *LifeService.UserInfo) (int32, error) {
	var HasUser bool
	SLOG.Debug("SignIn")
	_, err := imp.app.HasUser(Wx_id, &HasUser)
	
	if err != nil {
		return -1, err
	}
	
	if HasUser {
		iRet, err := imp.app.GetUserInfo(Wx_id, SRsp)
		if err != nil {
			SLOG.Error("Call error: ", err)
		} else {
			SLOG.Debug("Call success")
		}
		return iRet, nil
	}
	return 404, errors.New("User not found")
}

//GetUserPermissionInfo Get user permission info
func (imp *UserInfoServiceImp) GetUserPermissionInfo(Wx_id string) (int32, error) {
	return 0, nil
}

//GetGroupList Get group list
func (imp *UserInfoServiceImp) GetGroupList(GroupInfo *map[int32]string) (int32, error) {
	SLOG.Debug("getGroupInfo")
	iRet, err := imp.app.GetGroupInfo(GroupInfo)
	return iRet, err
}

//Test test
func (imp *UserInfoServiceImp) Test(TestStr *string) (int32, error) {
	*TestStr = "Test Successfull"
	return 0, nil
}
