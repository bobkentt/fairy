package controllers

import (
   "github.com/astaxie/beego"
   "fmt" 
)

type hiscontroller struct {
    beego.Controller
}

func (this *hiscontroller) Get() {
    fmt.Printf("Call his.go: Get\n")
    return
}

