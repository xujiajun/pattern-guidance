package builder


type Car struct {
	wheel string
	engine string
	body string
}


type Builder interface {
	BuildWheel() Builder
	BuildEngine() Builder
	BuildBody() Builder
	BuildCar() Car
}


type BuilderCarImpl struct {
	car Car
}

func (builder *BuilderCarImpl)BuildWheel() Builder {
	builder.car.wheel = "wheel1"
	return builder
}

func (builder *BuilderCarImpl)BuildEngine() Builder {
	builder.car.engine = "engine1"
	return builder
}

func (builder *BuilderCarImpl)BuildBody() Builder {
	builder.car.body = "body1"
	return builder
	}

func (builder *BuilderCarImpl)BuildCar() Car   {
	return builder.car
}

type Director struct {
	builder Builder
}

func NewDirector(builder Builder) *Director  {
	return &Director{builder:builder}
}
