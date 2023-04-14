package instapaper

type ResponseUser struct {
	Type     string `json:"type"`
	UserID   int    `json:"user_id"`
	Username string `json:"username"`
}

type ResponseBookmark struct {
	Type        string `json:"type"`
	BookmarkID  int    `json:"bookmark_id"`
	URL         string `json:"url"`
	Title       string `json:"title"`
	Description string `json:"description"`
}
