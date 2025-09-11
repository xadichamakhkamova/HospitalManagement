---------------- Appointment CRUD ----------------

-- name: CreateAppointment :one 
INSERT INTO appointments (
    doctor_id,
    patient_id,
    appointment_date
)
VALUES ($1, $2, $3)
RETURNING
    id,
    doctor_id,
    patient_id,
    appointment_date,
    created_at,
    updated_at;

-- name: GetAppointmentById :one
SELECT 
    id,
    doctor_id,
    patient_id,
    appointment_date,
    created_at,
    updated_at
FROM appointments
WHERE id=$1 
  AND deleted_at IS NULL;

-- name: ListAppointments :many
SELECT 
    id,
    doctor_id,
    patient_id,
    appointment_date,
    created_at,
    updated_at,
    COUNT(*) OVER() AS total_count
FROM appointments
WHERE deleted_at IS NULL 
  AND ($1::date IS NULL OR DATE(appointment_date) = DATE($1::date))
ORDER BY appointment_date DESC
LIMIT $2 
OFFSET ($3 - 1) * $2;


-- name: UpdateAppointment :one
UPDATE appointments
SET
    doctor_id = $2,
    patient_id = $3,
    appointment_date = $4,
    updated_at = $5
WHERE id = $1
  AND deleted_at IS NULL
RETURNING
    id,
    doctor_id,
    patient_id,
    appointment_date,
    created_at,
    updated_at;

-- name: DeleteAppointment :exec
UPDATE appointments
SET deleted_at = $2
WHERE id = $1;

---------------- Prescription CRUD ----------------

-- name: CreatePrescription :one
INSERT INTO prescriptions
    (
        doctor_id,
        patient_id,
        case_history,
        medication,
        description
    )
VALUES($1, $2, $3, $4, $5)
RETURNING
    id,
    doctor_id,
    patient_id,
    case_history,
    medication,
    description,
    created_at,
    updated_at; 

-- name: GetPrescriptionById :one 
SELECT 
    id,
    doctor_id,
    patient_id,
    case_history,
    medication,
    description,
    created_at,
    updated_at
FROM prescriptions
WHERE id=$1 
  AND deleted_at IS NULL;

-- name: ListPrescriptions :many 
SELECT 
    id,
    doctor_id,
    patient_id,
    case_history,
    medication,
    description,
    created_at,
    updated_at,
    COUNT(*) OVER() AS total_count
FROM prescriptions
WHERE deleted_at IS NULL 
ORDER BY created_at DESC
LIMIT $1
OFFSET ($2 - 1) * $1;


-- name: UpdatePrescription :one 
UPDATE prescriptions 
SET 
    doctor_id = $2,
    patient_id = $3,
    case_history = $4,
    medication = $5,
    description= $6,
    updated_at = $7
WHERE id = $1 
    AND deleted_at IS NULL
RETURNING
    id,
    doctor_id,
    patient_id,
    case_history,
    medication,
    description,
    created_at,
    updated_at; 

-- name: DeletePrescription :exec
UPDATE prescriptions
SET deleted_at = $2
WHERE id = $1;