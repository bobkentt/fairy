package err 

const (
    ERR_OK                   = 0     //成功
    ERR_NOERROR              = 10000 //成功
    ERR_INTERNAL             = 10001 //系统内部错误
    ERR_BAD_REQUEST          = 10002 //请求body无效
    ERR_INVALID_PARAMS       = 10003 //请求URL参数无效
    ERR_AUTHENTICATE         = 10004 //认证失败
    ERR_AUTHORIZE            = 10005 //授权失败
    ERR_SIGN                 = 10006 //签名失败
    ERR_INVALID_TOKEN        = 10007 //TOKEN无效
    ERR_EXIST                = 10008 //资源已存在
    ERR_NOT_EXIST            = 10009 //资源不存在
    ERR_QUOTA_EXCEEDED       = 10010 //配额超限
    ERR_TIMEOUT              = 20000 //超时
    ERR_OTHER                = 20040 //其它错误
)

type ApiStatus struct {
    ErrNo  int    `json:"errno"`
    ErrMsg string `json:"errmsg,omitempty"`
}

type ApiResponse struct {
    ErrNo  int         `json:"errno"`
    ErrMsg string      `json:"errmsg,omitempty"`
    Data   interface{} `json:"data,omitempty"`
}

