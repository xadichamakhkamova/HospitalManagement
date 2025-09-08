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