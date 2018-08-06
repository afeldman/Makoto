package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"

	kpc "github.com/afeldman/Makoto/kpc"
)

var pversion = &cobra.Command{
	Use:   "version [KPC PACKAGE NAME]",
	Short: "show the version of the package",
	Long: `
each package should have a version. The version is an
indicator of the compability to other software packages.

E.g. I build a software using Karel String Package version 0.3.4,
that does not mean, that my Software runs with package version 0.4.1

So please make shure all your packages uses a version.
`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			log.Fatal("no package in argument")
		}

		kpg := kpc.GetKPC(args[0])
		if kpg != nil {
			fmt.Println(kpg.Version)
		} else {
			log.Fatalln("The requested Package is not detecable")
		}
	},
}
