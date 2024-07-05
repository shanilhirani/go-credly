/*
Copyright © 2024 Shanil Hirani
*/
package cmd

import (
	"fmt"
	"log"

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
		log.Fatalf("Error: Credly Username/ID not supplied")
	}
	param := args[0]

	// Create an HTTP client
	client := fetch.NewClient(nil)

	// Call the Fetch function
	credlyData, err := client.Fetch(param)
	if err != nil {
		log.Fatalf("Error: Call to Fetch encountered an error: %v", err)
	}

	// Filter the data
	filteredBadges, err := fetch.FilterData(param, credlyData)
	if err != nil {
		log.Fatalf("Error: Failed to filter data: %v", err)
	}

	// Print the filtered badges
	for _, badge := range filteredBadges {
		fmt.Printf("Badge Name: %s\n", badge.Name)
		fmt.Printf("Badge Description: %s\n", badge.Description)
		fmt.Printf("Badge Image URL: %s\n", badge.ImageURL)
		fmt.Printf("Badge URL: %s\n\n", badge.URL)
	}
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
