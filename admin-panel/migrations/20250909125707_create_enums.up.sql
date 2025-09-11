CREATE TYPE bed_status AS ENUM (
    'BED_STATUS_AVAILABLE',
    'BED_STATUS_OCCUPIED',
    'BED_STATUS_RESERVED',
    'BED_STATUS_MAINTENANCE'
);

CREATE TYPE bed_type AS ENUM (
    'GENERAL',     
    'PERSONAL',   
    'ICU',         
    'SURGICAL'      
);