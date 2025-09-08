CREATE TYPE profession_type AS ENUM (
    'DOCTOR',
    'NURSE',
    'SURGEON'
    'LAB_ASSISTANT',
    'PHARMACIST',
    'RECEPTIONIST'
);

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