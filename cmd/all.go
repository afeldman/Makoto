package cmd

import (
	"fmt"

	"github.com/afeldman/Makoto/makoto"
	"github.com/spf13/cobra"
)

var all = &cobra.Command{
	Use:   "all",
	Short: "List all available packages",
	Long: `List all packages currently registered in the local Makoto database.

Each package is displayed with its name and version (name@version).`,
	RunE: func(cmd *cobra.Command, args []string) error {
		pkgs := makoto.All()
		if len(pkgs) == 0 {
			fmt.Println("no packages registered")
			return nil
		}

		for _, kpc := range pkgs {
			fmt.Printf("%s@%s\n", kpc.Name, kpc.Version)
		}
		return nil
	},
}

func init() {
	Makoto.AddCommand(all)
}
