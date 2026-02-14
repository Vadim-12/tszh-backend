-- Align organizations table with current GORM models

-- =========================
-- short_title
-- =========================

-- Заполняем NULL значения, если они есть
UPDATE organizations
SET short_title = full_title
WHERE short_title IS NULL;

-- Убираем UNIQUE-ограничение
ALTER TABLE organizations
    DROP CONSTRAINT IF EXISTS organizations_short_title_key;

-- Делаем NOT NULL
ALTER TABLE organizations
    ALTER COLUMN short_title SET NOT NULL;

-- Обычный индекс для поиска
CREATE INDEX IF NOT EXISTS idx_organizations_short_title
    ON organizations (short_title);

-- =========================
-- email
-- =========================

-- Разрешаем NULL (Email *string)
ALTER TABLE organizations
    ALTER COLUMN email DROP NOT NULL;

-- =========================
-- inn
-- =========================
-- Проверка: 10 или 12 цифр (РФ)
ALTER TABLE organizations
    DROP CONSTRAINT IF EXISTS chk_organizations_inn_len_digits;

ALTER TABLE organizations
    ADD CONSTRAINT chk_organizations_inn_len_digits
    CHECK (inn ~ '^[0-9]{10}([0-9]{2})?$');