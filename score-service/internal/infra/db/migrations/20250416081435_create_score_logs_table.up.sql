CREATE TABLE score_logs (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    game VARCHAR(255) NOT NULL,
    round INT NOT NULL,
    winner_id BIGINT NOT NULL, -- 外鍵，指向 products.id
    loser_id BIGINT NOT NULL,  -- 外鍵，指向 products.id
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);

-- trigger function that updates the updated_at field whenever the row is updated
CREATE OR REPLACE FUNCTION update_updated_at_column()
RETURNS TRIGGER AS $$
BEGIN
   NEW.updated_at = CURRENT_TIMESTAMP;
   RETURN NEW;
END;
$$ language 'plpgsql';

-- trigger on score_logs table for each row update
CREATE TRIGGER update_score_logs_updated_at
BEFORE UPDATE ON score_logs
FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();