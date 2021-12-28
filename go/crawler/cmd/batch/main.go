package main

import (
	"context"
	"os"

	initSeller "github.com/panforyou/seller-finding/go/crawler/cmd/batch/init-seller"

	"github.com/spf13/cobra"
)

func main() {
	ctx := context.Background()

	rootCmd := &cobra.Command{
		Use:   "batch",
		Short: "バッチ処理用コマンド",
		Long:  `バッチ処理をローカルから走らせる場合に使用してください`,
	}

	rootCmd.AddCommand(initSeller.Cmd(ctx))
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
