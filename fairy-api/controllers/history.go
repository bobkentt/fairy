package controllers

import (
   "fmt" 
)

type HistoryController struct {
    BaseController	
}

func (c *HistoryController) Get() {
    fmt.Printf("Call GetHistoryList\n")
    return
}

func (c *HistoryController) GetHistoryItem() {
}
