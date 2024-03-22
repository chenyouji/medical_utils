package utils

import (
	"github.com/smartwalle/alipay/v3"
	"net/url"
)

var client *alipay.Client

type AliPay struct {
	AppID        string `json:"app_id"`
	PrivateKey   string `json:"private_key"`
	AliPublicKey string `json:"ali_public_key"`
	NotifyURL    string `json:"notify_url"`
	ReturnURL    string `json:"return_url"`
}

func InitAliPay(pay *AliPay) {
	var err error
	client, err = alipay.New(pay.AppID, pay.PrivateKey, false)
	if err != nil {
		panic(err)
	}
	err = client.LoadAliPayPublicKey(pay.AliPublicKey)
	if err != nil {
		panic(err)
	}
}

func GeneralPayURL(p alipay.TradePagePay) (string, error) {
	/*
		alipay.TradePagePay{
			Trade: alipay.Trade{
				NotifyURL:   ,
				ReturnURL:   ,
				Subject:     ,
				OutTradeNo:  ,
				TotalAmount: ,
				ProductCode: "FAST_INSTANT_TRADE_PAY",
			},
		}
	*/
	var res, err = client.TradePagePay(p)
	if err != nil {
		return "", err
	}

	// generate payURL string
	var payURL = res.String()
	return payURL, nil
}

func DecodePayNotify(values url.Values) (*alipay.Notification, error) {
	return client.DecodeNotification(values)
}

/*
TradeStatusWaitBuyerPay TradeStatus = "WAIT_BUYER_PAY" //（交易创建，等待买家付款）
TradeStatusClosed       TradeStatus = "TRADE_CLOSED"   //（未付款交易超时关闭，或支付完成后全额退款）
TradeStatusSuccess      TradeStatus = "TRADE_SUCCESS"  //（交易支付成功）
TradeStatusFinished     TradeStatus = "TRADE_FINISHED" //（交易结束，不可退款）
*/
