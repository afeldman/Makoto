package cmd

import (
	"encoding/json"
	"fmt"
	"log"

	kpc "github.com/afeldman/Makoto"
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
		if len(args) != 1 {
			log.Fatal("no package in argument")
		}

		kpg := kpc.GetKPC(args[0])
		if kpg != nil {
			con, err := json.Marshal(&(kpg.Conflicts))
			if err != nil {
				log.Fatalln(err)
			}
			fmt.Println(string(con))
		} else {
			log.Fatalln("can not find the requested Package")
		}
	},
}
