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
	a :=new(alipay.Client)
	a.Request(*s)

}

func Test_reflect(t *testing.T)  {
	s :=new(alipay.ClientParams)
	s.NewClient()
	log.Println(reflect.TypeOf(*s).NumField())
}