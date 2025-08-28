// Copyright (c) 2019-2025   Beijing Tingyu Technology Co., Ltd.
// Copyright (c) 2025        Lybic Development Team <team@lybic.ai, lybic@tingyutech.com>
// Copyright (c) 2025        Lu Yicheng <luyicheng@tingyutech.com>
//
// These Terms of Service ("Terms") set forth the rules governing your access to and use of the website lybic.ai
// ("Website"), our web applications, and other services (collectively, the "Services") provided by Beijing Tingyu
// Technology Co., Ltd. ("Company," "we," "us," or "our"), a company registered in Haidian District, Beijing. Any
// breach of these Terms may result in the suspension or termination of your access to the Services.
// By accessing and using the Services and/or the Website, you represent that you are at least 18 years old,
// acknowledge that you have read and understood these Terms, and agree to be bound by them. By using or accessing
// the Services and/or the Website, you further represent and warrant that you have the legal capacity and authority
// to agree to these Terms, whether as an individual or on behalf of a company. If you do not agree to all of these
// Terms, do not access or use the Website or Services.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package lybic

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

func NewFinishedAction() *FinishedAction {
	return &FinishedAction{
		Type: "finished",
	}
}

type FailedAction struct {
	Type    string  `json:"type"` // set to failed
	Message *string `json:"message,omitempty"`
	CallId  *string `json:"callId,omitempty"`
}

func NewFailedAction() *FailedAction {
	return &FailedAction{
		Type: "failed",
	}
}
