package routers

import (
	"fairy/fairy-api/controllers"
	"github.com/astaxie/beego"
)

func init() {
    // bless
    beego.Router("/api/v1/bless", &controllers.BlessController{}, "get:Get")    

    // usr
    beego.Router("/api/v1/usr", &controllers.UsrController{}, "post:Post")
    beego.Router("/api/v1/usr/cheer", &controllers.UsrController{}, "post:Cherr")
    beego.Router("/api/v1/usr/share", &controllers.UsrController{}, "post:Share")

    // habit
    beego.Router("/api/v1/habit", &controllers.HabitController{}, "post:Post")

    // coach
    beego.Router("/api/v1/coach/link", &controllers.CoachController{}, "post:Link")
    beego.Router("/api/v1/coach/report", &controllers.CoachController{}, "post:Report")

    // margin
    beego.Router("/api/v1/margin/recharge", &controllers.MarginController{}, "post:Recharge")
    beego.Router("/api/v1/margin/refund", &controllers.MarginController{}, "post:Refund")
    // history
    beego.Router("/api/v1/history", &controllers.HistoryController{}, "get:Get")
    beego.Router("/api/v1/history/item", &controllers.HistoryController{}, "get:GetHistoryItem")

    // resource
    beego.Router("/api/v1/resource/img", &controllers.ResourceController{}, "post:UploadImg")
}
