package platform

import (
	"context"
	"fmt"
	"net"

	"google.golang.org/grpc"
)

type platformServer struct {
	server *grpc.Server
	port   string
}

// NewServer returns a platformServer with port and server populated
func NewServer(port string) *platformServer {
	return &platformServer{port: port, server: grpc.NewServer()}
}

// Start Registers and runs the platformServer
func (p *platformServer) Start() error {
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%s", p.port))
	if err != nil {
		return err
	}

	RegisterPlatformServer(p.server, p)
	fmt.Printf("running on port %s\n", p.port)
	go p.server.Serve(lis)
	return nil
}

// Stop stops the grpc server
func (p *platformServer) Stop() error {
	p.server.Stop()
	return nil
}

// RegisterModule returns the Details associated with a running module
func (p *platformServer) RegisterModule(ctx context.Context, module *Module) (*Details, error) {
	// TODO - any actual module setup required on the platform

	return &Details{Details: fmt.Sprintf("register %s on port %s", module.Name, module.Port)}, nil
}
