---------------- Patient CRUD ----------------

-- name: CreatePatient :one
INSERT INTO patients
    (
        full_name, 
        email,
        password,
        address,
        phone_number,
        gender,
        birth_date,
        blood_group
    )
VALUES($1, $2, $3, $4, $5, $6, $7, $8)
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
    created_at,
    updated_at;

-- name: GetPatientById :one
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
    created_at,
    updated_at 
FROM 
    patients
WHERE id = $1 
    AND deleted_at IS NULL; 

-- name: ListPatients :many
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
    created_at,
    updated_at,
    COUNT(*) OVER() AS total_count
FROM 
    patients
WHERE deleted_at IS NULL
    AND (
        $1::text=''
        OR LOWER(full_name) LIKE LOWER(CONCAT('%', $1::text, '%')) 
        OR LOWER(email) LIKE LOWER(CONCAT('%', $1::text, '%')) 
        OR LOWER(gender) LIKE LOWER(CONCAT('%', $1::text, '%'))
    )
ORDER BY 
    created_at DESC 
LIMIT $2
OFFSET ($3 - 1) * $2; 

-- name: UpdatePatient :one 
UPDATE patients
SET 
    full_name = $2, 
    email = $3,
    password = $4,
    address = $5,
    phone_number = $6,
    gender = $7,
    birth_date = $8,
    blood_group = $9,
    updated_at = $10 
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
    created_at,
    updated_at;  

-- name: DeletePatient :exec
UPDATE patients
SET deleted_at = $2
WHERE id = $1;




