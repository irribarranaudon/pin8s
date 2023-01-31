package kubectlexecutor

type Pin8sCommander interface {
	SelectConfigContext()
}

type commander struct {
}

func NewCommander() Pin8sCommander {
	return &commander{}
}
