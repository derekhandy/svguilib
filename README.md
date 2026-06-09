# svguilib


<b> svguilib </b> is a Go library for quickly composing small Fyne desktop GUIs from declarative buttons, labels, input fields, bound data, and layout options. Cross-platform support.

Requires <b> Go 1.26 </b> or newer and <b> Fyne.io </b> desktop build requirements.


```go
// Usage Example

package main

import (
	. "github.com/derekhandy/svguilib"
)

func main() {
	gui := GUI{
		Name: "example",
		Components: IComponents{
			Labels: []ILabel{
				{Text: "example", Bold: true},
			},
			Entries: []IInputField{
				{Label: "Name"},
			},
			Buttons: []IButton{
				{Text: "Greet"},
			},
		},
		Options: IOptions{
			Size:       IVector2{X: 360, Y: 160},
			Resizeable: true,
			Order:      []string{"labels", "entries", "buttons"},
			ShowTitles: false,
		},
	}

	gui.Components.Buttons[0].Command = func() {
		name := gui.GetBoundData(0)
		gui.Components.Labels[0].Object.SetText("Hello " + name)
	}

	gui.StartGUI()
}
```

## Install

Download from repository or add the module to an existing Go project:

```bash
go get github.com/derekhandy/svguilib
go mod tidy
```

*The 'Fyne.io' dependency's compile time can be 5+ minutes long on first build.
It's normal for the first 'go run .' or 'go build' to stall for longer than expected.*

*GUI builds depend on Fyne's native desktop requirements. Build on the target operating system, or use a configured Fyne cross-compilation toolchain for Windows and Mac OS release binaries.*

## Use

Create a `GUI`, define component slices, set layout options, attach button commands, then call `StartGUI`.

```go
gui := GUI{
	Name: "tool",
	Components: IComponents{
		Buttons: []IButton{{Text: "Run"}},
		Labels:  []ILabel{{Text: "ready", Bold: true}},
		Entries: []IInputField{{Label: "Input"}},
	},
	Options: IOptions{
		Size:       IVector2{X: 480, Y: 180},
		Resizeable: true,
		Order:      []string{"labels", "entries", "buttons"},
		ShowTitles: false,
	},
}

gui.Components.Buttons[0].Command = func() {
	gui.Components.Labels[0].Object.SetText(gui.GetBoundData(0))
}

gui.StartGUI()
```

## API

```go
// Root GUI definition.
type GUI struct {
	Name       string
	Components IComponents
	Data       IData
	Options    IOptions
}

// Window and layout options.
type IOptions struct {
	Resizeable bool
	Size       IVector2
	Order      []string
	ShowTitles bool
}

// Component groups.
type IComponents struct {
	Buttons []IButton
	Labels  []ILabel
	Entries []IInputField
}

// Starts the Fyne app and opens the window.
gui.StartGUI()
```

## Components

```go
// Button with display text, callback, and populated Fyne object.
type IButton struct {
	Text    string
	Command func()
	Object  *widget.Button
}

// Label with display text, bold style, optional size, and populated Fyne object.
type ILabel struct {
	Text   string
	Bold   bool
	Size   IVector2
	Object *widget.Label
}

// Entry with placeholder label and populated Fyne object.
type IInputField struct {
	Label  string
	Object *widget.Entry
}
```

## Examples

```go
// Read entry data from the first input field.

value := gui.GetBoundData(0)
```

```go
// Set entry data programmatically.

gui.SetBoundData(0, "new value")
```

```go
// Show generated titles above each component group.

gui.Options.ShowTitles = true
```

```go
// Change layout order.

gui.Options.Order = []string{"entries", "buttons", "labels"}
```

## Info

Supported layout order values are:

```text
buttons
labels
entries
```

If `Options.Order` is empty, the default order is:

```go
[]string{"buttons", "labels", "entries"}
```

`GetBoundData` returns an empty string when the index does not exist. `SetBoundData` ignores out-of-range indexes.

## NOTICE

<b> svguilib is a lightweight wrapper around Fyne widgets and layout containers. Calling applications are responsible for validating user input, handling long-running work off the UI path, and packaging native GUI builds for each supported operating system.</b>
