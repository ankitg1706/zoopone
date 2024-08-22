package store

import (
	"testing"

	"github.com/ankitg1706/zoopone/model"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestSetLimitAndPage(t *testing.T) {
	// Set up the PostgreSQL database connection for testing
	dsn := "host=localhost user=ankit password=password dbname=manage port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to connect database: %v", err)
	}

	tests := []struct {
		name   string
		filter map[string]interface{}
		want   int
	}{
		{
			name: "Default Limit and Page",
			filter: map[string]interface{}{
				model.DataPerPage: "",
				model.PageNumber:        "",
			},
			want: 10, // Default limit is 10
		},
		{
			name: "Custom Limit and Page",
			filter: map[string]interface{}{
				model.DataPerPage: "10",
				model.PageNumber:        "1",
			},
			want: 5, // Custom limit is 5
		},
		{
			name: "Page parse error",
			filter: map[string]interface{}{
				"dataPerPage": "5",
				"page":        "some value",
			},
			want: 5, // Custom limit is 5
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			query := db.Model(&struct{}{})
			setLimitAndPage(tt.filter, query)

			assert.NotNil(t, query, "Expected query object to be set")
		})
	}
}

func TestSetDateRangeFilter(t *testing.T) {
	// Set up the PostgreSQL database connection for testing
	dsn := "host=localhost user=ankit password=password dbname=manage port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to connect database: %v", err)
	}

	tests := []struct {
		name   string
		filter map[string]interface{}
		want   string
	}{
		{
			name: "Valid Date Range",
			filter: map[string]interface{}{
				model.StartDate: "2006-01-02 15:04:05.000 -0700",
				model.EndDate:   "2006-01-02 15:04:05.000 -0700",
			},
			want: "2023-01-01", // Expected start date
		},
		{
			name: "Invalid Start Date",
			filter: map[string]interface{}{
				model.StartDate: "invalid-date",
				model.EndDate:   "2023-12-31",
			},
			want: "", // Expected empty due to parsing error
		},
		{
			name: "No Date Range",
			filter: map[string]interface{}{
				model.StartDate: "",
				model.EndDate:   "",
			},
			want: "", // Expected empty due to no dates
		},
		{
			name: "Invalid time parse startdate Date Range",
			filter: map[string]interface{}{
				model.StartDate: "2023-01-01010101",
				model.EndDate:   "2006-01-02 15:04:05.000 -0700",
			},
			want: "", // Expected start date
		},
		{
			name: "Invalid time parse enddate Date Range",
			filter: map[string]interface{}{
				model.StartDate: "2006-01-02 15:04:05.000 -0700",
				model.EndDate:   "2023-12-30101011",
			},
			want: "", // Expected start date
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			query := db.Model(&struct{}{})
			setDateRangeFilter(tt.filter, query)

			assert.NotNil(t, query, "Expected query object to be set")
		})
	}
}
