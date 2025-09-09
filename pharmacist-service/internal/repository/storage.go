package repository

import (
	"context"
	"database/sql"
	"github.com/xadichamakhkamova/HospitalManagement/livestream-protos"
	"debt-service/internal/storage"
	"debt-service/logger"
	"strconv"
	"time"

	"github.com/google/uuid"
)

type DebtREPO struct {
	queries *storage.Queries
}

func NewDebtSqlc(db *sql.DB) *storage.Queries {
	return storage.New(db)
}