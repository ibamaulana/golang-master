package model

import "time"

type Penilaian struct {
	ID          int64     `json:"id"`
	SekolahID   int64     `json:"sekolah_id"`
	MuridID     int64     `json:"murid_id"`
	KelasID     int64     `json:"kelas_id"`
	Tahun       int64     `json:"tahun"`
	Semester    int64     `json:"semester"`
	Date        time.Time `json:"date"`
	IndikatorID int64     `json:"indikator_id"`
	Nilai       int64     `json:"nilai"`
	Keterangan  string    `json:"keterangan"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (u *Penilaian) TableName() string {
	return "penilaian_harian"
}
