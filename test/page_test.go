package test

import (
	"testing"
	"github.com/go-fmt/gopay/alipay"
	"reflect"
	"log"
)

func Test_sign(t *testing.T)  {
	s :=new(alipay.ClientParams)
	s.NewClient()
	s.AppId = "safasdasda"
	s.NotifyUrl = "http://192.168.1.1/notify"
	s.ReturnUrl = "http://192.168.1.1/return"
	s.SignKey("F:\\cacert.pem")
}

func Test_reflect(t *testing.T)  {
	s :=new(alipay.ClientParams)
	s.NewClient()
	log.Println(reflect.TypeOf(*s).NumField())
}

func Test_sort(t *testing.T)  {
	m :=make(map[string]string)
	m["name"] ="ajfasfa"
	m["job"] = "fasdasd"
	m["hhah"] = "dasfas"
	for k,v :=range m {
		log.Println(k,v)
	}
}