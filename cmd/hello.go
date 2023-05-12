/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var (
	name string
	red  = color.BgRed
)

// helloCmd represents the hello command
var helloCmd = &cobra.Command{
	Use:   "hello",
	Short: "Says hello",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if name == "" {
			fmt.Println("hello")
		} else {
			//colorPrint.PrintWarning("hello " + name)
			Writer.PrintWarning("hello" + name)
			Writer.PrintError("in the working copy of 'cmd/decrypt.go', LF will be replaced by CRLF the next time Git touches it")
			Writer.PrintMessage("hello " + name)
		}
	},
}

func init() {
	rootCmd.AddCommand(helloCmd)
	helloCmd.Flags().StringVarP(&name, "name", "n", "", "Your name")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// helloCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// helloCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
