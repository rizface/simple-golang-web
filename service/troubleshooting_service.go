package service

import (
	"context"
	"pbl-orkom/model/domain"
	"pbl-orkom/model/web"
)

type TroubleshootingService interface {
	Get(ctx context.Context) []domain.Troubleshooting
	GetById(ctx context.Context, idTrouble int)
	SaveForm(ctx context.Context) []domain.Component
	Save(ctx context.Context, request web.TroubleshootRequest) bool
	Update(ctx context.Context, idTrouble int, request web.TroubleshootRequest) bool
	Delete(ctx context.Context, idTrouble int) bool
}
