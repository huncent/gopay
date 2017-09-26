package alipay

import (
	"net/http"
)



type Notify interface {
	notify(w http.ResponseWriter, r *http.Request)
}

type NotifyParams struct {
}

type PagePayParams struct {
	OutTradeNo         string
	ProductCode        string
	TotalAmount        int
	Subject            string
	Body               string
	GoodsDetail        string
	PassbackParams     string
	ExtendParams       string
	GoodsType          string
	TimeoutExpress     string
	EnablePayChannels  string
	DisablePayChannels string
	AuthToken          string
	QrPayMode          string
	QrcodeWidth        string
}


