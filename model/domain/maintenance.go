package domain

type MaintenanceRequest struct {
	Id            int
	Motherboard   string
	VendorRam     string
	JumlahRam     string
	GraphicCard   string
	NIC           string
	SistemOperasi string
	ArchOS        string
	TglMasuk string
}
