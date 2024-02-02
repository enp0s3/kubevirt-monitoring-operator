package main

import (
	"fmt"
	"github.com/enp0s3/kubevirt-monitoring-operator/pkg/monitoring_controller"
)

func main() {
	fmt.Println("starting monitoring controller")
	ctrl := monitoring_controller.MonitoringController{}
	ctrl.Start()
}
