package controllers

import (
   "fmt" 

   "fairy/storage"
   "fairy/config"
)

type MenuController struct {
    BaseController	
}

func (this *MenuController) GetMenuList() {
    fmt.Printf("Call GetMenuList\n")
    _, produts, err := storage.GetProduct()
    if err != nil {
        this.errResponse(config.ERR_INTERNAL, err.Error(), 500, nil)    
    }
    this.okResponse(0, products)
}

