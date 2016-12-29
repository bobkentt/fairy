package main

import (
    "flag"
    "fmt"
    "os"

	"github.com/astaxie/beego"

	_ "fairy/fairy-api/routers"
    "fairy/config"
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
    //fmt.Printf("LoadConfig (%s) success: (%v)\n", *configFile, config.Config)

    // server run
	beego.Run()
}

