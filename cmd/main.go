package main

import (
	"github.com/skbt-ecom/logging"
	"time"
)

func main() {
	log := logging.InitLogger()
	log.AddGraylogHook("localhost:12201", "local")
	//log.SetFormatter()

	for {
		log.WithExtraField("Hello", "World").Infof("adasdasdasdas")
		time.Sleep(3 * time.Second)
	}

}
