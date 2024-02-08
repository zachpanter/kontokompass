package models

import "time"

type csvRow struct {
	PostDate         time.Time `json:"post_date"`
	Description      string    `json:"description"`
	Debit            float64   `json:"debit"`
	Credit           float64   `json:"credit"`
	Balance          float64   `json:"balance"`
	ClassificationID int       `json:"classification_id"`
}
