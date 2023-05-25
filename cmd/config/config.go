package config

import (
	"pain/pin8s/asker"
	"pain/pin8s/client"

	"github.com/spf13/cobra"
)

func NewCmdConfig(c *client.K8sClient) *cobra.Command {

	runner := NewConfigRunner(c)

	configCmd := &cobra.Command{
		Use:   "context",
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
