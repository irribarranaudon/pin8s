package pod

import (
	"errors"
	"pain/pin8s/asker"
	"pain/pin8s/client"
	"strings"

	"github.com/spf13/cobra"
)

func NewCmdPod(c *client.K8sClient) *cobra.Command {

	runner := NewConfigPod(c)

	podsCmd := &cobra.Command{
		Use:   "pod",
		Args:  args,
		Short: "Command for k8s pods",
		Long:  `Command for utility pods operations`,
		Run:   runner.run,
	}

	podsCmd.Flags().StringVarP(&runner.options.Example, "term", "t", "", "only list pods that match the given term")
	return podsCmd
}

type podRunner struct {
	client  *client.K8sClient
	options *PodOptions
}

type PodOptions struct {
	Example string
}

func NewConfigPod(c *client.K8sClient) *podRunner {
	podOptions := &PodOptions{}
	return &podRunner{client: c,
		options: podOptions,
	}
}

func args(cmd *cobra.Command, args []string) error {
	if len(args) > 1 || len(args) == 0 {
		return errors.New("a namespace argument is required")
	}
	return nil
}

func (r *podRunner) run(cmd *cobra.Command, args []string) {
	resultLine := r.client.Pod.ListPods(args[0])
	var filteredPods []string
	for _, v := range resultLine {
		if strings.Contains(v, r.options.Example) {
			filteredPods = append(filteredPods, v)
		}
	}
	podanswer := asker.ShowSurvey(filteredPods, "Pod")
	actions := []string{"delete", "logs"}
	actionanswer := asker.ShowSurvey(actions, "Action")

	r.excecutePodAction(args[0], podanswer.Pod, actionanswer.Action, actions)
}

func (r *podRunner) excecutePodAction(ns, podname, answer string, actions []string) {
	switch answer {
	case actions[1]: //logs
		r.client.Pod.Logs(ns, podname)
	default: //delete
		r.client.Pod.DeletePod(ns, podname)
	}
}

func (o *PodOptions) PodOptionsSetter(cmd *cobra.Command, args []string) error {
	s, err := cmd.Flags().GetString("t")
	if err != nil {
		panic(err)
	}
	o.Example = s
	return nil
}
