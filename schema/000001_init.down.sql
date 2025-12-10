-- 000001_init.down.sql

DROP TABLE IF EXISTS users_and_building_units;
DROP TABLE IF EXISTS refresh_tokens;
DROP TABLE IF EXISTS building_units;
DROP TABLE IF EXISTS buildings;
DROP TABLE IF EXISTS organizations_staff;
DROP TABLE IF EXISTS organizations;
DROP TABLE IF EXISTS users;

-- EXTENSION pgcrypto можно не удалять, обычно её оставляют
-- DROP EXTENSION IF EXISTS "pgcrypto";