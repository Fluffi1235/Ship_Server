package radar

import (
	"context"
	ssov1 "diplomServer/gen"
	"google.golang.org/grpc"
	"sync"
)

type ServerRadar struct {
	ssov1.UnimplementedRadarServer
	Radar           *RadarImpl
	mu              sync.Mutex
	ClientsRadarDep map[string]ssov1.Radar_GetRadarInfoServer
}

func Register(gRPC *grpc.Server, radar *RadarImpl) {
	serverRadar := &ServerRadar{
		Radar:           radar,
		ClientsRadarDep: make(map[string]ssov1.Radar_GetRadarInfoServer),
	}
	go serverRadar.SendRadarInfoStream()
	go serverRadar.Radar.ShipMovement()
	go serverRadar.Radar.ObjectMovement("Авианесущий крейсер проекта 1143.5", "Союзник")
	go serverRadar.Radar.ObjectMovement("Атомные ракетный крейсер проект 1144", "Враг")
	go serverRadar.Radar.ObjectMovement("Ракетный крейсер проект 1164", "Враг")
	go serverRadar.Radar.ObjectMovement("Миноносец проект 956", "Враг")

	ssov1.RegisterRadarServer(
		gRPC,
		serverRadar,
	)
}

func (s *ServerRadar) GetRadarInfo(req *ssov1.UserName, stream ssov1.Radar_GetRadarInfoServer) error {
	s.mu.Lock()
	s.ClientsRadarDep[req.UserName] = stream
	s.mu.Unlock()

	<-stream.Context().Done()
	s.mu.Lock()
	delete(s.ClientsRadarDep, req.UserName)
	s.mu.Unlock()

	return nil
}

func (s *ServerRadar) SendRadarInfoStream() {
	for {
		select {
		case obj := <-s.Radar.InfoObject:
			for _, v := range s.ClientsRadarDep {
				err := v.Send(obj)
				if err != nil {

				}
			}
		case obj := <-s.Radar.InfoShip:
			for _, v := range s.ClientsRadarDep {
				err := v.Send(obj)
				if err != nil {

				}
			}
		}
	}
}

func (s *ServerRadar) ChangeShipParameters(ctx context.Context, req *ssov1.UpdateShipParameters) (*ssov1.Empty, error) {
	if req.TypeParam == "speed" {
		s.Radar.Ship.Speed += req.Value
	} else {
		s.Radar.Ship.Angel += req.Value
	}

	return &ssov1.Empty{}, nil
}
