package cmd

import (
	"log"

	"github.com/spf13/cobra"

	kpc "github.com/afeldman/kpc"
)

var contains = &cobra.Command{
	Use:   "contains [KPC PACKAGE NAME]",
	Short: "is the package part of the kpc set",
	Long: `
Befor you can start using Makoto with all power,
then you have to make shure the package is installed.
This method is going to show if a package is containt in the
set of known kpc projects.

Attention:
The name of a project can be different from the filename.
This function compaires the project name, not the file name.

AUTHOR:
	Anton Feldmann <anton.feldmann@gmail.com>
`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			log.Fatal("no package in argument")
		}

		if !kpc.ContainsPackage(args[0]) {
			log.Println("unable to find KPC: ", args[0])
		} else {
			log.Println("found KPC: ", args[0])
		}
	},
}
