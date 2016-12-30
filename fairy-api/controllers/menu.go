package controllers

import (
   "fmt" 

   "fairy/storage"
)

type MenuController struct {
    BaseController	
}

func (this *MenuController) GetMenuList() {
    fmt.Printf("Call GetMenuList\n")
    storage.GetProduct()
    return
}

