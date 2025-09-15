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
        $1::search IS NULL
        OR LOWER(name) LIKE LOWER(CONCAT('%', $1::search, '%'))
    ) 
ORDER BY 
    created_at DESC 
LIMIT $2
OFFSET ($3 - 1) * $2; 

-- name: UpdateDepartment :one 
UPDATE departments 
SET 
    name = $2,
    number = $3, 
    description = $4,
    updated_at = $5
WHERE id = $1 
    AND deleted_at IS NULL
RETURNING
    id,
    name,
    number,
    description,
    created_at,
    updated_at; 

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
    personal_id,
    department_number
    ) 
VALUES ($1, $2)
RETURNING 
    id,
    personal_id,
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
    d.id,
    d.personal_id,
    p.profession,
    p.full_name,
    p.email,
    p.password,
    p.address,
    p.phone_number,
    d.department_number,
    d.created_at,
    d.updated_at
FROM doctors d
JOIN personals p ON d.personal_id = p.id
WHERE d.id = $1
  AND d.deleted_at IS NULL
  AND p.deleted_at IS NULL;

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
        $1::search IS NULL 
        OR LOWER(profession) LIKE LOWER(CONCAT('%', $1::search, '%'))
        OR LOWER(full_name) LIKE LOWER(CONCAT('%', $1::search, '%'))
    ) 
ORDER BY 
    created_at DESC 
LIMIT $2
OFFSET ($3 - 1) * $2; 

-- name: ListDoctors :many
SELECT 
    d.id,
    d.personal_id,
    p.profession,
    p.full_name,
    p.email,
    p.password,
    p.address,
    p.phone_number,
    d.department_number,
    d.created_at,
    d.updated_at,
    COUNT(*) OVER() AS total_count
FROM doctors d
JOIN personals p ON d.personal_id = p.id
WHERE d.deleted_at IS NULL AND p.deleted_at IS NULL
  AND (
      $1::search IS NULL 
      OR LOWER(p.profession) LIKE LOWER(CONCAT('%', $1::search, '%'))
      OR LOWER(p.full_name) LIKE LOWER(CONCAT('%', $1::search, '%'))
  )
ORDER BY d.created_at DESC
LIMIT $2
OFFSET ($3 - 1) * $2;


-- name: UpdatePersonal :one
UPDATE personals
SET 
    profession = $2,
    full_name = $3,
    email = $4,
    password = $5,
    address = $6,
    phone_number = $7,
    updated_at = $8
WHERE id = $1 
    AND deleted_at IS NULL
RETURNING
    id,
    profession,
    full_name,
    email,
    password,
    address,
    phone_number,
    created_at,
    updated_at; ;

-- name: UpdateDoctor :one
UPDATE doctors d
SET 
    department_number = $2,
    updated_at = $3
WHERE d.id = $1
  AND d.deleted_at IS NULL
RETURNING 
    d.id,
    d.personal_id,
    d.department_number,
    d.created_at,
    d.updated_at;

-- name: DeletePersonal :exec 
UPDATE personals 
SET deleted_at = $2 
WHERE id = $1;

-- name: DeleteDoctor :exec
UPDATE doctors
SET deleted_at = $2
WHERE id = $1;

