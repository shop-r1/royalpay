package royalpay

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

type Pay struct {
	partnerCode    string
	credentialCode string
	sign           string
	query          string
}

func NewPay(partner, credential string) *Pay {
	return &Pay{
		partnerCode:    partner,
		credentialCode: credential,
	}
}

func (p *Pay) makeSign() {
	nonceStr := GetRandomString(24)
	timeStr := getTime()
	p.sign = sign(p.partnerCode, p.credentialCode, nonceStr, timeStr)
	p.makeQuery(timeStr, nonceStr)
}

func (p *Pay) makeQuery(timeStr, nonceStr string) {
	v := url.Values{}
	v.Add("time", timeStr)
	v.Add("nonce_str", nonceStr)
	v.Add("sign", p.sign)
	p.query = v.Encode()
}

func (p *Pay) setHeader() {

}

func (p *Pay) PayUrlSign(uri string) string {
	p.makeSign()
	return uri + "?" + p.query
}

//二维码订单
func (p *Pay) QrcodeOrder(orderId string, body *Body) (*Result, error) {
	uri := "https://mpay.royalpay.com.au/api/v1.0/gateway/partners/%s/orders/%s"
	uri = fmt.Sprintf(uri, p.partnerCode, orderId)
	return p.execPut(uri, body)
}

//
func (p *Pay) QrcodeNativeOrder(orderId string, body *Body) (*Result, error) {
	uri := "https://mpay.royalpay.com.au/api/v1.0/gateway/partners/%s/native_orders/%s"
	uri = fmt.Sprintf(uri, p.partnerCode, orderId)
	return p.execPut(uri, body)
}

func (p *Pay) execPut(uri string, body *Body) (*Result, error) {
	var err error
	p.makeSign()
	uri += "?" + p.query
	var rb []byte
	rb, err = json.Marshal(body)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(http.MethodPut, uri, bytes.NewBuffer(rb))
	if err != nil {
		return nil, err
	}
	defer req.Body.Close()
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Do(req)
	rb = make([]byte, 0)
	rb, _ = ioutil.ReadAll(resp.Body)
	var result Result
	err = json.Unmarshal(rb, &result)
	if err != nil {
		return nil, err
	}
	if result.ReturnCode != "SUCCESS" {
		return nil, errors.New(result.ReturnMsg)
	}
	return &result, err
}
