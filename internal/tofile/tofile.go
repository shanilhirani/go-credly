// Package tofile outputs the results into markdown
package tofile

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/shanilhirani/go-credly/internal/fetch"
)

const (
	defaultContent = "<!--START_SECTION:go-credly-->%s<!--END_SECTION:go-credly-->"
	startMarker    = "<!--START_SECTION:go-credly-->"
	endMarker      = "<!--END_SECTION:go-credly-->"
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
	file, err := createOrOpenFile(filename)
	if err != nil {
		return false, err
	}
	defer file.Close() //nolint: errcheck //no comment

	err = replaceContent(file, badges)
	if err != nil {
		return false, err
	}

	return true, nil
}

func createOrOpenFile(filename string) (*os.File, error) {
	filePath := filepath.Clean(fmt.Sprintf("%s.md", filename))

	// Check if the file exists
	_, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		// Create a new file if it doesn't exist
		return os.Create(filePath)
	} else if err != nil {
		return nil, err
	}

	// Open the file for reading and writing
	return os.OpenFile(filePath, os.O_RDWR, 0o600)
}

func replaceContent(file *os.File, badges []fetch.FilteredBadge) error {
	// Read the existing file content
	existingContent, err := readFileContent(file)
	if err != nil {
		return err
	}

	// Replace the content between the markers
	updatedContent := replaceContentBetweenMarkers(existingContent, badges)

	// Truncate the file and write the updated content
	err = truncateAndWriteContent(file, updatedContent)
	if err != nil {
		return err
	}

	return nil
}

func readFileContent(file *os.File) (string, error) {
	scanner := bufio.NewScanner(file)
	var content strings.Builder
	for scanner.Scan() {
		content.WriteString(scanner.Text() + "\n")
	}
	if err := scanner.Err(); err != nil {
		return "", err
	}
	return content.String(), nil
}

func replaceContentBetweenMarkers(existingContent string, badges []fetch.FilteredBadge) string {
	startIndex := strings.Index(existingContent, startMarker)
	endIndex := strings.LastIndex(existingContent, endMarker)

	if startIndex == -1 || endIndex == -1 {
		// If markers are not found, create the default content with markers
		return fmt.Sprintf(defaultContent, generateBadgeContent(badges))
	}

	prefix := existingContent[:startIndex+len(startMarker)]
	suffix := existingContent[endIndex:]
	replacementContent := generateBadgeContent(badges)

	return prefix + replacementContent + suffix
}

func generateBadgeContent(badges []fetch.FilteredBadge) string {
	var content strings.Builder
	for _, badge := range badges {
		badgeInfo := []string{
			fmt.Sprintf("\n## %s\n", badge.BadgeName),
			fmt.Sprintf("%s\n", badge.BadgeDescription),
			fmt.Sprintf("![%s](%s)\n", badge.BadgeName, badge.BadgeImageURL),
			fmt.Sprintf("Link: [%s](%s)\n\n", badge.BadgeName, badge.BadgeURL),
		}

		for _, line := range badgeInfo {
			content.WriteString(line)
		}
		content.WriteString("\n")
	}
	return content.String()
}

func truncateAndWriteContent(file *os.File, content string) error {
	// Truncate the file to reset its content
	err := file.Truncate(0)
	if err != nil {
		return err
	}

	// Seek to the beginning of the file
	_, err = file.Seek(0, 0)
	if err != nil {
		return err
	}

	// Write the updated content to the file
	_, err = file.WriteString(content)
	if err != nil {
		return err
	}

	return nil
}
