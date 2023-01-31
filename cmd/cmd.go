package cmd

import (
	"fmt"
	"os"
	"pain/pin8s/kubectlexecutor"
	"pain/pin8s/util"

	"github.com/spf13/cobra"
)

type pin8sCmd struct {
	Commander kubectlexecutor.Pin8sCommander
}

func NewCmd(commander kubectlexecutor.Pin8sCommander) *pin8sCmd {
	return &pin8sCmd{Commander: commander}
}

func (cmd *pin8sCmd) Run() {
	rootCmd := cmd.NewPpin8sCommand()
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func (cmd *pin8sCmd) NewPpin8sCommand() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "pin8s",
		Short: "Hugo is a very fast static site generator",
		Long:  `A basic command tool utility for interactive basic operations with kubectl`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("este es el terrible inicio")
		},
	}

	groups := util.CommandGroups{
		{
			Message: "Basic Commands",
			Commands: []*cobra.Command{
				cmd.NewCmdConfig(),
			},
		},
	}

	groups.Add(rootCmd)

	return rootCmd
}
