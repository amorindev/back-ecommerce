package mongo

import (
	authPort "com.fernando/pkg/app/auth/port"
	otpPort "com.fernando/pkg/app/otp-codes/port"
	rolePort "com.fernando/pkg/app/role/port"
	userPort "com.fernando/pkg/app/user/port"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

var _ authPort.AuthTransaction = &Transaction{}

type Transaction struct {
	Client   *mongo.Client
	AuthRepo authPort.AuthRepo
	UserRepo userPort.UserRepo
	RoleRepo rolePort.RoleRepo
	OtpRepo  otpPort.OtpRepo
}

func NewTransaction(
	client *mongo.Client, 
	authRepo authPort.AuthRepo, 
	userRepo userPort.UserRepo, 
	roleRepo rolePort.RoleRepo, 
	otpRepo otpPort.OtpRepo,
) *Transaction {
	return &Transaction{
		Client:   client,
		AuthRepo: authRepo,
		UserRepo: userRepo,
		RoleRepo: roleRepo,
		OtpRepo:  otpRepo,
	}
}
