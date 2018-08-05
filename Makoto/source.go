package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"

	kpc "github.com/afeldman/Makoto"
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
		if len(args) != 1 {
			log.Fatal("no package in argument")
		}

		kpg := kpc.GetKPC(args[0])
		if kpg != nil {
			fmt.Print(kpg.Source)
		} else {
			log.Fatalln("The requested pacakge can nor be found")
		}
	},
}
