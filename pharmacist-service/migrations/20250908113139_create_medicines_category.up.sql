CREATE TABLE IF NOT EXISTS medicines_category(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name medicine_category NOT NULL,
    description TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);
