DROP TRIGGER IF EXISTS tg_before_update_at ON users;

DROP FUNCTION IF EXISTS fc_update_at;

DROP INDEX IF EXISTS idx_user_name;
DROP INDEX IF EXISTS idx_user_status;
DROP INDEX IF EXISTS idx_user_level;
DROP INDEX IF EXISTS idx_user_status_create_at;

DROP TABLE IF EXISTS users;