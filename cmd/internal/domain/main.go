package main

import (
	"log"

	"github.com/mirno/goexplorer/internal/drivers/driver1"
	"github.com/mirno/goexplorer/internal/drivers/driver2"
	"github.com/mirno/goexplorer/pkg/config"
)

func main() {
    conf, err := config.InitializeAppConfig("config.dt.yaml")
    if err != nil {
        log.Fatalf("Failed to load configuration: %v", err)
    }
	log.Printf("Loaded Config: %+v\n", conf)

    for name, driverConfig := range conf.Drivers {
        switch name {
        case "driver1":
            d1 := driver1.Driver1{Config: driverConfig}
			log.Printf("Loaded Config: %+v\n", conf)
            d1.PerformTask()
        case "driver2":
            d2 := driver2.Driver2{Config: driverConfig}
			log.Printf("Loaded Config: %+v\n", conf)
            d2.PerformTask()
        }
    }
}