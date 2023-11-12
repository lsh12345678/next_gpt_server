package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var (
	ConfigPath string
)

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&ConfigPath, "config", "./plato.yaml", "config file (default is ./plato.yaml)")
}

var rootCmd = &cobra.Command{
	Use:   "next_gpt",
	Short: "这是一个简单易用的GPT小程序",
	Run:   GPT,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func GPT(cmd *cobra.Command, args []string) {

}

func initConfig() {

}
