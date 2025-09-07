---------------- Medicine CRUD ----------------

-- name: CreateMedicine :one
INSERT INTO medicines
    (
        name,
        category,
        description,
        price,
        company,
        status
    )
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING
    id, 
    name,
    category,
    description,
    price,
    company,
    status,
    created_at,
    updated_at;

-- name: GetMedicineById :one
SELECT 
    id,
    name,
    category,
    company,
    status,
    price,
    created_at,
    updated_at, 
FROM 
    medicines
WHERE id=$1
    AND deleted_at IS NULL;

-- name: ListMedicines :many
SELECT
    id,
    name,
    category,
    company,
    status,
    price,
    created_at,
    updated_at,
    COUNT(*) OVER() AS total_count  -- total count filtering boyicha
FROM
    medicines
WHERE
    deleted_at IS NULL
    (:search IS NULL OR -- :search - frontenddan kelgan parametr. NULL bolsa, barcha datalarni qaytaradi. Agar bosh bolmasa name, category, company ustunlarida qidirishni boshlaydi.
        LOWER(name) LIKE LOWER(CONCAT('%', :search, '%')) OR --LOWER - barcha harflarni kichkina qilib oladi. 
        LOWER(category) LIKE LOWER(CONCAT('%', :search, '%')) OR -- CONCAT('%', :search, '%') - example: %aspirin% qilib qidiradi, yani matn ichida qayerdan bolmasin topadi.
        LOWER(company) LIKE LOWER(CONCAT('%', :search, '%'))
    )
    AND (:status IS NULL OR status = :status) -- agar stus NULL bolsa barcha yozuvlar olinadi, bolmasa taqqoslaydi.
ORDER BY
    created_at DESC -- yaratilgan sana boyicha teskari tartibda chiqaradi, natijada eng yangilari birinchi chiqadi.
LIMIT :limit -- faqat n ta row olinadi
OFFSET :offset; --m ta row o‘tkazib, keyingilarini oladi

-- name: UpdateMedicine :one
UPDATE medicines
SET
    name = $2,
    category = $3,
    description = $4,
    price = $5,
    company = $6,
    status =$7,
    updated_at=$8
WHERE id = $1
    AND deleted_at IS NULL
RETURNING
    id, 
    name,
    category,
    description,
    price,
    company,
    status,
    created_at,
    updated_at;

-- name: DeleteMedicine :exec
UPDATE medicines
SET deleted_at = $2
WHERE id = $1;

---------------- Medicine Category CRUD ----------------

-- name: CreateMedicineCategory :one
INSERT INTO medicine_categories
    (
        name,
        description
    )
VALUES ($1, $2)
RETURNING
    id,
    name,
    description,
    created_at,
    updated_at;

-- name: GetMedicineCategoryById :one
SELECT 
    name,
    description,
    created_at,
    updated_at,
FROM 
    medicine_categories
WHERE id=$1;

-- name: ListMedicineCategories :many
SELECT
    name,
    description,
    created_at,
    updated_at,
    COUNT(*) OVER() AS total_count
FROM 
    medicine_categories
WHERE 
    deleted_at IS NULL
    (
        :search IS NULL OR
        LOWER(name) LIKE LOWER(CONCAT('%', :search, '%'))
    )
ORDER BY
    created_at DESC
LIMIT :limit 
OFFSET :offset; 

-- name: UpdateMedicineCategory :one 
UPDATE medicine_categories
SET 
    name = $2,
    category = $3,
    description = $4,
    updated_at=$5
WHERE id=$1
    AND deleted_at IS NULL
RETURNING
    id,
    name,
    description,
    created_at,
    updated_at;

-- name DeleteMedicineCategory :exec
UPDATE medicine_categories
SET deleted_at = $2
WHERE id = $1;
