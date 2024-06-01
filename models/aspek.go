package models

import (
    "time"
)

type Aspek struct {
    ID         uint      `json:"id" gorm:"primary_key"`
    NamaAspek  string    `json:"nama_aspek"`
    Kode       string    `json:"kode"`
    KelasID    uint      `json:"kelas_id"`
    CreatedAt  time.Time `json:"created_at"`
    UpdatedAt  time.Time `json:"updated_at"`
}

func (Aspek) TableName() string {
    return "aspek"
}
