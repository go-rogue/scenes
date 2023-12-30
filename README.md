# Scenes

<a href="https://pkg.go.dev/github.com/go-rogue/scenes?tab=doc"><img src="https://godoc.org/github.com/golang/gddo?status.svg" alt="GoDoc"></a>
[![Software License](https://img.shields.io/badge/license-MIT-brightgreen.svg?style=flat-square)](LICENSE)

This package provides a Scene interface and Scene Director struct for use in game dev. The stack it uses is thread safe.

## Usage

The Scene Director operates on structs implementing `IScene` however this interface intentionally doesn't provide an `update` or `draw` methods. That is up to you to implement in your extension to `IScene`.

For example you might have your own `Scene` interface such as:

```go
type Scene interface {
    scenes.IScene
    Draw()
    HandleEvent(ev tcell.Event) bool
}
```

You would then cast the return of `PeekState()` to your required interface like so:

```go
scene := director.PeekState().(Scene)
scene.Draw()
```
