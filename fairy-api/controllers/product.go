package controllers

import (
   "fmt"

   "fairy/storage"
   "fairy/config"
)

type ProductController struct {
    BaseController
}

func (this *ProductController) Get() {
    fmt.Printf("Call Product Get\n")

    // parameter parse
    product_id, err := this.GetInt("product_id")
    if err != nil {
        this.errResponse(config.ERR_INVALID_PARAMS, "invaild request product_id", 400, nil)
    }
    product, err := storage.GetProduct(product_id)
    if err != nil {
        this.errResponse(config.ERR_INTERNAL, err.Error(), 500, nil)
    }
    this.okResponse(0, product)
}

func (this *ProductController) Post() {
    fmt.Printf("Call Product Post\n")

    var pt Storage.Product
    json.Unmarshal(this.Ctx.Input.RequestBody, &pt)

    err := storage.SetProduct(pt)
    if err != nil {
        this.errResponse(config.ERR_INTERNAL, err.Error(), 500, nil)
    }
    this.okResponse(0, nil)
}

func (this *ProductController) Delete() {
    fmt.Printf("Call Product Delete\n")

    // parameter parse
    product_id, err := this.GetInt("product_id")
    if err != nil {
        this.errResponse(config.ERR_INVALID_PARAMS, "invaild request product_id", 400, nil)
    }

    err := storage.RemoveProduct(product_id)
    if err != nil {
        this.errResponse(config.ERR_INTERNAL, err.Error(), 500, nil)
    }
    this.okResponse(0, nil)
}
