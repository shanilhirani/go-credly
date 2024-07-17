// Package types for credly data
package types

import "time"

// CredlyData data structure to be used for processing
type CredlyData struct {
	Data []struct {
		ID                string      `json:"id"`
		ExpiresAtDate     string      `json:"expires_at_date"`
		IssuedAtDate      string      `json:"issued_at_date"`
		IssuedTo          string      `json:"issued_to"`
		Locale            string      `json:"locale"`
		Public            bool        `json:"public"`
		State             string      `json:"state"`
		TranslateMetadata bool        `json:"translate_metadata"`
		AcceptedAt        time.Time   `json:"accepted_at"`
		ExpiresAt         time.Time   `json:"expires_at"`
		IssuedAt          time.Time   `json:"issued_at"`
		LastUpdatedAt     time.Time   `json:"last_updated_at"`
		UpdatedAt         time.Time   `json:"updated_at"`
		EarnerPath        string      `json:"earner_path"`
		EarnerPhotoURL    interface{} `json:"earner_photo_url"`
		IsPrivateBadge    bool        `json:"is_private_badge"`
		UserIsEarner      bool        `json:"user_is_earner"`
		Issuer            struct {
			Summary  string `json:"summary"`
			Entities []struct {
				Label   string `json:"label"`
				Primary bool   `json:"primary"`
				Entity  struct {
					Type                           string `json:"type"`
					ID                             string `json:"id"`
					Name                           string `json:"name"`
					URL                            string `json:"url"`
					VanityURL                      string `json:"vanity_url"`
					InternationalizeBadgeTemplates bool   `json:"internationalize_badge_templates"`
					ShareToZiprecruiter            bool   `json:"share_to_ziprecruiter"`
					Verified                       bool   `json:"verified"`
				} `json:"entity"`
			} `json:"entities"`
		} `json:"issuer"`
		BadgeTemplate struct {
			ID                              string  `json:"id"`
			Description                     string  `json:"description"`
			GlobalActivityURL               string  `json:"global_activity_url"`
			EarnThisBadgeURL                *string `json:"earn_this_badge_url"`
			EnableEarnThisBadge             bool    `json:"enable_earn_this_badge"`
			EnableDetailAttributeVisibility bool    `json:"enable_detail_attribute_visibility"`
			Name                            string  `json:"name"`
			Public                          bool    `json:"public"`
			RecipientType                   string  `json:"recipient_type"`
			VanitySlug                      string  `json:"vanity_slug"`
			ShowBadgeLmi                    bool    `json:"show_badge_lmi"`
			ShowSkillTagLinks               bool    `json:"show_skill_tag_links"`
			Translatable                    bool    `json:"translatable"`
			Level                           string  `json:"level"`
			TimeToEarn                      string  `json:"time_to_earn"`
			Cost                            string  `json:"cost"`
			TypeCategory                    string  `json:"type_category"`
			Image                           struct {
				ID  string `json:"id"`
				URL string `json:"url"`
			} `json:"image"`
			ImageURL              string `json:"image_url"`
			URL                   string `json:"url"`
			OwnerVanitySlug       string `json:"owner_vanity_slug"`
			BadgeTemplateEarnable bool   `json:"badge_template_earnable"`
			Recommendable         bool   `json:"recommendable"`
			Issuer                struct {
				Summary  string `json:"summary"`
				Entities []struct {
					Label   string `json:"label"`
					Primary bool   `json:"primary"`
					Entity  struct {
						Type                           string `json:"type"`
						ID                             string `json:"id"`
						Name                           string `json:"name"`
						URL                            string `json:"url"`
						VanityURL                      string `json:"vanity_url"`
						InternationalizeBadgeTemplates bool   `json:"internationalize_badge_templates"`
						ShareToZiprecruiter            bool   `json:"share_to_ziprecruiter"`
						Verified                       bool   `json:"verified"`
					} `json:"entity"`
				} `json:"entities"`
			} `json:"issuer"`
			RelatedBadgeTemplates []struct {
				ID    string `json:"id"`
				Name  string `json:"name"`
				Image struct {
					ID  string `json:"id"`
					URL string `json:"url"`
				} `json:"image"`
				ImageURL string `json:"image_url"`
				URL      string `json:"url"`
			} `json:"related_badge_templates"`
			Alignments              []interface{} `json:"alignments"`
			BadgeTemplateActivities []struct {
				ID                      string      `json:"id"`
				ActivityType            string      `json:"activity_type"`
				RequiredBadgeTemplateID interface{} `json:"required_badge_template_id"`
				Title                   string      `json:"title"`
				URL                     string      `json:"url"`
			} `json:"badge_template_activities"`
			Endorsements []interface{} `json:"endorsements"`
			Skills       []struct {
				ID         string `json:"id"`
				Name       string `json:"name"`
				VanitySlug string `json:"vanity_slug"`
			} `json:"skills"`
			Language string `json:"language,omitempty"`
		} `json:"badge_template"`
		Image struct {
			ID  string `json:"id"`
			URL string `json:"url"`
		} `json:"image"`
		ImageURL string `json:"image_url"`
		Evidence []struct {
			ID          string `json:"id"`
			Type        string `json:"type"`
			Title       string `json:"title,omitempty"`
			Description string `json:"description,omitempty"`
			Name        string `json:"name,omitempty"`
			Values      []struct {
				Key   string      `json:"key"`
				Value string      `json:"value"`
				URL   interface{} `json:"url"`
			} `json:"values,omitempty"`
		} `json:"evidence"`
		Recommendations []interface{} `json:"recommendations"`
		Language        string        `json:"language,omitempty"`
	} `json:"data"`
	Metadata struct {
		Count           int         `json:"count"`
		CurrentPage     int         `json:"current_page"`
		TotalCount      int         `json:"total_count"`
		TotalPages      int         `json:"total_pages"`
		Per             int         `json:"per"`
		PreviousPageURL interface{} `json:"previous_page_url"`
		NextPageURL     interface{} `json:"next_page_url"`
	} `json:"metadata"`
}
