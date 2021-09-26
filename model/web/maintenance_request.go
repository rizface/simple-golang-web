package web

type MaintenanceRequest struct {
	Motherboard   string `validate:"required"`
	VendorRam     string `validate:"required"`
	JumlahRam     string `validate:"required"`
	GraphicCard   string `validate:"required"`
	NIC           string `validate:"required"`
	SistemOperasi string `validate:"required"`
	ArchOS        string `validate:"required"`
}
