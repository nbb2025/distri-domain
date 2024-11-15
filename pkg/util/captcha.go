package util

import (
	"fmt"
	"github.com/mojocn/base64Captcha"
)

type CaptchaResult struct {
	Id          string `json:"ID"`
	Image       string `json:"Image"`
	VerifyValue string `json:"Answer"`
}

// 默认存储10240个验证码，每个验证码10分钟过期
var store = base64Captcha.DefaultMemStore

// GenerateCaptcha 生成图片验证码
func GenerateCaptcha() (CaptchaResult, error) {
	// 生成默认数字
	//driver := base64Captcha.DefaultDriverDigit
	// 此尺寸的调整需要根据网站进行调试，链接：
	// https://captcha.mojotv.cn/
	driver := base64Captcha.NewDriverDigit(70, 130, 4, 0.4, 100)
	// 生成base64图片
	captcha := base64Captcha.NewCaptcha(driver, store)
	// 获取
	id, b64s, ans, err := captcha.Generate()
	if err != nil {
		fmt.Println("Register GetCaptchaPhoto get base64Captcha has err:", err)
		return CaptchaResult{}, err
	}
	store.Set(id, ans)
	captchaResult := CaptchaResult{Id: id, Image: b64s, VerifyValue: ans}
	return captchaResult, nil
}

// 校验图片验证码,并清除内存空间
func VerifyCaptcha(id string, value string) bool {
	// done -- 只要id存在，就会校验并清除，无论校验的值是否成功, 所以同一id只能校验一次
	if id == "" || value == "" {
		// 注意：id,b64s是空 也会返回true 需要在加判断
		return false
	}
	verifyResult := store.Verify(id, value, true)
	return verifyResult
}
