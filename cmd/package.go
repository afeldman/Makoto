package cmd

import "github.com/spf13/cobra"

var kpc_cmd = &cobra.Command{
    	   Use:   "package",
    	   Short: "command works on packages",
    	   Long:  `
To get the subprograms working with the kcp file I used this the package command.
There is a logic seperation between the Makoto base and the KPC package information

AUTHOR:
	Anton Feldmann <anton.feldmann@gmail.com>
`,
}

func init(){
     kpc_cmd.AddCommand(authors)
     kpc_cmd.AddCommand(discription)
     kpc_cmd.AddCommand(pversion)
     kpc_cmd.AddCommand(url)
     kpc_cmd.AddCommand(source)
     kpc_cmd.AddCommand(requirements)
     kpc_cmd.AddCommand(conflicts)
     kpc_cmd.AddCommand(contains)
     kpc_cmd.AddCommand(prefix)

     kpc_cmd.AddCommand(L)
     kpc_cmd.AddCommand(C)
     kpc_cmd.AddCommand(I)
     kpc_cmd.AddCommand(T)

     kpc_cmd.AddCommand(c)
     kpc_cmd.AddCommand(i)
     kpc_cmd.AddCommand(l)
     kpc_cmd.AddCommand(t)
}
