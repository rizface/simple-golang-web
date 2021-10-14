package domain

type DashboardData struct {
	Maintenance int `json:"maintenance"`
	Troubleshooting int `json:"troubleshooting"`
	Admin int `json:"admin"`
}
