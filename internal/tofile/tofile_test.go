package tofile

import (
	"bytes"
	"strings"
	"testing"

	"github.com/shanilhirani/go-credly/internal/fetch"
)

func TestBadgeWriter_WriteBadges(t *testing.T) {
	badges := []fetch.FilteredBadge{
		{
			BadgeName:        "Test Badge",
			BadgeDescription: "A test badge",
			BadgeImageURL:    "https://example.com/badge.png",
			BadgeURL:         "https://example.com/badge",
		},
	}

	buf := &bytes.Buffer{}
	bw := NewBadgeWriter(buf)

	err := bw.WriteBadges(badges)
	if err != nil {
		t.Fatalf("WriteBadges failed: %v", err)
	}

	expected := `<!--START_SECTION:go-credly-->
## Test Badge
A test badge
![Test Badge](https://example.com/badge.png)
Link: [Test Badge](https://example.com/badge)


<!--END_SECTION:go-credly-->`

	if buf.String() != expected {
		t.Errorf("WriteBadges output doesn't match expected.\nGot:\n%s\nWant:\n%s", buf.String(), expected)
	}
}

func TestBadgeWriter_UpdateContent(t *testing.T) {
	existingContent := `# My Badges

<!--START_SECTION:go-credly-->
Old content
<!--END_SECTION:go-credly-->

Footer`

	badges := []fetch.FilteredBadge{
		{
			BadgeName:        "New Badge",
			BadgeDescription: "A new badge",
			BadgeImageURL:    "https://example.com/new-badge.png",
			BadgeURL:         "https://example.com/new-badge",
		},
	}

	reader := strings.NewReader(existingContent)
	buf := &bytes.Buffer{}
	bw := NewBadgeWriter(buf)

	err := bw.UpdateContent(reader, badges)
	if err != nil {
		t.Fatalf("UpdateContent failed: %v", err)
	}

	expected := `# My Badges

<!--START_SECTION:go-credly-->
## New Badge
A new badge
![New Badge](https://example.com/new-badge.png)
Link: [New Badge](https://example.com/new-badge)


<!--END_SECTION:go-credly-->

Footer`

	if buf.String() != expected {
		t.Errorf("UpdateContent output doesn't match expected.\nGot:\n%s\nWant:\n%s", buf.String(), expected)
	}
}

func TestGenerateBadgeContent(t *testing.T) {
	badges := []fetch.FilteredBadge{
		{
			BadgeName:        "Badge 1",
			BadgeDescription: "Description 1",
			BadgeImageURL:    "https://example.com/badge1.png",
			BadgeURL:         "https://example.com/badge1",
		},
		{
			BadgeName:        "Badge 2",
			BadgeDescription: "Description 2",
			BadgeImageURL:    "https://example.com/badge2.png",
			BadgeURL:         "https://example.com/badge2",
		},
	}

	content := generateBadgeContent(badges)

	expected := `
## Badge 1
Description 1
![Badge 1](https://example.com/badge1.png)
Link: [Badge 1](https://example.com/badge1)



## Badge 2
Description 2
![Badge 2](https://example.com/badge2.png)
Link: [Badge 2](https://example.com/badge2)


`

	if content != expected {
		t.Errorf("generateBadgeContent output doesn't match expected.\nGot:\n%s\nWant:\n%s", content, expected)
	}
}

func TestReplaceContentBetweenMarkers(t *testing.T) {
	testCases := []struct {
		name            string
		existingContent string
		badges          []fetch.FilteredBadge
		expected        string
	}{
		{
			name:            "No existing markers",
			existingContent: "Some content without markers",
			badges: []fetch.FilteredBadge{
				{BadgeName: "Test Badge", BadgeDescription: "Test Description", BadgeImageURL: "https://example.com/test.png", BadgeURL: "https://example.com/test"},
			},
			expected: "<!--START_SECTION:go-credly-->\n## Test Badge\nTest Description\n![Test Badge](https://example.com/test.png)\nLink: [Test Badge](https://example.com/test)\n\n\n<!--END_SECTION:go-credly-->",
		},
		{
			name: "Existing markers",
			existingContent: `# My Badges

<!--START_SECTION:go-credly-->
Old content
<!--END_SECTION:go-credly-->

Footer`,
			badges: []fetch.FilteredBadge{
				{BadgeName: "New Badge", BadgeDescription: "New Description", BadgeImageURL: "https://example.com/new.png", BadgeURL: "https://example.com/new"},
			},
			expected: `# My Badges

<!--START_SECTION:go-credly-->
## New Badge
New Description
![New Badge](https://example.com/new.png)
Link: [New Badge](https://example.com/new)


<!--END_SECTION:go-credly-->

Footer`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := replaceContentBetweenMarkers(tc.existingContent, tc.badges)
			if result != tc.expected {
				t.Errorf("replaceContentBetweenMarkers output doesn't match expected.\nGot:\n%s\nWant:\n%s", result, tc.expected)
			}
		})
	}
}
