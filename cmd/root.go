package cmd

import (
	"fmt"
	"os"
	"pain/pin8s/client"

	"github.com/spf13/cobra"
)

type pin8sCmd struct {
	client *client.K8sClient
}

func newPpin8sCommand(client *client.K8sClient) *pin8sCmd {
	return &pin8sCmd{
		client: client,
	}
}

func Execute() {

	rootCmd := loadCommand()
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

//conn := client.InitConnection(k8sCfg)

func loadCommand() *cobra.Command {
	client, err := client.NewK8sClient()

	if err != nil {
		panic(err)
	}

	cmd := newPpin8sCommand(client)

	rootCmd := &cobra.Command{
		Use:   "pin8s",
		Short: "Hugo is a very fast static site generator",
		Long:  `A basic command tool utility for interactive basic operations with kubectl`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Iniciando")
		},
	}

	groups := CommandGroups{
		{
			Message: "Basic Commands",
			Commands: []*cobra.Command{
				cmd.NewCmdConfig(),
				cmd.NewCmdPod(),
			},
		},
	}

	groups.Add(rootCmd)

	return rootCmd
}
