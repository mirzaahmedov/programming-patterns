package vehicle

type Vehicle struct {
	make     string
	model    string
	warranty string
}

func NewVehicle() Vehicle {
	return Vehicle{
		make:     "Chevrolet",
		model:    "Spark",
		warranty: "None",
	}
}

func (v Vehicle) WithMake(make string) Vehicle {
	v.make = make
	return v
}
func (v Vehicle) WithModel(model string) Vehicle {
	v.model = model
	return v
}
func (v Vehicle) WithWarranty(warranty string) Vehicle {
	v.warranty = warranty
	return v
}
