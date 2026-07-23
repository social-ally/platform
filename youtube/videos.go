// Code generated from social_media_api_catalog.json; DO NOT EDIT.

package youtube

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"net/url"
)

// Videos provides access to videos endpoints.
type videos struct {
	client *youTubeClient
}

// NewVideos creates a Videos endpoint group using client.
func NewVideos(client *youTubeClient) *videos {
	return &videos{client: client}
}

type (
	RequestUploadVideoQuery struct {
		Part       any        `json:"part"`
		UploadType UploadType `json:"uploadType"`
	}

	RequestUploadVideoHeaders struct {
		ContentType          any `json:"Content-Type"`
		XUploadContentType   any `json:"X-Upload-Content-Type"`
		XUploadContentLength int `json:"X-Upload-Content-Length"`
	}

	RequestUploadVideoBodySnippet struct {
		Title           string   `json:"title"`
		Description     string   `json:"description"`
		Tags            []string `json:"tags"`
		CategoryId      string   `json:"categoryId"`
		DefaultLanguage *string  `json:"defaultLanguage"`
	}

	RequestUploadVideoBodyStatus struct {
		PrivacyStatus           PrivacyStatus `json:"privacyStatus"`
		PublishAt               any           `json:"publishAt"`
		SelfDeclaredMadeForKids *bool         `json:"selfDeclaredMadeForKids"`
		ContainsSyntheticMedia  *bool         `json:"containsSyntheticMedia"`
	}

	RequestUploadVideoBody struct {
		Snippet RequestUploadVideoBodySnippet `json:"snippet"`
		Status  RequestUploadVideoBodyStatus  `json:"status"`
	}

	RequestUploadVideo struct {
		Query   RequestUploadVideoQuery   `json:"query"`
		Headers RequestUploadVideoHeaders `json:"headers"`
		Body    RequestUploadVideoBody    `json:"body"`
		Media   io.Reader                 `json:"-"`
	}

	ResponseUploadVideoSuccessSnippet struct {
	}

	ResponseUploadVideoSuccessStatus struct {
	}

	ResponseUploadVideoSuccess struct {
		ID      string                            `json:"id"`
		Kind    any                               `json:"kind"`
		Snippet ResponseUploadVideoSuccessSnippet `json:"snippet"`
		Status  ResponseUploadVideoSuccessStatus  `json:"status"`
	}

	ResponseUploadVideo struct {
		Success ResponseUploadVideoSuccess `json:"success"`
	}

	RequestListVideosQuery struct {
		Part       any     `json:"part"`
		ID         any     `json:"id"`
		Mine       *bool   `json:"mine"`
		MaxResults int     `json:"maxResults"`
		PageToken  *string `json:"pageToken"`
	}

	RequestListVideos struct {
		Query RequestListVideosQuery `json:"query"`
	}

	ResponseListVideosSuccessItemsItemSnippet struct {
	}

	ResponseListVideosSuccessItemsItemStatus struct {
	}

	ResponseListVideosSuccessItemsItemStatistics struct {
		ViewCount    string `json:"viewCount"`
		LikeCount    string `json:"likeCount"`
		CommentCount string `json:"commentCount"`
	}

	ResponseListVideosSuccessItemsItem struct {
		ID         string                                       `json:"id"`
		Snippet    ResponseListVideosSuccessItemsItemSnippet    `json:"snippet"`
		Status     ResponseListVideosSuccessItemsItemStatus     `json:"status"`
		Statistics ResponseListVideosSuccessItemsItemStatistics `json:"statistics"`
	}

	ResponseListVideosSuccess struct {
		Items         []ResponseListVideosSuccessItemsItem `json:"items"`
		NextPageToken *string                              `json:"nextPageToken"`
	}

	ResponseListVideos struct {
		Success ResponseListVideosSuccess `json:"success"`
	}

	RequestUpdateVideoQuery struct {
		Part any `json:"part"`
	}

	RequestUpdateVideoBodySnippet struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		CategoryId  string `json:"categoryId"`
	}

	RequestUpdateVideoBodyStatus struct {
		PrivacyStatus PrivacyStatus `json:"privacyStatus"`
		PublishAt     any           `json:"publishAt"`
	}

	RequestUpdateVideoBody struct {
		ID      string                        `json:"id"`
		Snippet RequestUpdateVideoBodySnippet `json:"snippet"`
		Status  RequestUpdateVideoBodyStatus  `json:"status"`
	}

	RequestUpdateVideo struct {
		Query RequestUpdateVideoQuery `json:"query"`
		Body  RequestUpdateVideoBody  `json:"body"`
	}

	ResponseUpdateVideoSuccessSnippet struct {
	}

	ResponseUpdateVideoSuccessStatus struct {
	}

	ResponseUpdateVideoSuccess struct {
		ID      string                            `json:"id"`
		Snippet ResponseUpdateVideoSuccessSnippet `json:"snippet"`
		Status  ResponseUpdateVideoSuccessStatus  `json:"status"`
	}

	ResponseUpdateVideo struct {
		Success ResponseUpdateVideoSuccess `json:"success"`
	}

	RequestDeleteVideoQuery struct {
		ID string `json:"id"`
	}

	RequestDeleteVideo struct {
		Query RequestDeleteVideoQuery `json:"query"`
	}

	ResponseDeleteVideoSuccess struct {
		HttpStatus int `json:"http_status"`
	}

	ResponseDeleteVideo struct {
		Success ResponseDeleteVideoSuccess `json:"success"`
	}
)

// UploadVideo calls POST https://www.googleapis.com/upload/youtube/v3/videos.
func (s *videos) UploadVideo(ctx context.Context, request *RequestUploadVideo) (*ResponseUploadVideo, error) {
	if request == nil {
		return nil, ErrNilEndpointRequest
	}
	query := url.Values{"part": {stringValue(request.Query.Part)}}
	body := any(request.Body)
	contentType := ""
	if request.Media != nil {
		query.Set("uploadType", "multipart")
		var err error
		body, contentType, err = multipartVideoBody(request.Body, request.Media, request.Headers.XUploadContentType)
		if err != nil {
			return nil, err
		}
	} else if request.Query.UploadType != "" {
		query.Set("uploadType", stringValue(request.Query.UploadType))
	}
	httpRequest, err := s.client.authenticatedRequest(ctx, http.MethodPost, UploadBaseURL+"/videos?"+query.Encode(), body)
	if err != nil {
		return nil, err
	}
	setUploadHeaders(httpRequest, request.Headers)
	if contentType != "" {
		httpRequest.Header.Set("Content-Type", contentType)
	}
	response := new(ResponseUploadVideo)
	if err := s.client.Do(httpRequest, &response.Success); err != nil {
		return nil, err
	}
	return response, nil
}

// ListVideos calls GET https://www.googleapis.com/youtube/v3/videos.
func (s *videos) ListVideos(ctx context.Context, request *RequestListVideos) (*ResponseListVideos, error) {
	if request == nil {
		return nil, ErrNilEndpointRequest
	}
	query := url.Values{"part": {stringValue(request.Query.Part)}, "id": {stringValue(request.Query.ID)}}
	if request.Query.Mine != nil {
		query.Set("mine", stringValue(*request.Query.Mine))
	}
	if request.Query.MaxResults != 0 {
		query.Set("maxResults", stringValue(request.Query.MaxResults))
	}
	if request.Query.PageToken != nil {
		query.Set("pageToken", *request.Query.PageToken)
	}
	httpRequest, err := s.client.authenticatedRequest(ctx, http.MethodGet, BaseURL+"/videos?"+query.Encode(), nil)
	if err != nil {
		return nil, err
	}
	response := new(ResponseListVideos)
	if err := s.client.Do(httpRequest, &response.Success); err != nil {
		return nil, err
	}
	return response, nil
}

// UpdateVideo calls PUT https://www.googleapis.com/youtube/v3/videos.
func (s *videos) UpdateVideo(ctx context.Context, request *RequestUpdateVideo) (*ResponseUpdateVideo, error) {
	if request == nil {
		return nil, ErrNilEndpointRequest
	}
	query := url.Values{"part": {stringValue(request.Query.Part)}}
	httpRequest, err := s.client.authenticatedRequest(ctx, http.MethodPut, BaseURL+"/videos?"+query.Encode(), request.Body)
	if err != nil {
		return nil, err
	}
	response := new(ResponseUpdateVideo)
	if err := s.client.Do(httpRequest, &response.Success); err != nil {
		return nil, err
	}
	return response, nil
}

// DeleteVideo calls DELETE https://www.googleapis.com/youtube/v3/videos.
func (s *videos) DeleteVideo(ctx context.Context, request *RequestDeleteVideo) (*ResponseDeleteVideo, error) {
	if request == nil {
		return nil, ErrNilEndpointRequest
	}
	if request.Query.ID == "" {
		return nil, ErrMissingID
	}
	httpRequest, err := s.client.authenticatedRequest(ctx, http.MethodDelete, BaseURL+"/videos?id="+url.QueryEscape(request.Query.ID), nil)
	if err != nil {
		return nil, err
	}
	response := &ResponseDeleteVideo{Success: ResponseDeleteVideoSuccess{HttpStatus: http.StatusNoContent}}
	if err := s.client.Do(httpRequest, nil); err != nil {
		return nil, err
	}
	return response, nil
}

func setUploadHeaders(request *http.Request, headers RequestUploadVideoHeaders) {
	if headers.ContentType != nil {
		request.Header.Set("Content-Type", stringValue(headers.ContentType))
	}
	if headers.XUploadContentType != nil {
		request.Header.Set("X-Upload-Content-Type", stringValue(headers.XUploadContentType))
	}
	if headers.XUploadContentLength != 0 {
		request.Header.Set("X-Upload-Content-Length", stringValue(headers.XUploadContentLength))
	}
}

func multipartVideoBody(metadata RequestUploadVideoBody, media io.Reader, mediaType any) (io.Reader, string, error) {
	var buffer bytes.Buffer
	writer := multipart.NewWriter(&buffer)
	metadataHeader := textproto.MIMEHeader{"Content-Type": {"application/json; charset=UTF-8"}}
	metadataPart, err := writer.CreatePart(metadataHeader)
	if err != nil {
		return nil, "", err
	}
	if err := json.NewEncoder(metadataPart).Encode(metadata); err != nil {
		return nil, "", err
	}
	typeValue := "application/octet-stream"
	if mediaType != nil {
		typeValue = stringValue(mediaType)
	}
	mediaPart, err := writer.CreatePart(textproto.MIMEHeader{"Content-Type": {typeValue}})
	if err != nil {
		return nil, "", err
	}
	if _, err := io.Copy(mediaPart, media); err != nil {
		return nil, "", err
	}
	if err := writer.Close(); err != nil {
		return nil, "", err
	}
	return &buffer, "multipart/related; boundary=" + writer.Boundary(), nil
}
