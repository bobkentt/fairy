package controllers

import (
   "fmt"

   "fairy/storage"
   "fairy/config"
)

type OrderController struct {
    BaseController
}

func (this *OrderController) GetOrderList() {
    fmt.Printf("Call GetOrderList\n")

    // Set default offset & limit
    offset, err:= this.GetInt("offset")
    if err != nil {
        offset = 0
    }
    limit, err := this.GetInt("limit")
    if err != nil {
        limit = 20
    }
    uid, err := this.GetUsrId()
    if err != nil {
        // jump to login
    }

    list, err := storage.GetOrderList(uid, offset, limit)
    if err != nil {
        this.errResponse(config.ERR_INTERNAL, err.Error(), 500, nil)
    }
    this.okResponse(0, list)
}

func (this *OrderController) Post() {
    fmt.Printf("Call Order Post\n")

    err := this.CheckUsrInfo()
    if err != nil {
        // jump to login
    }
    uid, err := this.GetUsrId()
    if err != nil {
        this.errResponse(config.ERR_INTERNAL, err.Error(), 500, nil)
    }
    var ob storage.Orders
    json.Unmarshal(this.Ctx.Input.RequestBody, &ob)

    err := storage.CreateOrder(uid, &ob)
    if err != nil {
        this.errResponse(config.ERR_INTERNAL, err.Error(), 500, nil)
    }
    this.okResponse(0, nil)
}

func (this *OrderController) Get() {
    fmt.Printf("Call Order Get\n")

    err := this.CheckUsrInfo()
    if err != nil {
        // jump to login
    }
    order_id, err:= this.GetInt("order_id")
    if err != nil {
        this.errResponse(config.ERR_INVALID_PARAMS, "invaild request order_id", 400, nil)
    }
    p, err := storage.GetOrderDetail(order_id)
    if err != nil {
        this.errResponse(config.ERR_INTERNAL, err.Error(), 500, nil)
    }
    this.okResponse(0, p)
}

func (this *OrderController) Delete() {
    fmt.Printf("Call Order Delete\n")

    err := this.CheckUsrInfo()
    if err != nil {
        // jump to login
    }

    order_id, err:= this.GetInt("order_id")
    if err != nil {
        this.errResponse(config.ERR_INVALID_PARAMS, "invaild request order_id", 400, nil)
    }
    err := storage.RemoveOrder(order_id)
    if err != nil {
        this.errResponse(config.ERR_INTERNAL, err.Error(), 500, nil)
    }
    this.okResponse(0,"success")
}
