package cmd

import (
	"fmt"
	"log"
	"path"

	kpc "github.com/afeldman/kpc"
	"github.com/spf13/cobra"
)

var (
	L = &cobra.Command{
		Use:   "L [KPC PACKAGE NAME]",
		Short: "Library installation path",
		Long: `
If you work with different projects, it might be
important to know where the sourcefiles (kl - Files)
for this projects are stored.

AUTHOR:
	Anton Feldmann <anton.feldmann@gmail.com>
`,
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) != 1 {
				log.Fatal("no package in argument")
			}

			kpg := kpc.GetKPC(args[0])
			if kpg != nil {
				fmt.Println(fmt.Sprintf("-I=%s", path.Join(kpg.Prefix, kpg.SrcDir)))
			} else {
				log.Fatalln("The requested Package is not detectable")
			}
		},
	}

	C = &cobra.Command{
		Use:   "C [KPC PACKAGE NAME]",
		Short: "installation path to constant parameter file",
		Long: `
This is just because my Karel Programming flavor. I like to put all information
in seperate files. This program lists you all directories where the files
with the constant parameter declaration are stored.
Because the seperation of all information into different files is my programing
style this tag can left empty, if you are using all information in the sorcefile
(*.kl - File).
If you like to have all information in the dedecated file, then you should put
the path into this field

AUTHOR:
	Anton Feldmann <anton.feldmann@gmail.com>
`,
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) != 1 {
				log.Fatal("no package in argument")
			}

			kpg := kpc.GetKPC(args[0])
			if kpg != nil {
				fmt.Println(fmt.Sprintf("-I=%s", path.Join(kpg.Prefix, kpg.ConstDir)))
			} else {
				log.Fatalln("The requested package can not be found")
			}
		},
	}

	I = &cobra.Command{
		Use:   "I [KPC PACKAGE NAME]",
		Short: "include path to installed karel include files",
		Long: `
The existens of this flag has the same reason, I discused on the "C" parameter.
This tag exists because of my programing style.
But if you use karel header files, please write here the path down, where the
files are installed on yout machine.

AUTHOR:
	Anton Feldmann <anton.feldmann@gmail.com>
`,
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) != 1 {
				log.Fatal("no package in argument")
			}

			kpg := kpc.GetKPC(args[0])
			if kpg != nil {
				fmt.Println(fmt.Sprintf("-I=%s", path.Join(kpg.Prefix, kpg.IncludeDir)))
			} else {
				log.Fatalln("The requested package can not be found")
			}

		},
	}

	T = &cobra.Command{
		Use:   "T [KPC PACKAGE NAME]",
		Short: "installation file path for this package",
		Long: `
I do also seperate the types from the karel source.
So if you do so too, then you should write here the path
to the files. If is tag contains the path to the karel type
files, then we are going to use the files for futher projects.

AUTHOR:
	Anton Feldmann <anton.feldmann@gmail.com>
`,
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) != 1 {
				log.Fatal("no package in argument")
			}

			kpg := kpc.GetKPC(args[0])
			if kpg != nil {
				fmt.Println(fmt.Sprintf("-I=%s", path.Join(kpg.Prefix, kpg.TypeDir)))
			} else {
				log.Fatalln("The requested package can not be found")
			}
		},
	}
)
