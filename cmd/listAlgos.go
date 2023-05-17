/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// listAlgosCmd represents the listAlgos command
var listAlgosCmd = &cobra.Command{
	Use:   "listAlgos",
	Short: "Lists available encryption algorithms.",
	Long:  `This command lists the available encryption algorithms supported by the LockFile project. It provides a list of encryption algorithms that can be used with the encrypt and decrypt commands.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("listAlgos called")
	},
}

func init() {
	rootCmd.AddCommand(listAlgosCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listAlgosCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listAlgosCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
