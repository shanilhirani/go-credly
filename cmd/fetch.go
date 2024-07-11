// Package cmd fetch
/*
Copyright © 2024 Shanil Hirani
*/
package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"

	"github.com/shanilhirani/go-credly/internal/fetch"
	"github.com/shanilhirani/go-credly/internal/tofile"
)

var (
	includeExpired bool
	outFile        string
)

var fetchCmd = &cobra.Command{
	Use:   "fetch",
	Short: "Fetch user badges from Credly API",
	Long:  `This command fetches user badges from the Credly API using the provided user ID or username.`,
	PreRun: func(cmd *cobra.Command, _ []string) {
		if cmd.Flags().Changed("include-expired") {
			log.Println("Including expired badges")
		}
		if cmd.Flags().Changed("out-file") {
			if outFile == "" {
				log.Println("WARNING: empty string provided. defaulting to BADGES")
				err := cmd.Flags().Set("out-file", "BADGES")
				if err != nil {
					log.Fatalf("ERROR: Failed to set default output file name: %v", err)
				}
			}
		}
	},
	Run: fetchRun,
}

func fetchRun(cmd *cobra.Command, args []string) {
	var param string

	switch {
	case len(args) > 0:
		param = args[0]
	default:
		log.Fatalf("Error: Please enter your credly user ID or username.")
		return
	}

	// Create an HTTP client
	client := fetch.NewClient(nil)

	// Call the Fetch function
	credlyData, err := client.Fetch(param)
	if err != nil {
		log.Fatalf("Error: Call to Fetch encountered an error: %v", err)
	}

	// Filter the data
	filteredBadges, err := fetch.FilterData(param, credlyData, includeExpired)
	if err != nil {
		log.Fatalf("Error: Failed to filter data: %v", err)
	}

	// Print the filtered badges
	if !cmd.Flags().Changed("out-file") {
		log.Printf("Displaying Credly Badges to stdout")
		for _, badge := range filteredBadges {
			fmt.Printf("Badge Name: %s\n", badge.BadgeName)
			fmt.Printf("Badge Description: %s\n", badge.BadgeDescription)
			fmt.Printf("Badge Image URL: %s\n", badge.BadgeImageURL)
			fmt.Printf("Badge URL: %s\n\n", badge.BadgeURL)
		}
	}

	// Write to file
	if cmd.Flags().Changed("out-file") {
		log.Printf("Writing Credly Badges to %s.md", outFile)
		writeToFile, err := tofile.ToFile(outFile, filteredBadges)
		if err != nil {
			log.Printf("Error: Failed to write to file: %v", err)
		} else if writeToFile {
			log.Printf("Credly Badges written to %s.md", outFile)
		}
	}
}

func init() { //nolint:gochecknoinits // required by cobra
	rootCmd.AddCommand(fetchCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// versionCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	fetchCmd.Flags().BoolVarP(&includeExpired, "include-expired", "e", false, "Include expired badges")
	fetchCmd.Flags().StringVarP(&outFile, "out-file", "o", outFile, "Write results to a file with Markdown extension. e.g BADGES.md")
}
