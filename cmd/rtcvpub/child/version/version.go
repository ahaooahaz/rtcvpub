package version

import (
	"github.com/ahaooahaz/rtcvpub/version"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "version",
	Short: "show version",
	Long:  "show version",
	Run: func(cmd *cobra.Command, args []string) {
		version.PrintFullVersionInfo()
	},
}
