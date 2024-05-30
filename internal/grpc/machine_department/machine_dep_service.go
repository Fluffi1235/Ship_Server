package machine_department

import (
	"context"
	ssov1 "diplomServer/gen"
	"go.uber.org/zap"
	"time"
)

type MachineImpl struct {
	log            *zap.SugaredLogger
	MachineService MachineDep
	tokenTTl       time.Duration
}

type MachineDep interface {
	GetInfoMachineDep(ctx context.Context) (*ssov1.InfoMachineDep, error)
	ChangeInfoEngine(ctx context.Context, req *ssov1.Engine) error
	ChangeCoolingSystem(ctx context.Context, req *ssov1.CoolingSystem) error
	ChangeFuelSystem(ctx context.Context, req *ssov1.FuelSystem) error
	ChangeGenerator(ctx context.Context, req *ssov1.Generator) error
}

type MachineDepData interface {
	GetInfoMachineDep(ctx context.Context) (*ssov1.InfoMachineDep, error)
	ChangeInfoEngine(ctx context.Context, req *ssov1.Engine) error
	ChangeCoolingSystem(ctx context.Context, req *ssov1.CoolingSystem) error
	ChangeFuelSystem(ctx context.Context, req *ssov1.FuelSystem) error
	ChangeGenerator(ctx context.Context, req *ssov1.Generator) error
}

func New(log *zap.SugaredLogger, machineService MachineDep, tokenTTL time.Duration) *MachineImpl {
	return &MachineImpl{
		log:            log,
		MachineService: machineService,
		tokenTTl:       tokenTTL,
	}
}

func (m *MachineImpl) GetInfoMachineDep(ctx context.Context) (*ssov1.InfoMachineDep, error) {
	infoEng, err := m.MachineService.GetInfoMachineDep(ctx)
	if err != nil {
		return nil, err
	}

	return infoEng, nil
}

func (m *MachineImpl) ChangeInfoEngine(ctx context.Context, req *ssov1.Engine) error {
	err := m.MachineService.ChangeInfoEngine(ctx, req)
	if err != nil {
		return err
	}

	return nil
}

func (m *MachineImpl) ChangeCoolingSystem(ctx context.Context, req *ssov1.CoolingSystem) error {
	err := m.MachineService.ChangeCoolingSystem(ctx, req)
	if err != nil {
		return err
	}

	return nil
}

func (m *MachineImpl) ChangeChangeFuelSystem(ctx context.Context, req *ssov1.FuelSystem) error {
	err := m.MachineService.ChangeFuelSystem(ctx, req)
	if err != nil {
		return err
	}

	return nil
}

func (m *MachineImpl) ChangeGenerator(ctx context.Context, req *ssov1.Generator) error {
	err := m.MachineService.ChangeGenerator(ctx, req)
	if err != nil {
		return err
	}

	return nil
}
