package controllers

import (
	"github.com/astaxie/beego"
)

type MarginController struct {
	beego.Controller
}

func (c *MarginController) Recharge() {
}

func (c *MarginController) Refund() {
}
