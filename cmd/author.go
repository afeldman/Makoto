package cmd

import (
	"encoding/json"
	"fmt"
	"log"

	kpc "github.com/afeldman/kpc"
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
		if len(args) != 1 {
			log.Fatal("no package in argument")
		}

		kpg := kpc.GetKPC(args[0])
		if kpg != nil {
			author, err := json.Marshal(&(kpg.Authors))
			if err != nil {
				log.Fatalln(err)
			}
			fmt.Println(string(author))
		} else {
			log.Fatalln("can not find the requested package")
		}
	},
}
