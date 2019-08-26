package main

import (
	"time"
	"errors"

	"LifeService"
)

//UserInfoServiceImp Interface implement
type UserInfoServiceImp struct {
	App *LifeService.DataService
	Obj string
}

//SignUp Create a new account
func (imp *UserInfoServiceImp) SignUp(Wx_id string, UserInfo *LifeService.UserInfo, RetCode *int32) (int32, error) {
	currentTime := time.Now().Format("2006-01-02 15:04:05")

	UserInfo.Group = 0
	UserInfo.Registration_time = currentTime
	
	iRet, err := imp.App.CreateUser(Wx_id, UserInfo)
	if err != nil {
		log.Error("Create user error with error message: ", err)
		*RetCode = 404
	}else{
		log.Debug("Create success")
		*RetCode = 200
	}
	return iRet, nil
}

//SignIn Judge if the user exist. If exist, return user info, otherwise, return error message
func (imp *UserInfoServiceImp) SignIn(Wx_id string, SRsp *LifeService.UserInfo) (int32, error) {
	var hasUser *bool
	imp.App.HasUser(Wx_id, hasUser)
	if *hasUser {
		_, err := imp.App.GetUserInfo(Wx_id, SRsp)
		if err != nil {
			log.Error("Call error: ", err)
		}else{
			log.Debug("Call success")
		}
		return 0, nil
	}
	return -1, errors.New("User not found")
}

//GetUserPermissionInfo Get user permission info
func (imp *UserInfoServiceImp) GetUserPermissionInfo(Wx_id string) (int32, error) {
	return 0, nil
}

//GetGroupList Get group list
func (imp *UserInfoServiceImp) GetGroupList(GroupInfo *map[int32]string) (int32, error) {
	iRet, err := imp.App.GetGroupInfo(GroupInfo)
	return iRet, err
}

//Test test
func (imp *UserInfoServiceImp) Test(TestStr *string) (int32, error) {
	*TestStr = "Test Successfull"
	return 0, nil
}