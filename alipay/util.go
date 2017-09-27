package alipay

import (
	"reflect"
	"sort"
	"log"
	"crypto/x509"
	"crypto/rsa"
	"encoding/pem"
	"errors"
	"crypto/rand"
	"encoding/base64"
	"strings"
	"encoding/json"
	"crypto/sha1"
	"crypto"
	"fmt"
)

func Struct2Map(obj interface{}) map[string]interface{} {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	var data = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		data[t.Field(i).Name] = v.Field(i).Interface()
	}
	return data
}

func SortMap(m map[string]interface{})  []string{
	var keys []string
	for k := range m {
		keys = append(keys,k)
	}
	log.Println(keys)
	sort.Strings(keys)
	return keys
}

func SHA1WithRSAEncrypt(origData []byte,publicKey []byte) ([]byte, error) {
	block,_ := pem.Decode([]byte(`-----BEGIN RSA PRIVATE KEY-----
MIICWwIBAAKBgQCsAocILoBYZqhMYDg40AFZTUiFjcSxwrUXOF0rgw8hh98tP+Ox
4awokuF+FJn8qN/9k9gFz9j7zM694vNv976W60k2ye6uiQdQy/gOJh+ciFME3kAH
QoyvuItKRec+3cEhwblpuY7Gchk0mk22WXWyAHuNGXcCMGSJo9ugGnPDUwIDAQAB
AoGARDgcZdpLfMP6K5Bdu+qDHm/QO2emgvm96J+qE/++mIXStZeJLptaNB1M4Tw6
dkJj06Y3Htb4L6ViuVyxP875/yqKocNur4KeoTLC/t+9L7f6jey7GCvWlCcpp97A
NiVCPILG+7Py2+xGyv0tQT+98yJqTb0yA0nIsmq1XpmjkQECQQDUmVH8L6mJvyzj
xDA32jLKoQLWyJTvzkBxeTnsuTL6GqXEcj9HN0Q1qEyTl/DcN5YfACL4XevV1NNV
s6CrKEojAkEAzx/5d3aAI3dg9VHfAcWdiUZ5uTbx+qpDfZ4CtoyhEfolRGTiy2SC
y9GDrtpY3NO1mm+D5lIM3HPbTgAxtvm9EQJADmI5K8jFva4TiW1ynbTDjvYJzSJR
AVCBB6xeAOgezNEUug/IvDa/BKpYU/wJrbyNCZfmxnqE87ise7Xlfu8A5QJBAJTP
Bx9SLvvMMAfwm0UdonJXBOsR08Zg/35HwPFAlhRhYNcDmIHCo8olq/M7Am8dV7Mt
/VjDiGP2hRBESXOJd9ECPxeHMhK4prFTN1N/+fHcnbB61P89mIsH5ff38D9uEwa/
sLqen0K1ibjhgbcs0LjoVklK9fxJ5AK7SwU0Oxaq0w==
-----END RSA PRIVATE KEY-----
`))
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	log.Println(string(block.Bytes))
	h := sha1.New()
	h.Write(origData)
	digest := h.Sum(nil)
	s, err := rsa.SignPKCS1v15(nil, privateKey, crypto.SHA1, digest)
	if err != nil {
		fmt.Errorf("rsaSign SignPKCS1v15 error")
		return nil, err
	}
	return s, nil
}

func RsaDecrypt(ciphertext []byte,privateKey []byte) ([]byte, error) {
	block, _ := pem.Decode(privateKey)
	if block == nil {
		return nil, errors.New("private key error!")
	}
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	return rsa.DecryptPKCS1v15(rand.Reader, priv, ciphertext)
}

func Base64Encode(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

func UrlWithOutEncode(m map[string]string)  string{
	var slice []string
	for k,v :=range m {
		if k == "sign" {
			continue
		}
		slice = append(slice,k+"="+v)
	}
	return strings.Join(slice,"&")
}

func serialize(c interface{}) map[string]string{
	s, _ := json.Marshal(c)
	m := make(map[string]string)
	err := json.Unmarshal(s, &m)
	if err != nil {
		log.Println(err.Error())
	}
	return m
}