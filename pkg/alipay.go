package pkg

import (
	"fmt"
	"github.com/smartwalle/alipay/v3"
	"github.com/yuanzhichao-star/public_pkg/config"
)

func PayMoney(orderSn string, money string) string {
	appId := config.AppCong.AliPay.AppId
	var privateKey = config.AppCong.PrivateKey // 必须，上一步中使用 RSA签名验签工具 生成的私钥
	var client, err = alipay.New(appId, privateKey, false)
	if err != nil {
		fmt.Println(err)
		return err.Error()
	}

	var p = alipay.TradeWapPay{}
	p.NotifyURL = "http://pigglobal.com"
	p.ReturnURL = "http://xxx"
	p.Subject = "支付车费"
	p.OutTradeNo = orderSn
	p.TotalAmount = money
	p.ProductCode = "QUICK_WAP_WAY"

	url, err := client.TradeWapPay(p)
	if err != nil {
		fmt.Println(err)
	}

	// 这个 payURL 即是用于打开支付宝支付页面的 URL，可将输出的内容复制，到浏览器中访问该 URL 即可打开支付页面。
	var payURL = url.String()
	return payURL
}
