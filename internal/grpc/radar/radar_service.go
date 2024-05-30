package radar

import (
	ssov1 "diplomServer/gen"
	"go.uber.org/zap"
	"math"
	"math/rand"
	"sync"
	"time"
)

type RadarImpl struct {
	log        *zap.SugaredLogger
	mu         sync.Mutex
	InfoShip   chan *ssov1.Object
	InfoObject chan *ssov1.Object
	Ship       ssov1.Object
}

func New(log *zap.SugaredLogger) *RadarImpl {
	return &RadarImpl{
		log:        log,
		InfoShip:   make(chan *ssov1.Object),
		InfoObject: make(chan *ssov1.Object),
	}
}

func (s *RadarImpl) ObjectMovement(name, status string) {
	x := rand.Intn(500) + 300
	y := rand.Intn(300) + 300
	speed := 1
	angle := rand.Intn(360)

	point := ssov1.Object{
		Name:  name,
		Type:  status,
		X:     float32(x),
		Y:     float32(y),
		Speed: float32(speed),
		Angel: float32(angle),
	}

	for {
		s.InfoObject <- &point

		point.X, point.Y = BeyondMap(point.X, point.Y)

		// Конвертируем угол из градусов в радианы
		rad := float64(point.Angel) * (math.Pi / 180.0)
		point.X += float32(point.Speed) * float32(math.Sin(rad))
		point.Y -= float32(point.Speed) * float32(math.Cos(rad))

		if math.Sqrt(math.Pow(float64(point.X-s.Ship.X), 2)+math.Pow(float64(point.Y-s.Ship.Y), 2)) < 170 {
			point.InRangeShip = true
		} else {
			point.InRangeShip = false
		}

		probability := rand.Intn(100)

		if probability < 10 {
			point.Angel -= float32(rand.Intn(10))
		} else if probability < 20 {
			point.Angel += float32(rand.Intn(10))
		}

		point.Angel = ChangeAngel(point.Angel)

		time.Sleep(200 * time.Millisecond)
	}
}

func (s *RadarImpl) ShipMovement() {
	x := rand.Intn(500) + 300
	y := rand.Intn(300) + 300

	s.Ship = ssov1.Object{
		Name:  "Ship",
		Type:  "main",
		X:     float32(x),
		Y:     float32(y),
		Speed: 0,
		Angel: 0,
	}

	for {
		s.InfoShip <- &s.Ship
		s.Ship.X, s.Ship.Y = BeyondMap(s.Ship.X, s.Ship.Y)

		rad := float64(s.Ship.Angel) * (math.Pi / 180.0)
		s.Ship.X += float32(s.Ship.Speed) * float32(math.Sin(rad))
		s.Ship.Y -= float32(s.Ship.Speed) * float32(math.Cos(rad))

		s.Ship.Angel = ChangeAngel(s.Ship.Angel)

		time.Sleep(200 * time.Millisecond)
	}
}

func ChangeAngel(angel float32) float32 {
	angel = float32(math.Mod(float64(angel), 360.0))
	if angel < 0 {
		angel += 360.0
	}

	return angel
}

func BeyondMap(x, y float32) (float32, float32) {
	if x <= 0 {
		x = 1200
	}

	if y <= 0 {
		y = 1200
	}

	if x > 2000 {
		x = 0
	}

	if y > 2000 {
		y = 0
	}

	return x, y
}
