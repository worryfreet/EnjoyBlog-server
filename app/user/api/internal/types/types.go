// Code generated by goctl. DO NOT EDIT.
package types

type User struct {
	Id          int64  `json:"id"`
	UserId      string `json:"userId"`
	Username    string `json:"username"`
	Email       string `json:"email"`
	Phone       string `json:"phone"`
	Password    string `json:"-"`
	Avatar      string `json:"avatar"`
	NickName    string `json:"nickName"`
	Description string `json:"description"`
	CreateTime  string `json:"createTime"`
	UpdateTime  string `json:"updateTime"`
}

type RegisterReq struct {
	Username    string `json:"username"`
	Email       string `json:"email"`
	Phone       string `json:"phone"`
	Password    string `json:"password"`
	Avatar      string `json:"avatar"`
	NickName    string `json:"nickName"`
	Description string `json:"description"`
}

type LoginReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResp struct {
	User
	AccessToken  string `json:"access_token"`
	AccessExpire int64  `json:"access_expire"`
}

type UpdatePwdReq struct {
	Passwd    string `json:"passwd"`
	NewPasswd string `json:"newPasswd"`
	Forget    bool   `json:"forget"`
}

type UpdateInfoReq struct {
	Username    string `json:"username"`
	Phone       string `json:"phone"`
	Avatar      string `json:"avatar"`
	NickName    string `json:"nickName"`
	Description string `json:"description"`
}
