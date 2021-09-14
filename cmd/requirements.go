package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

var requirements = &cobra.Command{
	Use:   "requirements [KPC PACKAGE NAME]",
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
		if len(args) != 1 {
			log.Fatal("no package in argument")
		}

		/*kpg := kpc.GetKPC(args[0])
		if kpg != nil {
			buff, err := json.Marshal(&kpg.Requirements)
			if err != nil {
				log.Fatalln(err)
			}
			fmt.Println(string(buff))
		} else {
			log.Fatalln("can not find required package")
		}*/
	},
}
