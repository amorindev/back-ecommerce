package interceptor

import (
	"context"

	"google.golang.org/grpc"
)

// ? Puedo llamarlo diferente?
func UnaryServerInterceptor(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
	//fmt.Printf("UnaryServerInterceptor PRE %s", info.FullMethod)
	m, err := handler(ctx, req)
	if err != nil {
		return nil, err
	}
	//fmt.Printf("UnaryServerInterceptor POST %s", info.FullMethod)
	return m, nil
}

// lo demás esta en el post de push notifications golang grpc

/* func methodRequiresAuthentication(fullMethod string) bool {
	return false
}

func extractMethodName(fullMethod string) string{
	// define the prefix to remove
	prefix := ""
} */
