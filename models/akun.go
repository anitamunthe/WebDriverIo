// models/akun.go

package models

type Akun struct {
	ID       int    `json:"id"`
	Nama     string `json:"nama"`
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
	// Properti tambahan sesuai kebutuhan autentikasi
	// ...
}
