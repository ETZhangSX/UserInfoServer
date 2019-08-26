package main

import (
	"github.com/TarsCloud/TarsGo/tars"
	"github.com/TarsCloud/TarsGo/tars/util/rogger"

	"LifeService"
)

var comm *tars.Communicator
var log *rogger.Logger

func NewUserInfoServiceImp() *UserInfoServiceImp { //Init service
	app := new(LifeService.DataService)
	obj := "LifeService.DataServer.Obj"
	comm.StringToProxy(obj, app)
	return &UserInfoServiceImp{
		App: app,
		Obj: obj,
	}
}

func main() { //Init servant
	comm = tars.NewCommunicator()
	imp := NewUserInfoServiceImp()                                    		//New Imp
	app := new(LifeService.UserInfoService)                                 //New init the A Tars
	cfg := tars.GetServerConfig()                               			//Get Config File Object
	app.AddServant(imp, cfg.App+"."+cfg.Server+".UserInfoServiceObj") 		//Register Servant
	tars.Run()
}
