// models/poin_aspek.go

package models

type PoinAspek struct {
    ID        uint   `json:"id" gorm:"primary_key"`
    NamaPoin string `json:"nama_poin"`
    AspekID   uint   `json:"aspek_id"`
    Aspek     Aspek  `json:"aspek"`
}

// TableName untuk menyesuaikan nama tabel
func (PoinAspek) TableName() string {
    return "poin_aspek"
}
