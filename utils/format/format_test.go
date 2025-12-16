package format

import "testing"

func TestFormatCBRDate(t *testing.T) {
	// Given
	tests := []struct {
		name string
		date string
		want string
	}{
		{
			name: "valid RFC3339",
			date: "2025-12-14T11:30:00+03:00",
			want: "14 December 2025 11:30",
		},
		{
			name: "invalid date",
			date: "invalid-date",
			want: "invalid-date",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// When
			got := FormatCBRDate(tt.date)

			// Then
			if got != tt.want {
				t.Fatalf("want %q, got %q", tt.want, got)
			}
		})
	}
}
