package cmd

import (
	"fmt"

	"github.com/afeldman/Makoto/makoto"
	"github.com/spf13/cobra"
)

var description = &cobra.Command{
	Use:   "desc [KPC PACKAGE NAME] [VERSION]",
	Short: "Show package description",
	Long: `Each package should have a description that indicates
to the user whether the package is usable for their programming needs.

Author:
  Anton Feldmann <anton.feldmann@gmail.com>
`,
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
			fmt.Println(*kpc.KPC.GetDescription())
		} else {
			fmt.Println("package not found")
		}
		return nil
	},
}

func init() {
	kpc_cmd.AddCommand(description)
}
