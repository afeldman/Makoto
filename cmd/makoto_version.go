package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// Makoto CLI version (semantic versioning)
const MakotoVersion = "0.3.3"

var version = &cobra.Command{
	Use:   "version",
	Short: "Print the Makoto version",
	Long: `Show the Makoto CLI version.

	Notes:
	- Makoto follows semantic versioning (MAJOR.MINOR.PATCH).
	- KPC files are written in TOML (not YAML).
	- Backwards compatibility of the CLI and the KPC TOML schema may change
	  across MAJOR versions.

	Author:
	  Anton Feldmann <anton.feldmann@gmail.com>
`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Makoto version:", MakotoVersion)
	},
}

func GetVersion() string {
	return MakotoVersion
}

func init() {
	Makoto.AddCommand(version)
}
