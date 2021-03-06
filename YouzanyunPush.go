package YouzanyunPush

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

/*
Event-Sign：防伪签名 ：MD5(client_id+entity+client_secret) ; 其中 entity 是从 RequestBody 读取的内容（详情见下文解析示例）
详细推荐文档：https://doc.youzanyun.com/resource/develop-guide/41355/41536
*/
// 申明全局变量 client_id,client_secret
var clientId,clientSecret string

type YouzanClient struct {
	ClientId     string
	ClientSecret string
}
type RetMessage struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
}

func md5sign(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}
// 计算签名 client_id + entity + client_secret 用户计算防伪签名 event-Sign

func (client *YouzanClient) Verifysign(req *http.Request) (err error) {
	eventSign := req.Header.Get("Event-Sign")
	if eventSign == "" {
		return errors.New("fail,no Event_Sign")
	}
	reqBody, err := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	if err != nil {
		return err
	}

	md5string := client.ClientId + string(reqBody) + client.ClientSecret
	fmt.Println(md5sign(md5string))
	if md5sign(md5string) == eventSign {
		return nil
	} else {
		return errors.New("sign_fail")
	}
}

func (client *YouzanClient)YouzanPush(w http.ResponseWriter, r *http.Request) {
	err := client.Verifysign(r)
	w.Header().Set("Content-Type", "application/json")
	if err == nil {
		// 校验消息合法以后，返回接收成功标识，然可以处理自己的逻辑解析body
		json.NewEncoder(w).Encode(RetMessage{"200", "success"})
	} else {
		json.NewEncoder(w).Encode(RetMessage{"200", err.Error()})
	}

}