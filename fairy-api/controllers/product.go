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
    merchant_id, err := this.GetInt("product_id")
    if err != nil {
        this.errResponse(config.ERR_INVALID_PARAMS, "invaild request merchant_id", 400, nil)    
    }
    
    _, product, err := storage.GetProduct()
    if err != nil {
        this.errResponse(config.ERR_INTERNAL, err.Error(), 500, nil)    
    }
    this.okResponse(0, products)
}

