-- Rollback organizations alignment

-- Удаляем индекс
DROP INDEX IF EXISTS idx_organizations_short_title;

-- Возвращаем NOT NULL для email
ALTER TABLE organizations
    ALTER COLUMN email SET NOT NULL;

-- Убираем CHECK для ИНН
ALTER TABLE organizations
    DROP CONSTRAINT IF EXISTS chk_organizations_inn_len_digits;

-- Возвращаем nullable short_title
ALTER TABLE organizations
    ALTER COLUMN short_title DROP NOT NULL;

-- Возвращаем UNIQUE
ALTER TABLE organizations
    ADD CONSTRAINT organizations_short_title_key UNIQUE (short_title);