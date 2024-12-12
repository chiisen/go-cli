package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var name string

var rootCmd = &cobra.Command{
	Use: "hello",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("hello ", name)
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

// 這行代碼的功能是為 rootCmd 命令添加一個名為 name 的字符串標誌（flag）。
// 這個標誌允許用戶在命令行中指定一個名稱，如果用戶沒有提供，則默認值為 "world"。
func init() {
	// 添加一個名為 "name" 的字符串標誌，默認值為 "world"
	rootCmd.Flags().StringVarP(&name, "name", "", "world", "")
}
