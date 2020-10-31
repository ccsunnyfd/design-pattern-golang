package main

import "fmt"

// IPay interface
type IPay interface {
	transfer(float64) error
	setCheckMode(iSecurityCheckMode)
}

// alipay exact pay
type alipay struct {
	secMode iSecurityCheckMode
}

func (aPay *alipay) setCheckMode(secMode iSecurityCheckMode) {
	aPay.secMode = secMode
}

func (aPay *alipay) transfer(amount float64) error {
	aPay.secMode.securityCheck()
	fmt.Printf("支付宝转账%f。。。\n", amount)
	return nil
}

// weixinPay exact pay
type weixinPay struct {
	secMode iSecurityCheckMode
}

func (wPay *weixinPay) setCheckMode(secMode iSecurityCheckMode) {
	wPay.secMode = secMode
}

func (wPay *weixinPay) transfer(amount float64) error {
	wPay.secMode.securityCheck()
	fmt.Printf("微信转账%f。。。\n", amount)
	return nil
}

// iSecurityCheckMode interface
type iSecurityCheckMode interface {
	securityCheck() error
}

// fingerPrintMode exact checkMode
type fingerPrintMode struct {
}

func (fiMode *fingerPrintMode) securityCheck() error {
	fmt.Println("指纹识别")
	return nil
}

// faceMode exact checkMode
type faceMode struct {
}

func (faMode *faceMode) securityCheck() error {
	fmt.Println("人脸识别")
	return nil
}

// main
func main() {
	fiCheck := &fingerPrintMode{}
	faCheck := &faceMode{}
	var pay IPay
	pay = &alipay{}
	pay.setCheckMode(fiCheck)
	pay.transfer(1000)

	fmt.Print("\n===============\n")

	pay = &weixinPay{}
	pay.setCheckMode(faCheck)
	pay.transfer(1500)
}
