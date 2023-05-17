/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// decryptCmd represents the decrypt command
var decryptCmd = &cobra.Command{
	Use:   "decrypt",
	Short: "Decrypts an encrypted file using algorithm and key.",
	Long:  `This command allows you to decrypt an encrypted file using the corresponding decryption algorithm and key. It takes the encrypted file path, decryption algorithm, decryption key, and output file path as arguments. The command decrypts the input file and saves the decrypted version to the output file.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("decrypt called")
	},
}

func init() {
	rootCmd.AddCommand(decryptCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// decryptCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// decryptCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
