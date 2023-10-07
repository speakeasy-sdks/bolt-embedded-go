// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package boltembeddedgo

import (
	"bytes"
	"context"
	"fmt"
	"github.com/speakeasy-sdks/bolt-embedded-go/pkg/models/operations"
	"github.com/speakeasy-sdks/bolt-embedded-go/pkg/models/sdkerrors"
	"github.com/speakeasy-sdks/bolt-embedded-go/pkg/models/shared"
	"github.com/speakeasy-sdks/bolt-embedded-go/pkg/utils"
	"io"
	"net/http"
	"strings"
)

// oAuth - Interact with Shopper data by completing the Bolt OAuth process.
type oAuth struct {
	sdkConfiguration sdkConfiguration
}

func newOAuth(sdkConfig sdkConfiguration) *oAuth {
	return &oAuth{
		sdkConfiguration: sdkConfig,
	}
}

// OAuthToken - OAuth Token Endpoint
// Endpoint for receiving access, ID, and refresh tokens from Bolt's OAuth server.
//
// To use this endpoint, first use the Authorization Code Request flow by using the `authorization_code` Grant Type (`grant_type`). Then, in the event that you would need a second or subsequent code, use the `refresh_token` value returned from a successful request as the `refresh_token` input value in your subsequent `refresh_token` Grant Type (`grant_type`) request.
//
//	**Reminder - the Content-Type of this request must be application/x-www-form-urlencoded**
func (s *oAuth) OAuthToken(ctx context.Context, request operations.OAuthTokenRequest) (*operations.OAuthTokenResponse, error) {
	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	url := strings.TrimSuffix(baseURL, "/") + "/v1/oauth/token"

	bodyReader, reqContentType, err := utils.SerializeRequestBody(ctx, request, false, true, "RequestBody", "form", `request:"mediaType=application/x-www-form-urlencoded"`)
	if err != nil {
		return nil, fmt.Errorf("error serializing request body: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "POST", url, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("user-agent", s.sdkConfiguration.UserAgent)

	req.Header.Set("Content-Type", reqContentType)

	utils.PopulateHeaders(ctx, req, request)

	client := s.sdkConfiguration.SecurityClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.OAuthTokenResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}

	rawBody, err := io.ReadAll(httpRes.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}
	httpRes.Body.Close()
	httpRes.Body = io.NopCloser(bytes.NewBuffer(rawBody))
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out shared.OAuthTokenResponse
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			res.OAuthTokenResponse = &out
		default:
			return nil, sdkerrors.NewSDKError(fmt.Sprintf("unknown content-type received: %s", contentType), httpRes.StatusCode, string(rawBody), httpRes)
		}
	case httpRes.StatusCode == 400:
		fallthrough
	case httpRes.StatusCode == 403:
		fallthrough
	case httpRes.StatusCode == 422:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out shared.ErrorsOauthServerResponse
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			res.ErrorsOauthServerResponse = &out
		default:
			return nil, sdkerrors.NewSDKError(fmt.Sprintf("unknown content-type received: %s", contentType), httpRes.StatusCode, string(rawBody), httpRes)
		}
	}

	return res, nil
}
