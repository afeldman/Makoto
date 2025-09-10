package cmd

import "github.com/spf13/cobra"

var kpc_cmd = &cobra.Command{
	Use:   "package",
	Short: "command works on packages",
	Long: `
To get the subprograms working with the kcp file I used this the package command.
There is a logic seperation between the Makoto base and the KPC package information

AUTHOR:
	Anton Feldmann <anton.feldmann@gmail.com>
`,
}

func init() {
	Makoto.AddCommand(kpc_cmd)
}
