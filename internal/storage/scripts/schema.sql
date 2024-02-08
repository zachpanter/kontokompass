CREATE TABLE IF NOT EXISTS transaction (
                                    id serial PRIMARY KEY,
                                    postdate timestamp NOT NULL,
                                    description text NOT NULL,
                                    debit float4,
                                    credit float4,
                                    balance float4,
                                    classification_text text NOT NULL,
                                    classification_id int
);
ALTER TABLE ONLY transaction ADD CONSTRAINT "classification_id_fkey" FOREIGN KEY ("classification_id") REFERENCES classification("id");


CREATE TABLE IF NOT EXISTS classification (
                                        id serial PRIMARY KEY,
                                        name text NOT NULL
);