package main

import (
    "github.com/TarsCloud/TarsGo/tars"
    "github.com/TarsCloud/TarsGo/tars/util/rogger"

    "LifeService"
)

// comm 定义communicator
var comm *tars.Communicator

//SLOG 日志打印
var SLOG = rogger.GetLogger("ServerLog")

func main() { //Init servant
    comm = tars.NewCommunicator()                                     //初始化communicator
    imp := new(UserInfoServiceImp)                                    //New Imp
    imp.init()
    app := new(LifeService.UserInfoService)                           //New init the A Tars
    cfg := tars.GetServerConfig()                                     //Get Config File Object
    app.AddServant(imp, cfg.App+"."+cfg.Server+".UserInfoServiceObj") //Register Servant
    tars.Run()
}
