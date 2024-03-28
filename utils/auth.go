package utils

import (
	"encoding/json"
	"errors"
	"github.com/astaxie/beego/httplib"
)

type Auth struct {
	AppCode string `json:"app_code"`
}

func Authentication(name, idCard string, a *Auth) error {
	res, err := httplib.Post("https://eid.shumaidata.com/eid/check?idcard="+idCard+"&name="+name).Header("Authorization", "APPCODE "+a.AppCode).String()
	if err != nil {
		return err
	}
	var r result
	_ = json.Unmarshal([]byte(res), &r)
	if r.Code != "0" || r.Result.Res != "1" {
		return errors.New("实名认证失败")
	}
	return nil
}

type result struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Result  struct {
		Name        string `json:"name"`
		Idcard      string `json:"idcard"`
		Res         string `json:"res"`
		Description string `json:"description"`
		Sex         string `json:"sex"`
		Birthday    string `json:"birthday"`
		Address     string `json:"address"`
	} `json:"result"`
}
