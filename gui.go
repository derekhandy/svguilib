//
//												svguilib @ v1.1.3
//
//									MIT License, Copyright (c) 2026 Derek Handy
//							Project can be found at: https://github.com/derekhandy/svguilib
//

package svguilib

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

type GUI struct {
	Name       string
	Components IComponents
	Data       IData
	Options    IOptions
}

type IOptions struct {
	Resizeable bool
	Size       IVector2
	Order      []string
	ShowTitles bool
}

type IComponents struct {
	Buttons []IButton
	Labels  []ILabel
	Entries []IInputField
}

type IButton struct {
	Text    string
	Command func()
	Object  *widget.Button
}

type ILabel struct {
	Text   string
	Bold   bool
	Size   IVector2
	Object *widget.Label
}

type IInputField struct {
	Label  string
	Object *widget.Entry
}

type IData struct {
	Data      []IVar
	BoundData []IBoundData
}

type IBoundData struct {
	Data binding.String
}

type IVar struct {
	Name  string
	T     any
	Value interface{}
}

type IVector2 struct {
	X float32
	Y float32
}

func (gui *GUI) StartGUI() {
	components := &gui.Components
	data := &gui.Data
	options := &gui.Options

	a := app.New()
	w := a.NewWindow(gui.Name)

	w.Resize(fyne.Size{Width: options.Size.X, Height: options.Size.Y})
	w.SetFixedSize(!options.Resizeable)
	w.CenterOnScreen()
	w.SetMaster()

	var buttonLabel *widget.Label
	var labelLabel *widget.Label
	var inputLabel *widget.Label

	if options.ShowTitles {
		if len(components.Buttons) > 0 {
			buttonLabel = widget.NewLabel("Commands: ")
		}
		if len(components.Labels) > 0 {
			labelLabel = widget.NewLabel("Labels: ")
		}
		if len(components.Entries) > 0 {
			inputLabel = widget.NewLabel("Input Fields: ")
		}
	}

	buttons := gui.MakeButtons()
	labels := gui.MakeLabels()
	entryContainers, boundEntryData := gui.MakeEntries()

	data.BoundData = make([]IBoundData, len(boundEntryData))
	for i := range boundEntryData {
		data.BoundData[i] = IBoundData{Data: boundEntryData[i]}
	}

	content := gui.BuildOrderedLayout(buttons, labels, entryContainers, buttonLabel, labelLabel, inputLabel)
	w.SetContent(content)

	w.ShowAndRun()
}

func (gui *GUI) MakeButtons() []fyne.CanvasObject {
	// Create buttons
	buttons := make([]*widget.Button, len(gui.Components.Buttons))
	for i := range gui.Components.Buttons {
		buttons[i] = widget.NewButton(gui.Components.Buttons[i].Text, func() { gui.Components.Buttons[i].Command() })
	}

	// Create button objects
	buttonObjects := make([]fyne.CanvasObject, len(buttons))
	for i, b := range buttons {
		buttonObjects[i] = b
	}

	// Set button objects -> GUI object
	for i, b := range buttons {
		gui.Components.Buttons[i].Object = b
	}

	return buttonObjects
}

func (gui *GUI) MakeLabels() []fyne.CanvasObject {
	// Create labels
	labels := make([]*widget.Label, len(gui.Components.Labels))
	for i := range gui.Components.Labels {
		labels[i] = widget.NewLabelWithStyle(gui.Components.Labels[i].Text, fyne.TextAlignCenter, fyne.TextStyle{Bold: gui.Components.Labels[i].Bold})
		if gui.Components.Labels[i].Size.X > 0 {
			labels[i].Resize(fyne.Size{Width: gui.Components.Labels[i].Size.X, Height: gui.Components.Labels[i].Size.Y})
		}
	}

	// Create label objects
	labelObjects := make([]fyne.CanvasObject, len(labels))
	for i, l := range labels {
		labelObjects[i] = l
	}

	// Set label objects -> GUI object
	for i, l := range labels {
		gui.Components.Labels[i].Object = l
	}

	return labelObjects
}

func (gui *GUI) MakeEntries() ([]fyne.CanvasObject, []binding.String) {
	// Create entries
	entries := make([]*widget.Entry, len(gui.Components.Entries))
	boundEntryData := make([]binding.String, len(gui.Components.Entries))
	entryContainers := make([]fyne.CanvasObject, len(gui.Components.Entries))
	for i := range gui.Components.Entries {
		entryData := binding.NewString()
		boundEntryData[i] = entryData
		entries[i] = widget.NewEntryWithData(entryData)
		entries[i].SetPlaceHolder(gui.Components.Entries[i].Label)
		width := float32(150)
		entryContainers[i] = container.NewGridWrap(fyne.Size{Width: width, Height: 36}, entries[i])
		entries[i].Resize(fyne.Size{Width: width, Height: 36})
	}

	// Set entry objects -> GUI object
	for i, e := range entries {
		gui.Components.Entries[i].Object = e
	}

	return entryContainers, boundEntryData
}

func (gui *GUI) BuildOrderedLayout(
	buttons []fyne.CanvasObject,
	labels []fyne.CanvasObject,
	entries []fyne.CanvasObject,
	buttonLabel *widget.Label,
	labelLabel *widget.Label,
	inputLabel *widget.Label,
) fyne.CanvasObject {
	var content []fyne.CanvasObject

	// If ComponentOrder is empty, use default order
	order := gui.Options.Order
	if len(order) == 0 {
		order = []string{"buttons", "labels", "entries"}
	}

	// Map to track which components have been added
	added := make(map[string]bool)

	// Build content based on ComponentOrder
	for _, componentType := range order {
		switch componentType {
		case "buttons":
			if !added["buttons"] && len(buttons) > 0 {
				if gui.Options.ShowTitles && buttonLabel != nil {
					content = append(content, container.NewCenter(buttonLabel))
				}
				content = append(content,
					container.NewCenter(container.NewHBox(buttons...)),
				)
				added["buttons"] = true
			}
		case "labels":
			if !added["labels"] && len(labels) > 0 {
				if gui.Options.ShowTitles && labelLabel != nil {
					content = append(content, container.NewCenter(labelLabel))
				}
				content = append(content,
					container.NewCenter(container.NewHBox(labels...)),
				)
				added["labels"] = true
			}
		case "entries":
			if !added["entries"] && len(entries) > 0 {
				if gui.Options.ShowTitles && inputLabel != nil {
					content = append(content, container.NewCenter(inputLabel))
				}
				content = append(content,
					container.NewCenter(container.NewHBox(entries...)),
				)
				added["entries"] = true
			}
		}
	}

	// Add any remaining components that weren't in ComponentOrder
	if !added["buttons"] && len(buttons) > 0 {
		if gui.Options.ShowTitles && buttonLabel != nil {
			content = append(content, container.NewCenter(buttonLabel))
		}
		content = append(content,
			container.NewCenter(container.NewHBox(buttons...)),
		)
	}
	if !added["labels"] && len(labels) > 0 {
		if gui.Options.ShowTitles && labelLabel != nil {
			content = append(content, container.NewCenter(labelLabel))
		}
		content = append(content,
			container.NewCenter(container.NewHBox(labels...)),
		)
	}
	if !added["entries"] && len(entries) > 0 {
		if gui.Options.ShowTitles && inputLabel != nil {
			content = append(content, container.NewCenter(inputLabel))
		}
		content = append(content,
			container.NewCenter(container.NewHBox(entries...)),
		)
	}

	return container.NewVBox(content...)
}

func (gui *GUI) GetBoundData(index int) string {
	if index < 0 || index >= len(gui.Data.BoundData) {
		return ""
	}

	boundData, _ := gui.Data.BoundData[index].Data.Get()
	return boundData
}

func (gui *GUI) SetBoundData(index int, value string) {
	if index < 0 || index >= len(gui.Data.BoundData) {
		return
	}

	gui.Data.BoundData[index].Data.Set(value)
}
