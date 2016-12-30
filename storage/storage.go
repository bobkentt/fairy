package storage 

import (
    "fmt"
    "github.com/astaxie/beego/orm"
    _ "github.com/go-sql-driver/mysql"
)

type Product struct {
    product_id uint32
    product_name   string
    product_type   uint32
    price  uint32 
    //brief  string
    //detail string
}

func NewInstance() {
    // set default database
    //orm.RegisterDataBase("default", "mysql", "root:wTJWj12XmIN9@/182.92.200.3:3306?charset=utf8", 30)
    dbUser := "root"    
    dbName := "fairy" //数据库名字
    dbPwd := "wTJWj12XmIN9"    
    dbLink := fmt.Sprintf("%s:%s@/%s?charset=utf8", dbUser, dbPwd, dbName)  
    //orm.RegisterDataBase("default", "mysql", "root:@/182.92.200.3:3306?charset=utf8", 30)
    orm.RegisterDataBase("default", "mysql", dbLink) 
}

func GetProduct() {
    o := orm.NewOrm()
    var pt Product
    num, err := o.Raw("SELECT product_id, product_name, product_type FROM v1_product WHERE product_id = ?", 1).QueryRows(&pt)
    if err != nil {
        fmt.Printf("getProduct err : %s", err)
    }
    fmt.Printf("getProduct success num = (%v), product = (%v)", num, pt)
}
