package main

import (
    "flag"
    "fmt"
    "os"

	"github.com/astaxie/beego"

    "fairy/config"
	_ "fairy/fairy-api/routers"
    "fairy/storage"
)

func main() {
    // param parse
    var configFile = flag.String("c", "./conf/conf.json", "Config file")
    flag.Parse()

    // config load
    err := config.LoadConfig(*configFile)
    if err != nil {
        fmt.Printf("LoadConfig (%s) failed: (%s)\n", *configFile, err)
        os.Exit(1)
    }

    storage.NewInstance(&config.Config.Mysql)

    // server run
	beego.Run()
}

