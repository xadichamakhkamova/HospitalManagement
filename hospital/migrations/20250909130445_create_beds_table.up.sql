CREATE TABLE IF NOT EXISTS beds(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    bed_number BIGINT NOT NULL,
    bed_type bed_type NOT NULL,
    description TEXT NOT NULL,
    status bed_status NOT NULL DEFAULT 'bed_status_available',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);