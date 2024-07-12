package fetch

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_parseExpiresAtDate(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name          string
		expiresAtDate string
		expect        time.Time
		assert        assert.ComparisonAssertionFunc
		wantErr       assert.ErrorAssertionFunc
	}{
		{
			name:          "empty string",
			expiresAtDate: "",
			expect:        time.Time{},
			assert:        assert.Equal,
			wantErr:       assert.Error,
		},
		{
			name:          "valid date",
			expiresAtDate: "2022-01-01",
			expect:        time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
			assert:        assert.Equal,
			wantErr:       assert.NoError,
		},
		{
			name:          "invalid date",
			expiresAtDate: "invalid",
			expect:        time.Time{},
			assert:        assert.Equal,
			wantErr:       assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseExpiresAtDate(tt.expiresAtDate)
			tt.assert(t, tt.expect, got)
			tt.wantErr(t, err)
		})
	}
}
