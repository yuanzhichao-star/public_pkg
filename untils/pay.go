package untils

import "github.com/yuanzhichao-star/public_pkg/pkg"

type Pay interface {
	Alipay(orderSn string, money string) string
}
type Ali struct {
}

func (a *Ali) Alipay(orderSn string, money string) string {
	payMoney := pkg.PayMoney(orderSn, money)
	return payMoney
}
