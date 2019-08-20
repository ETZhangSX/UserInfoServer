package main

import (
	"errors"

	"LifeService"
)

//UserInfoServiceImp Interface implement
type UserInfoServiceImp struct {
	App *LifeService.DataService
	Obj string
}

//SignUp Create a new account
func (imp *UserInfoServiceImp) SignUp(wx_id string, userInfo *LifeService.UserInfo, RetCode *int32) (int32, error) {
	iRet, err := imp.App.CreateUser(wx_id, userInfo)
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
func (imp *UserInfoServiceImp) SignIn(wx_id string, sRsp *LifeService.UserInfo) (int32, error) {
	var hasUser *bool
	imp.App.HasUser(wx_id, hasUser)
	if *hasUser {
		_, err := imp.App.GetUserInfo(wx_id, sRsp)
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
func (imp *UserInfoServiceImp) GetUserPermissionInfo(wx_id string) (int32, error) {
	return 0, nil
}

//GetGroupList Get group list
func (imp *UserInfoServiceImp) GetGroupList(groupInfo *map[int32]string) (int32, error) {
	iRet, err := imp.App.GetGroupInfo(groupInfo)
	return iRet, err
}
