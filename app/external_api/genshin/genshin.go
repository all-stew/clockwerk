package genshin

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/rand"
	"time"

	"github.com/imroc/req"
)

const (
	AppVersion       = "2.2.1"
	UserAgentVersion = "2.3.0"
	ClientType       = "5"
	Referer          = "https://webstatic.mihoyo.com/bbs/event/signin-ys/index.html"
	UserAgent        = "Mozilla/5.0 (iPad; CPU OS 13_6 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) miHoYoBBS/" + UserAgentVersion
	ActId            = "e202009291139501"
)

type GenshinCookie struct {
	CookieToken string `json:"cookie_token"`
	AccountId   string `json:"account_id"`
}

type GenshinAccountResp struct {
	RetCode int                `json:"retcode"`
	Message string             `json:"message"`
	Data    GenshinAccountData `json:"data"`
}

type GenshinAccountData struct {
	List []GenshinAccount `json:"list"`
}

type GenshinAccount struct {
	GameBiz    string `json:"game_biz"`
	Region     string `json:"region"`
	GameUid    string `json:"game_uid"`
	Nickname   string `json:"nickname"`
	Level      int    `json:"level"`
	IsChosen   bool   `json:"is_chosen"`
	RegionName string `json:"region_name"`
	IsOfficial bool   `json:"is_official"`
}

func GetGenshinCookie(cookie string) (GenshinCookie, error) {
	var data GenshinCookie
	err := json.Unmarshal([]byte(cookie), &data)
	return data, err
}

func GetUserGameRoleByCookie(cookie string) (*req.Resp, error) {
	// 获取角色列表接口
	uri := "https://external_api-takumi.mihoyo.com/binding/external_api/getUserGameRolesByCookie?game_biz=hk4e_cn"
	header := req.Header{
		"Cookie": cookie,
	}

	resp, err := req.Get(uri, header)
	//logger.Logf("resp is %s", resp.String())
	return resp, err
}

func getSignInfo() {
	uri := fmt.Sprintf("https://external_api-takumi.mihoyo.com/event/bbs_sign_reward/info?act_id=%s&region=%s&uid=%s", ActId, "cn_qd01", "genshin accountid")
	header := req.Header{
		"Cookie": "cookiestring",
	}
	res, err := req.Get(uri, header)

	fmt.Println(res)
	fmt.Println(err)
}

func sign() {
	requestJson := map[string]interface{}{
		"act_id": ActId,
		"region": "cn_qd01",
		"uid":    "genshin accountid",
	}
	uri := "https://external_api-takumi.mihoyo.com/event/bbs_sign_reward/sign"
	header := req.Header{
		"Content-Type":      "application/json",
		"x-rpc-device_id":   "F84E53D45BFE4424ABEA9D6F0205FF4A",
		"x-rpc-app_version": AppVersion,
		"x-rpc-client_type": ClientType,
		"Cookie":            "cookiestring",
		"Referer":           Referer,
		"DS":                getDs(),
		"User-Agent":        UserAgent,
	}

	jsonBody := req.BodyJSON(requestJson)
	res, err := req.Post(uri, header, jsonBody)

	fmt.Println(res.Request().Header)
	fmt.Println(res.Request().Body)
	fmt.Println(res)
	fmt.Println(err)
}

func getDs() string {
	t := time.Now().Unix()
	rand.Seed(time.Now().UnixNano())
	r := rand.Intn(999999)
	ms := fmt.Sprintf("salt=cx2y9z9a29tfqvr1qsq6c7yz99b5jsqt&t=%d&r=%d", t, r)
	md5Str := Md5(ms)
	return fmt.Sprintf("%d,%d,%s", t, r, md5Str)
}

func Md5(str string) string {
	m := md5.New()
	m.Write([]byte(str))
	return hex.EncodeToString(m.Sum(nil))
}
