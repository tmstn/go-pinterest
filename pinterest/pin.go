package pinterest

/*
	Pin API
*/

type PinResource Resource

func newPinResource(cli *Client) *PinResource {
	return &PinResource{Cli: cli}
}

// Pin represents the pin info.
type Pin struct {
	ID              *string           `json:"id"`
	CreatedAt       *string           `json:"created_at"`
	Link            *string           `json:"link"`
	Title           *string           `json:"title"`
	Description     *string           `json:"description"`
	DominantColor   *string           `json:"dominant_color"`
	AltText         *string           `json:"alt_text"`
	CreativeTypes   []PinCreativeType `json:"creative_types"`
	BoardID         *string           `json:"board_id"`
	BoardSectionID  *string           `json:"board_section_id"`
	BoardOwner      *BoardOwner       `json:"board_owner"`
	IsOwner         bool              `json:"is_owner"`
	Media           *Media            `json:"media"`
	ParentPinID     *string           `json:"parent_pin_id"`
	IsStandard      bool              `json:"is_standard"`
	HasBeenPromoted bool              `json:"has_been_promoted"`
	Note            string            `json:"note"`
}

func (p Pin) String() string {
	return Stringify(p)
}

// PinsResponse represents the response for list pins
type PinsResponse struct {
	Items    []*Pin  `json:"items"`
	Bookmark *string `json:"bookmark"`
}

func (p PinsResponse) String() string {
	return Stringify(p)
}

// ListPinsType represents the pin type of list pins options
type ListPinsType string

func (l ListPinsType) Private() bool {
	return l == PrivatePin
}

func (l ListPinsType) String() string {
	return string(l)
}

const (
	PrivatePin ListPinsType = "PRIVATE"
)

// PinCreativeType represents the creative type of list pins options
type PinCreativeType string

func (l PinCreativeType) Regular() bool {
	return l == RegularPin
}

func (l PinCreativeType) Video() bool {
	return l == VideoPin
}

func (l PinCreativeType) Shopping() bool {
	return l == ShoppingPin
}

func (l PinCreativeType) Carousel() bool {
	return l == CarouselPin
}

func (l PinCreativeType) MaxVideo() bool {
	return l == MaxVideoPin
}

func (l PinCreativeType) ShopThePin() bool {
	return l == ShopThePinPin
}

func (l PinCreativeType) Collection() bool {
	return l == CollectionPin
}

func (l PinCreativeType) Idea() bool {
	return l == IdeaPin
}

func (l PinCreativeType) String() string {
	return string(l)
}

const (
	RegularPin    PinCreativeType = "REGULAR"
	VideoPin      PinCreativeType = "VIDEO"
	ShoppingPin   PinCreativeType = "SHOPPING"
	CarouselPin   PinCreativeType = "CAROUSEL"
	MaxVideoPin   PinCreativeType = "MAX_VIDEO"
	ShopThePinPin PinCreativeType = "SHOP_THE_PIN"
	CollectionPin PinCreativeType = "COLLECTION"
	IdeaPin       PinCreativeType = "IDEA"
)

// ListPinsFilter represents the pin filter of list pins options
type ListPinsFilter string

func (l ListPinsFilter) ExcludeNative() bool {
	return l == ExcludeNative
}

func (l ListPinsFilter) ExcludeRepins() bool {
	return l == ExcludeRepins
}

func (l ListPinsFilter) HasBeenPromoted() bool {
	return l == HasBeenPromoted
}

func (l ListPinsFilter) String() string {
	return string(l)
}

const (
	ExcludeNative   ListPinsFilter = "exclude_native"
	ExcludeRepins   ListPinsFilter = "exclude_repins"
	HasBeenPromoted ListPinsFilter = "has_been_promoted"
)

// ListPinsOptions specifies the optional parameters to the List Pins method
type ListPinsOptions struct {
	Bookmark      string            `url:"bookmark,omitempty"`
	PageSize      int               `url:"page_size,omitempty"`
	PinFilter     ListPinsFilter    `url:"pin_filter,omitempty"`
	PinType       ListPinsType      `url:"pin_type,omitempty"`
	CreativeTypes []PinCreativeType `url:"pin_type,omitempty"`
}

// PinMediaSourceType represents the source type of pin media
type CreatePinMediaSourceType string

func (c CreatePinMediaSourceType) VideoID() bool {
	return c == VideoID
}

func (c CreatePinMediaSourceType) ImageURL() bool {
	return c == ImageURL
}

func (c CreatePinMediaSourceType) MultipleImageURLs() bool {
	return c == MultipleImageURLs
}

func (c CreatePinMediaSourceType) ImageBase64() bool {
	return c == ImageBase64
}

func (c CreatePinMediaSourceType) MultipleImageBase64() bool {
	return c == MultipleImageBase64
}

const (
	VideoID             CreatePinMediaSourceType = "video_id"
	ImageURL            CreatePinMediaSourceType = "image_url"
	MultipleImageURLs   CreatePinMediaSourceType = "multiple_image_urls"
	ImageBase64         CreatePinMediaSourceType = "image_base64"
	MultipleImageBase64 CreatePinMediaSourceType = "multiple image_base64"
)

// CreatePinMediaContentType represents the content type of pin media
type CreatePinMediaContentType string

func (c CreatePinMediaContentType) ImageJpeg() bool {
	return c == ImageJpeg
}

func (c CreatePinMediaContentType) ImagePng() bool {
	return c == ImagePng
}

func (c CreatePinMediaContentType) String() string {
	return string(c)
}

const (
	ImageJpeg CreatePinMediaContentType = "image/jpeg"
	ImagePng  CreatePinMediaContentType = "image/jpeg"
)

// CreatePinMediaImageItem represents the pin image item info
type CreatePinMediaImageItem struct {
	Title       string                    `json:"title,omitempty"`
	Description string                    `json:"description,omitempty"`
	Link        string                    `json:"link,omitempty"`
	Url         string                    `json:"url,omitempty"`
	ContentType CreatePinMediaContentType `json:"content_type,omitempty"`
	Data        string                    `json:"data,omitempty"`
}

// CreatePinMediaSourceOpts represents the parameters for pin media resource
type CreatePinMediaSourceOpts struct {
	SourceType            CreatePinMediaSourceType  `json:"source_type"`
	CoverImageURL         string                    `json:"cover_image_url,omitempty"`
	CoverImageContentType CreatePinMediaContentType `json:"cover_image_content_type,omitempty"`
	CoverImageData        string                    `json:"cover_image_data,omitempty"`
	MediaID               string                    `json:"media_id,omitempty"`
	IsStandard            bool                      `json:"is_standard"`
	Url                   string                    `json:"url,omitempty"`
	ContentType           CreatePinMediaContentType `json:"content_type,omitempty"`
	Data                  string                    `json:"data,omitempty"`
	Items                 []CreatePinMediaImageItem `json:"items,omitempty"`
	Index                 int                       `json:"index,omitempty"`
}

// CreatePinOpts represents the parameters for create a pin
type CreatePinOpts struct {
	Link           string                   `json:"link,omitempty"`
	Title          string                   `json:"title,omitempty"`
	Description    string                   `json:"description,omitempty"`
	DominantColor  string                   `json:"dominant_color,omitempty"`
	AltText        string                   `json:"alt_text,omitempty"`
	BoardID        string                   `json:"board_id"`
	BoardSectionID string                   `json:"board_section_id,omitempty"`
	MediaSource    CreatePinMediaSourceOpts `json:"media_source"`
	ParentPinID    string                   `json:"parent_pin_id,omitempty"`
	Note           string                   `json:"note,omitempty"`
}

// UpdatePinCarouselSlot represents the parameters for pin carousel slot
type UpdatePinCarouselSlotOpts struct {
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	Link        string `json:"link,omitempty"`
}

// UpdatePinOpts represents the parameters for update a pin
type UpdatePinOpts struct {
	AltText        string                      `json:"alt_text,omitempty"`
	BoardID        string                      `json:"board_id"`
	BoardSectionID string                      `json:"board_section_id,omitempty"`
	Description    string                      `json:"description,omitempty"`
	Link           string                      `json:"link,omitempty"`
	Title          string                      `json:"title,omitempty"`
	CarouselSlots  []UpdatePinCarouselSlotOpts `json:"carousel_slots,omitempty"`
	Note           string                      `json:"note,omitempty"`
}

// CreatePin Create a Pin on a board or board section owned by the "operation user_account".
// Refer: https://developers.pinterest.com/docs/api/v5/#operation/pins/get
func (r *PinResource) CreatePin(args CreatePinOpts) (*Pin, *APIError) {
	path := "/pins"

	resp := new(Pin)
	err := r.Cli.DoPost(path, args, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// getPinOpts represents the parameters for get pin
type getPinOpts struct {
	AdAccountID string `url:"ad_account_id,omitempty"`
}

// ListPins Get a list of the Pins owned by the "operation user_account".
// Refer: https://developers.pinterest.com/docs/api/v5/#operation/pins/list
func (r *BoardResource) ListPins(args ListPinsOptions) (*PinsResponse, *APIError) {
	path := "/pins"

	resp := new(PinsResponse)
	err := r.Cli.DoGet(path, args, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// GetPin Get a Pin owned by the "operation user_account" - or on a group board that has been shared with this account.
// Refer: https://developers.pinterest.com/docs/api/v5/#operation/pins/get
func (r *PinResource) GetPin(pinID, adAccountID string) (*Pin, *APIError) {
	path := "/pins/" + pinID

	resp := new(Pin)
	var err *APIError
	if adAccountID != "" {
		params := getPinOpts{AdAccountID: adAccountID}
		err = r.Cli.DoGet(path, params, resp)
	} else {
		err = r.Cli.DoGet(path, nil, resp)
	}
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// UpdatePin Update a pin owned by the "operation user_account".
// Refer: https://developers.pinterest.com/docs/api/v5/#operation/pins/update
func (r *PinResource) UpdatePin(pinID string, args UpdatePinOpts) (*Pin, *APIError) {
	path := "/pins/" + pinID

	resp := new(Pin)
	err := r.Cli.DoPatch(path, args, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// DeletePin Delete a Pins owned by the "operation user_account" - or on a group board that has been shared with this account.
// Refer: https://developers.pinterest.com/docs/api/v5/#operation/pins/delete
func (r *PinResource) DeletePin(pinID string) *APIError {
	path := "/pins/" + pinID
	err := r.Cli.DoDelete(path, nil)
	if err != nil {
		return err
	}
	return nil
}
