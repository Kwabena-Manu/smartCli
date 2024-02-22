/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "smartCli",
	Short: "Interact with the backend with this cli tool",
	Long: `smartCli allows you to interact with the api and execute
	functions that allow you to generate and configure api keys, login, logout, select projects
	simulate and build projects and persform several other operations with the backend`,
	Version:          "0.1",
	PersistentPreRun: authenticateUser,

	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.smartCli.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	versionTemplate := `{{printf "%s: %s - version %s\n" .Name .Short .Version}}`

	rootCmd.SetVersionTemplate((versionTemplate))
}

type CredentialStruct struct {
	AccessKey    string
	AccessSecret string
}

func authenticateUser(cmd *cobra.Command, args []string) {

	//Exempting Login command from this prerun function
	if cmd.Name() == "login" {
		return
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
			fmt.Println("You have to login my boy!")
			os.Exit(1)
		}

		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	file, err := os.ReadFile(authFilepath)

	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			fmt.Println("You have to login my boy!")
			os.Exit(1)
		}
	}

	if len(file) == 0 {
		fmt.Println("You have to login my boy!")
		os.Exit(1)
	}

	credentials := &CredentialStruct{}
	if err = json.Unmarshal(file, credentials); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	accessKeyExists := credentials.AccessKey
	accessSecretExists := credentials.AccessSecret
	if accessKeyExists == "" || accessSecretExists == "" {
		fmt.Println("You have to login my boy!")
		os.Exit(1)
	}
}
