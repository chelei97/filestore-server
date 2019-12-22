package handler

import (
	dblayer "filestore-server/db"
	"filestore-server/util"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	pwd_salt = "*#890;"
)

func SignUpHandler(w http.ResponseWriter, r *http.Request)  {
	if r.Method == http.MethodGet {
		data, err := ioutil.ReadFile("./static/view/signup.html")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Write(data)
		return
	}

	r.ParseForm()
	username := r.Form.Get("username")
	passwd := r.Form.Get("password")

	if len(username) < 3 || len(passwd) < 5 {
		w.Write([]byte("Invalid parameter"))
		return
	}
	encPassWd := util.Sha1([]byte(passwd + pwd_salt))

	suc := dblayer.UserSignUp(username, encPassWd)
	if suc {
		w.Write([]byte("SUCCESS"))
	}else {
		w.Write([]byte("FAILED"))
	}
}

//SignInHandle: 登录接口
func SignInHandle(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		data, err := ioutil.ReadFile("./static/view/signin.html")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Write(data)
		return
	}
	r.ParseForm()
	username := r.Form.Get("username")
	passwd := r.Form.Get("password")
	encPassWd := util.Sha1([]byte(passwd + pwd_salt))
	//校验密码
	pwdChecked := dblayer.UserSignIn(username, encPassWd)
	if !pwdChecked {
		w.Write([]byte("FAILED"))
		return
	}
	//生成token
	token := GenToken(username)
	upRes := dblayer.UpdateToken(username, token)
	if !upRes {
		w.Write([]byte("FAILED"))
		return
	}
	//登录成功后重定向到首页
	resp := util.RespMsg{
		Code: 0,
		Msg:  "OK",
		Data: struct {
			Location string
			Username string
			Token string
		}{
			Location:"http://" + r.Host + "/static/view/home.html",
			Username:username,
			Token:token,
		},
	}
	w.Write(resp.JSONBytes())
}

//UserInfoHandler: 查询用户信息
func UserInfoHandler (w http.ResponseWriter, r *http.Request) {
	//解析请求参数
	r.ParseForm()
	username := r.Form.Get("username")
	//查询用户信息
	user, err := dblayer.GetUserInfo(username)
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		return
	}
	//组装并相应用户数据
	resp := util.RespMsg{
		Code: 0,
		Msg:  "OK",
		Data: user,
	}
	w.Write(resp.JSONBytes())
}

//GenToken: 生成token
func GenToken(username string) string {
	//40bits : md5(username + timestamp + token_salt) + timestamp[:8]

	ts := fmt.Sprintf("%x", time.Now().Unix())
	tokenPrefix := util.MD5([]byte(username + ts + "_tokensalt"))

	return tokenPrefix + ts[:8]
}

//IsTokenValid: 判断token的有效性
func IsTokenValid(token string) bool {
	//判断是否过期
	//从数据库中查询username的token信息
	//比较是否一致
	return true
}