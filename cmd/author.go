package cmd

import (
	"fmt"

	"github.com/afeldman/Makoto/makoto"
	"github.com/spf13/cobra"
)

var authors = &cobra.Command{
	Use:   "author [KPC PACKAGE NAME] [VERSION]",
	Short: "Show authors of a KPC package",
	Long: `Each KPC package should have at least one author.
Authors are responsible for the package's further progress and are
the contact if you run into problems using the package.

Author:
  Anton Feldmann <anton.feldmann@gmail.com>`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 || len(args) > 2 {
			return cmd.Usage()
		}

		var kpc *makoto.KPC_DB_Entry
		if len(args) == 1 {
			kpc = makoto.Latest(args[0])
		} else {
			kpc = makoto.Get(args[0], args[1])
		}

		if kpc == nil {
			fmt.Println("package not found")
			return nil
		}

		if len(kpc.KPC.Authors) == 0 {
			fmt.Println("no authors")
			return nil
		}

		for _, author := range kpc.KPC.Authors {
			if author.Email != "" {
				fmt.Printf("%s <%s>\n", author.Name, author.Email)
			} else {
				fmt.Println(author.Name)
			}
		}
		return nil
	},
}

func init() {
	kpc_cmd.AddCommand(authors)
}
