package main

import (
	"errors"
	"strconv"
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
	comm.StringToProxy(imp.objName, imp.app)
}

//SignUp Create a new account
func (imp *UserInfoServiceImp) SignUp(Wx_id string, UserInfo *LifeService.UserInfo, RetCode *int32) (int32, error) {
	currentTime := time.Now().Format("2006-01-02 15:04:05")

	UserInfo.Group = 0
	UserInfo.Registration_time = currentTime

	iRet, err := imp.app.CreateUser(Wx_id, UserInfo)
	if err != nil {
		SLOG.Error("Create user error with error message: ", err)
		*RetCode = 400
	} else {
		if iRet == 0 {
			SLOG.Debug("Create success")
			*RetCode = 200
		} else {
			SLOG.Debug("Create user fail: user exist.")
			*RetCode = 300
		}
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

//IsClubManager 是否是社团管理员
func (imp *UserInfoServiceImp) IsClubManager(WxId string, ClubId string, isClubManager *bool) (int32, error) {
	var sTableName = "club_managers"
	var sCondition = "where `wx_id='" + WxId + "' and `club_id`=" + ClubId
	var count int32
	_,err := imp.app.GetRecordCount(sTableName, sCondition, &count)

	if err != nil {
		SLOG.Error("UserInfoServer::IsClubManager error")
	} else {
		SLOG.Debug("UserInfoServer::IsClubManager: count=" + strconv.Itoa(int(count)))
		if count > 0 {
			*isClubManager = true
		} else if count == 0 {
			*isClubManager = false
		} else {
			*isClubManager = false
			return -1, errors.New("Get user manager info failed")
		}
	}

	return 0, nil
}

//IsInClub 用户是否在社团中或已经申请社团
func (imp *UserInfoServiceImp) IsInClub(WxId string, ClubId string, justInClub bool, isIn *bool) (int32, error) {
	var sTableName = "apply_for_club"
	var sCondition = "where `user_id`='" + WxId + "' and `club_id`=" + ClubId + " and `apply_status`"
	if justInClub {
		sCondition += "=0"
	} else {
		sCondition += "!=2"
	}
	var count int32
	_, err := imp.app.GetRecordCount(sTableName, sCondition, &count)

	if err != nil {
		SLOG.Error("UserInfoServer::IsInClub error")
	} else {
		SLOG.Debug("UserInfoServer::IsInClub: " + strconv.Itoa(int(count)))
		if count > 0 {
			*isIn = true
		} else if count == 0 {
			*isIn = false
		} else {
			*isIn = false
			return -1, errors.New("Get user club info failed")
		}
	}
	return 0, nil
}

//Test test
func (imp *UserInfoServiceImp) Test(TestStr *string) (int32, error) {
	*TestStr = "Test Successfull"
	return 0, nil
}
