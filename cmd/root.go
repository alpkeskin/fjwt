// fjwt v1.0.0
// Author: Alp Keskin
// Github: github.com/alpkeskin
// Linkedin: linkedin.com/in/alpkeskin
package cmd

import (
	"log"
	"os"

	"github.com/alpkeskin/fjwt/cmd/cracker"
	"github.com/alpkeskin/fjwt/cmd/utils.go"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:                   "fjwt [JWT] -w [WORDLIST]",
	DisableFlagsInUseLine: true,
	Short:                 "An another JWT cracker but really fast!",
	Example:               "fjwt ey... -w wordlist.txt -t 10",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			log.Println("Missing JWT token!")
			cmd.Help()
			os.Exit(1)
		}
		if *utils.Wordlist == "" {
			log.Println("Missing wordlist!")
			cmd.Help()
			os.Exit(1)
		}
		cracker.Handler(args[0], *utils.Wordlist, *utils.Threads)
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	utils.Wordlist = rootCmd.Flags().StringP("wordlist", "w", "", "Wordlist path")
	utils.Threads = rootCmd.Flags().IntP("threads", "t", 10, "Number of threads")
	rootCmd.Version = utils.Version
}
