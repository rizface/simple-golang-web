package web

type DetailMaintenance struct {
	IdMaintenance    int64
	NamaPetugas      string `validate:"required"`
	LamaPengerjaan   int    `validate: "required"`
	InformasiLainnya string
	DetailPengerjaan string
}
