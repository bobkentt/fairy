package controllers

import (
   "fmt" 
)

type HistoryController struct {
    BaseController	
}

func (c *HistoryController) GetHistoryList() {
    fmt.Printf("Call GetHistoryList\n")
    return
}

func (c *HistoryController) GetHistoryItem() {
}
