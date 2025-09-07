---------------- Donor CRUD ----------------

-- name: CreateDonor :one
INSERT INTO donors
    (
        full_name, 
        email,
        password,
        address,
        phone_number,
        gender,
        birth_date,
        blood_group,
        weight,
        health_condition
    )
VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
RETURNING
    id,
    full_name, 
    email,
    password,
    address,
    phone_number,
    gender,
    birth_date,
    blood_group,
    weight,
    health_condition,
    created_at,
    updated_at;

-- name: GetDonorById :one 
SELECT 
    id,
    full_name, 
    email,
    password,
    address,
    phone_number,
    gender,
    birth_date,
    blood_group,
    last_donation,
    donation_count,
    (weight >= 50 AND health_condition = 'HEALTHY') AS is_eligible,
    last_checkup_date,
    weight,
    health_condition,
    donation_location,
    created_at,
    updated_at 
FROM 
    donors
WHERE id=$1 
    AND deleted_at IS NULL; 

-- name: ListDonors :many 
SELECT 
    id,
    full_name, 
    email,
    password,
    address,
    phone_number,
    gender,
    birth_date,
    blood_group,
    last_donation,
    donation_count,
    (weight >= 50 AND health_condition = 'HEALTHY') AS is_eligible,
    last_checkup_date,
    weight,
    health_condition,
    donation_location,
    created_at,
    updated_at,
    COUNT(*) OVER() AS total_count
FROM 
    donors 
WHERE 
    deleted_at IS NULL
    AND (
        :search IS NULL 
        OR LOWER(full_name) LIKE LOWER(CONCAT('%', :search, '%')) 
        OR LOWER(email) LIKE LOWER(CONCAT('%', :search, '%')) 
        OR LOWER(gender) LIKE LOWER(CONCAT('%', :search, '%')) 
        OR LOWER(blood_group) LIKE LOWER(CONCAT('%', :search, '%'))
    )
    AND (
        :only_eligible = FALSE 
        OR (weight >= 50 AND health_condition = 'HEALTHY') -- if equal to DEFAULT
    )
ORDER BY 
    created_at DESC 
LIMIT :limit
OFFSET (:page - 1) * :limit; 

-- name: UpdateDonor :one 
UPDATE donors
SET 
    full_name = $2, 
    email = $3,
    password = $4,
    address = $5,
    phone_number = $6,
    gender = $7,
    birth_date = $8,
    blood_group = $9,
    weight = $10,
    health_condition = $11,
    updated_at = $12
WHERE id = $1
    AND deleted_at IS NULL;  

-- name: DeleteDonor :exec
UPDATE donors
SET deleted_at = $2
WHERE id = $1;


-- name: RegisterDonation :one
UPDATE donors
SET 
    last_donation = NOW(),
    donation_count = donation_count + 1,
    donation_location = $2,
    updated_at = NOW()
WHERE id = $1
    AND deleted_at IS NULL
RETURNING 
    last_donation,
    donation_count,
    (weight >= 50 AND health_condition = 'HEALTHY') AS is_eligible;

-- name: RegisterCheckup :one
UPDATE donors
SET 
    last_checkup_date = NOW(),
    updated_at = NOW()
WHERE id = $1
    AND deleted_at IS NULL
RETURNING 
    last_checkup_date,
    (weight >= 50 AND health_condition = 'HEALTHY') AS is_eligible;


---------------- Bed Management----------------

-- name: AssignPatientToBed :one
INSERT INTO bed_management 
    (
        bed_id, 
        patient_id, 
        status
    )
VALUES ($1, $2, 'BED_OCCUPIED')
RETURNING 
    id, 
    bed_id, 
    patient_id, 
    status, 
    assigned_at, 
    updated_at;

-- name: ReleaseBed :one
UPDATE bed_management
SET 
    status = 'BED_AVAILABLE', 
    updated_at = NOW()
WHERE bed_id = $1 
    AND status = 'BED_OCCUPIED'
RETURNING 
    id, 
    bed_id, 
    patient_id, 
    status, 
    assigned_at, 
    updated_at;

-- name: ReserveBed :one
INSERT INTO bed_management 
    (
        bed_id, 
        patient_id, 
        status
    )
VALUES ($1, $2, 'BED_RESERVED')
RETURNING 
    id, 
    bed_id, 
    patient_id, 
    status, 
    assigned_at, 
    updated_at;

-- name: GetBedById :one
SELECT 
    id,
    bed_id,
    patient_id,
    status,
    assigned_at,
    updated_at
FROM bed_management
WHERE bed_id = $1
ORDER BY 
    updated_at DESC
LIMIT 1;

-- name: ListBeds :many
SELECT 
    id,
    bed_id,
    patient_id,
    status,
    assigned_at,
    updated_at,
    COUNT(*) OVER() AS total_count
FROM 
    bed_management
WHERE 
    (  
        :search IS NULL
        OR LOWER(status) LIKE LOWER(CONCAT('%', :search, '%'))
    ) 
ORDER BY 
    updated_at DESC 
LIMIT :limit
OFFSET (:page - 1) * :limit; 

