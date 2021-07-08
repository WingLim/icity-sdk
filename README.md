# iCity SDK

一个基于 Golang 的 iCity SDK。

## 使用

```go
package main

import  icity "github.com/WingLim/icity-sdk"

func main() {
	user := icity.Login("username", "password")
	user.NewDiary("", "Todoy~", icity.Public)
}
```

## API

### 登陆/登出
1. 登陆
```go
Login(username, password string, saveCookies bool) *User
```

2. 登出
```go
Logout(user *User) error
```

### 日记

```go
type Diary struct {
	// 日记 ID
	ID       string 
	// 发布的用户的昵称
	Nickname string 
	// 用户 ID
	UserId   string
	
	// 日记标题
	Title    string 
	// 日记内容
	Content  string 
	// 发布地点
	Location string 
	// 发布时间
	Date     time.Time
}
```

#### 日记权限
```go
const (
    // 公开
    Public DiaryPrivacy = iota + 1
    // 仅好友
    OnlyFriend
    // 私密
    Private
)
```

1. 发布日记
```go
user.NewDiary(title, content string, privacy DiaryPrivacy) Response

// 发布状态
Response.Success

// 发布的日记的 ID
Response.ActivityToken
```

2. 删除日记
```go
user.DeleteDiary(diaryID string) Response

// 删除状态
Response.Success

// 删除的日记的 ID
Response.ActivityToken
```

3. 喜欢日记
```go
// true: 成功
// false: 失败
user.Like(diaryID string) bool
```

4. 取消喜欢
```go
// true: 成功
// false: 失败 
user.UnLike(diaryID string) bool
```

### 评论

```go
type Comment struct {
	// 发布评论的用户的昵称
	Nickname string
	// 用户
	UserId   string
	// 评论内容
	Content  string 
	// 发布时间
	Date     time.Time
}
```

1. 发布评论
```go
user.NewComment(diaryID, comment string) Response

// 发布状态
Response.Success

// 发布的评论的 ID
Response.ActivityToken
```

2. 删除评论
```go
user.DeleteComment(commentID, diaryID string) Response

// 删除状态
Response.Success

// 删除的评论的 ID
Response.ActivityToken
```

3. 回复评论
```go
user.ReplyComment(userID, diaryID, comment string) Response

// 发布状态
Response.Success

// 发布的评论的 ID
Response.ActivityToken
```

4. 获取评论
```go
user.GetComments(diaryID string) []Comment
```

### 用户
```go
type User struct {
	// 用户名
	Username string
	// 用户密码
	Password string
	
	// 昵称
	Nickname string
	// 关于我
	Bio      string
	// 所在地
	Location string
    
	// 隐私设置
	SettingsPrivacy SettingsPrivacy
}
```

1. 设置昵称
```go
// true: 成功
// false: 失败 
user.SetNickName(nickname string) bool
```

2. 设置简介
```go
// true: 成功
// false: 失败 
user.SetBio(bio string) bool
```

3. 设置所在地
```go
// true: 成功
// false: 失败 
user.SetLocation(location string) bool
```

### 社交账号
1. 微博
```go
// true: 成功
// false: 失败 
user.SetSocialWeibo(weibo string) bool
```
...

更多社交账号设置请看 [settings_social.go](https://github.com/WingLim/icity-sdk/blob/main/settings_social.go)

### 隐私设置

```go
type SettingsPrivacy struct {
	// 新日记默认权限
	DefaultPrivacy DiaryPrivacy

	// 隐身模式
	InvisibleMode InvisibleType

	// 日记是否发布到世界
	UnWorld WorldType
    
	// 谁可以看
	// 关于我
	AboutMe ViewAccess

	// 我的朋友列表
	Followings ViewAccess

	// 我的关注者列表
	Followers ViewAccess
    
	// 互动设置
	// 评论我的日记
	AllowComment ViewAccess

	// 喜欢我的日记
	AllowLike ViewAccess

	// 给我发私信
	AllowMessage ViewAccess
}
```
