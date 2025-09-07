CREATE TYPE medicine_category AS ENUM (
    'tablet',
    'syrup',
    'injection',
    'ointment',
    'drops',
    'cream',
    'capsule'
);

CREATE TYPE medicine_status AS ENUM (
    'MEDICINE_STATUS_UNSPECIFIED',
    'MEDICINE_STATUS_AVAILABLE',
    'MEDICINE_STATUS_OUT_OF_STOCK',
    'MEDICINE_STATUS_EXPIRED',
    'MEDICINE_STATUS_DISCONTINUED'
);

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

CREATE TABLE IF NOT EXISTS medicine_categories(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name medicine_category NOT NULL,
    description TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);