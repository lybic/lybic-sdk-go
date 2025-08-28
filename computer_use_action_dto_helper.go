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
	"fmt"

	"github.com/lybic/lybic-sdk-go/pkg/json"
)

func unmarshalLength(data interface{}) (Length, error) {
	lengthMap, ok := data.(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("expected a map for Length, got %T", data)
	}

	typeVal, ok := lengthMap["type"].(string)
	if !ok {
		return nil, fmt.Errorf("missing 'type' in Length object")
	}

	raw, err := json.Marshal(lengthMap)
	if err != nil {
		return nil, err
	}

	switch typeVal {
	case "/":
		var fracLen FractionalLength
		if err := json.Unmarshal(raw, &fracLen); err != nil {
			return nil, err
		}
		return &fracLen, nil
	case "px":
		var pixLen PixelLength
		if err := json.Unmarshal(raw, &pixLen); err != nil {
			return nil, err
		}
		return &pixLen, nil
	default:
		return nil, fmt.Errorf("unknown Length type: %s", typeVal)
	}
}

func unmarshalMousePointActionBase(value map[string]interface{}) (*mousePointActionBase, error) {
	base := &mousePointActionBase{}
	var err error
	if v, ok := value["x"]; ok {
		base.X, err = unmarshalLength(v)
		if err != nil {
			return nil, err
		}
	}
	if v, ok := value["y"]; ok {
		base.Y, err = unmarshalLength(v)
		if err != nil {
			return nil, err
		}
	}
	if v, ok := value["holdKey"].(string); ok {
		base.HoldKey = &v
	}
	if v, ok := value["callId"].(string); ok {
		base.CallId = &v
	}
	return base, nil
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
	// This method is required to satisfy the Length interface.
	// Unmarshalling is handled by the parent struct's UnmarshalJSON.
	// todo: return nil?

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
	base, err := unmarshalMousePointActionBase(value)
	if err != nil {
		return err
	}
	m.X = base.X
	m.Y = base.Y
	m.HoldKey = base.HoldKey
	m.CallId = base.CallId
	if v, ok := value["button"].(float64); ok {
		m.Button = int(v)
	}
	return nil
}

func (MouseClickAction) _internalComputerUseActionDtoActionOneOf() {}

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
	base, err := unmarshalMousePointActionBase(value)
	if err != nil {
		return err
	}
	m.X = base.X
	m.Y = base.Y
	m.HoldKey = base.HoldKey
	m.CallId = base.CallId
	if v, ok := value["button"].(float64); ok {
		m.Button = int(v)
	}
	return nil
}

func (MouseDoubleClickAction) _internalComputerUseActionDtoActionOneOf() {}

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
	base, err := unmarshalMousePointActionBase(value)
	if err != nil {
		return err
	}
	m.X = base.X
	m.Y = base.Y
	m.HoldKey = base.HoldKey
	m.CallId = base.CallId
	return nil
}

func (MouseMoveAction) _internalComputerUseActionDtoActionOneOf() {}

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
	base, err := unmarshalMousePointActionBase(value)
	if err != nil {
		return err
	}
	m.X = base.X
	m.Y = base.Y
	m.HoldKey = base.HoldKey
	m.CallId = base.CallId
	if v, ok := value["stepVertical"].(float64); ok {
		m.StepVertical = int(v)
	}
	if v, ok := value["stepHorizontal"].(float64); ok {
		m.StepHorizontal = int(v)
	}
	return nil
}

func (MouseScrollAction) _internalComputerUseActionDtoActionOneOf() {}

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
	if v, ok := value["startX"]; ok {
		startX, err := unmarshalLength(v)
		if err != nil {
			return err
		}
		m.StartX = startX
	}
	if v, ok := value["startY"]; ok {
		startY, err := unmarshalLength(v)
		if err != nil {
			return err
		}
		m.StartY = startY
	}
	if v, ok := value["endX"]; ok {
		endX, err := unmarshalLength(v)
		if err != nil {
			return err
		}
		m.EndX = endX
	}
	if v, ok := value["endY"]; ok {
		endY, err := unmarshalLength(v)
		if err != nil {
			return err
		}
		m.EndY = endY
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
func (FinishedAction) _internalComputerUseActionDtoActionOneOf() {}

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
func rawMessageToComputerUseActionDtoActionOneOf(rawAction json.RawMessage) (ComputerUseActionDtoActionOneOf, error) {
	var base struct {
		Type string `json:"type"`
	}
	if err := json.Unmarshal(rawAction, &base); err != nil {
		return nil, fmt.Errorf("failed to unmarshal action type: %w", err)
	}

	var action ComputerUseActionDtoActionOneOf
	switch base.Type {
	case "mouse:click":
		action = &MouseClickAction{}
	case "mouse:doubleClick":
		action = &MouseDoubleClickAction{}
	case "mouse:move":
		action = &MouseMoveAction{}
	case "mouse:scroll":
		action = &MouseScrollAction{}
	case "mouse:drag":
		action = &MouseDragAction{}
	case "keyboard:type":
		action = &KeyboardTypeAction{}
	case "keyboard:hotkey":
		action = &KeyboardHotkeyAction{}
	case "screenshot":
		action = &ScreenshotAction{}
	case "wait":
		action = &WaitAction{}
	case "finished":
		action = &FinishedAction{}
	case "failed":
		action = &FailedAction{}
	default:
		return nil, fmt.Errorf("unknown action type: %s", base.Type)
	}

	if err := json.Unmarshal(rawAction, action); err != nil {
		return nil, fmt.Errorf("failed to unmarshal action of type %s: %w",
			base.Type, err)
	}
	return action, nil
}
