/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login with your email and password",
	Run: func(cmd *cobra.Command, args []string) {
		accessKey, _ := cmd.Flags().GetString("key")
		accessSecret, _ := cmd.Flags().GetString("secret")
		fmt.Println("login called")
		testPrepare(accessKey, accessSecret)
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// loginCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// loginCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	loginCmd.Flags().StringP("key", "k", "", "Access Key ID for authentication (required)")
	loginCmd.MarkFlagRequired("key")
	loginCmd.Flags().StringP("secret", "s", "", "Access Key Secret  for authentication (required)")
	loginCmd.MarkFlagRequired("secret")
}

// For testing purposes. Testing creating a credentials file in the user root
func testPrepare(aKey string, aSecret string) {

	//Do some string validation here
	validateConfigs(aKey, aSecret)

	dummyCredentials := CredentialStruct{
		AccessKey:    aKey,
		AccessSecret: aSecret,
	}

	js, err := json.Marshal(dummyCredentials)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	authFilepath := filepath.Join(homeDir, ".smartCli", ".config.json")
	// fmt.Println(authFilepath)
	_, err = os.Stat(authFilepath)

	if err != nil {
		if os.IsNotExist(err) {
			fmt.Fprintln(os.Stderr, "Configuration file doesn't exist")
			fmt.Fprintln(os.Stderr, "Creating configuration file")
		} else {

			fmt.Fprintln(os.Stderr, err)
		}

	}

	err = os.WriteFile(authFilepath, js, 0644)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		fmt.Fprintln(os.Stderr, "Failure creating configuration file")
		os.Exit(1)
	}

	fmt.Fprintf(os.Stderr, "Created configuration file at %s", authFilepath)

}

func validateConfigs(aKey string, aSecret string) (valKey string, valSecret string) {

	//Does nothing now
	return aKey, aSecret

}
