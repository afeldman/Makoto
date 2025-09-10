package cmd

import (
	"fmt"

	"github.com/afeldman/Makoto/makoto"
	"github.com/spf13/cobra"
)

var (
	Version  = "dev"
	pversion = &cobra.Command{
		Use:   "version [KPC PACKAGE NAME] [VERSION]",
		Short: "Show the version of a KPC package",
		Long: `Each KPC package should have a version. The version is an
indicator of compatibility with other software packages.

For example, if you build software using Karel String Package version 0.3.4,
that does not guarantee it will run with version 0.4.1.

So please make sure all your packages declare a version.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 || len(args) > 2 {
				return cmd.Usage()
			}

			var kpc *makoto.KPC_DB_Entry
			if len(args) == 1 {
				kpc = makoto.Latest(args[0])
			} else {
				kpc = makoto.Get(args[0], args[1])
			}

			if kpc != nil {
				fmt.Println(kpc.KPC.Version)
			} else {
				fmt.Println("package not found")
			}
			return nil
		},
	}
)

func init() {
	kpc_cmd.AddCommand(pversion)
}
