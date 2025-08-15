package untils

import (
	"github.com/yuanzhichao-star/public_pkg/pkg"
)

// 阿里云
type Sms interface {
	AliYun(mobile, code string) error
}
type AliSms struct {
}

func (as *AliSms) AliYun(mobile, code string) error {
	sms, err := pkg.SendSms(mobile, code)
	if err != nil {
		return err
	}
	if *sms.Body.Code != "OK" {
		return err
	}
	return nil
}
