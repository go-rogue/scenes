package scenes

type Director struct {
	scenes     *Stack
	ShouldQuit bool
}

func (e *Director) PushState(scene IScene) error {
	err := scene.Pushed(e)
	if err != nil {
		return err
	}
	e.scenes.Push(scene)
	return nil
}

func (e *Director) PopState() error {
	err := e.PeekState().(IScene).Popped(e)
	if err != nil {
		return err
	}
	_, popErr := e.scenes.Pop()
	return popErr
}

func (e *Director) ChangeState(scene IScene) error {
	if e.scenes.Len() > 0 {
		err := e.PopState()
		if err != nil {
			return err
		}
	}
	return e.PushState(scene)
}

func (e *Director) PeekState() interface{} {
	if e.scenes.Len() == 0 {
		return nil
	}
	return e.scenes.Peek()
}

func NewDirector(initialScene IScene) *Director {
	d := &Director{
		scenes: newStack(),
	}
	d.PushState(initialScene)
	return d
}
