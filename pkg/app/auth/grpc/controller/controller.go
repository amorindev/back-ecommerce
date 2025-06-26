package controller

import (
	authPb "com.fernando/pkg/app/auth/grpc/gen"
	authPort "com.fernando/pkg/app/auth/port"
	"google.golang.org/grpc"
)

type Controller struct {
	AuthSrv authPort.AuthSrv
	authPb.UnimplementedAuthServiceServer
}

// * Cambie handler a controller para grpc
// como estamos en handler solo handler authHandler
func NewController(serv *grpc.Server, authSrv authPort.AuthSrv) {
	c := &Controller{
		AuthSrv: authSrv,
	}
	// * registrar
	authPb.RegisterAuthServiceServer(serv, c)
}

// * Metadata gRPC
/* func Test(w http.ResponseWriter) (interface{}, error) {
	me := metadata.New(map[string]string{
		"admin": "test",
	})
	metaData := map[string][]string{
		"admin": {"fernan", "Alvaro"},
	}
	//w.Header().Set("Authorization", "tes")
	ctx := metadata.NewOutgoingContext(context.Background(), metaData)
	//metadata.NewIncomingContext()

	headers, ok := metadata.FromIncomingContext(context.Background())
	//metadata.FromOutgoingContext()
	if !ok {
		return nil, errors.New("no metadata found in context")
	}
	tokens := headers.Get("jwt")
	if len(tokens) < 1 {
		return nil, errors.New("no token found in metadata")
	}

	metadata.Pairs("test", "test")

	return nil, nil

} */
