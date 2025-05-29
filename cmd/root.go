package cmd

import (
	"fmt"
	"os"

	"github.com/ljgago/shortlink-go/pkg/log"
	"github.com/spf13/cobra"
)

// serverCmd represents the root
var rootCmd = &cobra.Command{
	Use:   "shortlink",
	Short: "Shortlink is a short link generator",
	Long:  `A simple short link generator self-hosted`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}

func Execute() {
	log.Setup(false, "debug", []string{})

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
