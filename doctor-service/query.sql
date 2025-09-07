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
    a.id,
    a.doctor_id,
    d.full_name AS doctor_name,
    a.patient_id
    p.full_name AS patient_name,
    a.appointment_date,
    a.created_at,
    a.updated_at 
FROM appointments a 
JOIN doctors d ON a.doctor_id = d.id
JOIN patients p ON a.patient_id = p.id 
WHERE a.id=$1 
    AND deleted_at IS NULL; 

-- name: ListAppointments :many
SELECT 
    a.id,
    a.doctor_id,
    d.full_name AS doctor_name,
    a.patient_id
    p.full_name AS patient_name,
    a.appointment_date,
    a.created_at,
    a.updated_at,
    COUNT(*) OVER() AS total_count
FROM appointments a 
JOIN doctors d ON a.doctor_id = d.id
JOIN patients p ON a.patient_id = p.id 
WHERE deleted_at IS NULL 
    AND (:date IS NULL OR DATE(a.appointment_date) = DATE(:date)) -- DATE() drops minutes
ORDER BY a.appointment_date DESC
LIMIT :limit 
OFFSET :offset;

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
    pr.id,
    pr.doctor_id,
    d.full_name AS doctor_name,
    pr.patient_id,
    p.full_name AS patient_name,
    pr.case_history,
    pr.medication,
    pr.description,
    pr.created_at,
    pr.updated_at; 
FROM prescriptions pr 
JOIN doctors d ON pr.doctor_id = d.id
JOIN patients p ON pr.patient_id = p.id
WHERE id=$1 
    AND deleted_at IS NULL; 

-- name: ListPrescriptions :many 
SELECT 
    pr.id,
    pr.doctor_id,
    d.full_name AS doctor_name,
    pr.patient_id,
    p.full_name AS patient_name,
    pr.case_history,
    pr.medication,
    pr.description,
    pr.created_at,
    pr.updated_at; 
FROM prescriptions pr
JOIN doctors d ON pr.doctor_id = d.id
JOIN patients p ON pr.patient_id = p.id
WHERE deleted_at IS NULL 
ORDER BY created_at DESC
LIMIT :limit
OFFSET :offset;  

-- name: UpdatePrescription :one 
UPDATE prescriptions 
SET 
    doctor_id = $2,
    patient_id = $3,
    case_history = $4,
    medication = $5,
    description= $6,
    updated_at = $7; 
WHERE id = $1 
    AND deleted_at IS NULL; 

-- name: DeletePrescription :exec
UPDATE prescriptions
SET deleted_at = $2
WHERE id = $1;