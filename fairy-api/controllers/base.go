package controllers

import (
    log "code.google.com/p/log4go"
    "github.com/astaxie/beego"
)

type Response struct {
    ErrNo  int         `json:"errno"`
    ErrMsg string      `json:"errmsg,omitempty"`
    Data   interface{} `json:"data,omitempty"`
}

type BaseController struct {
    beego.Controller
}

func (this *BaseController) errResponse(errno int, errmsg string, httpcode int, data interface{}) {
    resp := Response{
        ErrNo:  errno,
        ErrMsg: errmsg,
    }
    if data != nil {
        resp.Data = data
    }
    this.Data["json"] = &resp
    if httpcode == 500 {
        log.Warn("%s %s, (%d, %s)", this.Ctx.Input.Method(), this.Ctx.Input.URL(), errno, errmsg)
    }
    this.Ctx.Output.SetStatus(httpcode)
    this.ServeJSON()
}

func (this *BaseController) okResponse(errno int, data interface{}) {
    resp := Response{
        ErrNo: errno,
        Data:  data,
    }
    this.Data["json"] = &resp
    this.ServeJSON()
}
