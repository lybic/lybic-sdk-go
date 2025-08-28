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

// Package json provides JSON encoding and decoding functionalities.
//
//	If you need a high-performance json processing library, you can manually set the following variables
//	example1: use jsoniter (github.com/json-iterator/go) instead of the standard library.
//	json.Marshal = jsoniter.Marshal
//	json.Unmarshal = jsoniter.Unmarshal
//	json.MarshalIndent = jsoniter.MarshalIndent
//	json.NewDecoder = jsoniter.NewDecoder
//	json.NewEncoder = jsoniter.NewEncoder
//	example2: use sonic (github.com/bytedance/sonic) instead of the standard library.
//	so = sonic.ConfigStd
//	json.Marshal = so.Marshal
//	json.Unmarshal = so.Unmarshal
//	json.MarshalIndent = so.MarshalIndent
//	json.NewDecoder = so.NewDecoder
//	json.NewEncoder = so.NewEncoder
package json

import "encoding/json"

var (
	// Marshal is exported by pkg/json package.
	Marshal = json.Marshal
	// Unmarshal is exported by pkg/json package.
	Unmarshal = json.Unmarshal
	// MarshalIndent is exported by pkg/json package.
	MarshalIndent = json.MarshalIndent
	// NewDecoder is exported by pkg/json package.
	NewDecoder = json.NewDecoder
	// NewEncoder is exported by pkg/json package.
	NewEncoder = json.NewEncoder
)

type (
	RawMessage = json.RawMessage
)
