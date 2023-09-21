package utils

import (
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"meetingBooking/repository/cache"
	"reflect"
)

func GetValidMsg(err error, obj interface{}) string {
	//obj为结构体指针
	getObj := reflect.TypeOf(obj)
	//断言为具体的类型，err是一个接口
	if errs, ok := err.(validator.ValidationErrors); ok {
		for _, e := range errs {
			if f, exist := getObj.Elem().FieldByName(e.Field()); exist {
				return f.Tag.Get("msg") //错误信息不需要全部返回，当找到第一个错误的信息时，就可以结束
			}
		}
	}
	return err.Error()
}

func ValidEmail(email string, reqCaptcha string) error {
	captcha, err := cache.RedisGetKey(fmt.Sprintf("captcha_%s", email))
	if err != nil {
		return errors.New("验证码失效")
	}
	if captcha != reqCaptcha {
		return errors.New("验证码不正确")
	}
	return nil
}
