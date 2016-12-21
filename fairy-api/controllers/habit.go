package controllers

import (
    log "code.google.com/p/log4go"
	"github.com/astaxie/beego"
)

const (
    // habit types
    TYPE_USR_DEFINED         = 0
    TYPE_SMOKE               = 1 
    TYPE_MASTURBATION        = 2
    TYPE_SLEEP_DELAY         = 3
    TYPE_TALK_DIRTY          = 5
)

type HabitPara struct {
    Name string
    Type int 
}

type HabitController struct {
    BaseController	
}

var globalHabitTypes map[int]bool

func init() {
    globalHabitTypes = make(map[int]int)
    globalHabitTypes = map[int]int{
        TYPE_USR_DEFINED:  true,  
        TYPE_SMOKE:        true,  
        TYPE_MASTURBATION: true,
        TYPE_SLEEP_DELAY:  true,  
        TYPE_TALK_DIRTY:   true,  
    }  
}

func (c *HabitController) Addhabit() {
    var data HabitPara
    body := this.Ctx.Input.RequestBody
    if err := json.Unmarshal([]byte(body), &data); err != nil {
        this.errResponse(err.ERR_BAD_REQUEST, "invalid request", 400, nil)
        log.Warn("add habit: read request body failed: (%s)", err)
        return
    }
    name := data.Name
    type := data.Type
    if len(name) == 0 || len(name) > 64 {
        this.errResponse(err.ERR_INVALID_PARAMS, "invalid habit name length", 400, nil)
        return
    }
    
    _, err := storage.Instance.SetAdd("db_topic_"+appid, topic)
    if err != nil {
        this.errResponse(err.ERR_INTERNAL, "storage I/O failed", 500, nil)
        return
    }

    this.okResponse(0, nil)
}
