package storage 

import (
    "fmt"
    "github.com/astaxie/beego/orm"
    _ "github.com/go-sql-driver/mysql"

    "fairy/config"
)

type Product struct {
    Product_id uint32
    Product_name   string
    Product_type   uint32
    //price  uint32 
    //brief  string
    //detail string
}

func NewInstance(cfg *config.MysqlConfig) {
    // set default database
    link := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", cfg.Usr, cfg.Passwd, cfg.Server, cfg.Name)
    orm.RegisterDataBase("default", "mysql", link, 30, 30)
}

func GetProduct() (int64, []Product) {
    o := orm.NewOrm()
    var pts []Product
    num, err := o.Raw("SELECT product_id, product_name, product_type FROM v1_product WHERE product_id = ?", 1).QueryRows(&pts)
    if err != nil {
        fmt.Printf("getProduct err : %s", err)
    }
    return num, pts
}
