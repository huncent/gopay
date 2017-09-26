package alipay

import (
	"net/url"
	"time"
	"log"
	"encoding/json"
)

const (
	RequestUrl = "https://openapi.alipay.com/gateway.do"
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
	url.Values
}

func (c *ClientParams) NewClient() *ClientParams{
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

func (c *Client) Request(param ClientParams){
	s,_:=json.Marshal(param)
	m :=make(map[string]string)
	r :=json.Unmarshal(s,&m)
	for k,v :=range r {
		c.Values.Add(k,v.(string))
	}
	log.Println(c.Values.Encode())
}

func (c ClientParams) sign(secret string) string {
	return ""
}
