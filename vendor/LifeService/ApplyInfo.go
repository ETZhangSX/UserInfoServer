//Package LifeService comment
// This file war generated by tars2go 1.1
// Generated from DataService.tars
package LifeService

import (
	"fmt"
	"github.com/TarsCloud/TarsGo/tars/protocol/codec"
)

//ApplyInfo strcut implement
type ApplyInfo struct {
	Apply_id   string `json:"apply_id"`
	Apply_time string `json:"apply_time"`
	Wx_id      string `json:"wx_id"`
	Club_id    string `json:"club_id"`
	User_name  string `json:"user_name"`
	Club_name  string `json:"club_name"`
	Avatar_url string `json:"avatar_url"`
}

func (st *ApplyInfo) resetDefault() {
}

//ReadFrom reads  from _is and put into struct.
func (st *ApplyInfo) ReadFrom(_is *codec.Reader) error {
	var err error
	var length int32
	var have bool
	var ty byte
	st.resetDefault()

	err = _is.Read_string(&st.Apply_id, 0, true)
	if err != nil {
		return err
	}

	err = _is.Read_string(&st.Apply_time, 1, true)
	if err != nil {
		return err
	}

	err = _is.Read_string(&st.Wx_id, 2, true)
	if err != nil {
		return err
	}

	err = _is.Read_string(&st.Club_id, 3, true)
	if err != nil {
		return err
	}

	err = _is.Read_string(&st.User_name, 4, false)
	if err != nil {
		return err
	}

	err = _is.Read_string(&st.Club_name, 5, false)
	if err != nil {
		return err
	}

	err = _is.Read_string(&st.Avatar_url, 6, false)
	if err != nil {
		return err
	}

	_ = length
	_ = have
	_ = ty
	return nil
}

//ReadBlock reads struct from the given tag , require or optional.
func (st *ApplyInfo) ReadBlock(_is *codec.Reader, tag byte, require bool) error {
	var err error
	var have bool
	st.resetDefault()

	err, have = _is.SkipTo(codec.STRUCT_BEGIN, tag, require)
	if err != nil {
		return err
	}
	if !have {
		if require {
			return fmt.Errorf("require ApplyInfo, but not exist. tag %d", tag)
		}
		return nil

	}

	st.ReadFrom(_is)

	err = _is.SkipToStructEnd()
	if err != nil {
		return err
	}
	_ = have
	return nil
}

//WriteTo encode struct to buffer
func (st *ApplyInfo) WriteTo(_os *codec.Buffer) error {
	var err error

	err = _os.Write_string(st.Apply_id, 0)
	if err != nil {
		return err
	}

	err = _os.Write_string(st.Apply_time, 1)
	if err != nil {
		return err
	}

	err = _os.Write_string(st.Wx_id, 2)
	if err != nil {
		return err
	}

	err = _os.Write_string(st.Club_id, 3)
	if err != nil {
		return err
	}

	err = _os.Write_string(st.User_name, 4)
	if err != nil {
		return err
	}

	err = _os.Write_string(st.Club_name, 5)
	if err != nil {
		return err
	}

	err = _os.Write_string(st.Avatar_url, 6)
	if err != nil {
		return err
	}

	return nil
}

//WriteBlock encode struct
func (st *ApplyInfo) WriteBlock(_os *codec.Buffer, tag byte) error {
	var err error
	err = _os.WriteHead(codec.STRUCT_BEGIN, tag)
	if err != nil {
		return err
	}

	st.WriteTo(_os)

	err = _os.WriteHead(codec.STRUCT_END, 0)
	if err != nil {
		return err
	}
	return nil
}
