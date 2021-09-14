package cmd

import (
	"fmt"
	"log"

	"github.com/afeldman/Makoto/makoto"
	"github.com/spf13/cobra"
)

var source = &cobra.Command{
	Use:   "source [KPC PACKAGE NAME]",
	Short: "Source repo url",
	Long: `
The package URL can differ from the actual repository url.
This tag makes clear where to find the sourcecode of the
requested package.

I did not restrict the useage of a sorce management tool.
It is up to you to set the url and use yout prefered
scm to download the code and build it.

AUTHOR:
	Anton Feldmann <anton.feldmann@gmail.com>
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
			fmt.Println(kpc.KPC.Main)
		}
	},
}
