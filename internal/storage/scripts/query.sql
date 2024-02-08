-- name: InsertTransaction :exec
INSERT INTO transaction (postdate, description, debit, credit, balance, classification_text) VALUES ($1, $2, $3, $4, $5, $6);