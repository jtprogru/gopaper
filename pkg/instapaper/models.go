package instapaper

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/dghubble/oauth1"
)

type Bookmark struct {
	BookmarkID string `json:"bookmark_id"`
	Title      string `json:"title"`
	URL        string `json:"url"`
}

func (a *api) GetUnreadBookmarks() ([]Bookmark, error) {
	config := oauth1.NewConfig(a.consumerKey, a.consumerSecret)
	token := oauth1.NewToken(a.accessToken, a.accessSecret)
	httpClient := config.Client(oauth1.NoContext, token)

	response, err := httpClient.Get("https://www.instapaper.com/api/1/bookmarks/list?limit=20&folder_id=unread")
	if err != nil {
		return nil, fmt.Errorf("failed to fetch bookmarks: %w", err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch bookmarks: HTTP status %d", response.StatusCode)
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var bookmarks []Bookmark
	err = json.Unmarshal(body, &bookmarks)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal bookmarks: %w", err)
	}

	return bookmarks, nil
}
