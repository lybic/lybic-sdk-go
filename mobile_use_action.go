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
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/lybic/lybic-sdk-go/pkg/json"
)

func (c *client) ParseMobileUseModelTextOutput(ctx context.Context, modelType string, dto ParseTextRequestDto) (*MobileUseActionResponseDto, error) {
	url := fmt.Sprintf("/api/mobile-use/parse/%s", modelType)
	c.config.Logger.Info("Sending request to parse mobile use action", "url:", url, "dto:", dto)
	resp, err := c.request(ctx, http.MethodPost, url, nil, dto)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		c.config.Logger.Errorf("failed to parse mobile use: %s", resp.Status)
		// Log the response body for debugging purposes
		body, _ := io.ReadAll(resp.Body)
		c.config.Logger.Errorf("Response body: %s", body)
		return nil, fmt.Errorf("failed to parse mobile use: %s", resp.Status)
	}

	var actions MobileUseActionResponseDto
	if err := json.NewDecoder(resp.Body).Decode(&actions); err != nil {
		c.config.Logger.Errorf("failed to decode response body: %v", err)
		return nil, err
	}

	return &actions, nil
}
