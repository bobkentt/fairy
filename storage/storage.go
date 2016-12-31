package storage 

import (
    "fmt"

    "github.com/astaxie/beego/orm"
    _ "github.com/go-sql-driver/mysql"
    log "code.google.com/p/log4go"

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

func GetProduct(product_id int64) (int64, []Product, error) {
    var pts []Product
    o := orm.NewOrm()
    num, err := o.Raw("SELECT product_id, product_name, product_type FROM v1_product WHERE product_id = ?", product_id).QueryRows(&pts)
    if err != nil {
        log.Warn("GetProduct failed err = %s, product_id = %d", err, product_id)
        return 0, pts, err
    }
    return num, pts, nil
}
