# OAuth
(*OAuth*)

## Overview

Interact with Shopper data by completing the Bolt OAuth process.


### Available Operations

* [OAuthToken](#oauthtoken) - OAuth Token Endpoint

## OAuthToken

Endpoint for receiving access, ID, and refresh tokens from Bolt's OAuth server. 

To use this endpoint, first use the Authorization Code Request flow by using the `authorization_code` Grant Type (`grant_type`). Then, in the event that you would need a second or subsequent code, use the `refresh_token` value returned from a successful request as the `refresh_token` input value in your subsequent `refresh_token` Grant Type (`grant_type`) request.

 **Reminder - the Content-Type of this request must be application/x-www-form-urlencoded**


### Example Usage

```go
package main

import(
	"context"
	"log"
	boltembeddedgo "github.com/speakeasy-sdks/bolt-embedded-go"
	"github.com/speakeasy-sdks/bolt-embedded-go/pkg/models/shared"
	"github.com/speakeasy-sdks/bolt-embedded-go/pkg/models/operations"
)

func main() {
    s := boltembeddedgo.New(
        boltembeddedgo.WithSecurity(shared.Security{
            OAuth: boltembeddedgo.String(""),
        }),
    )

    ctx := context.Background()
    res, err := s.OAuth.OAuthToken(ctx, operations.OAuthTokenRequest{
        RequestBody: operations.CreateOAuthTokenRequestBodyOAuthTokenInput(
                shared.OAuthTokenInput{
                    ClientID: "marvelous East female",
                    ClientSecret: "systems",
                    Code: "Northwest Pants Strategist",
                    GrantType: shared.OAuthTokenInputGrantTypeAuthorizationCode,
                    Scope: shared.OAuthTokenInputScopeOpenid,
                },
        ),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.OAuthTokenResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `request`                                                                    | [operations.OAuthTokenRequest](../../models/operations/oauthtokenrequest.md) | :heavy_check_mark:                                                           | The request object to use for the request.                                   |


### Response

**[*operations.OAuthTokenResponse](../../models/operations/oauthtokenresponse.md), error**

