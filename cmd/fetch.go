/*
Copyright © 2024 Shanil Hirani
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/shanilhirani/go-credly/internal/fetch"
)

var fetchCmd = &cobra.Command{
	Use:   "fetch",
	Short: "Fetch user badges from Credly API",
	Long:  `This command fetches user badges from the Credly API using the provided user ID or username.`,
	Run:   fetchRun,
}

func fetchRun(cmd *cobra.Command, args []string) {
	// Check if the required argument (user ID or username) is provided
	if len(args) < 1 {
		cmd.Help()
		os.Exit(1)
	}
	param := args[0]

	// Create an HTTP client
	client := fetch.NewClient(nil)

	// Call the Fetch function
	result, err := client.Fetch(param)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// Print the result
	fmt.Printf("%+v\n", result)
}

func init() {
	rootCmd.AddCommand(fetchCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// versionCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// versionCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	// rootCmd.Flags().StringP("username", "u", "", "Your Credly username/id")

}
