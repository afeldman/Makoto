package cmd

import (
	"fmt"
	"log"

	"github.com/afeldman/Makoto/makoto"
	"github.com/spf13/cobra"
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
		if len(args) > 2 || len(args) == 0 {
			log.Fatal("usage: Makoto description pk_name pk_version")
		}

		var kpc *makoto.KPC_DB_Entry
		if len(args) < 2 {
			kpc = makoto.Latest(args[0])
		} else {
			kpc = makoto.Get(args[0], args[1])
		}

		if kpc != nil {
			fmt.Println(kpc.KPC.Version)
		}
	},
}
