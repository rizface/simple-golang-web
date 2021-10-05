package service

import (
	"context"
	"pbl-orkom/model/domain"
	"pbl-orkom/model/web"
)

type AdminService interface{
	Get(ctx context.Context) []domain.User
	Save(ctx context.Context, request web.UserRequest) bool
	UpdateForm(ctx context.Context, idAdmin int) domain.User
	Update(ctx context.Context, idAdmin int, request web.UserRequest) bool
	Delete(ctx context.Context, idAdmin int) bool
}
