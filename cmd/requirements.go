package cmd

import (
	"fmt"

	"github.com/afeldman/Makoto/makoto"
	"github.com/spf13/cobra"
)

var requirements = &cobra.Command{
	Use:   "req [KPC PACKAGE NAME] [VERSION]",
	Short: "List package requirements",
	Long: `Packages can declare requirements (dependencies) on other KPC packages.

This command shows the required packages and their versions as declared
in the KPC file. Useful when preparing to fetch and build all dependencies.

Author:
  Anton Feldmann <anton.feldmann@gmail.com>`,
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

		if kpc == nil {
			fmt.Println("package not found")
			return nil
		}

		if len(kpc.KPC.Requirements) == 0 {
			fmt.Println("no requirements")
			return nil
		}

		for _, req := range kpc.KPC.Requirements {
			fmt.Printf("%s@%s\n", req.Name, req.Version)
		}
		return nil
	},
}

func init() {
	kpc_cmd.AddCommand(requirements)
}
