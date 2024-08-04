package data

import (
	"database/sql"
	"github.com/google/uuid"
	"time"
)

type User struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  password  `json:"-"`
	Activated bool      `json:"activated"`
	CreatedAt time.Time `json:"-"`
	Version   uuid.UUID `json:"version"`
}

type password struct {
	text *string
	hash []byte
}

type UserModel struct {
	db *sql.DB
}
