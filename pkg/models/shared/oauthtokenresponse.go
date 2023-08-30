// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// OAuthTokenResponse - OAuth token response.
type OAuthTokenResponse struct {
	// An access token you can use to make requests on behalf of a Bolt Account.
	AccessToken *string `json:"access_token,omitempty"`
	// Access token’s expiration in seconds.
	ExpiresIn *int64 `json:"expires_in,omitempty"`
	// A JWT token issued when the request includes the scope open_id.
	IDToken *string `json:"id_token,omitempty"`
	// A refresh token you can use to issue a brand new access token without obtaining a new authorization code.
	RefreshToken *string `json:"refresh_token,omitempty"`
	// The scope granted to the refresh token. Currently this will always be bolt.account.view.
	RefreshTokenScope *string `json:"refresh_token_scope,omitempty"`
	// The scope granted to access token, depending on the scope granted to the authorization code as well as the scope parameter. Options include `bolt.account.manage`, `bolt.account.view`, `openid`.
	Scope *string `json:"scope,omitempty"`
	// The token_type will always be bearer.
	TokenType *string `json:"token_type,omitempty"`
}

func (o *OAuthTokenResponse) GetAccessToken() *string {
	if o == nil {
		return nil
	}
	return o.AccessToken
}

func (o *OAuthTokenResponse) GetExpiresIn() *int64 {
	if o == nil {
		return nil
	}
	return o.ExpiresIn
}

func (o *OAuthTokenResponse) GetIDToken() *string {
	if o == nil {
		return nil
	}
	return o.IDToken
}

func (o *OAuthTokenResponse) GetRefreshToken() *string {
	if o == nil {
		return nil
	}
	return o.RefreshToken
}

func (o *OAuthTokenResponse) GetRefreshTokenScope() *string {
	if o == nil {
		return nil
	}
	return o.RefreshTokenScope
}

func (o *OAuthTokenResponse) GetScope() *string {
	if o == nil {
		return nil
	}
	return o.Scope
}

func (o *OAuthTokenResponse) GetTokenType() *string {
	if o == nil {
		return nil
	}
	return o.TokenType
}
