-- name: InsertTransaction :exec
INSERT INTO transaction (ta_postdate, ta_description, ta_debit, ta_credit, ta_balance, ta_classification_text) VALUES ($1, $2, $3, $4, $5, $6);