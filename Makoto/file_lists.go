package cmd

import (
	"encoding/json"
	"fmt"
	"log"

	kpc "github.com/afeldman/Makoto"
	"github.com/spf13/cobra"
)

var (
	c = &cobra.Command{
		Use:   "c [KPC PACKAGE NAME]",
		Short: "config file list",
		Long: `
If you use different config information to build your KAREL program,
then you need a list of all files containing the constant information.
this program gives you a list of all constant fails belonging to the required
package.

AUTHOR:
	Anton Feldmann <anton.feldmann@gmail.com>
`,
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) != 1 {
				log.Fatal("no package in argument")
			}

			kpg := kpc.GetKPC(args[0])
			if kpg != nil {
				buff, err := json.Marshal(&kpg.Consts)
				if err != nil {
					log.Fatalln(err)
				}
				fmt.Println(string(buff))
			} else {
				log.Fatalln("Requested Package not found")
			}
		},
	}

	l = &cobra.Command{
		Use:   "l [KPC PACKAGE NAME]",
		Short: "list of all source files",
		Long: `
Different to C does Karel not build libraries.
So you are going to have access to the source files.

To get the list of all source files belonging to the
requestd package, the "l" comand is used.

AUTHOR:
	Anton Feldmann <anton.feldmann@gmail.com>
`,
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) != 1 {
				log.Fatal("no package in argument")
			}

			kpg := kpc.GetKPC(args[0])
			if kpg != nil {
				buff, err := json.Marshal(&kpg.Libs)
				if err != nil {
					log.Fatalln(err)
				}
				fmt.Println(string(buff))
			} else {
				log.Fatalln("Requested Package not found")
			}
		},
	}

	t = &cobra.Command{
		Use:   "t [KPC PACKAGE NAME]",
		Short: "list all type files",
		Long: `
To have a list of all type files belonging to the requested
package this command can be used. For your next project it
is important to use the same types.
This command list all the type files the requested project installed
on your computer

AUTHOR:
	Anton Feldmann <anton.feldmann@gmail.com>
`,
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) != 1 {
				log.Fatal("no package in argument")
			}

			kpg := kpc.GetKPC(args[0])
			if kpg != nil {
				buff, err := json.Marshal(&kpg.Types)
				if err != nil {
					log.Fatalln(err)
				}
				fmt.Println(string(buff))
			} else {
				log.Fatalln("Requested Package not found")
			}
		},
	}

	i = &cobra.Command{
		Use:   "i [KPC PACKAGE NAME]",
		Short: "list all karel header files ",
		Long: `
If you want to use the list of all headerfile the requestet package
contains, then this flag will list all karel header files that belongs
to the requested package

AUTHOR:
	Anton Feldmann <anton.feldmann@gmail.com>
`,
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) != 1 {
				log.Fatal("no package in argument")
			}

			kpg := kpc.GetKPC(args[0])
			if kpg != nil {
				buff, err := json.Marshal(&kpg.Includes)
				if err != nil {
					log.Fatalln(err)
				}
				fmt.Println(string(buff))
			} else {
				log.Fatalln("Requested Package not found")
			}
		},
	}
)
