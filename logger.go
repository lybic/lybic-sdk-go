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

// Logger is an interface for logging messages at different levels.
//
//		This interface defines methods for logging debug, info, warning, and error messages.
//	 Support: zap（SugaredLogger）、logrus、zerolog、slog
type Logger interface {
	Debug(...interface{})
	Info(...interface{})
	Warn(...interface{})
	Error(...interface{})

	Debugf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Warnf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
}

// NewEmptyLogger returns a Logger that does not log anything.
type emptyLogger struct{}

func (emptyLogger) Debug(...interface{}) {}

func (emptyLogger) Info(...interface{}) {}

func (emptyLogger) Warn(...interface{}) {}

func (emptyLogger) Error(...interface{}) {}

func (emptyLogger) Debugf(format string, args ...interface{}) {}

func (emptyLogger) Infof(format string, args ...interface{}) {}

func (emptyLogger) Warnf(format string, args ...interface{}) {}

func (emptyLogger) Errorf(format string, args ...interface{}) {}
