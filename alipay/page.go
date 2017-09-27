package alipay

import (
	"net/http"
)

type Notify interface {
	asyncNotify(w http.ResponseWriter, r *http.Request)
	syncNotify(w http.ResponseWriter, r *http.Request)
}

type NotifyPayParams struct {
	TradeNo           string `json:"trade_no"`
	AppId             string `json:"app_id"`
	OutTradeNo        string `json:"out_trade_no"`
	OutBizNo          string `json:"out_biz_no"`
	BuyerId           string `json:"buyer_id"`
	SellerId          string `json:"seller_id"`
	TradeStatus       string `json:"trade_status"`
	TotalAmount       float64 `json:"total_amount"`
	ReceiptAmount     float64 `json:"receipt_amount"`
	InvoiceAmount     float64 `json:"invoice_amount"`
	BuyerPayAmount    float64 `json:"buyer_pay_amount"`
	PointAmount       float64 `json:"point_amount"`
	RefundFee         float64 `json:"refund_fee"`
	Subject           float64 `json:"subject"`
	Body              string `json:"body"`
	GmtCreate         string `json:"gmt_create"`
	GmtPayment        string `json:"gmt_payment"`
	GmtRefund         string `json:"gmt_refund"`
	GmtClose          string `json:"gmt_close"`
	FundBillList      string `json:"fund_bill_list"`
	VoucherDetailList string `json:"voucher_detail_list"`
	PassbackParams    string `json:"passback_params"`
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
