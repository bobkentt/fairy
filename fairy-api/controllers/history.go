package controllers

import (
   "fmt" 
)

type HistoryController struct {
    BaseController	
}

func (this *HistoryController) Get() {
    fmt.Printf("Call GetHistoryList\n")
    this.okResponse(0, "你好，AAA-李阳\n")
    return
}

func (this *HistoryController) GetHistoryItem() {
}
