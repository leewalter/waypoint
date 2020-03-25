package cli

import (
	"github.com/mitchellh/devflow/internal/plugin"
	"github.com/mitchellh/devflow/sdk"
)

type PluginCommand struct {
	*baseCommand
}

func (c *PluginCommand) Run(args []string) int {
	plugin, ok := plugin.Builtins[args[0]]
	if !ok {
		panic("no such plugin")
	}

	// Run the plugin
	sdk.Main(plugin...)
	return 0
}

func (c *PluginCommand) Synopsis() string {
	return "Execute a built-in plugin."
}

func (c *PluginCommand) Help() string {
	return ""
}