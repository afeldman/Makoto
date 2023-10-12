package cmd

import (
	"log"
	"github.com/spf13/cobra"
)

var requirements = &cobra.Command{
	Use:   "req [KPC PACKAGE NAME]",
	Short: "list package requirements",
	Long: `
As imagne packages can need requirements. To have a list of
all necessary requirements, to build the requested package
this list can be used.
This list might be helpful if your are making a dowload to get
all the required packages and build them.

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
			for _, req := range kpc.KPC.Requirements {
				fmt.Println(req.Name + "@" + req.Version)
			}
		}*/
	},
}
