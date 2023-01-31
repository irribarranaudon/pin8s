package main

import (
	"pain/pin8s/cmd"
	"pain/pin8s/kubectlexecutor"
)

func main() {
	commander := kubectlexecutor.NewCommander()
	cmd := cmd.NewCmd(commander)
	cmd.Run()
}
