package pinterest

/*
	Search API
*/

type TermsResource Resource

func newTermsResource(cli *Client) *TermsResource {
	return &TermsResource{Cli: cli}
}

// RelatedTermsResourceOptions specifies the optional parameters to the list related terms method
type RelatedTermsResourceOptions struct {
	Terms []string `url:"terms,omitempty"`
}

// SuggestedTermsResourceOptions specifies the optional parameters to the list suggested terms method
type SuggestedTermsResourceOptions struct {
	Term  string `url:"term,omitempty"`
	Limit int    `url:"limit,omitempty"`
}

// RelatedTermsItem represents the related terms list item
type RelatedTermsItem struct {
	Term         *string  `json:"term"`
	RelatedTerms []string `json:"related_terms"`
}

func (r RelatedTermsItem) String() string {
	return Stringify(r)
}

// RelatedTermsResponse represents the response for list related terms
type RelatedTermsResponse struct {
	ID               *string             `json:"id"`
	RelatedTermCount *int                `json:"related_term_count"`
	RelatedTermsList []*RelatedTermsItem `json:"related_terms_list"`
}

func (r RelatedTermsResponse) String() string {
	return Stringify(r)
}

// ListRelatedTerms Get a list of terms logically related to each input term.
// Refer: https://developers.pinterest.com/docs/api/v5/#operation/terms_related/list
func (r *TermsResource) ListRelatedTerms(args RelatedTermsResourceOptions) (*RelatedTermsResponse, *APIError) {
	path := "/terms/related"

	resp := new(RelatedTermsResponse)
	err := r.Cli.DoGet(path, args, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// ListSuggestedTerms Get popular search terms that begin with your input term.
// Refer: https://developers.pinterest.com/docs/api/v5/#operation/terms_suggested/list
func (r *TermsResource) ListSuggestedTerms(args SuggestedTermsResourceOptions) ([]string, *APIError) {
	path := "/terms/suggested"

	resp := []string{}
	err := r.Cli.DoGet(path, args, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
