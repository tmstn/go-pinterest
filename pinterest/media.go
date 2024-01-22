package pinterest

/*
	Media API
*/

type MediaResource Resource

func newMediaResource(cli *Client) *MediaResource {
	return &MediaResource{Cli: cli}
}

// Image represents the image info
type Image struct {
	Width  *int    `json:"width"`
	Height *int    `json:"height"`
	Url    *string `json:"url"`
}

func (m Image) String() string {
	return Stringify(m)
}

const ImageMedia MediaType = "image"
const VideoMedia MediaType = "video"
const MultipleImageMedia MediaType = "multiple_images"
const MultipleVideoMedia MediaType = "multiple_videos"
const MultipleMixedMedia MediaType = "multiple_mixed"
const ImageMediaItem MediaItemType = "image"
const VideoMediaItem MediaItemType = "video"

type MediaType string

func (m MediaType) String() string {
	return string(m)
}

type MediaItemType string

func (m MediaItemType) String() string {
	return string(m)
}

// Media represents the media info
type Media struct {
	MediaItem
	Items     []*MediaItem `json:"items"`
	MediaType MediaType    `json:"media_type"`
}

// MediaItem represents the media item info
type MediaItem struct {
	ItemType MediaItemType `json:"item_type"`
	ImageItem
	VideoItem
}

// ImageItem represents the image item info
type ImageItem struct {
	Title       *string           `json:"title"`
	Description *string           `json:"description"`
	Link        *string           `json:"link"`
	Images      map[string]*Image `json:"images"`
}

// VideoItem represents the video item info
type VideoItem struct {
	CoverImageUrl string  `json:"cover_image_url"`
	VideoURL      *string `json:"video_url"`
	Duration      float64 `json:"duration"`
	Height        int     `json:"height"`
	Width         int     `json:"width"`
}

func (m Media) String() string {
	return Stringify(m)
}

// MediaUpload represents the media upload info.
type MediaUpload struct {
	MediaID   *string `json:"media_id"`
	MediaType *string `json:"media_type"`
	Status    *string `json:"status"`
}

func (m MediaUpload) String() string {
	return Stringify(m)
}

// MediaUploadsResponse represents the response for list media uploads.
type MediaUploadsResponse struct {
	Items    []*MediaUpload `json:"items"`
	Bookmark *string        `json:"bookmark"`
}

func (m MediaUploadsResponse) String() string {
	return Stringify(m)
}

// ListMediaUploads List media uploads filtered by given parameters.
// Refer: https://developers.pinterest.com/docs/api/v5/#operation/media/list
func (r *MediaResource) ListMediaUploads(args ListOptions) (*MediaUploadsResponse, *APIError) {
	path := "/media"

	resp := new(MediaUploadsResponse)
	err := r.Cli.DoGet(path, args, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// GetMediaUploadDetail Get details for a registered media upload, including its current status.
// Refer: https://developers.pinterest.com/docs/api/v5/#operation/media/get
func (r *MediaResource) GetMediaUploadDetail(mediaID string) (*MediaUpload, *APIError) {
	path := "/media/" + mediaID

	resp := new(MediaUpload)
	err := r.Cli.DoGet(path, nil, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// RegisterMediaUploadOpts represents the parameters for register media upload.
type RegisterMediaUploadOpts struct {
	MediaType string `json:"media_type"`
}

// RegisterMediaUploadResponse The response for register media upload.
type RegisterMediaUploadResponse struct {
	MediaID          *string           `json:"media_id"`
	MediaType        *string           `json:"media_type"`
	UploadURL        *string           `json:"upload_url"`
	UploadParameters map[string]string `json:"upload_parameters"`
}

func (m RegisterMediaUploadResponse) String() string {
	return Stringify(m)
}

// RegisterMediaUpload Register your intent to upload media.
// Refer: https://developers.pinterest.com/docs/api/v5/#operation/media/create
func (r *MediaResource) RegisterMediaUpload(args RegisterMediaUploadOpts) (*RegisterMediaUploadResponse, *APIError) {
	path := "/media"

	resp := new(RegisterMediaUploadResponse)
	err := r.Cli.DoPost(path, args, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
