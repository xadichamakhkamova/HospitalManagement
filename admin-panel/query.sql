---------------- Department CRUD ----------------

-- name: CreateDepartment :one
INSERT INTO departments 
    (
        name,
        number,
        description
    )
VALUES($1, $2, $3)
RETURNING 
    id,
    name,
    number,
    description,
    created_at,
    updated_at; 

-- name: GetDepartmentById :one 
SELECT 
    id,
    name,
    number,
    description,
    created_at,
    updated_at 
FROM 
    departments 
WHERE id = $1 
    AND deleted_at IS NULL;  

-- name: ListDepartments :many 
SELECT 
    id,
    name,
    number,
    description,
    created_at,
    updated_at,
    COUNT(*) OVER() AS total_count 
FROM 
    departments 
WHERE deleted_at IS NULL 
    AND (   
        :search IS NULL
        OR LOWER(name) LIKE LOWER(CONCAT('%', :search, '%'))
    ) 
ORDER BY 
    created_at DESC 
LIMIT :limit
OFFSET (:page - 1) * :limit; 

-- name: UpdateDepartment :one 
UPDATE departments 
SET 
    name = $1,
    number = $2, 
    description = $3 
WHERE id = $1 
    AND deleted_at IS NULL;

-- name: DeleteDepartment :exec
UPDATE departments
SET deleted_at = $2
WHERE id = $1;

---------------- Personal and Doctor CRUD ----------------

-- name: CreatePersonal :one
INSERT INTO personals 
    (
        profession,
        full_name,
        email,
        password,
        address,
        phone_number
    ) 
VALUES($1, $2, $3, $4, $5, $6)
RETURNING 
    id,
    profession,
    full_name,
    email,
    password,
    address,
    phone_number,
    created_at,
    updated_at; 

-- name: CreateDoctor :one 
INSERT INTO doctors 
    (
        profession,
        full_name,
        email,
        password,
        address,
        phone_number,
        department_number,
    ) 
VALUES($1, $2, $3, $4, $5, $6, $7)
RETURNING 
    id,
    profession,
    full_name,
    email,
    password,
    address,
    phone_number,
    department_number,
    created_at,
    updated_at; 

-- name: GetPersonalById :one 
SELECT 
    id,
    profession,
    full_name,
    email,
    password,
    address,
    phone_number,
    created_at,
    updated_at 
FROM 
    personals 
WHERE id = $1 
    AND deleted_at IS NULL; 

-- name: GetDoctorById :one 
SELECT 
    id,
    profession,
    full_name,
    email,
    password,
    address,
    phone_number,
    department_number,
    created_at,
    updated_at 
FROM 
    doctors
WHERE id = $1 
    AND deleted_at IS NULL;  

-- name: ListPersonals :many
SELECT 
    id,
    profession,
    full_name,
    email,
    password,
    address,
    phone_number,
    created_at,
    updated_at ,
    COUNT(*) OVER() AS total_count 
FROM 
    personals 
WHERE deleted_at IS NULL 
    AND (
        :search IS NULL 
        OR LOWER(profession) LIKE LOWER(CONCAT('%', :search, '%'))
        OR LOWER(full_name) LIKE LOWER(CONCAT('%', :search, '%'))
    ) 
ORDER BY 
    created_at DESC 
LIMIT :limit
OFFSET (:page - 1) * :limit; 

-- name: ListDoctors :many
SELECT 
    id,
    profession,
    full_name,
    email,
    password,
    address,
    phone_number,
    department_number,
    created_at,
    updated_at ,
    COUNT(*) OVER() AS total_count 
FROM 
    doctors
WHERE deleted_at IS NULL 
    AND (
        :search IS NULL 
        OR LOWER(profession) LIKE LOWER(CONCAT('%', :search, '%'))
        OR LOWER(full_name) LIKE LOWER(CONCAT('%', :search, '%'))
    ) 
ORDER BY 
    created_at DESC 
LIMIT :limit
OFFSET (:page - 1) * :limit; 

-- name: UpdatePersonal :one
UPDATE personals
SET 
    profession = $2,
    full_name = $3,
    email = $4,
    password = $5,
    address = %6,
    phone_number = %7,
    updated_at = %8
WHERE id = $1 
    AND deleted_at IS NULL;

-- name: UpdateDoctor :one
UPDATE doctors
SET 
    profession = $2,
    full_name = $3,
    email = $4,
    password = $5,
    address = %6,
    phone_number = %7,
    department_number = $8,
    updated_at = %9 
WHERE id = $1 
    AND deleted_at IS NULL;

-- name: DeletePersonal :exec 
UPDATE personals 
SET deleted_at = $2 
WHERE id = $1;

-- name: DeleteDoctors :exec 
UPDATE doctors
SET deleted_at = $2 
WHERE id = $1;

---------------- BED CRUD ----------------

-- name: CreateBed :one
INSERT INTO beds 
    (
       bed_number,
       bed_type,
       description 
    ) 
VALUES($1, $2, $3)
RETURNING 
    id,
    bed_number,
    description,
    status,
    created_at,
    updated_at; 

-- name: GetBedByID :one
SELECT 
    id,
    bed_number,
    bed_type,
    description,
    status,
    created_at,
    updated_at
FROM 
    beds
WHERE id = $1
  AND deleted_at IS NULL;

-- name: ListBeds :many
SELECT 
    id,
    bed_number,
    bed_type,
    description,
    status,
    created_at,
    updated_at,
    COUNT(*) OVER() AS total_count
FROM 
    beds
WHERE deleted_at IS NULL 
    AND (
        :search IS NULL
        OR LOWER(bed_number::text) LIKE LOWER(CONCAT('%', :search, '%'))
        OR LOWER(bed_type) LIKE LOWER(CONCAT('%', :search, '%'))
    )
  AND (:status IS NULL OR status = :status)
ORDER BY 
    created_at DESC
LIMIT :limit
OFFSET (:page - 1) * :limit;

-- name: UpdateBed :one
UPDATE beds
SET 
    bed_number = $2,
    bed_type = $3,
    description = $4,
    status = $5,
    updated_at =$6
WHERE id = $1
  AND deleted_at IS NULL
RETURNING 
    id,
    bed_number,
    bed_type,
    description,
    status,
    created_at,
    updated_at;

-- name: DeleteBed :exec
UPDATE beds
SET deleted_at = $2
WHERE id = $1;
