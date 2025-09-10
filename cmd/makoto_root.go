package cmd

import (
	"github.com/afeldman/Makoto/makoto"
	"github.com/spf13/cobra"
)

var (
	confFile string

	Makoto = &cobra.Command{
		Use:   "makoto",
		Short: "Karel package configurator (KPC/TOML)",
		Long: `Makoto - Karel Package Configurator

Manage and inspect KPC package metadata written in TOML.

Config:
  - Default config path: ~/.config/makoto/makoto.toml
  - Override with --config <path>

Common commands:
  makoto version
  makoto package version <name> [version]
  makoto package author  <name> [version]
  makoto all
`,
	}
)

func init() {
	// Global flag
	Makoto.PersistentFlags().StringVar(
		&confFile,
		"config",
		"",
		"Path to Makoto config file (default: $HOME/.config/makoto/makoto.toml)",
	)

	// Hook into cobra’s lifecycle
	cobra.OnInitialize(initConfig)

}

func Execute() {
	_ = Makoto.Execute()
}

func initConfig() {
	// Initialize Makoto config + DB
	var m *makoto.Makoto
	if confFile == "" {
		m = makoto.InitMakoto()
	} else {
		m = makoto.ConfigMakoto(confFile)
	}
	m.StoreMakoto()
	m.DBInit()
}
