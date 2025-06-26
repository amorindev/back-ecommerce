package service

import (
	branchioPort "com.fernando/pkg/branchio/port"
	emailPort "com.fernando/pkg/email/port"
)

var _ emailPort.EmailSrv = &EmailService{}

type EmailService struct {
	BranchioAdapter branchioPort.Adapter
	EmailAdapter    emailPort.EmailAdapter
}

func NewEmailSrv(branchioAdapter branchioPort.Adapter, emailAdapter emailPort.EmailAdapter) *EmailService {
	return &EmailService{
		BranchioAdapter: branchioAdapter,
		EmailAdapter:    emailAdapter,
	}
}
