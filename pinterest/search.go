package pinterest

/*
	Search API
*/

type SearchResource Resource

func newSearchResource(cli *Client) *SearchResource {
	return &SearchResource{Cli: cli}
}

// SearchResourceOpts specifies the optional parameters to various search methods
type SearchResourceOpts struct {
	ListOptions
	Query string `url:"query,omitempty"`
}

// SearchUsersBoards Search for boards for the "operation user_account". This includes boards of all board types.
// Refer: https://developers.pinterest.com/docs/api/v5/#operation/search_user_boards/get
func (r *SearchResource) SearchUsersBoards(args SearchResourceOpts) (*BoardsResponse, *APIError) {
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
func (r *SearchResource) SearchUsersPins(args SearchResourceOpts) (*PinsResponse, *APIError) {
	path := "/search/pins"

	resp := new(PinsResponse)
	err := r.Cli.DoGet(path, args, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// SearchOpts specifies the optional parameters to search partner pins
type SearchOpts struct {
	Term        string `url:"term,omitempty"`
	CountryCode string `url:"country_code,omitempty"`
	Bookmark    string `url:"bookmark,omitempty"`
	Locale      string `url:"locale,omitempty"`
	Limit       int    `url:"limit,omitempty"`
}

// SearchResult specifies the search response items
type SearchResult struct {
	Media       *Media  `json:"media,omitempty"`
	AltText     *string `json:"alt_text,omitempty"`
	Link        *string `json:"link,omitempty"`
	Title       *string `json:"title,omitempty"`
	Description *int    `json:"description,omitempty"`
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
func (r *SearchResource) SearchPartnerPins(args SearchOpts) (*SearchResponse, *APIError) {
	path := "/search/partner/pins"

	resp := new(SearchResponse)
	err := r.Cli.DoGet(path, args, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
