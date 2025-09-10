package cmd

import (
	"fmt"

	"github.com/afeldman/Makoto/makoto"
	"github.com/spf13/cobra"
)

var reject = &cobra.Command{
	Use:   "reject [PACKAGE NAME] [VERSION]",
	Short: "Remove a KPC package from the database",
	Long: `Remove a KPC package entry from the local database.

Examples:
  makoto reject hello_world 0.1.0

This will permanently remove the entry from the kpc.db file.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 2 {
			return cmd.Usage()
		}

		name := args[0]
		version := args[1]

		if err := makoto.Reject(name, version); err != nil {
			return fmt.Errorf("failed to remove %s@%s: %w", name, version, err)
		}

		fmt.Printf("✅ Removed %s@%s from database\n", name, version)
		return nil
	},
}

func init() {
	kpc_cmd.AddCommand(reject)
}
