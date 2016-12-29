package controllers

import (
   "fmt" 
)

type MenuController struct {
    BaseController	
}

func (this *MenuController) GetMenuList() {
    fmt.Printf("Call GetMenuList\n")
    return
}

