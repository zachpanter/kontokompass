CREATE TABLE IF NOT EXISTS category (
cat_id serial PRIMARY KEY,
cat_name text NOT NULL
);

CREATE TABLE IF NOT EXISTS transaction (
ta_id serial PRIMARY KEY,
ta_postdate timestamp NOT NULL,
ta_description text NOT NULL,
ta_debit float4,
ta_credit float4,
ta_balance float4 NOT NULL,
ta_classification_text text NOT NULL,
cat_id int
);
ALTER TABLE transaction
    ADD CONSTRAINT unique_transaction UNIQUE (ta_postdate, ta_description, ta_debit, ta_credit, ta_balance, ta_classification_text);

ALTER TABLE ONLY transaction
    ADD CONSTRAINT "cat_id_fkey" FOREIGN KEY ("cat_id") REFERENCES category("cat_id");
