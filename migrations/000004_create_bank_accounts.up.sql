CREATE TABLE IF NOT EXISTS bank_accounts (
    id          UUID PRIMARY KEY DEFAULT uuid_generate_v4(),

    account_id  UUID NOT NULL,

    name        VARCHAR(255) NOT NULL,

    updated_at  TIMESTAMP NOT NULL DEFAULT NOW(),
    created_at  TIMESTAMP NOT NULL DEFAULT NOW(),

    CONSTRAINT fk_account
        FOREIGN KEY(account_id)
            REFERENCES accounts(id)
            ON DELETE CASCADE
);

CREATE TRIGGER bank_accounts_updated_at
    BEFORE UPDATE ON bank_accounts
    FOR EACH ROW
EXECUTE PROCEDURE on_update_timestamp();
