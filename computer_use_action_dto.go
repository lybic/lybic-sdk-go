package lybic

import "github.com/lybic/lybic-sdk-go/pkg/json"

// ComputerUseActionDtoActionOneOf is an interface for types that can be used as lybic computer-use action
//
//	Implementations:
//
// MouseClickAction
// MouseDoubleClickAction
// MouseMoveAction
// MouseScrollAction
// MouseDragAction
// KeyboardTypeAction
// KeyboardHotkeyAction
// ScreenshotAction
// WaitAction
// FinishedAction
// FailedAction
type ComputerUseActionDtoActionOneOf interface {
	MarshalJSON() ([]byte, error)
	UnmarshalJSON(src []byte) error

	_internalComputerUseActionDtoActionOneOf()
}

// Length is an interface for types that can be used as lybic computer-use Length
//
//	Implementations:
//
// FractionalLength
// PixelLength
type Length interface {
	MarshalJSON() ([]byte, error)
	UnmarshalJSON(src []byte) error

	_internalLength()
}
type FractionalLength struct {
	Type        string `json:"type"` // set to `/`
	Numerator   int    `json:"numerator"`
	Denominator int    `json:"denominator"`
}

func (f FractionalLength) MarshalJSON() ([]byte, error) {
	var toSerialize map[string]interface{}
	toSerialize = map[string]interface{}{
		"type":        "/",
		"numerator":   f.Numerator,
		"denominator": f.Denominator,
	}
	return json.Marshal(toSerialize)
}
func (f *FractionalLength) UnmarshalJSON(src []byte) error {
	var value map[string]interface{}
	if err := json.Unmarshal(src, &value); err != nil {
		return err
	}
	if v, ok := value["type"].(string); ok {
		f.Type = v
	}
	if v, ok := value["numerator"].(float64); ok {
		f.Numerator = int(v)
	}
	if v, ok := value["denominator"].(float64); ok {
		f.Denominator = int(v)
	}
	return nil
}
func (f FractionalLength) _internalLength() {}
func NewPixelLength(value int) *PixelLength {
	return &PixelLength{
		Type:  "px",
		Value: value,
	}
}

type PixelLength struct {
	Type  string `json:"type"` // set to `px`
	Value int    `json:"value"`
}

func (p PixelLength) MarshalJSON() ([]byte, error) {
	var toSerialize map[string]interface{}
	toSerialize = map[string]interface{}{
		"type":  "px",
		"value": p.Value,
	}
	return json.Marshal(toSerialize)
}

func (p *PixelLength) UnmarshalJSON(src []byte) error {
	var value map[string]interface{}
	if err := json.Unmarshal(src, &value); err != nil {
		return err
	}
	if v, ok := value["type"].(string); ok {
		p.Type = v
	}
	if v, ok := value["value"].(float64); ok {
		p.Value = int(v)
	}
	return nil
}

func (p PixelLength) _internalLength() {}
func NewFractionalLength(numerator int, denominator int) *FractionalLength {
	return &FractionalLength{
		Type:        "/",
		Numerator:   numerator,
		Denominator: denominator,
	}
}

type MouseClickAction struct {
	Type    string  `json:"type"` // set to mouse:click
	X       Length  `json:"x"`
	Y       Length  `json:"y"`
	Button  int     `json:"button"`            // Mouse button flag combination. 1: left, 2: right, 4: middle, 8: back, 16: forward; add them together to press multiple buttons at once.
	HoldKey *string `json:"holdKey,omitempty"` // Key to hold down during click, in xdotool key syntax. Example: "ctrl", "alt", "alt+shift"
	CallId  *string `json:"callId,omitempty"`
}

func (m MouseClickAction) MarshalJSON() ([]byte, error) {
	var toSerialize map[string]interface{}
	toSerialize = map[string]interface{}{
		"type":   "mouse:click",
		"x":      m.X,
		"y":      m.Y,
		"button": m.Button,
	}
	if m.HoldKey != nil {
		toSerialize["holdKey"] = *m.HoldKey
	}
	if m.CallId != nil {
		toSerialize["callId"] = *m.CallId
	}
	return json.Marshal(toSerialize)
}

func (m *MouseClickAction) UnmarshalJSON(src []byte) error {
	var value map[string]interface{}
	if err := json.Unmarshal(src, &value); err != nil {
		return err
	}
	if v, ok := value["type"].(string); ok {
		m.Type = v
	}
	if v, ok := value["x"].(map[string]interface{}); ok {
		if typeVal, ok := v["type"].(string); ok {
			switch typeVal {
			case "/":
				var fracLen FractionalLength
				if err := json.Unmarshal(src, &fracLen); err != nil {
					return err
				}
				m.X = &fracLen
			case "px":
				var pixLen PixelLength
				if err := json.Unmarshal(src, &pixLen); err != nil {
					return err
				}
				m.X = &pixLen
			}
		}
	}
	if v, ok := value["y"].(map[string]interface{}); ok {
		if typeVal, ok := v["type"].(string); ok {
			switch typeVal {
			case "/":
				var fracLen FractionalLength
				if err := json.Unmarshal(src, &fracLen); err != nil {
					return err
				}
				m.Y = &fracLen
			case "px":
				var pixLen PixelLength
				if err := json.Unmarshal(src, &pixLen); err != nil {
					return err
				}
				m.Y = &pixLen
			}
		}
	}
	if v, ok := value["button"].(float64); ok {
		m.Button = int(v)
	}
	if v, ok := value["holdKey"].(string); ok {
		m.HoldKey = &v
	}
	if v, ok := value["callId"].(string); ok {
		m.CallId = &v
	}
	return nil
}

func (MouseClickAction) _internalComputerUseActionDtoActionOneOf() {}
func NewMouseClickAction(x Length, y Length, button int) *MouseClickAction {
	return &MouseClickAction{
		Type:   "mouse:click",
		X:      x,
		Y:      y,
		Button: button,
	}
}

type MouseDoubleClickAction struct {
	Type    string  `json:"type"` // set to mouse:doubleClick
	X       Length  `json:"x"`
	Y       Length  `json:"y"`
	Button  int     `json:"button"`            // Mouse button flag combination. 1: left, 2: right, 4: middle, 8: back, 16: forward; add them together to press multiple buttons at once.
	HoldKey *string `json:"holdKey,omitempty"` // Key to hold down during click, in xdotool key syntax. Example: "ctrl", "alt", "alt+shift"
	CallId  *string `json:"callId,omitempty"`
}

func (m MouseDoubleClickAction) MarshalJSON() ([]byte, error) {
	var toSerialize map[string]interface{}
	toSerialize = map[string]interface{}{
		"type":   "mouse:doubleClick",
		"x":      m.X,
		"y":      m.Y,
		"button": m.Button,
	}
	if m.HoldKey != nil {
		toSerialize["holdKey"] = *m.HoldKey
	}
	if m.CallId != nil {
		toSerialize["callId"] = *m.CallId
	}
	return json.Marshal(toSerialize)
}

func (m *MouseDoubleClickAction) UnmarshalJSON(src []byte) error {
	var value map[string]interface{}
	if err := json.Unmarshal(src, &value); err != nil {
		return err
	}
	if v, ok := value["type"].(string); ok {
		m.Type = v
	}
	if v, ok := value["x"].(map[string]interface{}); ok {
		if typeVal, ok := v["type"].(string); ok {
			switch typeVal {
			case "/":
				var fracLen FractionalLength
				if err := json.Unmarshal(src, &fracLen); err != nil {
					return err
				}
				m.X = &fracLen
			case "px":
				var pixLen PixelLength
				if err := json.Unmarshal(src, &pixLen); err != nil {
					return err
				}
				m.X = &pixLen
			}
		}
	}
	if v, ok := value["y"].(map[string]interface{}); ok {
		if typeVal, ok := v["type"].(string); ok {
			switch typeVal {
			case "/":
				var fracLen FractionalLength
				if err := json.Unmarshal(src, &fracLen); err != nil {
					return err
				}
				m.Y = &fracLen
			case "px":
				var pixLen PixelLength
				if err := json.Unmarshal(src, &pixLen); err != nil {
					return err
				}
				m.Y = &pixLen
			}
		}
	}
	if v, ok := value["button"].(float64); ok {
		m.Button = int(v)
	}
	if v, ok := value["holdKey"].(string); ok {
		m.HoldKey = &v
	}
	if v, ok := value["callId"].(string); ok {
		m.CallId = &v
	}
	return nil
}

func (MouseDoubleClickAction) _internalComputerUseActionDtoActionOneOf() {}
func NewMouseDoubleClickAction(x Length, y Length, button int) *MouseDoubleClickAction {
	return &MouseDoubleClickAction{
		Type:   "mouse:doubleClick",
		X:      x,
		Y:      y,
		Button: button,
	}
}

type MouseMoveAction struct {
	Type    string  `json:"type"` // set to mouse:move
	X       Length  `json:"x"`
	Y       Length  `json:"y"`
	HoldKey *string `json:"holdKey,omitempty"` // Key to hold down during click, in xdotool key syntax. Example: "ctrl", "alt", "alt+shift"
	CallId  *string `json:"callId,omitempty"`
}

func (m MouseMoveAction) MarshalJSON() ([]byte, error) {
	var toSerialize map[string]interface{}
	toSerialize = map[string]interface{}{
		"type": "mouse:move",
		"x":    m.X,
		"y":    m.Y,
	}
	if m.HoldKey != nil {
		toSerialize["holdKey"] = *m.HoldKey
	}
	if m.CallId != nil {
		toSerialize["callId"] = *m.CallId
	}
	return json.Marshal(toSerialize)
}

func (m *MouseMoveAction) UnmarshalJSON(src []byte) error {
	var value map[string]interface{}
	if err := json.Unmarshal(src, &value); err != nil {
		return err
	}
	if v, ok := value["type"].(string); ok {
		m.Type = v
	}
	if v, ok := value["x"].(map[string]interface{}); ok {
		if typeVal, ok := v["type"].(string); ok {
			switch typeVal {
			case "/":
				var fracLen FractionalLength
				if err := json.Unmarshal(src, &fracLen); err != nil {
					return err
				}
				m.X = &fracLen
			case "px":
				var pixLen PixelLength
				if err := json.Unmarshal(src, &pixLen); err != nil {
					return err
				}
				m.X = &pixLen
			}
		}
	}
	if v, ok := value["y"].(map[string]interface{}); ok {
		if typeVal, ok := v["type"].(string); ok {
			switch typeVal {
			case "/":
				var fracLen FractionalLength
				if err := json.Unmarshal(src, &fracLen); err != nil {
					return err
				}
				m.Y = &fracLen
			case "px":
				var pixLen PixelLength
				if err := json.Unmarshal(src, &pixLen); err != nil {
					return err
				}
				m.Y = &pixLen
			}
		}
	}
	if v, ok := value["holdKey"].(string); ok {
		m.HoldKey = &v
	}
	if v, ok := value["callId"].(string); ok {
		m.CallId = &v
	}
	return nil
}

func (MouseMoveAction) _internalComputerUseActionDtoActionOneOf() {}
func NewMouseMoveAction(x Length, y Length) *MouseMoveAction {
	return &MouseMoveAction{
		Type: "mouse:move",
		X:    x,
		Y:    y,
	}
}

type MouseScrollAction struct {
	Type           string  `json:"type"` // set to mouse:scroll
	X              Length  `json:"x"`
	Y              Length  `json:"y"`
	StepVertical   int     `json:"stepVertical"`
	StepHorizontal int     `json:"stepHorizontal"`
	HoldKey        *string `json:"holdKey,omitempty"` // Key to hold down during click, in xdotool key syntax. Example: "ctrl", "alt", "alt+shift"
	CallId         *string `json:"callId,omitempty"`
}

func (m MouseScrollAction) MarshalJSON() ([]byte, error) {
	var toSerialize map[string]interface{}
	toSerialize = map[string]interface{}{
		"type":           "mouse:scroll",
		"x":              m.X,
		"y":              m.Y,
		"stepVertical":   m.StepVertical,
		"stepHorizontal": m.StepHorizontal,
	}
	if m.HoldKey != nil {
		toSerialize["holdKey"] = *m.HoldKey
	}
	if m.CallId != nil {
		toSerialize["callId"] = *m.CallId
	}
	return json.Marshal(toSerialize)
}

func (m *MouseScrollAction) UnmarshalJSON(src []byte) error {
	var value map[string]interface{}
	if err := json.Unmarshal(src, &value); err != nil {
		return err
	}
	if v, ok := value["type"].(string); ok {
		m.Type = v
	}
	if v, ok := value["x"].(map[string]interface{}); ok {
		if typeVal, ok := v["type"].(string); ok {
			switch typeVal {
			case "/":
				var fracLen FractionalLength
				if err := json.Unmarshal(src, &fracLen); err != nil {
					return err
				}
				m.X = &fracLen
			case "px":
				var pixLen PixelLength
				if err := json.Unmarshal(src, &pixLen); err != nil {
					return err
				}
				m.X = &pixLen
			}
		}
	}
	if v, ok := value["y"].(map[string]interface{}); ok {
		if typeVal, ok := v["type"].(string); ok {
			switch typeVal {
			case "/":
				var fracLen FractionalLength
				if err := json.Unmarshal(src, &fracLen); err != nil {
					return err
				}
				m.Y = &fracLen
			case "px":
				var pixLen PixelLength
				if err := json.Unmarshal(src, &pixLen); err != nil {
					return err
				}
				m.Y = &pixLen
			}
		}
	}
	if v, ok := value["stepVertical"].(float64); ok {
		m.StepVertical = int(v)
	}
	if v, ok := value["stepHorizontal"].(float64); ok {
		m.StepHorizontal = int(v)
	}
	if v, ok := value["holdKey"].(string); ok {
		m.HoldKey = &v
	}
	if v, ok := value["callId"].(string); ok {
		m.CallId = &v
	}
	return nil
}

func (MouseScrollAction) _internalComputerUseActionDtoActionOneOf() {}
func NewMouseScrollAction(x Length, y Length, stepVertical int, stepHorizontal int) *MouseScrollAction {
	return &MouseScrollAction{
		Type:           "mouse:scroll",
		X:              x,
		Y:              y,
		StepVertical:   stepVertical,
		StepHorizontal: stepHorizontal,
	}
}

type MouseDragAction struct {
	Type    string  `json:"type"` // set to mouse:drag
	StartX  Length  `json:"startX"`
	StartY  Length  `json:"startY"`
	EndX    Length  `json:"endX"`
	EndY    Length  `json:"endY"`
	HoldKey *string `json:"holdKey,omitempty"` // Key to hold down during click, in xdotool key syntax. Example: "ctrl", "alt", "alt+shift"
	CallId  *string `json:"callId,omitempty"`
}

func (m MouseDragAction) MarshalJSON() ([]byte, error) {
	var toSerialize map[string]interface{}
	toSerialize = map[string]interface{}{
		"type":   "mouse:drag",
		"startX": m.StartX,
		"startY": m.StartY,
		"endX":   m.EndX,
		"endY":   m.EndY,
	}
	if m.HoldKey != nil {
		toSerialize["holdKey"] = *m.HoldKey
	}
	if m.CallId != nil {
		toSerialize["callId"] = *m.CallId
	}
	return json.Marshal(toSerialize)
}

func (m *MouseDragAction) UnmarshalJSON(src []byte) error {
	var value map[string]interface{}
	if err := json.Unmarshal(src, &value); err != nil {
		return err
	}
	if v, ok := value["type"].(string); ok {
		m.Type = v
	}
	if v, ok := value["startX"].(map[string]interface{}); ok {
		if typeVal, ok := v["type"].(string); ok {
			switch typeVal {
			case "/":
				var fracLen FractionalLength
				if err := json.Unmarshal(src, &fracLen); err != nil {
					return err
				}
				m.StartX = &fracLen
			case "px":
				var pixLen PixelLength
				if err := json.Unmarshal(src, &pixLen); err != nil {
					return err
				}
				m.StartX = &pixLen
			}
		}
	}
	if v, ok := value["startY"].(map[string]interface{}); ok {
		if typeVal, ok := v["type"].(string); ok {
			switch typeVal {
			case "/":
				var fracLen FractionalLength
				if err := json.Unmarshal(src, &fracLen); err != nil {
					return err
				}
				m.StartY = &fracLen
			case "px":
				var pixLen PixelLength
				if err := json.Unmarshal(src, &pixLen); err != nil {
					return err
				}
				m.StartY = &pixLen
			}
		}
	}
	if v, ok := value["endX"].(map[string]interface{}); ok {
		if typeVal, ok := v["type"].(string); ok {
			switch typeVal {
			case "/":
				var fracLen FractionalLength
				if err := json.Unmarshal(src, &fracLen); err != nil {
					return err
				}
				m.EndX = &fracLen
			case "px":
				var pixLen PixelLength
				if err := json.Unmarshal(src, &pixLen); err != nil {
					return err
				}
				m.EndX = &pixLen
			}
		}
	}
	if v, ok := value["endY"].(map[string]interface{}); ok {
		if typeVal, ok := v["type"].(string); ok {
			switch typeVal {
			case "/":
				var fracLen FractionalLength
				if err := json.Unmarshal(src, &fracLen); err != nil {
					return err
				}
				m.EndY = &fracLen
			case "px":
				var pixLen PixelLength
				if err := json.Unmarshal(src, &pixLen); err != nil {
					return err
				}
				m.EndY = &pixLen
			}
		}
	}
	if v, ok := value["holdKey"].(string); ok {
		m.HoldKey = &v
	}
	if v, ok := value["callId"].(string); ok {
		m.CallId = &v
	}
	return nil
}

func (MouseDragAction) _internalComputerUseActionDtoActionOneOf() {}
func NewMouseDragAction(startX Length, startY Length, endX Length, endY Length) *MouseDragAction {
	return &MouseDragAction{
		Type:   "mouse:drag",
		StartX: startX,
		StartY: startY,
		EndX:   endX,
		EndY:   endY,
	}
}

type KeyboardTypeAction struct {
	Type                string  `json:"type"` // set to keyboard:type
	Content             string  `json:"content"`
	TreatNewLineAsEnter bool    `json:"treatNewLineAsEnter"` // Whether to treat line breaks as enter. If true, any line breaks(n) in content will be treated as enter key press, and content will be split into multiple lines.
	CallId              *string `json:"callId,omitempty"`
}

func (k KeyboardTypeAction) MarshalJSON() ([]byte, error) {
	var toSerialize map[string]interface{}
	toSerialize = map[string]interface{}{
		"type":                "keyboard:type",
		"content":             k.Content,
		"treatNewLineAsEnter": k.TreatNewLineAsEnter,
	}
	if k.CallId != nil {
		toSerialize["callId"] = *k.CallId
	}
	return json.Marshal(toSerialize)
}

func (k *KeyboardTypeAction) UnmarshalJSON(src []byte) error {
	var value map[string]interface{}
	if err := json.Unmarshal(src, &value); err != nil {
		return err
	}
	if v, ok := value["type"].(string); ok {
		k.Type = v
	}
	if v, ok := value["content"].(string); ok {
		k.Content = v
	}
	if v, ok := value["treatNewLineAsEnter"].(bool); ok {
		k.TreatNewLineAsEnter = v
	}
	if v, ok := value["callId"].(string); ok {
		k.CallId = &v
	}
	return nil
}

func (KeyboardTypeAction) _internalComputerUseActionDtoActionOneOf() {}
func NewKeyboardTypeAction(content string, treatNewLineAsEnter bool) *KeyboardTypeAction {
	return &KeyboardTypeAction{
		Type:                "keyboard:type",
		Content:             content,
		TreatNewLineAsEnter: treatNewLineAsEnter,
	}
}

type KeyboardHotkeyAction struct {
	Type     string  `json:"type"` // set to keyboard:hotkey
	Keys     string  `json:"keys"`
	Duration *int    `json:"duration,omitempty"` // Duration in milliseconds. If specified, the hotkey will be held for a while and then released.
	CallId   *string `json:"callId,omitempty"`
}

func (k KeyboardHotkeyAction) MarshalJSON() ([]byte, error) {
	var toSerialize map[string]interface{}
	toSerialize = map[string]interface{}{
		"type": "keyboard:hotkey",
		"keys": k.Keys,
	}
	if k.Duration != nil {
		toSerialize["duration"] = *k.Duration
	}
	if k.CallId != nil {
		toSerialize["callId"] = *k.CallId
	}
	return json.Marshal(toSerialize)
}

func (k *KeyboardHotkeyAction) UnmarshalJSON(src []byte) error {
	var value map[string]interface{}
	if err := json.Unmarshal(src, &value); err != nil {
		return err
	}
	if v, ok := value["type"].(string); ok {
		k.Type = v
	}
	if v, ok := value["keys"].(string); ok {
		k.Keys = v
	}
	if v, ok := value["duration"].(float64); ok {
		duration := int(v)
		k.Duration = &duration
	}
	if v, ok := value["callId"].(string); ok {
		k.CallId = &v
	}
	return nil
}

func (KeyboardHotkeyAction) _internalComputerUseActionDtoActionOneOf() {}
func NewKeyboardHotkeyAction(keys string) *KeyboardHotkeyAction {
	return &KeyboardHotkeyAction{
		Type: "keyboard:hotkey",
		Keys: keys,
	}
}

type ScreenshotAction struct {
	Type   string  `json:"type"` // set to screenshot
	CallId *string `json:"callId,omitempty"`
}

func (s ScreenshotAction) MarshalJSON() ([]byte, error) {
	var toSerialize map[string]interface{}
	toSerialize = map[string]interface{}{
		"type": "screenshot",
	}
	if s.CallId != nil {
		toSerialize["callId"] = *s.CallId
	}
	return json.Marshal(toSerialize)
}

func (s *ScreenshotAction) UnmarshalJSON(src []byte) error {
	var value map[string]interface{}
	if err := json.Unmarshal(src, &value); err != nil {
		return err
	}
	if v, ok := value["type"].(string); ok {
		s.Type = v
	}
	if v, ok := value["callId"].(string); ok {
		s.CallId = &v
	}
	return nil
}

func (ScreenshotAction) _internalComputerUseActionDtoActionOneOf() {}
func NewScreenshotAction() *ScreenshotAction {
	return &ScreenshotAction{
		Type: "screenshot",
	}
}

type WaitAction struct {
	Type     string  `json:"type"` // set to wait
	Duration int     `json:"duration"`
	CallId   *string `json:"callId,omitempty"`
}

func (w WaitAction) MarshalJSON() ([]byte, error) {
	var toSerialize map[string]interface{}
	toSerialize = map[string]interface{}{
		"type":     "wait",
		"duration": w.Duration,
	}
	if w.CallId != nil {
		toSerialize["callId"] = *w.CallId
	}
	return json.Marshal(toSerialize)
}

func (w *WaitAction) UnmarshalJSON(src []byte) error {
	var value map[string]interface{}
	if err := json.Unmarshal(src, &value); err != nil {
		return err
	}
	if v, ok := value["type"].(string); ok {
		w.Type = v
	}
	if v, ok := value["duration"].(float64); ok {
		w.Duration = int(v)
	}
	if v, ok := value["callId"].(string); ok {
		w.CallId = &v
	}
	return nil
}

func (WaitAction) _internalComputerUseActionDtoActionOneOf() {}
func NewWaitAction(duration int) *WaitAction {
	return &WaitAction{
		Type:     "wait",
		Duration: duration,
	}
}

type FinishedAction struct {
	Type    string  `json:"type"` // set to finished
	Message *string `json:"message,omitempty"`
	CallId  *string `json:"callId,omitempty"`
}

func (f FinishedAction) MarshalJSON() ([]byte, error) {
	var toSerialize map[string]interface{}
	toSerialize = map[string]interface{}{
		"type": "finished",
	}
	if f.Message != nil {
		toSerialize["message"] = *f.Message
	}
	if f.CallId != nil {
		toSerialize["callId"] = *f.CallId
	}
	return json.Marshal(toSerialize)
}

func (f *FinishedAction) UnmarshalJSON(src []byte) error {
	var value map[string]interface{}
	if err := json.Unmarshal(src, &value); err != nil {
		return err
	}
	if v, ok := value["type"].(string); ok {
		f.Type = v
	}
	if v, ok := value["message"].(string); ok {
		f.Message = &v
	}
	if v, ok := value["callId"].(string); ok {
		f.CallId = &v
	}
	return nil
}
func NewFinishedAction() *FinishedAction {
	return &FinishedAction{
		Type: "finished",
	}
}
func (FinishedAction) _internalComputerUseActionDtoActionOneOf() {}

type FailedAction struct {
	Type    string  `json:"type"` // set to failed
	Message *string `json:"message,omitempty"`
	CallId  *string `json:"callId,omitempty"`
}

func (f FailedAction) MarshalJSON() ([]byte, error) {
	var toSerialize map[string]interface{}
	toSerialize = map[string]interface{}{
		"type": "failed",
	}
	if f.Message != nil {
		toSerialize["message"] = *f.Message
	}
	if f.CallId != nil {
		toSerialize["callId"] = *f.CallId
	}
	return json.Marshal(toSerialize)
}

func (f *FailedAction) UnmarshalJSON(src []byte) error {
	var value map[string]interface{}
	if err := json.Unmarshal(src, &value); err != nil {
		return err
	}
	if v, ok := value["type"].(string); ok {
		f.Type = v
	}
	if v, ok := value["message"].(string); ok {
		f.Message = &v
	}
	if v, ok := value["callId"].(string); ok {
		f.CallId = &v
	}
	return nil
}

func (FailedAction) _internalComputerUseActionDtoActionOneOf() {}

func NewFailedAction() *FailedAction {
	return &FailedAction{
		Type: "failed",
	}
}
