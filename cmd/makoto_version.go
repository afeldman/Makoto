package cmd

import (
       "fmt"

       "github.com/spf13/cobra"
)


const MAKOTO_VERSION = "0.1.0"

var version = &cobra.Command{
    Use:   "version",
    Short: "print version number of Rossum",
    Long:  `
MAKOTO has a changing versions. This means during the work on MAKOTO,
the version number will change.
I look forward to be down compatible with all MAKOTO_VERSIONS.
But because Makoto works with YAML this statement should held true as long as
the YAML Syntax stays stable. If a big YAML change happen, we are going to see :)

AUTHOR:
	Anton Feldmann <anton.feldmann@gmail.com>
`,
    Run:   func(cmd *cobra.Command, args []string){
    	   fmt.Println("Makotos's Version is: ", MAKOTO_VERSION)
    },
}

func GetVersion() string{
	return MAKOTO_VERSION
}
