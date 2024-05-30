package postgreSQL

import (
	"context"
	ssov1 "diplomServer/gen"
	"fmt"
)

type Machine interface {
	GetInfoMachineDep(ctx context.Context) (*ssov1.InfoMachineDep, error)
	ChangeInfoEngine(ctx context.Context, req *ssov1.Engine) error
	ChangeCoolingSystem(ctx context.Context, req *ssov1.CoolingSystem) error
	ChangeFuelSystem(ctx context.Context, req *ssov1.FuelSystem) error
	ChangeGenerator(ctx context.Context, req *ssov1.Generator) error
}

func (s *Storage) GetInfoMachineDep(ctx context.Context) (*ssov1.InfoMachineDep, error) {
	const op = "repository.postgreSQL.GetInfoMachineDep"

	var infoMachineDep ssov1.InfoMachineDep

	// Запрос для получения данных из таблицы EngineStatus
	engineStatusQuery := `
		SELECT status, rpm, temperature, voltage
		FROM EngineStatus
		ORDER BY ID
	`
	// Получаем данные из таблицы EngineStatus
	var engineStatus []ssov1.Engine
	if err := s.db.SelectContext(ctx, &engineStatus, engineStatusQuery); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	// Запрос для получения данных из таблицы CoolingSystemStatus
	coolingSystemStatusQuery := `
		SELECT status, coolanttemperature, systempressure
		FROM CoolingSystemStatus
		ORDER BY ID
	`
	// Получаем данные из таблицы CoolingSystemStatus
	var coolingSystemStatus []ssov1.CoolingSystem
	if err := s.db.SelectContext(ctx, &coolingSystemStatus, coolingSystemStatusQuery); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	// Запрос для получения данных из таблицы GeneratorStatus
	generatorStatusQuery := `
		SELECT status, power, fuel, voltage
		FROM GeneratorStatus
		ORDER BY ID
	`
	// Получаем данные из таблицы GeneratorStatus
	var generatorStatus []ssov1.Generator
	if err := s.db.SelectContext(ctx, &generatorStatus, generatorStatusQuery); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	// Запрос для получения данных из таблицы GeneratorStatus
	fuelStatusQuery := `
		SELECT status, fuellevel, contaminantslevel, fuelfilterstatus, flowrate, leakdetection, fuelpumpstatus
    	FROM FuelSystemStatus
    	ORDER BY ID
	`
	// Получаем данные из таблицы GeneratorStatus
	var fuelSystemsStatus []ssov1.FuelSystem
	if err := s.db.SelectContext(ctx, &fuelSystemsStatus, fuelStatusQuery); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	// Наполняем структуру MachineDep данными из каждой таблицы
	infoMachineDep.Engine1 = &engineStatus[0]
	infoMachineDep.Engine2 = &engineStatus[1]
	infoMachineDep.CoolingSystem1 = &coolingSystemStatus[0]
	infoMachineDep.CoolingSystem2 = &coolingSystemStatus[1]
	infoMachineDep.Generator1 = &generatorStatus[0]
	infoMachineDep.Generator2 = &generatorStatus[1]
	infoMachineDep.FuelSystem1 = &fuelSystemsStatus[0]
	infoMachineDep.FuelSystem2 = &fuelSystemsStatus[1]

	return &infoMachineDep, nil
}

func (s *Storage) ChangeInfoEngine(ctx context.Context, req *ssov1.Engine) error {
	const op = "repository.postgreSQL.ChangeInfoEngine"

	query := `
		Update enginestatus set status = $1, rpm = $2, temperature = $3, voltage = $4 where id = $5
	`

	_, err := s.db.ExecContext(ctx, query, req.Status, req.Rpm, req.Temperature, req.Voltage, req.Id)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (s *Storage) ChangeCoolingSystem(ctx context.Context, req *ssov1.CoolingSystem) error {
	const op = "repository.postgreSQL.ChangeCoolingSystem"

	query := `
		Update coolingsystemstatus set status = $1, coolanttemperature = $2, systempressure = $3 where id = $4
	`

	_, err := s.db.ExecContext(ctx, query, req.Status, req.CoolantTemperature, req.SystemPressure, req.Id)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (s *Storage) ChangeFuelSystem(ctx context.Context, req *ssov1.FuelSystem) error {
	const op = "repository.postgreSQL.ChangeFuelSystem"

	query := `
		Update fuelsystemstatus set status = $1, fuellevel = $2, contaminantslevel = $3,
			fuelfilterstatus = $4, flowrate = $5, leakdetection = $6, fuelpumpstatus = $7                
		    where id = $8
	`

	_, err := s.db.ExecContext(ctx, query, req.Status, req.FuelLevel, req.ContaminantsLevel,
		req.FuelFilterStatus, req.FlowRate, req.LeakDetection, req.FuelPumpStatus, req.Id)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (s *Storage) ChangeGenerator(ctx context.Context, req *ssov1.Generator) error {
	const op = "repository.postgreSQL.ChangeFuelSystem"

	query := `
		Update generatorstatus set status = $1, power = $2, fuel = $3,
			voltage = $4 where id = $5
	`

	_, err := s.db.ExecContext(ctx, query, req.Status, req.Power, req.Fuel, req.Voltage, req.Id)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}
