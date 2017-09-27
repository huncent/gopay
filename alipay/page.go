package alipay

import (
	"net/http"
)



type Notify interface {
	asyncNotify(w http.ResponseWriter, r *http.Request)
	syncNotify(w http.ResponseWriter, r *http.Request)
}

type NotifyParams struct {
}

type PagePayParams struct {
	OutTradeNo         string `json:"out_trade_no"`
	ProductCode        string `json:"product_code"`
	TotalAmount        float64 `json:"total_amount"`
	Subject            string `json:"subject"`
	Body               string `json:"body"`
	GoodsDetail        string `json:"goods_detail"`
	PassbackParams     string `json:"passback_params"`
	ExtendParams       string `json:"extend_params"`
	GoodsType          string `json:"goods_type"`
	TimeoutExpress     string `json:"timeout_express"`
	EnablePayChannels  string `json:"enable_pay_channels"`
	DisablePayChannels string `json:"disable_pay_channels"`
	AuthToken          string `json:"auth_token"`
	QrPayMode          string `json:"qr_pay_mode"`
	QrcodeWidth        string `json:"qrcode_width"`
}


