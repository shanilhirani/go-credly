// Package tofile outputs the results into markdown
package tofile

import (
	"fmt"
	"io"
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

// BadgeWriter is a struct that writes the badges to a file
type BadgeWriter struct {
	writer io.Writer
}

// NewBadgeWriter creates a new instance of the BadgeWriter struct with the provided io.Writer
func NewBadgeWriter(w io.Writer) *BadgeWriter {
	return &BadgeWriter{writer: w}
}

// WriteBadges writes the badges to the io.Writer
func (bw *BadgeWriter) WriteBadges(badges []fetch.FilteredBadge) error {
	content := generateBadgeContent(badges)
	_, err := fmt.Fprintf(bw.writer, defaultContent, content)
	return err
}

// UpdateContent updates the content between the markers in the io.Writer with the provided badges
func (bw *BadgeWriter) UpdateContent(r io.Reader, badges []fetch.FilteredBadge) error {
	existingContent, err := io.ReadAll(r)
	if err != nil {
		return err
	}

	updatedContent := replaceContentBetweenMarkers(string(existingContent), badges)
	_, err = bw.writer.Write([]byte(updatedContent))
	return err
}

func generateBadgeContent(badges []fetch.FilteredBadge) string {
	var content strings.Builder
	for _, badge := range badges {
		fmt.Fprintf(&content, "\n## %s\n", badge.BadgeName)
		fmt.Fprintf(&content, "%s\n", badge.BadgeDescription)
		fmt.Fprintf(&content, "![%s](%s)\n", badge.BadgeName, badge.BadgeImageURL)
		fmt.Fprintf(&content, "Link: [%s](%s)\n\n", badge.BadgeName, badge.BadgeURL)
		content.WriteString("\n")
	}
	return content.String()
}

func replaceContentBetweenMarkers(existingContent string, badges []fetch.FilteredBadge) string {
	startIndex := strings.Index(existingContent, startMarker)
	endIndex := strings.LastIndex(existingContent, endMarker)

	if startIndex == -1 || endIndex == -1 {
		return fmt.Sprintf(defaultContent, generateBadgeContent(badges))
	}

	prefix := existingContent[:startIndex+len(startMarker)]
	suffix := existingContent[endIndex:]
	replacementContent := generateBadgeContent(badges)

	return prefix + replacementContent + suffix
}

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
	filePath := filepath.Clean(fmt.Sprintf("%s.md", filename))

	file, err := createOrOpenFile(filePath)
	if err != nil {
		return false, err
	}
	defer file.Close()

	bw := NewBadgeWriter(file)
	err = bw.UpdateContent(file, badges)
	if err != nil {
		return false, err
	}

	return true, nil
}

func createOrOpenFile(filePath string) (*os.File, error) {
	_, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		return os.Create(filePath)
	} else if err != nil {
		return nil, err
	}

	return os.OpenFile(filePath, os.O_RDWR|os.O_TRUNC, 0o600)
}
