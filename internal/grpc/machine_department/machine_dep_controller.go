package machine_department

import (
	"context"
	ssov1 "diplomServer/gen"
	"google.golang.org/grpc"
	"sync"
)

type ServerMachineDep struct {
	ssov1.UnimplementedMachineDepartmentServer
	Machine           *MachineImpl
	Ship              ssov1.Object
	mu                sync.Mutex
	ClientsMachineDep map[string]ssov1.MachineDepartment_GetInfoMachineDepServer
	NewInfo           chan bool
}

func Register(gRPC *grpc.Server, machine *MachineImpl) {
	server := &ServerMachineDep{
		Machine:           machine,
		NewInfo:           make(chan bool),
		ClientsMachineDep: make(map[string]ssov1.MachineDepartment_GetInfoMachineDepServer),
	}

	go server.SendMachineInfoStream()

	ssov1.RegisterMachineDepartmentServer(
		gRPC,
		server,
	)
}

func (s *ServerMachineDep) GetInfoMachineDep(user *ssov1.UserName, stream ssov1.MachineDepartment_GetInfoMachineDepServer) error {
	s.mu.Lock()
	s.ClientsMachineDep[user.UserName] = stream
	s.mu.Unlock()
	s.NewInfo <- true

	<-stream.Context().Done()
	s.mu.Lock()
	delete(s.ClientsMachineDep, user.UserName)
	s.mu.Unlock()

	return nil

}

func (s *ServerMachineDep) SendMachineInfoStream() {
	for {
		select {
		case <-s.NewInfo:
			resp, err := s.Machine.GetInfoMachineDep(context.Background())
			if err != nil {
				// Обработайте ошибку (например, логирование)
			}

			for userName, client := range s.ClientsMachineDep {
				err = client.Send(resp)
				if err != nil {
					delete(s.ClientsMachineDep, userName)
				}
			}
		}
	}
}

func (s *ServerMachineDep) ChangeInfoEngine(ctx context.Context, req *ssov1.Engine) (*ssov1.Empty, error) {
	err := s.Machine.ChangeInfoEngine(ctx, req)
	if err != nil {
		return nil, err
	}

	s.NewInfo <- true

	return nil, nil
}

func (s *ServerMachineDep) ChangeInfoCoolingSystem(ctx context.Context, req *ssov1.CoolingSystem) (*ssov1.Empty, error) {
	err := s.Machine.ChangeCoolingSystem(ctx, req)
	if err != nil {
		return nil, err
	}

	s.NewInfo <- true

	return nil, nil
}

func (s *ServerMachineDep) ChangeInfoGenerator(ctx context.Context, req *ssov1.Generator) (*ssov1.Empty, error) {
	err := s.Machine.ChangeGenerator(ctx, req)
	if err != nil {
		return nil, err
	}

	s.NewInfo <- true

	return nil, nil
}

func (s *ServerMachineDep) ChangeInfoFuelSystem(ctx context.Context, req *ssov1.FuelSystem) (*ssov1.Empty, error) {
	err := s.Machine.ChangeChangeFuelSystem(ctx, req)
	if err != nil {
		return nil, err
	}

	s.NewInfo <- true

	return nil, nil
}
