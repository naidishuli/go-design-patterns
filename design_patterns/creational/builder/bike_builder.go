package builder

type BikeBuilder struct {

}

func (bb *BikeBuilder) SetWheels() BuildProcess {
	return nil
}
func (bb *BikeBuilder) SetSeats() BuildProcess {
	return nil
}
func (bb *BikeBuilder) SetStructure() BuildProcess {
	return nil
}
func (bb *BikeBuilder) Build() VehicleProduct {
	return VehicleProduct{}
}
