package stubs

// TokenResponse is the response when getting the access token
func TokenResponse() []byte {
	return []byte(`
{
   "access_token":"19077204-9d0a-31fa-85cf-xxxxxxxxxx",
   "scope":"am_application_scope default",
   "token_type":"Bearer",
   "expires_in":2496
}
`)
}
