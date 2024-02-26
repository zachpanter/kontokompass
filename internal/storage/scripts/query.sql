-- name: TransactionInsert :exec
INSERT INTO transaction (ta_postdate, ta_description, ta_debit, ta_credit, ta_balance, ta_classification_text)
VALUES ($1, $2, $3, $4, $5, $6)
ON CONFLICT ON CONSTRAINT unique_transaction DO NOTHING;

-- name: TransactionSelect :one
SELECT ta_postdate, ta_description, ta_debit, ta_credit, ta_balance, ta_classification_text
FROM transaction
WHERE ta_id = $1;