// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package boltembeddedgo

import (
	"context"
	"fmt"
	"github.com/speakeasy-sdks/bolt-embedded-go/pkg/models/shared"
	"github.com/speakeasy-sdks/bolt-embedded-go/pkg/utils"
	"net/http"
	"time"
)

// ServerList contains the list of servers available to the SDK
var ServerList = []string{
	// The Production URL (Live Data).
	"https://api.bolt.com",
	// The Sandbox URL (Test Data).
	"https://api-sandbox.bolt.com",
	// The Staging URL (Staged Data).
	"https://api-staging.bolt.com",
}

// HTTPClient provides an interface for suplying the SDK with a custom HTTP client
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// String provides a helper function to return a pointer to a string
func String(s string) *string { return &s }

// Bool provides a helper function to return a pointer to a bool
func Bool(b bool) *bool { return &b }

// Int provides a helper function to return a pointer to an int
func Int(i int) *int { return &i }

// Int64 provides a helper function to return a pointer to an int64
func Int64(i int64) *int64 { return &i }

// Float32 provides a helper function to return a pointer to a float32
func Float32(f float32) *float32 { return &f }

// Float64 provides a helper function to return a pointer to a float64
func Float64(f float64) *float64 { return &f }

type sdkConfiguration struct {
	DefaultClient     HTTPClient
	SecurityClient    HTTPClient
	Security          func(context.Context) (interface{}, error)
	ServerURL         string
	ServerIndex       int
	Language          string
	OpenAPIDocVersion string
	SDKVersion        string
	GenVersion        string
	UserAgent         string
	RetryConfig       *utils.RetryConfig
}

func (c *sdkConfiguration) GetServerDetails() (string, map[string]string) {
	if c.ServerURL != "" {
		return c.ServerURL, nil
	}

	return ServerList[c.ServerIndex], nil
}

// BoltEmbeddedAPI - Embedded API Reference: Postman Collection:
//
// [![](https://run.pstmn.io/button.svg)](https://god.gw.postman.com/run-collection/9136127-55d2bde1-a248-473f-95b5-64cfd02fb445?action=collection%2Ffork&collection-url=entityId%3D9136127-55d2bde1-a248-473f-95b5-64cfd02fb445%26entityType%3Dcollection%26workspaceId%3D78beee89-4238-4c5f-bd1f-7e98978744b4#?env%5BBolt%20Sandbox%20Environment%5D=W3sia2V5IjoiYXBpX2Jhc2VfdXJsIiwidmFsdWUiOiJodHRwczovL2FwaS1zYW5kYm94LmJvbHQuY29tIiwidHlwZSI6ImRlZmF1bHQiLCJlbmFibGVkIjp0cnVlfSx7ImtleSI6InRrX2Jhc2UiLCJ2YWx1ZSI6Imh0dHBzOi8vc2FuZGJveC5ib2x0dGsuY29tIiwidHlwZSI6ImRlZmF1bHQiLCJlbmFibGVkIjp0cnVlfSx7ImtleSI6ImFwaV9rZXkiLCJ2YWx1ZSI6IjxyZXBsYWNlIHdpdGggeW91ciBCb2x0IFNhbmRib3ggQVBJIGtleT4iLCJ0eXBlIjoic2VjcmV0IiwiZW5hYmxlZCI6dHJ1ZX0seyJrZXkiOiJwdWJsaXNoYWJsZV9rZXkiLCJ2YWx1ZSI6IjxyZXBsYWNlIHdpdGggeW91ciBCb2x0IFNhbmRib3ggcHVibGlzaGFibGUga2V5PiIsInR5cGUiOiJkZWZhdWx0IiwiZW5hYmxlZCI6dHJ1ZX0seyJrZXkiOiJkaXZpc2lvbl9pZCIsInZhbHVlIjoiPHJlcGxhY2Ugd2l0aCB5b3VyIEJvbHQgU2FuZGJveCBwdWJsaWMgZGl2aXNpb24gSUQ+IiwidHlwZSI6ImRlZmF1bHQiLCJlbmFibGVkIjp0cnVlfV0=)
//
// ## About
// The Embedded API reference is a consolidation of critical APIs that a developer will use when integrating with Bolt's Embedded Accounts product suite.
type BoltEmbeddedAPI struct {
	// Create Embedded Accounts user flows for logged-in and guest experiences by interacting with and updating shopper data.
	//
	Account *account
	// Interact with Shopper data by completing the Bolt OAuth process.
	//
	OAuth *oAuth
	// Create and manage transactions for non credit card payments such as Paypal in your Embedded Accounts experience.
	//
	Payments *payments
	// A collection of endpoints that provide useful functionality to assist in testing your Bolt integration.
	//
	Testing *testing
	// Authorize credit card transactions and perform operations on those transactions with Bolt's transaction API.
	//
	Transactions *transactions

	sdkConfiguration sdkConfiguration
}

type SDKOption func(*BoltEmbeddedAPI)

// WithServerURL allows the overriding of the default server URL
func WithServerURL(serverURL string) SDKOption {
	return func(sdk *BoltEmbeddedAPI) {
		sdk.sdkConfiguration.ServerURL = serverURL
	}
}

// WithTemplatedServerURL allows the overriding of the default server URL with a templated URL populated with the provided parameters
func WithTemplatedServerURL(serverURL string, params map[string]string) SDKOption {
	return func(sdk *BoltEmbeddedAPI) {
		if params != nil {
			serverURL = utils.ReplaceParameters(serverURL, params)
		}

		sdk.sdkConfiguration.ServerURL = serverURL
	}
}

// WithServerIndex allows the overriding of the default server by index
func WithServerIndex(serverIndex int) SDKOption {
	return func(sdk *BoltEmbeddedAPI) {
		if serverIndex < 0 || serverIndex >= len(ServerList) {
			panic(fmt.Errorf("server index %d out of range", serverIndex))
		}

		sdk.sdkConfiguration.ServerIndex = serverIndex
	}
}

// WithClient allows the overriding of the default HTTP client used by the SDK
func WithClient(client HTTPClient) SDKOption {
	return func(sdk *BoltEmbeddedAPI) {
		sdk.sdkConfiguration.DefaultClient = client
	}
}

func withSecurity(security interface{}) func(context.Context) (interface{}, error) {
	return func(context.Context) (interface{}, error) {
		return &security, nil
	}
}

// WithSecurity configures the SDK to use the provided security details

func WithSecurity(security shared.Security) SDKOption {
	return func(sdk *BoltEmbeddedAPI) {
		sdk.sdkConfiguration.Security = withSecurity(security)
	}
}

// WithSecuritySource configures the SDK to invoke the Security Source function on each method call to determine authentication
func WithSecuritySource(security func(context.Context) (shared.Security, error)) SDKOption {
	return func(sdk *BoltEmbeddedAPI) {
		sdk.sdkConfiguration.Security = func(ctx context.Context) (interface{}, error) {
			return security(ctx)
		}
	}
}

func WithRetryConfig(retryConfig utils.RetryConfig) SDKOption {
	return func(sdk *BoltEmbeddedAPI) {
		sdk.sdkConfiguration.RetryConfig = &retryConfig
	}
}

// New creates a new instance of the SDK with the provided options
func New(opts ...SDKOption) *BoltEmbeddedAPI {
	sdk := &BoltEmbeddedAPI{
		sdkConfiguration: sdkConfiguration{
			Language:          "go",
			OpenAPIDocVersion: "1.0.1",
			SDKVersion:        "0.6.1",
			GenVersion:        "2.150.0",
			UserAgent:         "speakeasy-sdk/go 0.6.1 2.150.0 1.0.1 github.com/speakeasy-sdks/bolt-embedded-go",
		},
	}
	for _, opt := range opts {
		opt(sdk)
	}

	// Use WithClient to override the default client if you would like to customize the timeout
	if sdk.sdkConfiguration.DefaultClient == nil {
		sdk.sdkConfiguration.DefaultClient = &http.Client{Timeout: 60 * time.Second}
	}
	if sdk.sdkConfiguration.SecurityClient == nil {
		if sdk.sdkConfiguration.Security != nil {
			sdk.sdkConfiguration.SecurityClient = utils.ConfigureSecurityClient(sdk.sdkConfiguration.DefaultClient, sdk.sdkConfiguration.Security)
		} else {
			sdk.sdkConfiguration.SecurityClient = sdk.sdkConfiguration.DefaultClient
		}
	}

	sdk.Account = newAccount(sdk.sdkConfiguration)

	sdk.OAuth = newOAuth(sdk.sdkConfiguration)

	sdk.Payments = newPayments(sdk.sdkConfiguration)

	sdk.Testing = newTesting(sdk.sdkConfiguration)

	sdk.Transactions = newTransactions(sdk.sdkConfiguration)

	return sdk
}
