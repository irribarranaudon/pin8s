package cmd

import (
	"fmt"
	"os"
	"pain/pin8s/client"
	"pain/pin8s/cmd/config"
	"pain/pin8s/cmd/deploy"
	"pain/pin8s/cmd/pod"

	"github.com/spf13/cobra"
)

type Pin8sCmd struct {
	Client *client.K8sClient
}

func newPpin8sCommand(client *client.K8sClient) *Pin8sCmd {
	return &Pin8sCmd{
		Client: client,
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
		Short: "pin8s is a utility tool for basic k8s operations",
		Long:  `A basic command tool utility for interactive basic k8s operations`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Iniciando")
		},
	}

	groups := CommandGroups{
		{
			Message: "Basic Commands",
			Commands: []*cobra.Command{
				config.NewCmdConfig(cmd.Client),
				pod.NewCmdPod(cmd.Client),
				deploy.NewCmdDeploy(cmd.Client),
			},
		},
	}

	groups.Add(rootCmd)

	return rootCmd
}
