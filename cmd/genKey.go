/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// genKeyCmd represents the genKey command
var genKeyCmd = &cobra.Command{
	Use:   "genKey",
	Short: "Generates a random encryption key.",
	Long:  `This command generates a random encryption key. It takes the key length as an argument and generates a new encryption key of the specified length. The generated key can be used for encryption and decryption operations.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("genKey called")
	},
}

func init() {
	rootCmd.AddCommand(genKeyCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// genKeyCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// genKeyCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
