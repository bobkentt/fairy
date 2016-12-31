package storage 

import (
    "fmt"

    "github.com/astaxie/beego/orm"
    _ "github.com/go-sql-driver/mysql"
    log "code.google.com/p/log4go"

    "fairy/config"
)

type ProductBrief struct {
    Product_id     uint32
    Product_name   string
    Product_type   uint32
    price          uint32 
    img            string
    //brief          string
}

type Product struct {
    Product_id     uint32
    Product_name   string
    Product_type   uint32
    price          uint32 
    img            string
    details        string
}

func NewInstance(cfg *config.MysqlConfig) {
    // set default database
    link := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", cfg.Usr, cfg.Passwd, cfg.Server, cfg.Name)
    orm.RegisterDataBase("default", "mysql", link, 30, 30)
}

func GetProductList(merchant_id int) (int64, []ProductBrief, error) {
    var pts []ProductBrief
    o := orm.NewOrm()
    num, err := o.Raw("SELECT product_id, product_name, product_type, price, brief FROM v1_product WHERE merchant_id = ?", merchant_id).QueryRows(&pts)
    if err != nil {
        log.Warn("GetProduct failed err = %s, merchant_id = %d", err, merchant_id)
        return 0, pts, err
    }
    return num, pts, nil
}

func GetProduct(product_id int) (Product, error) {
    return _, nil
}
