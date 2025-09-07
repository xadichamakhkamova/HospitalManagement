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
    updated_at,
    deleted_at;

-- name: GetMedicineById :one
SELECT * FROM medicines
WHERE id=$1;

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
    COUNT(*) OVER() AS total_count  -- jami yozuvlar soni filter bo'yicha
FROM
    medicines
WHERE
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
WHERE id = $1
RETURNING
    id, 
    name,
    category,
    description,
    price,
    company,
    status,
    created_at,
    updated_at,
    deleted_at;

-- name: DeleteMedicine :one
UPDATE medicines
SET deleted_at = $2
WHERE id = $1;

