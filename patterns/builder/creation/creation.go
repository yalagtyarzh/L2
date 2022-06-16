package creation

// ManufacturingDirector - структура, которая позволяет менять строителя по необходимости
type ManufacturingDirector struct {
	builder BuildProcess
}

// SetBuilder присваивает определенного строителя директору
func (f *ManufacturingDirector) SetBuilder(b BuildProcess) {
	f.builder = b
}

// Construct создает объект, которым занимается заранее определенный строитель
func (f *ManufacturingDirector) Construct() {
	f.builder.SetSeats().SetStructure().SetWheels()
}

// BuildProcess - интерфейс, который указывает, какими методами должен обладать объект,
// чтобы быть строителем
type BuildProcess interface {
	SetWheels() BuildProcess
	SetSeats() BuildProcess
	SetStructure() BuildProcess
	GetVehicle() VehicleProduct
}

// VehicleProduct - объект, представляющий собой средство передвижения
type VehicleProduct struct {
	Wheels    int
	Seats     int
	Structure string
}

// CarBuilder - объект, который является строителем машин
type CarBuilder struct {
	v VehicleProduct
}

// SetWheels устанавливает определенное количество колес для средства передвижения,
// представляющее собой машину
func (c *CarBuilder) SetWheels() BuildProcess {
	c.v.Wheels = 4
	return c
}

// SetSeats устанавливает определенное количество сидений для средства передвижения,
// представляющее собой машину
func (c *CarBuilder) SetSeats() BuildProcess {
	c.v.Seats = 5
	return c
}

// SetStructure устанавливает структуру средства передвижения, представляющее собой машину
func (c *CarBuilder) SetStructure() BuildProcess {
	c.v.Structure = "Car"
	return c
}

// GetVehicle возвращает средство передвижения, представляющее собой машину
func (c *CarBuilder) GetVehicle() VehicleProduct {
	return c.v
}

// Далее по аналогии с мотоциклом и автобусом соответственно

type BikeBuilder struct {
	v VehicleProduct
}

func (b *BikeBuilder) SetWheels() BuildProcess {
	b.v.Wheels = 2
	return b
}

func (b *BikeBuilder) SetSeats() BuildProcess {
	b.v.Seats = 2
	return b
}

func (b *BikeBuilder) SetStructure() BuildProcess {
	b.v.Structure = "Motorbike"
	return b
}

func (b *BikeBuilder) GetVehicle() VehicleProduct {
	return b.v
}

type BusBuilder struct {
	v VehicleProduct
}

func (b *BusBuilder) SetWheels() BuildProcess {
	b.v.Wheels = 4 * 2
	return b
}
func (b *BusBuilder) SetSeats() BuildProcess {
	b.v.Seats = 30
	return b
}
func (b *BusBuilder) SetStructure() BuildProcess {
	b.v.Structure = "Bus"
	return b
}
func (b *BusBuilder) GetVehicle() VehicleProduct {
	return b.v
}
