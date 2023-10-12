package cmd

import (
	//"github.com/afeldman/Makoto/makoto"

	"github.com/spf13/cobra"
)

var (
	Makoto = &cobra.Command{
		Use:   "Makoto",
		Short: "Makoto is a karel package configurator",
		Long: `
	   	  MAKOTO A KAREL PACKAGE CONFIGURATOR
	   	  ===================================

Makoto is a Package Configurator for Karel files.
The idea to this tool was given using package config in linux.
The package configure it is easy to build your code because
the tool handels the setting for the compiler flags.

Beacause FANUC KAREL has a limitation on the number of lines in
the code I do have the programing stype to seperate all KAREL
information in dedecated files. To get the information together
this tool helps me.

AUTHOR:
	Anton Feldmann <anton.feldmann@gmail.com>

\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\`,
	}

	confFile string
)

func init() {
	cobra.OnInitialize(initConfig)

	Makoto.PersistentFlags().StringVar(&confFile, "config",
		"",
		"config file (default $HOME/.config/makoto/makoto.toml)")

	Makoto.AddCommand(version)
	Makoto.AddCommand(kpc_cmd)
	Makoto.AddCommand(all)
}

func Execute() {
	Makoto.Execute()
}

func initConfig() {

	/*var makoto_ makoto.Makoto
	if len(strings.TrimSpace(confFile)) == 0 {
		makoto_ = *makoto.Init()
	} else {
		makoto_ = *makoto.Config(confFile)
	}
	makoto_.Store()
	makoto_.DBInit()*/

}
