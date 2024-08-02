/*
 * @Date: 2024-07-16 19:09:53
 * @LastEditTime: 2024-07-29 16:49:35
 * @FilePath: \cobra\cmd\version.go
 * @description: 注释
 */
/*
 * @Author:
 * @Date: 2024-07-20 11:12:36
 * @LastEditors:
 * @LastEditTime: 2024-07-20 11:27:01
 * @Description:
 */
/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		//cmd当前指向自身
		author, err := cmd.Flags().GetString("author")
		if err != nil {
			fmt.Println("请输入作者信息")
		}
		fmt.Println("作者信息", author)
		fmt.Printf("当前接收的参数：%v\n", args)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// versionCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// versionCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	//                         详细命令   简短命令   默认值    注释
	versionCmd.Flags().StringP("author", "a", "Ellan-copy", "作者name")
}
