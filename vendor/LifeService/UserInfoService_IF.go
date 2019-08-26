//Package LifeService comment
// This file war generated by tars2go 1.1
// Generated from UserInfoService.tars
package LifeService

import (
	"context"
	"fmt"
	"github.com/TarsCloud/TarsGo/tars"
	m "github.com/TarsCloud/TarsGo/tars/model"
	"github.com/TarsCloud/TarsGo/tars/protocol/codec"
	"github.com/TarsCloud/TarsGo/tars/protocol/res/requestf"
	"github.com/TarsCloud/TarsGo/tars/util/current"
	"github.com/TarsCloud/TarsGo/tars/util/tools"
)

//UserInfoService struct
type UserInfoService struct {
	s m.Servant
}

//SignUp is the proxy function for the method defined in the tars file, with the context
func (_obj *UserInfoService) SignUp(Wx_id string, UserInfo *UserInfo, RetCode *int32, _opt ...map[string]string) (ret int32, err error) {

	var length int32
	var have bool
	var ty byte
	_os := codec.NewBuffer()
	err = _os.Write_string(Wx_id, 1)
	if err != nil {
		return ret, err
	}

	err = UserInfo.WriteBlock(_os, 2)
	if err != nil {
		return ret, err
	}

	var _status map[string]string
	var _context map[string]string
	if len(_opt) == 1 {
		_context = _opt[0]
	} else if len(_opt) == 2 {
		_context = _opt[0]
		_status = _opt[1]
	}
	_resp := new(requestf.ResponsePacket)
	ctx := context.Background()
	err = _obj.s.Tars_invoke(ctx, 0, "SignUp", _os.ToBytes(), _status, _context, _resp)
	if err != nil {
		return ret, err
	}
	_is := codec.NewReader(tools.Int8ToByte(_resp.SBuffer))
	err = _is.Read_int32(&ret, 0, true)
	if err != nil {
		return ret, err
	}

	err = _is.Read_int32(&(*RetCode), 3, true)
	if err != nil {
		return ret, err
	}

	_obj.setMap(len(_opt), _resp, _context, _status)
	_ = length
	_ = have
	_ = ty
	return ret, nil
}

//SignUpWithContext is the proxy function for the method defined in the tars file, with the context
func (_obj *UserInfoService) SignUpWithContext(ctx context.Context, Wx_id string, UserInfo *UserInfo, RetCode *int32, _opt ...map[string]string) (ret int32, err error) {

	var length int32
	var have bool
	var ty byte
	_os := codec.NewBuffer()
	err = _os.Write_string(Wx_id, 1)
	if err != nil {
		return ret, err
	}

	err = UserInfo.WriteBlock(_os, 2)
	if err != nil {
		return ret, err
	}

	var _status map[string]string
	var _context map[string]string
	if len(_opt) == 1 {
		_context = _opt[0]
	} else if len(_opt) == 2 {
		_context = _opt[0]
		_status = _opt[1]
	}
	_resp := new(requestf.ResponsePacket)
	err = _obj.s.Tars_invoke(ctx, 0, "SignUp", _os.ToBytes(), _status, _context, _resp)
	if err != nil {
		return ret, err
	}
	_is := codec.NewReader(tools.Int8ToByte(_resp.SBuffer))
	err = _is.Read_int32(&ret, 0, true)
	if err != nil {
		return ret, err
	}

	err = _is.Read_int32(&(*RetCode), 3, true)
	if err != nil {
		return ret, err
	}

	_obj.setMap(len(_opt), _resp, _context, _status)
	_ = length
	_ = have
	_ = ty
	return ret, nil
}

//SignIn is the proxy function for the method defined in the tars file, with the context
func (_obj *UserInfoService) SignIn(Wx_id string, SRsp *UserInfo, _opt ...map[string]string) (ret int32, err error) {

	var length int32
	var have bool
	var ty byte
	_os := codec.NewBuffer()
	err = _os.Write_string(Wx_id, 1)
	if err != nil {
		return ret, err
	}

	var _status map[string]string
	var _context map[string]string
	if len(_opt) == 1 {
		_context = _opt[0]
	} else if len(_opt) == 2 {
		_context = _opt[0]
		_status = _opt[1]
	}
	_resp := new(requestf.ResponsePacket)
	ctx := context.Background()
	err = _obj.s.Tars_invoke(ctx, 0, "SignIn", _os.ToBytes(), _status, _context, _resp)
	if err != nil {
		return ret, err
	}
	_is := codec.NewReader(tools.Int8ToByte(_resp.SBuffer))
	err = _is.Read_int32(&ret, 0, true)
	if err != nil {
		return ret, err
	}

	err = (*SRsp).ReadBlock(_is, 2, true)
	if err != nil {
		return ret, err
	}

	_obj.setMap(len(_opt), _resp, _context, _status)
	_ = length
	_ = have
	_ = ty
	return ret, nil
}

//SignInWithContext is the proxy function for the method defined in the tars file, with the context
func (_obj *UserInfoService) SignInWithContext(ctx context.Context, Wx_id string, SRsp *UserInfo, _opt ...map[string]string) (ret int32, err error) {

	var length int32
	var have bool
	var ty byte
	_os := codec.NewBuffer()
	err = _os.Write_string(Wx_id, 1)
	if err != nil {
		return ret, err
	}

	var _status map[string]string
	var _context map[string]string
	if len(_opt) == 1 {
		_context = _opt[0]
	} else if len(_opt) == 2 {
		_context = _opt[0]
		_status = _opt[1]
	}
	_resp := new(requestf.ResponsePacket)
	err = _obj.s.Tars_invoke(ctx, 0, "SignIn", _os.ToBytes(), _status, _context, _resp)
	if err != nil {
		return ret, err
	}
	_is := codec.NewReader(tools.Int8ToByte(_resp.SBuffer))
	err = _is.Read_int32(&ret, 0, true)
	if err != nil {
		return ret, err
	}

	err = (*SRsp).ReadBlock(_is, 2, true)
	if err != nil {
		return ret, err
	}

	_obj.setMap(len(_opt), _resp, _context, _status)
	_ = length
	_ = have
	_ = ty
	return ret, nil
}

//GetUserPermissionInfo is the proxy function for the method defined in the tars file, with the context
func (_obj *UserInfoService) GetUserPermissionInfo(Wx_id string, _opt ...map[string]string) (ret int32, err error) {

	var length int32
	var have bool
	var ty byte
	_os := codec.NewBuffer()
	err = _os.Write_string(Wx_id, 1)
	if err != nil {
		return ret, err
	}

	var _status map[string]string
	var _context map[string]string
	if len(_opt) == 1 {
		_context = _opt[0]
	} else if len(_opt) == 2 {
		_context = _opt[0]
		_status = _opt[1]
	}
	_resp := new(requestf.ResponsePacket)
	ctx := context.Background()
	err = _obj.s.Tars_invoke(ctx, 0, "GetUserPermissionInfo", _os.ToBytes(), _status, _context, _resp)
	if err != nil {
		return ret, err
	}
	_is := codec.NewReader(tools.Int8ToByte(_resp.SBuffer))
	err = _is.Read_int32(&ret, 0, true)
	if err != nil {
		return ret, err
	}

	_obj.setMap(len(_opt), _resp, _context, _status)
	_ = length
	_ = have
	_ = ty
	return ret, nil
}

//GetUserPermissionInfoWithContext is the proxy function for the method defined in the tars file, with the context
func (_obj *UserInfoService) GetUserPermissionInfoWithContext(ctx context.Context, Wx_id string, _opt ...map[string]string) (ret int32, err error) {

	var length int32
	var have bool
	var ty byte
	_os := codec.NewBuffer()
	err = _os.Write_string(Wx_id, 1)
	if err != nil {
		return ret, err
	}

	var _status map[string]string
	var _context map[string]string
	if len(_opt) == 1 {
		_context = _opt[0]
	} else if len(_opt) == 2 {
		_context = _opt[0]
		_status = _opt[1]
	}
	_resp := new(requestf.ResponsePacket)
	err = _obj.s.Tars_invoke(ctx, 0, "GetUserPermissionInfo", _os.ToBytes(), _status, _context, _resp)
	if err != nil {
		return ret, err
	}
	_is := codec.NewReader(tools.Int8ToByte(_resp.SBuffer))
	err = _is.Read_int32(&ret, 0, true)
	if err != nil {
		return ret, err
	}

	_obj.setMap(len(_opt), _resp, _context, _status)
	_ = length
	_ = have
	_ = ty
	return ret, nil
}

//GetGroupList is the proxy function for the method defined in the tars file, with the context
func (_obj *UserInfoService) GetGroupList(GroupInfo *map[int32]string, _opt ...map[string]string) (ret int32, err error) {

	var length int32
	var have bool
	var ty byte
	_os := codec.NewBuffer()

	var _status map[string]string
	var _context map[string]string
	if len(_opt) == 1 {
		_context = _opt[0]
	} else if len(_opt) == 2 {
		_context = _opt[0]
		_status = _opt[1]
	}
	_resp := new(requestf.ResponsePacket)
	ctx := context.Background()
	err = _obj.s.Tars_invoke(ctx, 0, "GetGroupList", _os.ToBytes(), _status, _context, _resp)
	if err != nil {
		return ret, err
	}
	_is := codec.NewReader(tools.Int8ToByte(_resp.SBuffer))
	err = _is.Read_int32(&ret, 0, true)
	if err != nil {
		return ret, err
	}

	err, have = _is.SkipTo(codec.MAP, 1, true)
	if err != nil {
		return ret, err
	}

	err = _is.Read_int32(&length, 0, true)
	if err != nil {
		return ret, err
	}
	(*GroupInfo) = make(map[int32]string)
	for i0, e0 := int32(0), length; i0 < e0; i0++ {
		var k0 int32
		var v0 string

		err = _is.Read_int32(&k0, 0, false)
		if err != nil {
			return ret, err
		}

		err = _is.Read_string(&v0, 1, false)
		if err != nil {
			return ret, err
		}

		(*GroupInfo)[k0] = v0
	}

	_obj.setMap(len(_opt), _resp, _context, _status)
	_ = length
	_ = have
	_ = ty
	return ret, nil
}

//GetGroupListWithContext is the proxy function for the method defined in the tars file, with the context
func (_obj *UserInfoService) GetGroupListWithContext(ctx context.Context, GroupInfo *map[int32]string, _opt ...map[string]string) (ret int32, err error) {

	var length int32
	var have bool
	var ty byte
	_os := codec.NewBuffer()
	var _status map[string]string
	var _context map[string]string
	if len(_opt) == 1 {
		_context = _opt[0]
	} else if len(_opt) == 2 {
		_context = _opt[0]
		_status = _opt[1]
	}
	_resp := new(requestf.ResponsePacket)
	err = _obj.s.Tars_invoke(ctx, 0, "GetGroupList", _os.ToBytes(), _status, _context, _resp)
	if err != nil {
		return ret, err
	}
	_is := codec.NewReader(tools.Int8ToByte(_resp.SBuffer))
	err = _is.Read_int32(&ret, 0, true)
	if err != nil {
		return ret, err
	}

	err, have = _is.SkipTo(codec.MAP, 1, true)
	if err != nil {
		return ret, err
	}

	err = _is.Read_int32(&length, 0, true)
	if err != nil {
		return ret, err
	}
	(*GroupInfo) = make(map[int32]string)
	for i1, e1 := int32(0), length; i1 < e1; i1++ {
		var k1 int32
		var v1 string

		err = _is.Read_int32(&k1, 0, false)
		if err != nil {
			return ret, err
		}

		err = _is.Read_string(&v1, 1, false)
		if err != nil {
			return ret, err
		}

		(*GroupInfo)[k1] = v1
	}

	_obj.setMap(len(_opt), _resp, _context, _status)
	_ = length
	_ = have
	_ = ty
	return ret, nil
}

//Test is the proxy function for the method defined in the tars file, with the context
func (_obj *UserInfoService) Test(TestStr *string, _opt ...map[string]string) (ret int32, err error) {

	var length int32
	var have bool
	var ty byte
	_os := codec.NewBuffer()

	var _status map[string]string
	var _context map[string]string
	if len(_opt) == 1 {
		_context = _opt[0]
	} else if len(_opt) == 2 {
		_context = _opt[0]
		_status = _opt[1]
	}
	_resp := new(requestf.ResponsePacket)
	ctx := context.Background()
	err = _obj.s.Tars_invoke(ctx, 0, "Test", _os.ToBytes(), _status, _context, _resp)
	if err != nil {
		return ret, err
	}
	_is := codec.NewReader(tools.Int8ToByte(_resp.SBuffer))
	err = _is.Read_int32(&ret, 0, true)
	if err != nil {
		return ret, err
	}

	err = _is.Read_string(&(*TestStr), 1, true)
	if err != nil {
		return ret, err
	}

	_obj.setMap(len(_opt), _resp, _context, _status)
	_ = length
	_ = have
	_ = ty
	return ret, nil
}

//TestWithContext is the proxy function for the method defined in the tars file, with the context
func (_obj *UserInfoService) TestWithContext(ctx context.Context, TestStr *string, _opt ...map[string]string) (ret int32, err error) {

	var length int32
	var have bool
	var ty byte
	_os := codec.NewBuffer()
	var _status map[string]string
	var _context map[string]string
	if len(_opt) == 1 {
		_context = _opt[0]
	} else if len(_opt) == 2 {
		_context = _opt[0]
		_status = _opt[1]
	}
	_resp := new(requestf.ResponsePacket)
	err = _obj.s.Tars_invoke(ctx, 0, "Test", _os.ToBytes(), _status, _context, _resp)
	if err != nil {
		return ret, err
	}
	_is := codec.NewReader(tools.Int8ToByte(_resp.SBuffer))
	err = _is.Read_int32(&ret, 0, true)
	if err != nil {
		return ret, err
	}

	err = _is.Read_string(&(*TestStr), 1, true)
	if err != nil {
		return ret, err
	}

	_obj.setMap(len(_opt), _resp, _context, _status)
	_ = length
	_ = have
	_ = ty
	return ret, nil
}

//SetServant sets servant for the service.
func (_obj *UserInfoService) SetServant(s m.Servant) {
	_obj.s = s
}

//TarsSetTimeout sets the timeout for the servant which is in ms.
func (_obj *UserInfoService) TarsSetTimeout(t int) {
	_obj.s.TarsSetTimeout(t)
}
func (_obj *UserInfoService) setMap(l int, res *requestf.ResponsePacket, ctx map[string]string, sts map[string]string) {
	if l == 1 {
		for k, _ := range ctx {
			delete(ctx, k)
		}
		for k, v := range res.Context {
			ctx[k] = v
		}
	} else if l == 2 {
		for k, _ := range ctx {
			delete(ctx, k)
		}
		for k, v := range res.Context {
			ctx[k] = v
		}
		for k, _ := range sts {
			delete(sts, k)
		}
		for k, v := range res.Status {
			sts[k] = v
		}
	}
}

//AddServant adds servant  for the service.
func (_obj *UserInfoService) AddServant(imp _impUserInfoService, obj string) {
	tars.AddServant(_obj, imp, obj)
}

//AddServant adds servant  for the service with context.
func (_obj *UserInfoService) AddServantWithContext(imp _impUserInfoServiceWithContext, obj string) {
	tars.AddServantWithContext(_obj, imp, obj)
}

type _impUserInfoService interface {
	SignUp(Wx_id string, UserInfo *UserInfo, RetCode *int32) (ret int32, err error)
	SignIn(Wx_id string, SRsp *UserInfo) (ret int32, err error)
	GetUserPermissionInfo(Wx_id string) (ret int32, err error)
	GetGroupList(GroupInfo *map[int32]string) (ret int32, err error)
	Test(TestStr *string) (ret int32, err error)
}
type _impUserInfoServiceWithContext interface {
	SignUp(ctx context.Context, Wx_id string, UserInfo *UserInfo, RetCode *int32) (ret int32, err error)
	SignIn(ctx context.Context, Wx_id string, SRsp *UserInfo) (ret int32, err error)
	GetUserPermissionInfo(ctx context.Context, Wx_id string) (ret int32, err error)
	GetGroupList(ctx context.Context, GroupInfo *map[int32]string) (ret int32, err error)
	Test(ctx context.Context, TestStr *string) (ret int32, err error)
}

func SignUp(ctx context.Context, _val interface{}, _os *codec.Buffer, _is *codec.Reader, withContext bool) (err error) {
	var length int32
	var have bool
	var ty byte
	var Wx_id string
	err = _is.Read_string(&Wx_id, 1, true)
	if err != nil {
		return err
	}
	var UserInfo UserInfo
	err = UserInfo.ReadBlock(_is, 2, true)
	if err != nil {
		return err
	}
	var RetCode int32
	if withContext == false {
		_imp := _val.(_impUserInfoService)
		ret, err := _imp.SignUp(Wx_id, &UserInfo, &RetCode)
		if err != nil {
			return err
		}

		err = _os.Write_int32(ret, 0)
		if err != nil {
			return err
		}
	} else {
		_imp := _val.(_impUserInfoServiceWithContext)
		ret, err := _imp.SignUp(ctx, Wx_id, &UserInfo, &RetCode)
		if err != nil {
			return err
		}

		err = _os.Write_int32(ret, 0)
		if err != nil {
			return err
		}
	}

	err = _os.Write_int32(RetCode, 3)
	if err != nil {
		return err
	}

	_ = length
	_ = have
	_ = ty
	return nil
}
func SignIn(ctx context.Context, _val interface{}, _os *codec.Buffer, _is *codec.Reader, withContext bool) (err error) {
	var length int32
	var have bool
	var ty byte
	var Wx_id string
	err = _is.Read_string(&Wx_id, 1, true)
	if err != nil {
		return err
	}
	var SRsp UserInfo
	if withContext == false {
		_imp := _val.(_impUserInfoService)
		ret, err := _imp.SignIn(Wx_id, &SRsp)
		if err != nil {
			return err
		}

		err = _os.Write_int32(ret, 0)
		if err != nil {
			return err
		}
	} else {
		_imp := _val.(_impUserInfoServiceWithContext)
		ret, err := _imp.SignIn(ctx, Wx_id, &SRsp)
		if err != nil {
			return err
		}

		err = _os.Write_int32(ret, 0)
		if err != nil {
			return err
		}
	}

	err = SRsp.WriteBlock(_os, 2)
	if err != nil {
		return err
	}

	_ = length
	_ = have
	_ = ty
	return nil
}
func GetUserPermissionInfo(ctx context.Context, _val interface{}, _os *codec.Buffer, _is *codec.Reader, withContext bool) (err error) {
	var length int32
	var have bool
	var ty byte
	var Wx_id string
	err = _is.Read_string(&Wx_id, 1, true)
	if err != nil {
		return err
	}
	if withContext == false {
		_imp := _val.(_impUserInfoService)
		ret, err := _imp.GetUserPermissionInfo(Wx_id)
		if err != nil {
			return err
		}

		err = _os.Write_int32(ret, 0)
		if err != nil {
			return err
		}
	} else {
		_imp := _val.(_impUserInfoServiceWithContext)
		ret, err := _imp.GetUserPermissionInfo(ctx, Wx_id)
		if err != nil {
			return err
		}

		err = _os.Write_int32(ret, 0)
		if err != nil {
			return err
		}
	}

	_ = length
	_ = have
	_ = ty
	return nil
}
func GetGroupList(ctx context.Context, _val interface{}, _os *codec.Buffer, _is *codec.Reader, withContext bool) (err error) {
	var length int32
	var have bool
	var ty byte
	var GroupInfo map[int32]string
	if withContext == false {
		_imp := _val.(_impUserInfoService)
		ret, err := _imp.GetGroupList(&GroupInfo)
		if err != nil {
			return err
		}

		err = _os.Write_int32(ret, 0)
		if err != nil {
			return err
		}
	} else {
		_imp := _val.(_impUserInfoServiceWithContext)
		ret, err := _imp.GetGroupList(ctx, &GroupInfo)
		if err != nil {
			return err
		}

		err = _os.Write_int32(ret, 0)
		if err != nil {
			return err
		}
	}

	err = _os.WriteHead(codec.MAP, 1)
	if err != nil {
		return err
	}
	err = _os.Write_int32(int32(len(GroupInfo)), 0)
	if err != nil {
		return err
	}
	for k2, v2 := range GroupInfo {

		err = _os.Write_int32(k2, 0)
		if err != nil {
			return err
		}

		err = _os.Write_string(v2, 1)
		if err != nil {
			return err
		}
	}

	_ = length
	_ = have
	_ = ty
	return nil
}
func Test(ctx context.Context, _val interface{}, _os *codec.Buffer, _is *codec.Reader, withContext bool) (err error) {
	var length int32
	var have bool
	var ty byte
	var TestStr string
	if withContext == false {
		_imp := _val.(_impUserInfoService)
		ret, err := _imp.Test(&TestStr)
		if err != nil {
			return err
		}

		err = _os.Write_int32(ret, 0)
		if err != nil {
			return err
		}
	} else {
		_imp := _val.(_impUserInfoServiceWithContext)
		ret, err := _imp.Test(ctx, &TestStr)
		if err != nil {
			return err
		}

		err = _os.Write_int32(ret, 0)
		if err != nil {
			return err
		}
	}

	err = _os.Write_string(TestStr, 1)
	if err != nil {
		return err
	}

	_ = length
	_ = have
	_ = ty
	return nil
}

//Dispatch is used to call the server side implemnet for the method defined in the tars file. withContext shows using context or not.
func (_obj *UserInfoService) Dispatch(ctx context.Context, _val interface{}, req *requestf.RequestPacket, resp *requestf.ResponsePacket, withContext bool) (err error) {
	_is := codec.NewReader(tools.Int8ToByte(req.SBuffer))
	_os := codec.NewBuffer()
	switch req.SFuncName {
	case "SignUp":
		err := SignUp(ctx, _val, _os, _is, withContext)
		if err != nil {
			return err
		}
	case "SignIn":
		err := SignIn(ctx, _val, _os, _is, withContext)
		if err != nil {
			return err
		}
	case "GetUserPermissionInfo":
		err := GetUserPermissionInfo(ctx, _val, _os, _is, withContext)
		if err != nil {
			return err
		}
	case "GetGroupList":
		err := GetGroupList(ctx, _val, _os, _is, withContext)
		if err != nil {
			return err
		}
	case "Test":
		err := Test(ctx, _val, _os, _is, withContext)
		if err != nil {
			return err
		}

	default:
		return fmt.Errorf("func mismatch")
	}
	var _status map[string]string
	s, ok := current.GetResponseStatus(ctx)
	if ok && s != nil {
		_status = s
	}
	var _context map[string]string
	c, ok := current.GetResponseContext(ctx)
	if ok && c != nil {
		_context = c
	}
	*resp = requestf.ResponsePacket{
		IVersion:     1,
		CPacketType:  0,
		IRequestId:   req.IRequestId,
		IMessageType: 0,
		IRet:         0,
		SBuffer:      tools.ByteToInt8(_os.ToBytes()),
		Status:       _status,
		SResultDesc:  "",
		Context:      _context,
	}
	return nil
}
