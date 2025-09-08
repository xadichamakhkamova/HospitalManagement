CREATE TABLE IF NOT EXISTS bed_management (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    bed_id UUID NOT NULL REFERENCES beds(id),
    patient_id UUID NOT NULL REFERENCES patients(id),
    status bed_status NOT NULL, 
    assigned_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

