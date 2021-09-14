package cmd

import (
	"fmt"
	"log"
	"strings"

	"github.com/afeldman/Makoto/makoto"
	"github.com/spf13/cobra"
)

var conflicts = &cobra.Command{
	Use:   "conflicts [KPC PACKAGE NAME]",
	Short: "known package conflicts",
	Long: `
Using software in big projects usually, somtimes coses error because
different software versions are not compatible over time. To avoid this
the field "conclicts lists a set of known conflicts with different software
packages.
Please feel free to include conflicts in that list, if you run into software version
conflicts.

AUTHOR:
	Anton Feldmann <anton.feldmann@gmail.com>
`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 2 || len(args) == 0 {
			log.Fatal("usage: Makoto conflicts pk_name pk_version")
		}

		var kpc *makoto.KPC_DB_Entry
		if len(args) < 2 {
			kpc = makoto.Latest(args[0])
		} else {
			kpc = makoto.Get(args[0], args[1])
		}

		if kpc != nil {
			for _, conflict := range kpc.KPC.Conflicts {
				fmt.Print(conflict.Name + "@" + strings.Join(conflict.Versions, ","))
			}
		}

	},
}
