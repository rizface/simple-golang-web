package web

type TroubleshootRequest struct{
	NamaCustomer string `validate:"required"`
	LamaPengerjaan int `validate:"required"`
	Biaya int `validate:"required"`
	Permasalahan string `validate:"required"`
	DetailPengerjaan string
	InformasiLainnya string
	ComponentId []string
}
