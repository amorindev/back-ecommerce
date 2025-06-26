package service

import (
	authPort "com.fernando/pkg/app/auth/port"
	roleRepo "com.fernando/pkg/app/role/port"
	sessionPort "com.fernando/pkg/app/session/port"
	userPort "com.fernando/pkg/app/user/port"
)

var _ sessionPort.SessionSrv = &Service{}

type Service struct {
	SessionRepository sessionPort.SessionRepo
	AuthRepository    authPort.AuthRepo
	UserRepository    userPort.UserRepo
	RoleRepository    roleRepo.RoleRepo
}

func NewSessionSrv(
	authRepo authPort.AuthRepo,
	userRepo userPort.UserRepo,
	roleRepo roleRepo.RoleRepo,
	sessionRepo sessionPort.SessionRepo,
) *Service {
	return &Service{
		SessionRepository: sessionRepo,
		AuthRepository:    authRepo,
		UserRepository:    userRepo,
		RoleRepository:    roleRepo,
	}
}
