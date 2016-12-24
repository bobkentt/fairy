package controllers

import (
    "encoding/json"
    "fairy/config"

    log "code.google.com/p/log4go"
)

const (
    // Habit types
    TYPE_USR_DEFINED  = 0
    TYPE_SMOKE        = 1
    TYPE_MASTURBATION = 2
    TYPE_SLEEP_DELAY  = 3
    TYPE_TALK_DIRTY   = 5
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
    globalHabitTypes = make(map[int]bool)
    globalHabitTypes = map[int]bool{
        TYPE_USR_DEFINED:  true,
        TYPE_SMOKE:        true,
        TYPE_MASTURBATION: true,
        TYPE_SLEEP_DELAY:  true,
        TYPE_TALK_DIRTY:   true,
    }
}

func (this *HabitController) Post() {
    var data HabitPara
    body := this.Ctx.Input.RequestBody
    if err := json.Unmarshal([]byte(body), &data); err != nil {
        this.errResponse(config.ERR_BAD_REQUEST, "invalid request", 400, nil)
        log.Warn("add habit: read request body failed: (%s)", err)
        return
    }
    pname := data.Name
    ptype := data.Type
    if len(pname) == 0 || len(pname) > 64 {
        this.errResponse(config.ERR_INVALID_PARAMS, "invalid habit name length", 400, nil)
        return
    }
    _, exists := globalHabitTypes[ptype]
    if exists {
        this.errResponse(config.ERR_INVALID_PARAMS, "invalid habit type", 400, nil)
        return
    }

    this.okResponse(0, nil)
}
