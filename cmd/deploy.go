package cmd

/*import (
	"log"
	"pain/pin8s/kubectlexecutor"

	"github.com/spf13/cobra"
)

func (p *pin8sCmd) NewCmdDeploys() *cobra.Command {
	configCmd := &cobra.Command{
		Use:   "deploys",
		Short: "deploys operations",
		Long:  `Command to execute deploys operations interactively`,
		Run:   NewDeployRunner(p.Commander).run,
	}
	return configCmd
}

type deployRunner struct {
	Commander kubectlexecutor.Pin8sCommander
}

func NewDeployRunner(c kubectlexecutor.Pin8sCommander) *deployRunner {
	return &deployRunner{Commander: c}
}

func (r *deployRunner) run(cmd *cobra.Command, args []string) {
	if len(args) < 1 || len(args) > 1 {
		log.Fatal("1 arg required")
	}
	r.Commander.ListPods(args[0])
}*/
