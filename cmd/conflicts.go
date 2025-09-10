package cmd

import (
	"fmt"

	"github.com/afeldman/Makoto/makoto"
	"github.com/spf13/cobra"
)

var conflicts = &cobra.Command{
	Use:   "conflicts [KPC PACKAGE NAME] [VERSION]",
	Short: "List known package conflicts",
	Long: `Large projects can suffer from incompatibilities across versions.
The "conflicts" field in the KPC describes packages or versions that
are known to be incompatible.

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

		if len(kpc.KPC.Conflicts) == 0 {
			fmt.Println("no conflicts")
			return nil
		}

		for _, c := range kpc.KPC.Conflicts {
			if len(c.Versions) > 0 {
				fmt.Printf("%s@%v\n", c.Name, c.Versions)
			} else {
				fmt.Printf("%s\n", c.Name)
			}
		}
		return nil
	},
}

func init() {
	kpc_cmd.AddCommand(conflicts)
}
