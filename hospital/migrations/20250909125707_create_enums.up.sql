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

CREATE TYPE gender_type AS ENUM (
    'MALE',
    'FEMALE'
);

CREATE TYPE blood_type AS ENUM (
    'A_POSITIVE',
    'A_NEGATIVE',
    'B_POSITIVE',
    'B_NEGATIVE',
    'AB_POSITIVE',
    'AB_NEGATIVE',
    'O_POSITIVE',
    'O_NEGATIVE'
);

CREATE TYPE health_condition AS ENUM (
    'HEALTHY',
    'MINOR_ILLNESS',
    'CHRONIC_DISEASE',
    'CRITICAL',
    'RECOVERING'
);

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

CREATE TYPE health_condition_type AS ENUM (
    'HEALTHY',
    'MINOR_ILLNESS',
    'CHRONIC_DISEASE',
    'CRITICAL_CONDITION',
    'RECOVERING'
);
