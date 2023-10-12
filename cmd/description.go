package cmd

import (
	"log"

	//"github.com/afeldman/Makoto/makoto"
	"github.com/spf13/cobra"
)

var discription = &cobra.Command{
	Use:   "desc [KPC PACKAGE NAME]",
	Short: "package description",
	Long: `
Each package should have a discription that indicats
to the possible user, if the package are usable for
his or her programming needs.

AUTHOR:
	Anton Feldmann <anton.feldmann@gmail.com>
`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 2 || len(args) == 0 {
			log.Fatal("usage: Makoto description pk_name pk_version")
		}

		/*var kpc *makoto.KPC_DB_Entry
		if len(args) < 2 {
			kpc = makoto.Latest(args[0])
		} else {
			kpc = makoto.Get(args[0], args[1])
		}

		if kpc != nil {
			fmt.Print(kpc.KPC.Description)
		}*/

	},
}
