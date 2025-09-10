package cmd

import (
	"fmt"

	"github.com/afeldman/Makoto/makoto"
	"github.com/spf13/cobra"
)

var prefix = &cobra.Command{
	Use:   "prefix [KPC PACKAGE NAME] [VERSION]",
	Short: "Show the install prefix of a KPC package",
	Long: `The install prefix defines the base directory where the package parts
are installed. Paths for {type, const, header, source, form, dict} files
are relative to this prefix.

This separation makes usage in build systems (e.g. Makefiles) more convenient.

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

		if kpc.KPC.Prefix == "" {
			fmt.Println("no prefix set")
			return nil
		}

		fmt.Println(kpc.KPC.Prefix)
		return nil
	},
}

func init() {
	kpc_cmd.AddCommand(prefix)
}
