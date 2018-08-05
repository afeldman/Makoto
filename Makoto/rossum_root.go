package cmd

import (
	"log"
	"path"

	homedir "github.com/atrox/homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/afeldman/go-util/env"
)

var (
	Rossum = &cobra.Command{
		Use:   "rossum",
		Short: "Rossum is a karel package configurator",
		Long: `
	   	  ROSSUM A KAREL PACKAGE CONFIGURATOR
	   	  ===================================

Rossum is a Package Configurator for Karel files.
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

	Rossum.PersistentFlags().StringVar(&confFile, "config", "", "config file (default $HOME/.config/rossum/rossum.yaml)")

	Rossum.AddCommand(version)
	Rossum.AddCommand(kpc_cmd)
	Rossum.AddCommand(all)
}

func Execute() {
	Rossum.Execute()
}

func initConfig() {
	if confFile != "" {
		viper.SetConfigFile(confFile)
	} else {

		home, err := homedir.Dir()
		if err != nil {
			log.Fatal(err)
		}

		confFile = path.Join(home, ".config", "rossum", "rossum.yaml")

		viper.AddConfigPath(path.Join(home, ".config", "rossum"))
		viper.SetConfigName("rossum")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Println(err)
	} else {
		if err := viper.Unmarshal(&rfg); err != nil {
			log.Fatal("unable to decode into the rossum configuration structure, %v", err)
		}
	}

	if len(rfg.RootDir) > 0 {
		rfg.init(env.GetEnv("KPC_PATH"))
	}
	rfg.save(confFile)

}