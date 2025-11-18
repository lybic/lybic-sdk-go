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

func (m MobileTapAction) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{
		"type": "mobile:tap",
		"x":    m.X,
		"y":    m.Y,
	}
	if m.CallId != nil {
		toSerialize["callId"] = *m.CallId
	}
	return json.Marshal(toSerialize)
}

func (m *MobileTapAction) UnmarshalJSON(src []byte) error {
	var value map[string]interface{}
	if err := json.Unmarshal(src, &value); err != nil {
		return err
	}
	if v, ok := value["type"].(string); ok {
		m.Type = v
	}
	if v, ok := value["x"]; ok {
		x, err := unmarshalLength(v)
		if err != nil {
			return err
		}
		m.X = x
	}
	if v, ok := value["y"]; ok {
		y, err := unmarshalLength(v)
		if err != nil {
			return err
		}
		m.Y = y
	}
	if v, ok := value["callId"].(string); ok {
		m.CallId = &v
	}
	return nil
}

func (MobileTapAction) __internalMobileUseActionDtoActionOneOf() {}

func (m MobileDoubleTapAction) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{
		"type": "mobile:doubleTap",
		"x":    m.X,
		"y":    m.Y,
	}
	if m.CallId != nil {
		toSerialize["callId"] = *m.CallId
	}
	return json.Marshal(toSerialize)
}

func (m *MobileDoubleTapAction) UnmarshalJSON(src []byte) error {
	var value map[string]interface{}
	if err := json.Unmarshal(src, &value); err != nil {
		return err
	}
	if v, ok := value["type"].(string); ok {
		m.Type = v
	}
	if v, ok := value["x"]; ok {
		x, err := unmarshalLength(v)
		if err != nil {
			return err
		}
		m.X = x
	}
	if v, ok := value["y"]; ok {
		y, err := unmarshalLength(v)
		if err != nil {
			return err
		}
		m.Y = y
	}
	if v, ok := value["callId"].(string); ok {
		m.CallId = &v
	}
	return nil
}

func (MobileDoubleTapAction) __internalMobileUseActionDtoActionOneOf() {}

func (m MobileSwipeAction) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{
		"type":     "mobile:swipe",
		"startX":   m.StartX,
		"startY":   m.StartY,
		"endX":     m.EndX,
		"endY":     m.EndY,
		"duration": m.Duration,
	}
	if m.CallId != nil {
		toSerialize["callId"] = *m.CallId
	}
	return json.Marshal(toSerialize)
}

func (m *MobileSwipeAction) UnmarshalJSON(src []byte) error {
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
	if v, ok := value["duration"].(float64); ok {
		m.Duration = int(v)
	}
	if v, ok := value["callId"].(string); ok {
		m.CallId = &v
	}
	return nil
}

func (MobileSwipeAction) __internalMobileUseActionDtoActionOneOf() {}

func (m MobileTypeAction) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{
		"type":    "mobile:type",
		"content": m.Content,
	}
	if m.CallId != nil {
		toSerialize["callId"] = *m.CallId
	}
	return json.Marshal(toSerialize)
}

func (m *MobileTypeAction) UnmarshalJSON(src []byte) error {
	var value map[string]interface{}
	if err := json.Unmarshal(src, &value); err != nil {
		return err
	}
	if v, ok := value["type"].(string); ok {
		m.Type = v
	}
	if v, ok := value["content"].(string); ok {
		m.Content = v
	}
	if v, ok := value["callId"].(string); ok {
		m.CallId = &v
	}
	return nil
}

func (MobileTypeAction) __internalMobileUseActionDtoActionOneOf() {}

func (m MobileHotkeyAction) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{
		"type": "mobile:hotkey",
		"key":  m.Key,
	}
	if m.CallId != nil {
		toSerialize["callId"] = *m.CallId
	}
	return json.Marshal(toSerialize)
}

func (m *MobileHotkeyAction) UnmarshalJSON(src []byte) error {
	var value map[string]interface{}
	if err := json.Unmarshal(src, &value); err != nil {
		return err
	}
	if v, ok := value["type"].(string); ok {
		m.Type = v
	}
	if v, ok := value["key"].(string); ok {
		m.Key = v
	}
	if v, ok := value["callId"].(string); ok {
		m.CallId = &v
	}
	return nil
}

func (MobileHotkeyAction) __internalMobileUseActionDtoActionOneOf() {}

func (m MobileHomeAction) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{
		"type": "mobile:home",
	}
	if m.CallId != nil {
		toSerialize["callId"] = *m.CallId
	}
	return json.Marshal(toSerialize)
}

func (m *MobileHomeAction) UnmarshalJSON(src []byte) error {
	var value map[string]interface{}
	if err := json.Unmarshal(src, &value); err != nil {
		return err
	}
	if v, ok := value["type"].(string); ok {
		m.Type = v
	}
	if v, ok := value["callId"].(string); ok {
		m.CallId = &v
	}
	return nil
}

func (MobileHomeAction) __internalMobileUseActionDtoActionOneOf() {}

func (m MobileBackAction) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{
		"type": "mobile:back",
	}
	if m.CallId != nil {
		toSerialize["callId"] = *m.CallId
	}
	return json.Marshal(toSerialize)
}

func (m *MobileBackAction) UnmarshalJSON(src []byte) error {
	var value map[string]interface{}
	if err := json.Unmarshal(src, &value); err != nil {
		return err
	}
	if v, ok := value["type"].(string); ok {
		m.Type = v
	}
	if v, ok := value["callId"].(string); ok {
		m.CallId = &v
	}
	return nil
}

func (MobileBackAction) __internalMobileUseActionDtoActionOneOf() {}

func (m MobileScreenshotAction) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{
		"type": "mobile:screenshot",
	}
	if m.CallId != nil {
		toSerialize["callId"] = *m.CallId
	}
	return json.Marshal(toSerialize)
}

func (m *MobileScreenshotAction) UnmarshalJSON(src []byte) error {
	var value map[string]interface{}
	if err := json.Unmarshal(src, &value); err != nil {
		return err
	}
	if v, ok := value["type"].(string); ok {
		m.Type = v
	}
	if v, ok := value["callId"].(string); ok {
		m.CallId = &v
	}
	return nil
}

func (MobileScreenshotAction) __internalMobileUseActionDtoActionOneOf() {}

func (m MobileWaitAction) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{
		"type":     "mobile:wait",
		"duration": m.Duration,
	}
	if m.CallId != nil {
		toSerialize["callId"] = *m.CallId
	}
	return json.Marshal(toSerialize)
}

func (m *MobileWaitAction) UnmarshalJSON(src []byte) error {
	var value map[string]interface{}
	if err := json.Unmarshal(src, &value); err != nil {
		return err
	}
	if v, ok := value["type"].(string); ok {
		m.Type = v
	}
	if v, ok := value["duration"].(float64); ok {
		m.Duration = int(v)
	}
	if v, ok := value["callId"].(string); ok {
		m.CallId = &v
	}
	return nil
}

func (MobileWaitAction) __internalMobileUseActionDtoActionOneOf() {}

func (m MobileFinishedAction) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{
		"type": "mobile:finished",
	}
	if m.Message != nil {
		toSerialize["message"] = *m.Message
	}
	if m.CallId != nil {
		toSerialize["callId"] = *m.CallId
	}
	return json.Marshal(toSerialize)
}

func (m *MobileFinishedAction) UnmarshalJSON(src []byte) error {
	var value map[string]interface{}
	if err := json.Unmarshal(src, &value); err != nil {
		return err
	}
	if v, ok := value["type"].(string); ok {
		m.Type = v
	}
	if v, ok := value["message"].(string); ok {
		m.Message = &v
	}
	if v, ok := value["callId"].(string); ok {
		m.CallId = &v
	}
	return nil
}

func (MobileFinishedAction) __internalMobileUseActionDtoActionOneOf() {}

func (m MobileFailedAction) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{
		"type": "mobile:failed",
	}
	if m.Message != nil {
		toSerialize["message"] = *m.Message
	}
	if m.CallId != nil {
		toSerialize["callId"] = *m.CallId
	}
	return json.Marshal(toSerialize)
}

func (m *MobileFailedAction) UnmarshalJSON(src []byte) error {
	var value map[string]interface{}
	if err := json.Unmarshal(src, &value); err != nil {
		return err
	}
	if v, ok := value["type"].(string); ok {
		m.Type = v
	}
	if v, ok := value["message"].(string); ok {
		m.Message = &v
	}
	if v, ok := value["callId"].(string); ok {
		m.CallId = &v
	}
	return nil
}

func (MobileFailedAction) __internalMobileUseActionDtoActionOneOf() {}

func rawMessageToMobileUseActionDtoActionOneOf(rawAction json.RawMessage) (MobileUseActionDtoActionOneOf, error) {
	var base struct {
		Type string `json:"type"`
	}
	if err := json.Unmarshal(rawAction, &base); err != nil {
		return nil, fmt.Errorf("failed to unmarshal action type: %w", err)
	}

	var action MobileUseActionDtoActionOneOf
	switch base.Type {
	case "mobile:tap":
		action = &MobileTapAction{}
	case "mobile:doubleTap":
		action = &MobileDoubleTapAction{}
	case "mobile:swipe":
		action = &MobileSwipeAction{}
	case "mobile:type":
		action = &MobileTypeAction{}
	case "mobile:hotkey":
		action = &MobileHotkeyAction{}
	case "mobile:home":
		action = &MobileHomeAction{}
	case "mobile:back":
		action = &MobileBackAction{}
	case "mobile:screenshot":
		action = &MobileScreenshotAction{}
	case "mobile:wait":
		action = &MobileWaitAction{}
	case "mobile:finished":
		action = &MobileFinishedAction{}
	case "mobile:failed":
		action = &MobileFailedAction{}
	default:
		return nil, fmt.Errorf("unknown action type: %s", base.Type)
	}

	if err := json.Unmarshal(rawAction, action); err != nil {
		return nil, fmt.Errorf("failed to unmarshal action of type %s: %w",
			base.Type, err)
	}
	return action, nil
}

func (m *MobileUseActionResponseDto) UnmarshalJSON(data []byte) error {
	var temp struct {
		Actions  []json.RawMessage `json:"actions"`
		Unknown  *string           `json:"unknown,omitempty"`
		Memory   *string           `json:"memory,omitempty"`
		Thoughts *string           `json:"thoughts,omitempty"`
	}

	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}
	m.Actions = make([]MobileUseActionResponseDtoActionsOneOf, len(temp.Actions))
	for i, tempAction := range temp.Actions {
		action, err := rawMessageToMobileUseActionDtoActionOneOf(tempAction)
		if err != nil {
			return err
		}
		m.Actions[i] = MobileUseActionResponseDtoActionsOneOf{MobileUseActionResponseDtoActionsOneOfInterface: action}
	}
	if temp.Unknown != nil {
		m.Unknown = temp.Unknown
	}
	if temp.Memory != nil {
		m.Memory = temp.Memory
	}
	if temp.Thoughts != nil {
		m.Thoughts = temp.Thoughts
	}

	return nil
}

func (MobileTapAction) _internalSandboxUseActionDtoActionOneOf() {}

func (MobileDoubleTapAction) _internalSandboxUseActionDtoActionOneOf()  {}
func (MobileSwipeAction) _internalSandboxUseActionDtoActionOneOf()      {}
func (MobileTypeAction) _internalSandboxUseActionDtoActionOneOf()       {}
func (MobileHotkeyAction) _internalSandboxUseActionDtoActionOneOf()     {}
func (MobileHomeAction) _internalSandboxUseActionDtoActionOneOf()       {}
func (MobileFinishedAction) _internalSandboxUseActionDtoActionOneOf()   {}
func (MobileFailedAction) _internalSandboxUseActionDtoActionOneOf()     {}
func (MobileWaitAction) _internalSandboxUseActionDtoActionOneOf()       {}
func (MobileScreenshotAction) _internalSandboxUseActionDtoActionOneOf() {}
func (MobileBackAction) _internalSandboxUseActionDtoActionOneOf()       {}
