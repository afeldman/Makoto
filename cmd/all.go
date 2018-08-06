package cmd

import (
	"fmt"

	kpc "github.com/afeldman/Makoto/kpc"
	"github.com/spf13/cobra"
)

var all = &cobra.Command{
	Use:   "all",
	Short: "List all available packages",
	Long: `
Makoto works an a list of central yaml files.
All files in this central foulder how have the ending kcp
and the name tag can be displayed.
Make shure your kpc file contains all information.
as more as information the project kpc file contains,
as more as usefull is Makoto

AUTHOR:
	anton feldmann <anton.feldmann@gmail.com>
`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(kpc.All())
	},
}
