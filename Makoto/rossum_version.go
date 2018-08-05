package cmd

import (
       "fmt"
       
       "github.com/spf13/cobra"
)


const ROSSUM_VERSION = "0.1.0"

var version = &cobra.Command{
    Use:   "version",
    Short: "print version number of Rossum",
    Long:  `
ROSSUM has a changing versions. This means during the work on ROSSUM,
the version number will change.
I look forward to be down compatible with all ROSSUM_VERSIONS.
But because ROSSUM works with YAML this statement should held true as long as
the YAML Syntax stays stable. If a big YAML change happen, we are going to see :)

AUTHOR:
	Anton Feldmann <anton.feldmann@gmail.com>
`,
    Run:   func(cmd *cobra.Command, args []string){
    	   fmt.Println("Rossum's Version is: ", ROSSUM_VERSION)
    },
}

func GetVersion() string{
     return ROSSUM_VERSION
}