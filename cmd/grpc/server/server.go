package server

/* import (
	"log"
	"net"

	PGConn "com.fernando/internal/postgresql"
	"com.fernando/internal/resend"
	authAdapter "com.fernando/pkg/app/auth/adapter"
	"com.fernando/pkg/app/auth/grpc/controller"
	authRepository "com.fernando/pkg/app/auth/repository/postgresql"
	authService "com.fernando/pkg/app/auth/service"
	authTransction "com.fernando/pkg/app/auth/transaction/postgresql"
	otpRepository "com.fernando/pkg/app/otp-codes/repository/postgresql"

	//otpMongo "com.fernando/pkg/app/otps-codes/repository/mongo"
	roleRepository "com.fernando/pkg/app/role/repository/postgresql"
	sessionRepository "com.fernando/pkg/app/session/repository/postgresql"
	sessionService "com.fernando/pkg/app/session/service"
	userRepository "com.fernando/pkg/app/user/repository/postgresql"
	branchioAdapter "com.fernando/pkg/branchio/adapter"
	resendAdapter "com.fernando/pkg/email/adapter/resend"
	emailService "com.fernando/pkg/email/service"
	"google.golang.org/grpc"
)

// gRPC o Grpc pra nombres
// Addr es publico ?, debería llamarlo port o addr, debe ser puntero
type gRPCServer struct {
	server *grpc.Server
	addr   string
}

// ? Caundo llamar resend por su nombre al paquete o usar dapter googel awsSes resend
func NewgRPCServer(addr string) *gRPCServer {
	grpcServer := grpc.NewServer(
	//grpc.UnaryInterceptor(interceptor.UnaryServerInterceptor),
	)
	//grpc.UnaryServerInterceptor referencia

	// register our grpc services

	// * Clients
	pgConn := PGConn.New().DB
	// ? name dtabase puedes tener varias bse de datos con postgresql?
	// si usamos los dos servidores al mismo tiempo estaremos creando dos clientes de resend
	resendClient := resend.NewClient()

	//googleVerifier, err := authAdapter.NewGoogleProvider()
	_, err := authAdapter.NewGoogleProvider()
	if err != nil {
		log.Fatal(err)
	}

	
		appleVerifier, err := authAdapter.NewAppleProvider("", "", "", "")
		if err != nil {
			log.Fatal(err)
		}

	// * Adapters
	// falta los .env de branch io y corregir
	brachioAdp := branchioAdapter.NewAdapter()
	resendAdp := resendAdapter.NewAdapter(resendClient)
	//authAdp := authAdapter.NewAdapter(googleVerifier.Verifier)
	authAdp := authAdapter.NewAdapter()

	// * Repositories
	authRepo := authRepository.NewRepository(pgConn)
	roleRepo := roleRepository.NewRepository(pgConn)
	userRepo := userRepository.NewRepository(pgConn)
	sessionRepo := sessionRepository.NewRepository(pgConn)
	otpRepo := otpRepository.NewRepository(pgConn)
	// usar postgresql
	//otpRepository := otpMongo.NewRepository(mongoConn.DB, otpColl)

	// * Transctions
	authTx := authTransction.NewTransaction(pgConn, authRepo, roleRepo)

	// * Services
	emailSrv := emailService.NewService(brachioAdp, resendAdp)
	sessionSrv := sessionService.NewService(authRepo, userRepo, roleRepo, sessionRepo)
	authSrv := authService.NewService(authRepo, roleRepo, userRepo, otpRepo, sessionRepo, emailSrv, sessionSrv, authTx, authAdp)

	controller.NewController(grpcServer, authSrv)

	// ? verificar si es ""
	serv := &gRPCServer{server: grpcServer, addr: addr}

	return serv
}
/*
func (serv *gRPCServer) Start() {
	lis, err := net.Listen("tcp", ":"+serv.addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("gRPC server running http://localhost:%s\n", serv.addr)
	log.Fatal(serv.server.Serve(lis))
} */


// ! esto es otro ver si se va a poner
/* func (serv *gRPCServer) Start(addr string) {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Fatal(serv.server.Serve(lis))
} */

// intentar que no se caiga los servidores
