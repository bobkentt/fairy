package err 

const (
    // controal api errno
    ERR_OK                   = 0     //成功
    ERR_NOERROR              = 10000 //成功
    ERR_INTERNAL             = 10001 //系统内部错误
    ERR_DEV_OFFLINE          = 10002 //设备离线
    ERR_BAD_REQUEST          = 10003 //请求body无效
    ERR_INVALID_PARAMS       = 10004 //请求URL参数无效
    ERR_AUTHENTICATE         = 10005 //认证失败
    ERR_AUTHORIZE            = 10006 //授权失败
    ERR_SIGN                 = 10007 //签名失败
    ERR_INVALID_TOKEN        = 10008 //TOKEN无效
    ERR_EXIST                = 10009 //资源已存在
    ERR_NOT_EXIST            = 10010 //资源不存在
    ERR_QUOTA_EXCEEDED       = 10011 //配额超限
    ERR_MQ_TIMEOUT           = 20000 //远程控制：MQ超时
    ERR_REMOTE_TIMEOUT       = 20001 //远程控制：设备端响应超时
    ERR_REMOTE_INVALID_REPLY = 20002 //远程控制：设备端响应无效
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

