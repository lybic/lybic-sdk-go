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

import (
	"github.com/lybic/lybic-sdk-go/pkg/json"
)

// MobileUseActionDtoActionOneOf is an interface for types that can be used as lybic mobile-use action
//
// Implementations:
//
// # MobileTapAction
//
// # MobileDoubleTapAction
//
// # MobileSwipeAction
//
// # MobileTypeAction
//
// # MobileHotkeyAction
//
// # MobileHomeAction
//
// # MobileBackAction
//
// # MobileScreenshotAction
//
// # MobileWaitAction
//
// # MobileFinishedAction
//
// # MobileFailedAction
type MobileUseActionDtoActionOneOf interface {
	MarshalJSON() ([]byte, error)
	UnmarshalJSON(src []byte) error

	// __internalMobileUseActionDtoActionOneOf is a dummy method to prevent external implementations
	__internalMobileUseActionDtoActionOneOf()
	// _internalSandboxUseActionDtoActionOneOf is a dummy method to prevent external implementations
	_internalSandboxUseActionDtoActionOneOf()
}

func TryUnmarshalToMobileUseActionDtoActionOneOf(data map[string]any) (MobileUseActionDtoActionOneOf, error) {
	marshal, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	return rawMessageToMobileUseActionDtoActionOneOf(marshal)
}

type MobileTapAction struct {
	Type   string  `json:"type"` // set to mobile:tap
	X      Length  `json:"x"`
	Y      Length  `json:"y"`
	CallId *string `json:"callId,omitempty"`
}

func NewMobileTapAction(x Length, y Length) *MobileTapAction {
	return &MobileTapAction{
		Type: "mobile:tap",
		X:    x,
		Y:    y,
	}
}

type MobileDoubleTapAction struct {
	Type   string  `json:"type"` // set to mobile:doubleTap
	X      Length  `json:"x"`
	Y      Length  `json:"y"`
	CallId *string `json:"callId,omitempty"`
}

func NewMobileDoubleTapAction(x Length, y Length) *MobileDoubleTapAction {
	return &MobileDoubleTapAction{
		Type: "mobile:doubleTap",
		X:    x,
		Y:    y,
	}
}

type MobileSwipeAction struct {
	Type     string  `json:"type"` // set to mobile:swipe
	StartX   Length  `json:"startX"`
	StartY   Length  `json:"startY"`
	EndX     Length  `json:"endX"`
	EndY     Length  `json:"endY"`
	Duration int     `json:"duration"`
	CallId   *string `json:"callId,omitempty"`
}

func NewMobileSwipeAction(startX, startY, endX, endY Length, duration int) *MobileSwipeAction {
	return &MobileSwipeAction{
		Type:     "mobile:swipe",
		StartX:   startX,
		StartY:   startY,
		EndX:     endX,
		EndY:     endY,
		Duration: duration,
	}
}

type MobileTypeAction struct {
	Type    string  `json:"type"` // set to mobile:type
	Content string  `json:"content"`
	CallId  *string `json:"callId,omitempty"`
}

func NewMobileTypeAction(content string) *MobileTypeAction {
	return &MobileTypeAction{
		Type:    "mobile:type",
		Content: content,
	}
}

type MobileHotkeyAction struct {
	Type   string  `json:"type"` // set to mobile:hotkey
	Key    string  `json:"key"`
	CallId *string `json:"callId,omitempty"`
}

func NewMobileHotkeyAction(key string) *MobileHotkeyAction {
	return &MobileHotkeyAction{
		Type: "mobile:hotkey",
		Key:  key,
	}
}

type MobileHomeAction struct {
	Type   string  `json:"type"` // set to mobile:home
	CallId *string `json:"callId,omitempty"`
}

func NewMobileHomeAction() *MobileHomeAction {
	return &MobileHomeAction{
		Type: "mobile:home",
	}
}

type MobileBackAction struct {
	Type   string  `json:"type"` // set to mobile:back
	CallId *string `json:"callId,omitempty"`
}

func NewMobileBackAction() *MobileBackAction {
	return &MobileBackAction{
		Type: "mobile:back",
	}
}

type MobileScreenshotAction struct {
	Type   string  `json:"type"` // set to mobile:screenshot
	CallId *string `json:"callId,omitempty"`
}

func NewMobileScreenshotAction() *MobileScreenshotAction {
	return &MobileScreenshotAction{
		Type: "mobile:screenshot",
	}
}

type MobileWaitAction struct {
	Type     string  `json:"type"` // set to mobile:wait
	Duration int     `json:"duration"`
	CallId   *string `json:"callId,omitempty"`
}

func NewMobileWaitAction(duration int) *MobileWaitAction {
	return &MobileWaitAction{
		Type:     "mobile:wait",
		Duration: duration,
	}
}

type MobileFinishedAction struct {
	Type    string  `json:"type"` // set to mobile:finished
	Message *string `json:"message,omitempty"`
	CallId  *string `json:"callId,omitempty"`
}

func NewMobileFinishedAction() *MobileFinishedAction {
	return &MobileFinishedAction{
		Type: "mobile:finished",
	}
}

type MobileFailedAction struct {
	Type    string  `json:"type"` // set to mobile:failed
	Message *string `json:"message,omitempty"`
	CallId  *string `json:"callId,omitempty"`
}

func NewMobileFailedAction() *MobileFailedAction {
	return &MobileFailedAction{
		Type: "mobile:failed",
	}
}
