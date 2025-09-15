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
    description,
    company,
    status,
    price,
    created_at,
    updated_at
FROM 
    medicines
WHERE id=$1
    AND deleted_at IS NULL;

-- name: ListMedicines :many
SELECT
    id,
    name,
    category,
    description,
    company,
    status,
    price,
    created_at,
    updated_at,
    COUNT(*) OVER() AS total_count  -- total count filtering boyicha
FROM
    medicines
WHERE deleted_at IS NULL
    AND (
        $1::text=''  -- :text - frontenddan kelgan parametr. '' bolsa, barcha datalarni qaytaradi. Agar bosh bolmasa name, category, company ustunlarida qidirishni boshlaydi.
        OR LOWER(name) LIKE LOWER(CONCAT('%', $1::text, '%')) --LOWER - barcha harflarni kichkina qilib oladi. 
        OR LOWER(category) LIKE LOWER(CONCAT('%', $1::text, '%')) -- CONCAT('%', :search, '%') - example: %aspirin% qilib qidiradi, yani matn ichida qayerdan bolmasin topadi.
        OR LOWER(company) LIKE LOWER(CONCAT('%', $1::text, '%'))
    )
    AND ($2::status IS NULL OR status = $2::status) -- agar stus NULL bolsa barcha yozuvlar olinadi, bolmasa taqqoslaydi.
ORDER BY
    created_at DESC -- yaratilgan sana boyicha teskari tartibda chiqaradi, natijada eng yangilari birinchi chiqadi.
LIMIT $3 -- faqat n ta row olinadi
OFFSET ($4 - 1) * $3; ; --m ta row o‘tkazib, keyingilarini oladi

-- name: UpdateMedicine :one
UPDATE medicines
SET
    name = $2,
    category = $3,
    description = $4,
    price = $5,
    company = $6,
    status = $7,
    updated_at = $8
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
