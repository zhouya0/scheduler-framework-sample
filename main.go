package main

import (
	"fmt"
	"github.com/zhouya0/sample-scheduler-framework/pkg/plugins"
	"github.com/zhouya0/sample-scheduler-framework/pkg/multipoint"
	"github.com/zhouya0/sample-scheduler-framework/pkg/qos"
	"k8s.io/component-base/logs"
	"k8s.io/kubernetes/cmd/kube-scheduler/app"
	"math/rand"
	"os"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	command := app.NewSchedulerCommand(
		app.WithPlugin(plugins.Name, plugins.New),
		app.WithPlugin(multipoint.Name, multipoint.New),
		app.WithPlugin(qos.Name, qos.New),
	)

	logs.InitLogs()
	defer logs.FlushLogs()

	if err := command.Execute(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

}
