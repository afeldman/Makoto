package cmd

import (
	
	"github.com/afeldman/Makoto/makoto"
	log "github.com/sirupsen/logrus"

	"github.com/spf13/cobra"
)

var authors = &cobra.Command{
	Use:   "author [KPC PACKAGE NAME]",
	Short: "show authors of kpc package",
	Long: `
Each KPC package should have atleased on author.
The auther is responceble for futher progres on the
software and the one you can conntact if you run into problems
using the package.

AUTHOR:
	anton feldmann <anton.feldmann@gmail.com>
`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 2 || len(args) == 0 {
			log.Fatalln("usage: makoto author pk_name pk_version")
		}

		var kpc *makoto.KPC_DB_Entry
		if len(args) < 2 {
			kpc = makoto.Latest(args[0])
		} else {
			kpc = makoto.Get(args[0], args[1])
		}

		if kpc != nil {
			/*for _, author := range kpc.Authors {
				fmt.Println(author.Name + " " + author.Email)
			}*/
		}
	},
}
