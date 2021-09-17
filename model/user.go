package model

// 用户结构体
type User struct {
	Id         int    `json:"id,omitempty"`
	NickName   string `json:"nickname,omitempty"`
	Email      string `json:"email,omitempty"`
	Password   string `json:"password,omitempty"`
	Avator     string `json:"avator,omitempty"`
	Birthday   string `json:"birthday,omitempty"`
	Competence string `json:"competence,omitempty"`
	Intro      string `json:"intro,omitempty"`
	Token      string `json:"token,omitempty"`
}

// 返回数据的用户结构体
type RespUser struct {
	Id         int    `json:"id,omitempty"`
	NickName   string `json:"name,omitempty"`
	Email      string `json:"email,omitempty"`
	Avator     string `json:"avator,omitempty"`
	Birthday   string `json:"birthday,omitempty"`
	Competence string `json:"competence,omitempty"`
	Intro      string `json:"intro,omitempty"`
	Token      string `json:"token,omitempty"`
}

// 将数据库User结构体转换成响应user结构体
func ConverUser(user User) RespUser {
	respuser := RespUser{
		Id:         user.Id,
		NickName:   user.NickName,
		Email:      user.Email,
		Avator:     user.Avator,
		Birthday:   user.Birthday,
		Competence: user.Competence,
		Intro:      user.Intro,
		Token:      user.Token,
	}

	return respuser
}
