package main

import (
	"gomodules.xyz/logs"
	"k8s.io/klog/v2"
)

func main() {
	logs.Init(nil, true)
	defer logs.FlushLogs()

	realMain()
}

func realMain() {
	klog.Infof("hello world")
	klog.V(3).Info("hello world - 3")
	klog.V(5).Info("hello world - 5")
}
