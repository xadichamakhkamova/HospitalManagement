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
        $1::search  IS NULL 
        OR LOWER(full_name) LIKE LOWER(CONCAT('%', $1::search , '%')) 
        OR LOWER(email) LIKE LOWER(CONCAT('%', $1::search , '%')) 
        OR LOWER(gender) LIKE LOWER(CONCAT('%', $1::search , '%')) 
        OR LOWER(blood_group) LIKE LOWER(CONCAT('%', $1::search , '%'))
    )
    AND (
        $2::only_eligible = FALSE 
        OR (weight >= 50 AND health_condition = 'HEALTHY') -- if equal to DEFAULT
    )
ORDER BY 
    created_at DESC 
LIMIT $3
OFFSET ($4 - 1) * $3; 

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
    AND deleted_at IS NULL
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

