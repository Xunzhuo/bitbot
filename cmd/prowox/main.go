package main

import (
	"github.com/xunzhuo/prowox/cmd/prowox/config"
	"github.com/xunzhuo/prowox/pkg/commands"
	"github.com/xunzhuo/prowox/pkg/core"
	"k8s.io/klog"
)

func main() {
	klog.Info("Starting Prowox ...")

	if err := config.InitConfig(); err != nil {
		klog.Error(err)
		return
	}

	if config.Get().TYPE == "schedule" {
		if err := commands.MergeAcceptedPRs(); err != nil {
			klog.Error(err)
		}
	}

	if config.Get().TYPE == "created" {
		commands.Greeting()
		commands.Help()
		if err := core.RunCommands(); err != nil {
			klog.Error(err)
		}
	}

	if config.Get().TYPE == "comment" {
		if err := core.RunCommands(); err != nil {
			klog.Error(err)
		}
	}
}
