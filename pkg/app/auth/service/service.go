package service

import (
	authPort "com.fernando/pkg/app/auth/port"
	otpPort "com.fernando/pkg/app/otp-codes/port"
	phonePort "com.fernando/pkg/app/phones/port"
	rolePort "com.fernando/pkg/app/role/port"
	sessionPort "com.fernando/pkg/app/session/port"
	userPort "com.fernando/pkg/app/user/port"
	emailPort "com.fernando/pkg/email/port"
	/* smsPort "com.fernando/pkg/sms/port" */
)

var _ authPort.AuthSrv = &Service{}

// authSrv o service - al igual en service y repo ver standar
// authRepo repo respository
type Service struct {
	AuthRepo    authPort.AuthRepo
	RoleRepo    rolePort.RoleRepo
	UserRepo    userPort.UserRepo
	OtpRepo     otpPort.OtpRepo
	SessionRepo sessionPort.SessionRepo
	PhoneRepo   phonePort.PhoneRepo
	AuthAdapter authPort.AuthAdapter
	EmailSrv    emailPort.EmailSrv
	/* SmsSrv      smsPort.SmsSrv */
	SessionSrv  sessionPort.SessionSrv
	AuthTx      authPort.AuthTransaction
}

// antes de los srv - , authAdapter authPort.AuthAdapter
func NewService(
	userRepo userPort.UserRepo,
	authRepo authPort.AuthRepo,
	roleRepo rolePort.RoleRepo,
	otpRepo otpPort.OtpRepo,
	sessionRepo sessionPort.SessionRepo,
	phoneRepo phonePort.PhoneRepo,
	sessionSrv sessionPort.SessionSrv,
	emailSrv emailPort.EmailSrv,
	/* smsSrv smsPort.SmsSrv, */
	authTx authPort.AuthTransaction,
	authAdapter authPort.AuthAdapter,
) *Service {
	return &Service{
		AuthRepo:    authRepo,
		RoleRepo:    roleRepo,
		UserRepo:    userRepo,
		OtpRepo:     otpRepo,
		SessionRepo: sessionRepo,
		PhoneRepo:   phoneRepo,
		AuthAdapter: authAdapter,
		EmailSrv:    emailSrv,
		/* SmsSrv:      smsSrv, */
		SessionSrv:  sessionSrv,
		AuthTx:      authTx,
	}
}
