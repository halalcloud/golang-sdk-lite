package oauth

type AuthorizeRequest struct {
	ClientId            string `protobuf:"bytes,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`                                    // Client ID
	ResponseType        string `protobuf:"bytes,2,opt,name=response_type,json=responseType,proto3" json:"response_type,omitempty"`                        // Response type, e.g., "code"
	RedirectUri         string `protobuf:"bytes,3,opt,name=redirect_uri,json=redirectUri,proto3" json:"redirect_uri,omitempty"`                           // Redirect URI
	Scope               string `protobuf:"bytes,4,opt,name=scope,proto3" json:"scope,omitempty"`                                                          // Scopes requested
	State               string `protobuf:"bytes,5,opt,name=state,proto3" json:"state,omitempty"`                                                          // State parameter for CSRF protection
	CodeChallenge       string `protobuf:"bytes,6,opt,name=code_challenge,json=codeChallenge,proto3" json:"code_challenge,omitempty"`                     // Code challenge for PKCE
	CodeChallengeMethod string `protobuf:"bytes,7,opt,name=code_challenge_method,json=codeChallengeMethod,proto3" json:"code_challenge_method,omitempty"` // Method for code challenge, e.g., "S256"
	Version             string `protobuf:"bytes,8,opt,name=version,proto3" json:"version,omitempty"`                                                      // Version of the OAuth protocol, e.g., "2.0"
	Nonce               string `protobuf:"bytes,9,opt,name=nonce,proto3" json:"nonce,omitempty"`                                                          // Nonce for additional security, used in OpenID Connect
	Device              string `protobuf:"bytes,10,opt,name=device,proto3" json:"device,omitempty"`                                                       // Device information, used for device authorization
	Mode                string `protobuf:"bytes,11,opt,name=mode,proto3" json:"mode,omitempty"`                                                           // Country code for the user
	DebugHost           string `protobuf:"bytes,12,opt,name=debug_host,json=debugHost,proto3" json:"debug_host,omitempty"`                                // Debug host for testing purposes
}

type AuthorizeResponse struct {
	Code        string `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`                                  // Authorization code
	RedirectUri string `protobuf:"bytes,2,opt,name=redirect_uri,json=redirectUri,proto3" json:"redirect_uri,omitempty"` // Redirect URI to send the user back to
	State       string `protobuf:"bytes,3,opt,name=state,proto3" json:"state,omitempty"`                                // State parameter to match the request
	Status      string `protobuf:"bytes,4,opt,name=status,proto3" json:"status,omitempty"`                              // Status of the authorization request
}

type TokenRequest struct {
	Code         string `protobuf:"bytes,3,opt,name=code,proto3" json:"code,omitempty"`                                     // Authorization code received from the authorization endpoint
	RedirectUri  string `protobuf:"bytes,4,opt,name=redirect_uri,json=redirectUri,proto3" json:"redirect_uri,omitempty"`    // Redirect URI used in the authorization request
	GrantType    string `protobuf:"bytes,5,opt,name=grant_type,json=grantType,proto3" json:"grant_type,omitempty"`          // Grant type, e.g., "authorization_code"
	CodeVerifier string `protobuf:"bytes,6,opt,name=code_verifier,json=codeVerifier,proto3" json:"code_verifier,omitempty"` // Code verifier for PKCE
}

type TokenResponse struct {
	AccessToken  string `protobuf:"bytes,1,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`    // Access token
	RefreshToken string `protobuf:"bytes,2,opt,name=refresh_token,json=refreshToken,proto3" json:"refresh_token,omitempty"` // Refresh token
	ExpiresIn    string `protobuf:"varint,3,opt,name=expires_in,json=expiresIn,proto3" json:"expires_in,omitempty"`         // Token expiration time in seconds
	TokenType    string `protobuf:"bytes,4,opt,name=token_type,json=tokenType,proto3" json:"token_type,omitempty"`          // Type of the token, e.g., "Bearer"
	Scope        string `protobuf:"bytes,5,opt,name=scope,proto3" json:"scope,omitempty"`                                   // Scopes granted
	ExpiresInTs  string `protobuf:"varint,6,opt,name=expires_in_ts,json=expiresInTs,proto3" json:"expires_in_ts,omitempty"` // string identity = 6; // User identity associated with the token
}
