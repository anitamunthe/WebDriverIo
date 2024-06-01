// models/kelas.go

package models

import (
	"time"
)

type Kelas struct {
	ID        uint       `json:"id" gorm:"primary_key"`
	Kode      string     `json:"kode"`
	NamaKelas string     `json:"nama_kelas"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	KelasAspek []KelasAspek `json:"kelas_aspek,omitempty" gorm:"foreignkey:KelasID"`
}

type KelasAspek struct {
	KelasID uint `json:"-"`
	AspekID uint `json:"aspek_id"`
}