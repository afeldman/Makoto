package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

var prefix = &cobra.Command{
	Use:   "prefix [KPC PACKAGE NAME]",
	Short: "the install prefix",
	Long: `
To save all information about the pathes - where the package parts are installed -
can be wery long. So I did put the prefix field into the kpc structure.

Please be aware, Makoto does not put the prefix together with the other pathes.
the pathes for the {type,const,header,source} files are relative to the prefix.
So to use the different pathes, get first the prefix and put the information together.
This makes the use in makefiles more comfortable.

AUTHOR:
	Anton Feldmann <anton.feldmann@gmail.com>
`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			log.Fatal("no package in argument")
		}

		/*kpg := kpc.GetKPC(args[0])
		if kpg != nil {
			fmt.Println(kpg.Prefix)
		} else {
			log.Fatalln("Can not find the required package")
		}*/
	},
}
