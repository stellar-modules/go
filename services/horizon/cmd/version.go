package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	apkg "github.com/stellar-modules/go/sdk/support/app"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "print horizon version",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(apkg.Version())
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
