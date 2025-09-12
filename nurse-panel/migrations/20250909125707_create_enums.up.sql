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

create type health_condition_type as enum (
    'healthy',
    'minor_illness',
    'chronic_disease',
    'critical_condition',
    'recovering'
);
