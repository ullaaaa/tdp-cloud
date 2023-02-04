package worker

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cli = &cobra.Command{
	Use:   "worker",
	Short: "注册子节点",
	Run: func(cmd *cobra.Command, args []string) {
		Execute()
	},
}

func WithCli() *cobra.Command {

	cli.Flags().StringP("remote", "r", "", "注册地址 (e.g. ws://{domain}/wsi/{appid}/worker)")

	viper.BindPFlag("worker.remote", cli.Flags().Lookup("remote"))

	return cli

}
