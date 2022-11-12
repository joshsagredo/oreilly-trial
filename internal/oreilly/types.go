package oreilly

const (
	requestAuthority    = "learning.oreilly.com"
	requestPragma       = "no-cache"
	requestCacheControl = "no-cache"
	requestAccept       = "application/json"
	requestUserAgent    = "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.4324.96 Safari/537.36"
	requestContentType  = "application/json"
	requestOrigin       = "https://learning.oreilly.com"
	requestSecFetchSite = "same-origin"
	requestSecFetchMode = "cors"
	requestSecFetchDest = "empty"
	requestReferer      = "https://learning.oreilly.com/p/register/"
	requestAcceptLang   = "en-US,en;q=0.9"
)

// successResponse is the DTO for remote API call
type successResponse struct {
	UserID string `json:"user_id"`
}
