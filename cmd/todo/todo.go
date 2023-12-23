/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package todo

import (
	"fmt"

	"github.com/spf13/cobra"
)

// todo/todoCmd represents the todo/todo command
var TodoCmd = &cobra.Command{
	Use:   "todo/todo",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("todo/todo called")
	},
}

func init() {

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// todo/todoCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// todo/todoCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
