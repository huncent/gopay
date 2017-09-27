package alipay

import (
	"net/url"
	"time"
	"log"
	"io/ioutil"
	"net/http"
	"strings"
)

const (
	RequestUrl = "openapi.alipay.com/gateway.do"
	Charset    = "UTF-8"
	Format     = "json"
)

type ClientParams struct {
	AppId      string `json:"app_id"`
	Method     string `json:"method"`
	Format     string `json:"format"`
	ReturnUrl  string `json:"return_url"`
	Charset    string `json:"charset"`
	SignType   string `json:"sign_type"`
	Sign       string `json:"sign"`
	Timestamp  string `json:"timestamp"`
	Version    string `json:"version"`
	NotifyUrl  string `json:"notify_url"`
	BizContent string `json:"biz_content"`
}

type NotifyParams struct {
	NotifyTime string `json:"notify_time"`
	NotifyType string `json:"notify_type"`
	NotifyId   string `json:"notify_id"`
	Charset    string `json:"charset"`
	Version    string `json:"version"`
	SignType   string `json:"sign_type"`
	Sign       string `json:"sign"`
	AuthAppId  string `json:"auth_app_id"`
}

type Client struct {
	url    url.Values
	client *http.Client
}

func (c *ClientParams) NewClient() *ClientParams {
	c.Method = RequestUrl
	if c.Format == "" {
		c.Format = Format
	}
	if c.Charset == "" {
		c.Charset = Charset
	}
	c.Timestamp = time.Now().Format("2006-01-02 15:04:05")
	c.SignType = "RSA"
	c.Version = "1.0"
	return c
}

func (c *Client) Request(param *ClientParams) {
	c.client = new(http.Client)
	c.url = make(url.Values)
	for k, v := range serialize(param) {
		c.url.Set(k, v)
	}
	req, err := http.NewRequest(http.MethodPost, "https://"+RequestUrl, strings.NewReader(c.url.Encode()))
	if err != nil {
		log.Println(err.Error())
	}
	resp, err := c.client.Do(req)
	if err != nil {
		log.Println(err.Error())
	}
	log.Println(ioutil.ReadAll(resp.Body))
}

func (c *ClientParams) SignKey(pemPath string) string {
	u := UrlWithOutEncode(serialize(c))
	log.Println(u)
	f, err := ioutil.ReadFile(pemPath)
	if err != nil {
		log.Println(err.Error())
	}
	code, err := SHA1WithRSAEncrypt([]byte(u), f)
	if err != nil {
		log.Println(err.Error())
	}
	log.Println(string(code))
	c.Sign = Base64Encode(code)
	log.Println(c.Sign)
	u = u + "&sign=" + c.Sign
	return Base64Encode([]byte(u))
}
