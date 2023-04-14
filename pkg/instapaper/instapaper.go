package instapaper

import (
	"errors"
	"net/http"
)

var (
	BaseUrl         = "https://www.instapaper.com"
	ApiVersion      = "1"
	AccessToken     = "oauth/access_token"
	LoginUrl        = "https://www.instapaper.com/user/login"
	RequestDelaySec = 0.5

	// General errors
	Error1040 = errors.New("Rate-limit exceeded")
	Error1041 = errors.New("Premium account required")
	Error1042 = errors.New("Application is suspended")

	// Bookmark errors
	Error1220 = errors.New("Domain requires full content to be supplied")
	Error1221 = errors.New("Domain has opted out of Instapaper compatibility")
	Error1240 = errors.New("Invalid URL specified")
	Error1241 = errors.New("Invalid or missing bookmark_id")
	Error1242 = errors.New("Invalid or missing folder_id")
	Error1243 = errors.New("Invalid or missing progress")
	Error1244 = errors.New("Invalid or missing progress_timestamp")
	Error1245 = errors.New("Private bookmarks require supplied content")
	Error1250 = errors.New("Unexpected error when saving bookmark")

	// Folder errors
	Error1251 = errors.New("User already has a folder with this title")
	Error1252 = errors.New("Cannot add bookmarks to this folder")

	// Operational errors
	Error1500 = errors.New("Unexpected service error")
	Error1550 = errors.New("Error generating text version of this URL")

	// Highlight Errors
	Error1600 = errors.New("Cannot create highlight with empty text")
	Error1601 = errors.New("Duplicate highlight")
)

type Api struct {
	Client *http.Client
}

func New() *Api {
	return &Api{}
}
