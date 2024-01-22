package pinterest

/*
	Search API
*/

type SearchResource Resource

func newSearchResource(cli *Client) *SearchResource {
	return &SearchResource{Cli: cli}
}

// SearchResourceOptions specifies the optional parameters to various search methods
type SearchResourceOptions struct {
	Bookmark string `url:"bookmark,omitempty"`
	PageSize int    `url:"page_size,omitempty"`
	Query    string `url:"query,omitempty"`
}

// SearchUsersBoards Search for boards for the "operation user_account". This includes boards of all board types.
// Refer: https://developers.pinterest.com/docs/api/v5/#operation/search_user_boards/get
func (r *SearchResource) SearchUsersBoards(args SearchResourceOptions) (*BoardsResponse, *APIError) {
	path := "/search/boards"

	resp := new(BoardsResponse)
	err := r.Cli.DoGet(path, args, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// SearchUsersPins Search for pins for the "operation user_account".
// Refer: https://developers.pinterest.com/docs/api/v5/#operation/search_user_pins/get
func (r *SearchResource) SearchUsersPins(args SearchResourceOptions) (*PinsResponse, *APIError) {
	path := "/search/pins"

	resp := new(PinsResponse)
	err := r.Cli.DoGet(path, args, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// SearchOptions specifies the optional parameters to search partner pins
type SearchOptions struct {
	Term        string `url:"term,omitempty"`
	CountryCode string `url:"country_code,omitempty"`
	Bookmark    string `url:"bookmark,omitempty"`
	Locale      string `url:"locale,omitempty"`
	Limit       int    `url:"limit,omitempty"`
}

// SearchResult specifies the search response items
type SearchResult struct {
	Media       *Media  `url:"media,omitempty"`
	AltText     *string `url:"alt_text,omitempty"`
	Link        *string `url:"link,omitempty"`
	Title       *string `url:"title,omitempty"`
	Description *int    `url:"description,omitempty"`
}

func (s SearchResult) String() string {
	return Stringify(s)
}

// SearchResponse represents the response for search partner pins
type SearchResponse struct {
	Items    []*SearchResult `json:"items"`
	Bookmark *string         `json:"bookmark"`
}

func (s SearchResponse) String() string {
	return Stringify(s)
}

// SearchPins Get the top n Pins by a given search term.
// Refer: https://developers.pinterest.com/docs/api/v5/#operation/search_partner_pins
func (r *SearchResource) SearchPartnerPins(args SearchOptions) (*SearchResponse, *APIError) {
	path := "/search/pins"

	resp := new(SearchResponse)
	err := r.Cli.DoGet(path, args, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
