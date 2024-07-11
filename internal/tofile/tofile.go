// Package tofile outputs the results into markdown
package tofile

import (
	"fmt"
	// "log"
	"os"
	"path/filepath"

	"github.com/shanilhirani/go-credly/internal/fetch"
)

// ToFile a function that takes the result of fetch.Filterbadges and writes to file
// named by user. It returns an error if any occurs.
// The file is written in the format of a markdown file with the following structure:
// # Badges
//
// ## Badge Name
// Badge Description
//
// ![Badge Name](Badge Image URL)
//
// Link: [Badge Name](Badge URL)
//
// ---
func ToFile(filename string, badges []fetch.FilteredBadge) (bool, error) {
	// log.Printf("Writing results to %s\n", filename)
	file, err := os.Create(filepath.Clean(fmt.Sprintf("%s.md", filename)))
	if err != nil {
		return false, err
	}
	defer file.Close() //nolint:errcheck // no comment

	_, err = file.WriteString("# Badges\n\n")
	if err != nil {
		return false, err
	}

	for _, badge := range badges {
		badgeInfo := []string{
			fmt.Sprintf("## %s\n", badge.BadgeName),
			fmt.Sprintf("%s\n\n", badge.BadgeDescription),
			fmt.Sprintf("![%s](%s)\n\n", badge.BadgeName, badge.BadgeImageURL),
			fmt.Sprintf("Link: [%s](%s)\n\n", badge.BadgeName, badge.BadgeURL),
			"---\n\n",
		}

		for _, line := range badgeInfo {
			_, err = file.WriteString(line)
			if err != nil {
				return false, err
			}
		}
	}

	// log.Printf("Results written to %s\n", filename)
	return true, nil
}
