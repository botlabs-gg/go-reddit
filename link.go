package reddit

import (
	"encoding/json"
	"fmt"
)

// Link contains information about a link.
type Link struct {
	ApprovedBy          string        `json:"approved_by"`
	Archived            bool          `json:"archived"`
	Author              string        `json:"author"`
	AuthorFlairCSSClass string        `json:"author_flair_css_class"`
	AuthorFlairText     string        `json:"author_flair_text"`
	BannedBy            string        `json:"banned_by"`
	Clicked             bool          `json:"clicked"`
	ContestMode         bool          `json:"contest_mode"`
	Created             int           `json:"created"`
	CreatedUtc          int           `json:"created_utc"`
	Distinguished       string        `json:"distinguished"`
	Domain              string        `json:"domain"`
	Downs               int           `json:"downs"`
	Edited              bool          `json:"edited"`
	Gilded              int           `json:"gilded"`
	Hidden              bool          `json:"hidden"`
	HideScore           bool          `json:"hide_score"`
	ID                  string        `json:"id"`
	IsSelf              bool          `json:"is_self"`
	Likes               bool          `json:"likes"`
	LinkFlairCSSClass   string        `json:"link_flair_css_class"`
	LinkFlairText       string        `json:"link_flair_text"`
	Locked              bool          `json:"locked"`
	Media               Media         `json:"media"`
	MediaEmbed          interface{}   `json:"media_embed"`
	ModReports          []interface{} `json:"mod_reports"`
	Name                string        `json:"name"`
	NumComments         int           `json:"num_comments"`
	NumReports          int           `json:"num_reports"`
	Over18              bool          `json:"over_18"`
	Permalink           string        `json:"permalink"`
	Quarantine          bool          `json:"quarantine"`
	RemovalReason       interface{}   `json:"removal_reason"`
	ReportReasons       []interface{} `json:"report_reasons"`
	Saved               bool          `json:"saved"`
	Score               int           `json:"score"`
	SecureMedia         interface{}   `json:"secure_media"`
	SecureMediaEmbed    interface{}   `json:"secure_media_embed"`
	SelftextHTML        string        `json:"selftext_html"`
	Selftext            string        `json:"selftext"`
	Stickied            bool          `json:"stickied"`
	Subreddit           string        `json:"subreddit"`
	SubredditID         string        `json:"subreddit_id"`
	SuggestedSort       string        `json:"suggested_sort"`
	Thumbnail           string        `json:"thumbnail"`
	Title               string        `json:"title"`
	URL                 string        `json:"url"`
	Ups                 int           `json:"ups"`
	UserReports         []interface{} `json:"user_reports"`
	Visited             bool          `json:"visited"`
}

type linkListing struct {
	Kind string `json:"kind"`
	Data struct {
		Modhash  string `json:"modhash"`
		Children []struct {
			Kind string `json:"kind"`
			Data Link   `json:"data"`
		} `json:"children"`
		After  string      `json:"after"`
		Before interface{} `json:"before"`
	} `json:"data"`
}

// DeleteLink deletes a link submitted by the currently authenticated user. Requires the 'edit' OAuth scope.
func (c *Client) DeleteLink(linkID string) error {
  return c.deleteThing(fmt.Sprintf("t3_%s", linkID))
}

// GetHotLinks retrieves a listing of hot links.
func (c *Client) GetHotLinks(subreddit string) ([]*Link, error) {
	return c.getLinks(subreddit, "hot")
}

// GetNewLinks retrieves a listing of new links.
func (c *Client) GetNewLinks(subreddit string) ([]*Link, error) {
	return c.getLinks(subreddit, "new")
}

// GetTopLinks retrieves a listing of top links.
func (c *Client) GetTopLinks(subreddit string) ([]*Link, error) {
	return c.getLinks(subreddit, "top")
}

func (c *Client) getLinks(subreddit string, sort string) ([]*Link, error) {
	url := fmt.Sprintf("%s/r/%s/%s.json", baseURL, subreddit, sort)
	resp, err := c.http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result linkListing
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	var links []*Link
	for _, link := range result.Data.Children {
		links = append(links, &link.Data)
	}

	return links, nil
}
