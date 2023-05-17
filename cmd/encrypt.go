/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// encryptCmd represents the encrypt command
var encryptCmd = &cobra.Command{
	Use:   "encrypt",
	Short: "Encrypts a file using specified algorithm and key.",
	Long:  `This command allows you to encrypt a file using a specified encryption algorithm and key. It takes the input file path, encryption algorithm, encryption key, and output file path as arguments. The command encrypts the input file and saves the encrypted version to the output file.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("encrypt called")
	},
}

func init() {
	rootCmd.AddCommand(encryptCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// encryptCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// encryptCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
