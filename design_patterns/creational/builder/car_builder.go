package builder

type CarBuilder struct{}

func (cb *CarBuilder) SetWheels() BuildProcess {
	return nil
}
func (cb *CarBuilder) SetSeats() BuildProcess {
	return nil
}
func (cb *CarBuilder) SetStructure() BuildProcess {
	return nil
}
func (cb *CarBuilder) Build() VehicleProduct {
	return VehicleProduct{}
}
