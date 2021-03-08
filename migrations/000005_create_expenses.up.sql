CREATE TABLE IF NOT EXISTS expenses (
    id          UUID PRIMARY KEY DEFAULT uuid_generate_v4(),

    bank_account_id  UUID NOT NULL,

    name        VARCHAR(255) NOT NULL,
    amount      NUMERIC(16, 2),

    updated_at  TIMESTAMP NOT NULL DEFAULT NOW(),
    created_at  TIMESTAMP NOT NULL DEFAULT NOW(),

    CONSTRAINT fk_bank_account
      FOREIGN KEY(bank_account_id)
          REFERENCES bank_accounts(id)
          ON DELETE CASCADE
);

CREATE TRIGGER expenses_updated_at
    BEFORE UPDATE ON expenses
    FOR EACH ROW
EXECUTE PROCEDURE on_update_timestamp();
