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
    Price          uint32
    Img            string
    Details        string
}

func NewInstance(cfg *config.MysqlConfig) {
    // set default database
    link := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", cfg.Usr, cfg.Passwd, cfg.Server, cfg.Name)
    orm.RegisterDataBase("default", "mysql", link, 30, 30)
}

func GetProductList(merchant_id int) (int64, []ProductBrief, error) {
    var pts []ProductBrief
    o := orm.NewOrm()
    num, err := o.Raw("SELECT product_id, product_name, product_type, price, img FROM v1_product WHERE merchant_id = ?", merchant_id).QueryRows(&pts)
    if err != nil {
        log.Warn("GetProduct failed err = %s, merchant_id = %d", err, merchant_id)
        return 0, pts, err
    }
    return num, pts, nil
}

func GetProduct(product_id int) (Product, error) {
    var pt Product
    o := orm.NewOrm()
    err := o.Raw("SELECT product_id, product_name, product_type, price, img, details FROM v1_product WHERE product_id = ?", product_id).QueryRow(&pt)
    if err != nil {
        log.Warn("GetProduct failed err = %s, product_id = %d", err, product_id)
        return pt, err
    }
    return pt, nil
}

func SetProduct(p* Product) error {
    if p.product_id == 0 {
        return createProduct(p)
    }
    return editProduct(p)
}

func createProduct(p* Product) error {
    return nil
}

func editProduct(p* Product) error {
    return nil
}

func RemoveProduct(product_id int) error {
    // update status into config.STATUS_DELETE 
    return nil
}

type Order struct {
    Order_id uint32
    Detail []stirng
}

func GetOrderList(uid int, offset int, limit int) ([]Order, error) {
    return nil, nil
}

func CreateOrder(uid int, order Order) (error) {
    return nil
}

func GetOrderDetail(order_id int) (Order, error) {
    var order Order
    return order, nil
}

func RemoveOrder(order_id int) (Order, error) {
    var order Order
    return order, nil
}
