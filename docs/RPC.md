# 用户信息服务 RPC接口
以下类型在tars协议文件中定义，参照[tars数据类型定义文档]()
```
UserInfo          用户信息
ErrorCode         服务错误码
```

## 接口列表
* [HasPhone](#interface-hasphone)
* [SignUp](#interface-signup)
* [SignIn](#interface-signin)
* [GetGroupList](#interface-getgrouplist)
* [IsClubManager](#interface-isclubmanager)
* [IsInClub](#interface-isinclub)
* [IsAppliedActivity](#interface-isappliedactivity)
* [Test](#interface-test)

## 接口详情

### <a id="interface-hasphone"> HasPhone
判断手机号是否存在

**定义**
```cpp
int HasPhone(string phone, out bool phoneExist)
```

**参数**

|**属性**|**类型**|**说明**|
|-|-|-|
|phone|string|手机号|

**返回值**

|**属性**|**类型**|**说明**|
|-|-|-|
|phoneExist|bool|true表示手机号存在;false表示不存在|

### <a id="interface-signup"> SignUp
注册

**定义**
```cpp
int SignUp(string wxId, UserInfo userInfo, out ErrorCode errCode)
```

**参数**

|**属性**|**类型**|**说明**|
|-|-|-|
|wxId|string|用户id|
|userInfo|UserInfo|用户信息|

**返回值**

|**属性**|**类型**|**说明**|
|-|-|-|
|errCode|ErrorCode|服务错误码|

### <a id="interface-signin"> SignIn
登录

**定义**
```cpp
int SignIn(string wxId, out UserInfo userInfo, out ErrorCode errCode)
```

**参数**

|**属性**|**类型**|**说明**|
|-|-|-|
|wxId|string|用户id|

**返回值**

|**属性**|**类型**|**说明**|
|-|-|-|
|userInfo|UserInfo|用户信息|
|errCode|ErrorCode|服务错误码|

### <a id="interface-getgrouplist"> GetGroupList
获取用户权限组信息

**定义**
```cpp
int GetGroupList(out map<int, string> groupInfo)
```

**参数**

无

**返回值**

|**属性**|**类型**|**说明**|
|-|-|-|
|groupInfo|map<int, string>|用户权限组信息|

### <a id="interface-isclubmanager"> IsClubManager
判断是否是社团管理员

**定义**
```cpp
int IsClubManager(string wxId, string clubId, out bool isClubManager)
```

**参数**

|**属性**|**类型**|**说明**|
|-|-|-|
|wxId|string|用户id|
|clubId|string|社团id|

**返回值**

|**属性**|**类型**|**说明**|
|-|-|-|
|isClubManager|bool|true表示用户是社团管理员; false表示不是|

### <a id="interface-isinclub"> IsInClub
判断用户是否在社团中或者已申请社团

**定义**
```cpp
int IsInClub(string wxId, string clubId, bool justInClub, out bool isIn)
```

**参数**

|**属性**|**类型**|**说明**|
|-|-|-|
|wxId|string|用户id|
|clubId|string|社团id|
|justInClub|bool|true表示只判断用户是否在社团中; false表示判断用户是否申请社团, 包括已经在社团中的用户|

**返回值**

|**属性**|**类型**|**说明**|
|-|-|-|
|isIn|bool|true表示用户在社团中或已申请社团; false表示不在|

### <a id="interface-isappliedactivity"> IsAppliedActivity
判断用户是否已经报名活动

**定义**
```cpp
int IsAppliedActivity(string wxId, string activityId, out bool isApplied)
```

**参数**

|**属性**|**类型**|**说明**|
|-|-|-|
|wxId|string|用户id|
|activityId|string|活动id|

**返回值**

|**属性**|**类型**|**说明**|
|-|-|-|
|isApplied|bool|true表示用户已经报名活动; false表示用户未报名|

### <a id="interface-test"> Test

**定义**
```cpp
int Test(out string testStr)
```

**参数**

无

**返回值**

|**属性**|**类型**|**说明**|
|-|-|-|
|testStr|string|测试字符串|
