package config

import (
	"context"
	"log"

	"k8s.io/client-go/tools/clientcmd"
)

type ConfigClient struct {
	ContextsOptions *ContextsOptions
	ctx             context.Context
}

type ContextsOptions struct {
	configAccess clientcmd.ConfigAccess
}

type Config interface {
	GetContexts() *ResultLine
	SelectContext(ctxname string)
}

type ResultLine struct {
	Selected string
	Names    []string
}

func NewConfigClient(pathOptions *clientcmd.PathOptions) *ConfigClient {
	options := &ContextsOptions{
		configAccess: pathOptions,
	}
	ctx := context.TODO()
	return &ConfigClient{
		ContextsOptions: options,
		ctx:             ctx,
	}
}

func (c *ConfigClient) GetContexts() *ResultLine {
	config, err := c.ContextsOptions.configAccess.GetStartingConfig()
	if err != nil {
		log.Fatal(err)
	}
	rl := &ResultLine{Selected: config.CurrentContext}
	var names []string
	for key := range config.Contexts {
		names = append(names, key)
	}
	rl.Names = names
	return rl
}

func (c *ConfigClient) SelectContext(ctxname string) {
	configAccess := c.ContextsOptions.configAccess
	config, err := configAccess.GetStartingConfig()
	if err != nil {
		log.Fatal(err)
	}

	_, exists := config.Contexts[ctxname]

	if !exists {
		log.Fatal("context does not exist")
	}

	config.CurrentContext = ctxname

	clientcmd.ModifyConfig(configAccess, *config, true)
}
