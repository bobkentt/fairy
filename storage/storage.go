package storage 

import (
    "fmt"
    "github.com/astaxie/beego/orm"
    _ "github.com/go-sql-driver/mysql"
)

func NewInstance() {
    // set default database
    orm.RegisterDataBase("fairy", "mysql", "root:root@/my_db?charset=utf8", 30)
    
}
