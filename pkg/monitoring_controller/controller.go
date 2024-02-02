package monitoring_controller

import (
	"fmt"
	"time"
)

type MonitoringController struct{}

func (mc *MonitoringController) Start() error {
	for {
		time.Sleep(time.Second)
		fmt.Println("Running execution loop")
	}
	return nil
}
