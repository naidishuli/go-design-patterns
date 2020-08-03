package builder

type VehicleProduct struct {
	Wheels    int
	Seats     int
	Structure string
}

// This interface defines the steps that are necessary to build a vehicle. Every
// builder must implement this interface if they are to be used by the manufacturing. On
// every Set step, we return the same build process, so we can chain various steps together in
// the same statement, as we'll see later. Finally, we'll need a GetVehicle method to retrieve
// the Vehicle instance from the builder:
type BuildProcess interface {
	SetWheels() BuildProcess
	SetSeats() BuildProcess
	SetStructure() BuildProcess
	Build() VehicleProduct
}


// The ManufacturingDirector director variable is the one in charge of accepting the
// builders. It has a Construct method that will use the builder that is stored in
// Manufacturing, and will reproduce the required steps. The SetBuilder method will
// allow us to change the builder that is being used in the Manufacturing director:
type ManufacturingDirector struct {
}

func (md *ManufacturingDirector) Construct() {

}

func (md *ManufacturingDirector) SetBuilder(b BuildProcess) {

}
