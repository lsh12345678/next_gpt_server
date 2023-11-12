package cmd

import "github.com/spf13/cobra"

func init() {
	rootCmd.AddCommand()
}

var gatewayCmd = &cobra.Command{
	Use: "gateway",
	Run: GatewayHandle,
}

func GatewayHandle(cmd *cobra.Command, args []string) {

}
