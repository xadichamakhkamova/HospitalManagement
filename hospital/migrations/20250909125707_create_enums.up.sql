CREATE TYPE medicine_category AS ENUM (
    'tablet',
    'syrup',
    'injection',
    'ointment',
    'drops',
    'cream',
    'capsule'
);

create type medicine_status as enum (
    'medicine_status_unspecified',
    'medicine_status_available',
    'medicine_status_out_of_stock',
    'medicine_status_expired',
    'medicine_status_discontinued'
);

create type gender_type as enum (
    'male',
    'female'
);

create type blood_type as enum (
    'a_positive',
    'a_negative',
    'b_positive',
    'b_negative',
    'ab_positive',
    'ab_negative',
    'o_positive',
    'o_negative'
);

create type profession_type as enum (
    'doctor',
    'nurse',
    'surgeon',
    'lab_assistant',
    'pharmacist',
    'receptionist'
);

create type bed_status as enum (
    'bed_status_available',
    'bed_status_occupied',
    'bed_status_reserved',
    'bed_status_maintenance'
);

create type bed_type as enum (
    'general',      -- oddiy palata
    'personal',     -- shaxsiy xona
    'icu',          -- reanimatsiya (intensive care unit)
    'surgical'      -- jarrohlik bo‘limi
);

create type health_condition_type as enum (
    'healthy',
    'minor_illness',
    'chronic_disease',
    'critical_condition',
    'recovering'
);
