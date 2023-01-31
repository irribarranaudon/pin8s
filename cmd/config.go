package cmd

import (
	"github.com/spf13/cobra"
)

func (p *pin8sCmd) NewCmdConfig() *cobra.Command {
	configCmd := &cobra.Command{
		Use:   "config",
		Short: "Config for k8s contexts",
		Long:  `Config for select k8s contexts interactively`,
		Run: func(cmd *cobra.Command, args []string) {
			p.Commander.SelectConfigContext()
		},
	}
	return configCmd
}
