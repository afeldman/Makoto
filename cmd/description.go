package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

var discription = &cobra.Command{
	Use:   "description [KPC PACKAGE NAME]",
	Short: "package description",
	Long: `
Each package should have a discription that indicats
to the possible user, if the package are usable for
his or her programming needs.

AUTHOR:
	Anton Feldmann <anton.feldmann@gmail.com>
`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			log.Fatal("no package in argument")
		}

	},
}
