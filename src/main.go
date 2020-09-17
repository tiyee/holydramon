package main

import (
	"flag"
	"fmt"
	"github.com/tiyee/holydramon/src/engine"
	"os"
)

var configFilePath string

func main() {
	e := engine.New()
	if err := e.InitConfig(&configFilePath); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	if err := e.InitLogger(); err != nil {
		defer func() {
			if err := e.Logger().Sync(); err != nil {
				fmt.Println(err.Error())
			}
		}()
		fmt.Println(err.Error())
		os.Exit(1)
	}
	if err := initRouter(e); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	if err := initHooks(e); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	if err := e.InitComponent(); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	if err := e.Run(); err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("hello world")
}

func init() {
	flag.StringVar(&configFilePath, "c", "", "-c : the config path")
	flag.Parse()
	if _, err := os.Stat(configFilePath); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

}
