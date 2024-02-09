/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Confiure an api key for your account",
	Long: `Configure an api key for your account.
	Requires authentication of your email and password`,
	// Args: cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("config called")
	},
}

func init() {
	apiCmd.AddCommand(configCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// configCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// configCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	configCmd.Flags().StringP("email", "e", "", "Email for authentication before configuration (required)")
	configCmd.MarkFlagRequired("email")
	configCmd.Flags().StringP("password", "p", "", "Password for authentication before configuration(required)")
	configCmd.MarkFlagRequired("password")
}
