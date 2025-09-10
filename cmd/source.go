package cmd

import (
	"fmt"

	"github.com/afeldman/Makoto/makoto"
	"github.com/spf13/cobra"
)

var source_cmd = &cobra.Command{
	Use:   "source [KPC PACKAGE NAME] [VERSION]",
	Short: "Show the source repository information of a KPC package",
	Long: `The package homepage can differ from the actual source repository URL.
This field shows where to find the source code of the requested package.

Makoto does not enforce any particular source management tool.
You can set the URL freely and use your preferred SCM (e.g. git) to download and build.

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

		repo := kpc.KPC.Repository
		fmt.Printf("URL:    %s\n", repo.URL)
		if repo.Tag != "" {
			fmt.Printf("Tag:    %s\n", repo.Tag)
		}
		if repo.Rev != "" {
			fmt.Printf("Rev:    %s\n", repo.Rev)
		}
		if repo.Branch != "" {
			fmt.Printf("Branch: %s\n", repo.Branch)
		}
		return nil
	},
}

func init() {
	kpc_cmd.AddCommand(source_cmd)
}
