package scenes

type IScene interface {
	Pushed(director *Director) error // Executed when the state is pushed onto the StateMachine stack
	Popped(director *Director) error // Executed when the state is popped off of the StateMachine stack
	Tick(dt float32)
	GetName() string
	ShouldQuit() bool
}

type Scene struct {
	Director *Director
	Name     string
}

func (s *Scene) Pushed(director *Director) error {
	s.Director = director
	return nil
}

func (s *Scene) Popped(director *Director) error {
	return nil
}

func (s *Scene) Tick(dt float32) {
	// abstract
}

func (s *Scene) GetName() string {
	return s.Name
}

// ShouldQuit is deprecated, use ShouldQuit on Director instead
func (s *Scene) ShouldQuit() bool {
	return s.Director.ShouldQuit
}
