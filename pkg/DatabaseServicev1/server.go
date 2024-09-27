package DatabaseServicev1

import (
	"google.golang.org/grpc"
)

type serverAPI struct {
	UnimplementedDatabaseServiceServer
}

func Register(gRPC *grpc.Server) {
	RegisterDatabaseServiceServer(gRPC, &serverAPI{})
}
