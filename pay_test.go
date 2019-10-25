package royalpay

import (
	"strconv"
	"testing"
	"time"
)

var (
	partnerCode    = "xxx"
	credentialCode = "xxx"
)

func TestPay_QrcodeOrder(t *testing.T) {
	type fields struct {
		partnerCode    string
		credentialCode string
		sign           string
		query          string
	}
	type args struct {
		orderId string
		body    *Body
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			"royalpay_test_wechat",
			fields{
				partnerCode:    partnerCode,
				credentialCode: credentialCode,
			},
			args{
				orderId: strconv.Itoa(int(time.Now().UnixNano())),
				body: &Body{
					Description: "test01",
					Price:       5,
					Currency:    "CNY",
					Channel:     Wechat,
					Operator:    "lwx",
					NotifyUrl:   "http://tenant.shop-r1.ngrok.jifenhua.cn/admin",
				},
			},
			false,
		},
		{
			"royalpay_test_alipay",
			fields{
				partnerCode:    partnerCode,
				credentialCode: credentialCode,
			},
			args{
				orderId: strconv.Itoa(int(time.Now().UnixNano())),
				body: &Body{
					Description: "test01",
					Price:       5,
					Currency:    "CNY",
					Channel:     Alipay,
					Operator:    "lwx",
					NotifyUrl:   "http://tenant.shop-r1.ngrok.jifenhua.cn/admin",
				},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Pay{
				partnerCode:    tt.fields.partnerCode,
				credentialCode: tt.fields.credentialCode,
				sign:           tt.fields.sign,
				query:          tt.fields.query,
			}
			gotResult, err := p.QrcodeOrder(tt.args.orderId, tt.args.body)
			if (err != nil) != tt.wantErr {
				t.Errorf("Pay.QrcodeOrder() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotResult == nil || gotResult.ReturnCode != "SUCCESS" {
				t.Errorf("Pay.QrcodeNativeOrder() fail ")
				return
			}
		})
	}
}

func TestPay_QrcodeNativeOrder(t *testing.T) {
	type fields struct {
		partnerCode    string
		credentialCode string
		sign           string
		query          string
	}
	type args struct {
		orderId string
		body    *Body
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			"royalpay_test_wechat",
			fields{
				partnerCode:    partnerCode,
				credentialCode: credentialCode,
			},
			args{
				orderId: strconv.Itoa(int(time.Now().UnixNano())),
				body: &Body{
					Description: "test01",
					Price:       5,
					Currency:    "CNY",
					Channel:     Wechat,
					Operator:    "lwx",
					NotifyUrl:   "http://tenant.shop-r1.ngrok.jifenhua.cn/admin",
				},
			},
			false,
		},
		{
			"royalpay_test_alipay",
			fields{
				partnerCode:    partnerCode,
				credentialCode: credentialCode,
			},
			args{
				orderId: strconv.Itoa(int(time.Now().UnixNano())),
				body: &Body{
					Description: "test01",
					Price:       5,
					Currency:    "CNY",
					Channel:     Alipay,
					Operator:    "lwx",
					NotifyUrl:   "http://tenant.shop-r1.ngrok.jifenhua.cn/admin",
				},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Pay{
				partnerCode:    tt.fields.partnerCode,
				credentialCode: tt.fields.credentialCode,
				sign:           tt.fields.sign,
				query:          tt.fields.query,
			}
			got, err := p.QrcodeNativeOrder(tt.args.orderId, tt.args.body)
			if (err != nil) != tt.wantErr {
				t.Errorf("Pay.QrcodeNativeOrder() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got == nil || got.ReturnCode != "SUCCESS" {
				t.Errorf("Pay.QrcodeNativeOrder() fail ")
				return
			}
		})
	}
}
