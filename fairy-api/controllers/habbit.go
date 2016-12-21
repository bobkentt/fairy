package controllers

import (
	"github.com/astaxie/beego"
)

type HabitController struct {
	beego.Controller
}

func (c *HabitController) Addhabbit() {
    body := this.Ctx.Input.RequestBody
    if err := json.Unmarshal([]byte(body), &data); err != nil {
        this.errResponse(cloud.ERR_BAD_REQUEST, "invalid request", 400, nil)
        log.Warn("add topic: read request body failed: (%s)", err)
        return
    }

}
