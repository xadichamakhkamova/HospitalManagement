CREATE TABLE IF NOT EXISTS medicines(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(255) NOT NULL,
    category medicine_category NOT NULL,
    description TEXT NOT NULL,
    price FLOAT NOT NULL,
    company VARCHAR(24) NOT NULL,
    status medicine_status NOT NULL DEFAULT 'medicine_status_unspecified',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);