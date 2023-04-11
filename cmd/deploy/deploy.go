package deploy

import (
	"errors"
	"pain/pin8s/asker"
	"pain/pin8s/client"
	"strings"

	"github.com/spf13/cobra"
)

func NewCmdDeploy(c *client.K8sClient) *cobra.Command {

	runner := NewDeployRunner(c)

	podsCmd := &cobra.Command{
		Use:   "deploy",
		Args:  args,
		Short: "Command for k8s deploys",
		Long:  `Command for utility deploy operations`,
		Run:   runner.run,
	}

	podsCmd.Flags().StringVarP(&runner.options.Example, "term", "t", "", "only list deploys that match the given term")
	return podsCmd
}

type deployRunner struct {
	client  *client.K8sClient
	options *DeployOptions
}

type DeployOptions struct {
	Example string
}

func NewDeployRunner(c *client.K8sClient) *deployRunner {
	deployOptions := &DeployOptions{}
	return &deployRunner{client: c,
		options: deployOptions,
	}
}

func (r *deployRunner) run(cmd *cobra.Command, args []string) {
	resultLine := r.client.Deploy.ListDeploys(args[0])
	var filteredDeploys []string
	for _, v := range resultLine {
		if strings.Contains(v, r.options.Example) {
			filteredDeploys = append(filteredDeploys, v)
		}
	}
	podanswer := asker.ShowSurvey(filteredDeploys, "Deploy")
	actions := []string{"delete"}
	actionanswer := asker.ShowSurvey(actions, "Action")

	r.excecuteDeployAction(args[0], podanswer.Pod, actionanswer.Action, actions)
}

func (r *deployRunner) excecuteDeployAction(ns, podname, answer string, actions []string) {
	r.client.Deploy.DeleteDeploy(ns, podname)
}

func args(cmd *cobra.Command, args []string) error {
	if len(args) > 1 || len(args) == 0 {
		return errors.New("requires at least one arg")
	}
	return nil
}

func (o *DeployOptions) DeployOptionsSetter(cmd *cobra.Command, args []string) error {
	s, err := cmd.Flags().GetString("t")
	if err != nil {
		panic(err)
	}
	o.Example = s
	return nil
}
