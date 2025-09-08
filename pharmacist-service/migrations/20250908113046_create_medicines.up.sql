CREATE TABLE IF NOT EXISTS medicines(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    category medicine_category NOT NULL,
    description TEXT NOT NULL,
    price NUMERIC(10,2) NOT NULL,
    company VARCHAR(24) NOT NULL,
    status medicine_status NOT NULL DEFAULT 'MEDICINE_STATUS_UNSPECIFIED',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);