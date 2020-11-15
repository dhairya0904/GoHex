package interceptor

import (
	"context"
	json "encoding/json"
	logger "rotteneggs/productservice/common/logging"

	"google.golang.org/grpc"
)

// ServerInterceptor - RPC interceptor for all the requests
func ServerInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	requestJSON, _ := json.Marshal(req)

	log := logger.NewLogger()
	log.Info(info.FullMethod)
	log.Info(string(requestJSON))

	return handler(ctx, req)
}
