package alipay

import (
	"net/url"
	"time"
	"log"
	"encoding/json"
)

const (
	RequestUrl = "openapi.alipay.com/gateway.do"
	Charset ="UTF-8"
	Format = "json"
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

type Client struct {
	url url.Values
}

func (c *ClientParams) NewClient(appId string,privateKey string) *ClientParams{
	c.Method = RequestUrl
	if c.Format =="" {
		c.Format = Format
	}
	if c.Charset =="" {
		c.Charset = Charset
	}
	c.Timestamp = time.Now().Format("2006-01-02 15:04:05")
	c.SignType = "RSA"
	c.Version ="1.0"
	return c
}

func (c *Client) Request(param *ClientParams){
	s,_:=json.Marshal(param)
	m :=make(map[string]string)
	err :=json.Unmarshal(s,&m)
	if err !=nil{
		log.Println(err.Error())
	}
	c.url = make(url.Values)
	for k,v :=range m {
		c.url.Set(k,v)
	}
	log.Println(c.url.Encode())
}

func (c ClientParams) sign(secret string) string {
	return ""
}
