package cmd

import (
	"pain/pin8s/asker"
	"pain/pin8s/client"

	"github.com/spf13/cobra"
)

func (cmd *pin8sCmd) NewCmdConfig() *cobra.Command {

	runner := NewConfigRunner(cmd.client)

	configCmd := &cobra.Command{
		Use:   "config",
		Short: "Config for k8s contexts",
		Long:  `Config for select k8s contexts interactively`,
		Run:   runner.run,
	}
	return configCmd
}

type configRunner struct {
	client *client.K8sClient
}

func NewConfigRunner(c *client.K8sClient) *configRunner {
	return &configRunner{client: c}
}

func (r *configRunner) run(cmd *cobra.Command, args []string) {
	resultLine := r.client.Config.GetContexts()
	answer := asker.ShowSurvey(resultLine.Names, "Context", asker.WithDefaultValue(resultLine.Selected))
	r.client.Config.SelectContext(answer.Context)
}
