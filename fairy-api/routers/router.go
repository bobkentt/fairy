package routers

import (
	"fairy/fairy-api/controllers"
	"github.com/astaxie/beego"
)

func init() {
    // bless
    beego.Router("/api/v1/bless", &controllers.BlessController{}, "get:Get")

    // menu
    beego.Router("/api/v1/menu/list", &controllers.MenuController{}, "get:GetMenuList")

    // product
    beego.Router("/api/v1/product", &controllers.ProductController{}, "get:Get,post:Post,delete:Delete")

    // order 
    beego.Router("/api/v1/order/list", &controllers.OrderController{}, "get:GetOrderList")
    beego.Router("/api/v1/order", &controllers.OrderController{}, "get:Get,post:Post,delete:Delete")

    // usr
    beego.Router("/api/v1/usr", &controllers.UsrController{}, "post:Post")
    beego.Router("/api/v1/usr/share", &controllers.UsrController{}, "post:Share")

    // resource
    beego.Router("/api/v1/resource/img", &controllers.ResourceController{}, "post:UploadImg")
}
