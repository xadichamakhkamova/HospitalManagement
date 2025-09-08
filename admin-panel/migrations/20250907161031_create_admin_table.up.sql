CREATE TYPE bed_status AS ENUM (
    'BED_STATUS_AVAILABLE',
    'BED_STATUS_OCCUPIED',
    'BED_STATUS_RESERVED',
    'BED_STATUS_MAINTENANCE'
);

CREATE TYPE bed_type AS ENUM (
    'GENERAL',      -- oddiy palata
    'PERSONAL',     -- shaxsiy xona
    'ICU',          -- reanimatsiya (Intensive Care Unit)
    'SURGICAL'      -- jarrohlik bo‘limi
);

CREATE TABLE IF NOT EXISTS beds(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    bed_number BIGINT NOT NULL,
    bed_type bed_type NOT NULLL,
    description TEXT NOT NULL,
    status bed_status NOT NULL DEFAULT 'BED_STATUS_AVAILABLE',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);
