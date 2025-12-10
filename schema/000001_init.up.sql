-- 000001_init.up.sql

-- UUID generator
CREATE EXTENSION IF NOT EXISTS "pgcrypto";

-- =========================
-- users
-- =========================
CREATE TABLE IF NOT EXISTS users
(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

    full_name      TEXT NOT NULL,
    email          TEXT UNIQUE,
    phone_number   TEXT NOT NULL UNIQUE,
    password_hash  TEXT NOT NULL,

    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

-- =========================
-- organizations
-- =========================
CREATE TABLE IF NOT EXISTS organizations
(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

    full_title     TEXT NOT NULL,
    short_title    TEXT UNIQUE,
    inn            TEXT NOT NULL UNIQUE,
    email          TEXT NOT NULL,
    legal_address  TEXT NOT NULL,

    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

-- =========================
-- organizations_staff (join users <-> organizations with role)
-- =========================
CREATE TABLE IF NOT EXISTS organizations_staff
(
    organization_id UUID NOT NULL,
    user_id         UUID NOT NULL,
    role            VARCHAR(32) NOT NULL,

    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),

    PRIMARY KEY (organization_id, user_id),

    CONSTRAINT fk_orgstaff_org
        FOREIGN KEY (organization_id)
        REFERENCES organizations(id)
        ON DELETE CASCADE,

    CONSTRAINT fk_orgstaff_user
        FOREIGN KEY (user_id)
        REFERENCES users(id)
        ON DELETE CASCADE
);

CREATE INDEX IF NOT EXISTS idx_orgstaff_org_id ON organizations_staff(organization_id);
CREATE INDEX IF NOT EXISTS idx_orgstaff_user_id ON organizations_staff(user_id);
-- индекс по role добавляй только если реально будешь часто фильтровать/группировать по роли

-- =========================
-- buildings
-- =========================
CREATE TABLE IF NOT EXISTS buildings
(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

    address          TEXT NOT NULL,
    floors           SMALLINT NOT NULL,
    cadastral_number TEXT NOT NULL UNIQUE,
    year_built       SMALLINT NOT NULL,
    building_type    VARCHAR(32) NOT NULL,
    entrances        SMALLINT NOT NULL,
    apartments       SMALLINT NULL,

    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

-- =========================
-- building_units
-- =========================
CREATE TABLE IF NOT EXISTS building_units
(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

    number       TEXT NOT NULL,
    unit_type    VARCHAR(32) NOT NULL,
    total_area   DOUBLE PRECISION NOT NULL,
    living_area  NUMERIC NULL,

    building_id UUID NOT NULL,

    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),

    CONSTRAINT fk_building_units_building
        FOREIGN KEY (building_id)
        REFERENCES buildings(id)
        ON DELETE CASCADE
);

CREATE INDEX IF NOT EXISTS idx_building_units_building_id ON building_units(building_id);
CREATE INDEX IF NOT EXISTS idx_building_units_unit_type ON building_units(unit_type);

-- =========================
-- refresh_tokens
-- =========================
CREATE TABLE IF NOT EXISTS refresh_tokens
(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

    user_id UUID NOT NULL,
    expires_at TIMESTAMPTZ NOT NULL,
    revoked BOOLEAN NOT NULL DEFAULT FALSE,

    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),

    CONSTRAINT fk_refresh_tokens_user
        FOREIGN KEY (user_id)
        REFERENCES users(id)
        ON DELETE CASCADE
);

CREATE INDEX IF NOT EXISTS idx_refresh_tokens_user_id ON refresh_tokens(user_id);
CREATE INDEX IF NOT EXISTS idx_refresh_tokens_expires_at ON refresh_tokens(expires_at);
CREATE INDEX IF NOT EXISTS idx_refresh_tokens_revoked ON refresh_tokens(revoked);

-- =========================
-- users_and_building_units (join users <-> building_units with role)
-- =========================
CREATE TABLE IF NOT EXISTS users_and_building_units
(
    user_id UUID NOT NULL,
    building_unit_id UUID NOT NULL,
    role VARCHAR(32) NOT NULL,

    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),

    PRIMARY KEY (user_id, building_unit_id),

    CONSTRAINT fk_uabu_user
        FOREIGN KEY (user_id)
        REFERENCES users(id)
        ON DELETE CASCADE,

    CONSTRAINT fk_uabu_building_unit
        FOREIGN KEY (building_unit_id)
        REFERENCES building_units(id)
        ON DELETE CASCADE
);

CREATE INDEX IF NOT EXISTS idx_uabu_user_id ON users_and_building_units(user_id);
CREATE INDEX IF NOT EXISTS idx_uabu_building_unit_id ON users_and_building_units(building_unit_id);
-- role индекс — по желанию, как и выше