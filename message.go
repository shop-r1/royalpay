package royalpay

type (
	Currency string
	Channel  string
)

const (
	AUD    Currency = "AUD"
	CNY    Currency = "CNY"
	Alipay Channel  = "Alipay"
	Wechat Channel  = "Wechat"
)

type Body struct {
	Description string   `json:"description"`
	Price       int      `json:"price"`
	Currency    Currency `json:"currency"`
	Channel     Channel  `json:"channel"`
	NotifyUrl   string   `json:"notify_url"`
	Operator    string   `json:"operator"`
}

type Result struct {
	ReturnCode     string `json:"return_code"`
	ResultCode     string `json:"result_code"`
	ReturnMsg      string `json:"return_msg"`
	Channel        string `json:"channel"`
	PartnerCode    string `json:"partner_code"`
	FullName       string `json:"full_name"`
	PartnerName    string `json:"partner_name"`
	OrderId        string `json:"order_id"`
	PartnerOrderId string `json:"partner_order_id"`
	CodeUrl        string `json:"code_url,omitempty"`
	QrcodeImg      string `json:"qrcode_img,omitempty"`
	PayUrl         string `json:"pay_url"`
}

var ErrCode = map[string]string{
	"ORDER_NOT_EXIST":  "订单不存在",
	"ORDER_MISMATCH":   "订单号与商户不匹配",
	"SYSTEMERROR":      "系统内部异常",
	"INVALID_SHORT_ID": "商户编码不合法或没有对应商户",
	"SIGN_TIMEOUT":     "签名超时，time字段与服务器时间相差超过5分钟",
	"INVALID_SIGN":     "签名错误",
	"PARAM_INVALID":    "参数不符合要求，具体细节可参考return_msg字段",
	"NOT_PERMITTED":    "未开通网关支付权限",
	"INVALID_CHANNEL":  "不合法的支付渠道名称，请检查大小写",
	"ORDER_PAID":       "单已支付",
}
