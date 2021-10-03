package service

import (
	"context"
	"pbl-orkom/model/domain"
	"pbl-orkom/model/web"
)

type LoginService interface{
	Login(ctx context.Context, request web.LoginRequest) domain.User
}