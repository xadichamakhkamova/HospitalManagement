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
        birtd_date,
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
    birtd_date,
    blood_group,
    created_at,
    updated_at;

-- name: GetPatientById :one
SELECT 
    full_name, 
    email,
    password,
    address,
    phone_number,
    gender,
    birtd_date,
    blood_group,
    created_at,
    updated_at 
FROM 
    patients
WHERE id=$1 
    AND deleted_at IS NULL; 

-- name: ListPatients :many
SELECT
    full_name, 
    email,
    password,
    address,
    phone_number,
    gender,
    birtd_date,
    blood_group,
    created_at,
    updated_at,
    COUNT(*) OVER() AS total_count
FROM 
    patients
WHERE  
    deleted_at IS NULL
    (:search IS NULL OR
        LOWER(gender) LIKE LOWER(CONCAT('%' :search, '%'))
    )
ORDER BY 
    created_at DESC 
LIMIT :limit
OFFSET :offset;

-- name: UpdatePatient :one 
UPDATE patients
SET 
    full_name=$2, 
    email=$3,
    password=$4,
    address=$5,
    phone_number=$6,
    gender=$7,
    birtd_date=$8,
    blood_group=$9,
    updated_at=$10 
WHERE id=$1
    AND deleted_at IS NULL;  

-- name: DeletePatient :exec
UPDATE patients
SET deleted_at = $2
WHERE id = $1;




