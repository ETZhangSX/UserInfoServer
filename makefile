APP       := LifeService
TARGET    := UserInfoServer
MFLAGS    :=
DFLAGS    :=
CONFIG    := client
STRIP_FLAG:= N
J2GO_FLAG:= 

libpath=${subst :, ,$(GOPATH)}
$(foreach path,$(libpath),$(eval -include $(path)/src/LifeService/UserInfoServer/makefile.tars))
