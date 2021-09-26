package service

import (
	"context"
	"pbl-orkom/model/domain"
	"pbl-orkom/model/web"
)

type MaintenanceService interface {
	Get(ctx context.Context) []domain.DetailJoinMaintenance
	//GetById(ctx context.Context, spek web.MaintenanceRequest) domain.DetailJoinMaintenance
	Save(ctx context.Context, spek web.MaintenanceRequest, detail web.DetailMaintenance) bool
	Update(ctx context.Context, idMaintenance int, spek web.MaintenanceRequest) bool
	UpdateDetail(ctx context.Context, idMaintenance int, detail web.DetailMaintenance) bool
	Delete(ctx context.Context, idMaintenance int) bool
	FormUpdate(ctx context.Context, idMaintenance int) domain.MaintenanceRequest
	FormUpdateDetail(ctx context.Context, idMaintenance int) domain.DetailMaintenance
	Export(ctx context.Context) string
	createExcel([]domain.DetailJoinMaintenance) string
}
