package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

var url = &cobra.Command{
	Use:   "url [KPC PACKAGE NAME]",
	Short: "the package homepage",
	Long: `
If necessary some projects have homepages to describe
the tool and the SDK or API, if it is a library.

AUTHOR:
	Anton Feldmann <anton.feldmann@gmail.com>
`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			log.Fatal("no package in argument")
		}

		/*kpg := kpc.GetKPC(args[0])
		if kpg != nil {
			fmt.Print(kpg.Homepage)
		} else {
			log.Fatalln("can not find the requested package")
		}*/
	},
}
