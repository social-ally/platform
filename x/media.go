package x

import (
	"bytes"
	"context"
	"io"
	"mime/multipart"
	"net/http"
)

// media provides access to media endpoints.
type media struct {
	client *xClient
}

// NewMedia creates a media endpoint group using client.
func NewMedia(client *xClient) *media {
	return &media{client: client}
}

type (
	RequestUploadMediaHeaders struct {
		ContentType any `json:"Content-Type"`
	}

	RequestUploadMediaBody struct {
		Media         any    `json:"media"`
		MediaCategory string `json:"media_category"`
		MediaType     string `json:"media_type"`
	}

	RequestUploadMedia struct {
		Headers RequestUploadMediaHeaders `json:"headers"`
		Body    RequestUploadMediaBody    `json:"body"`
	}

	ResponseUploadMediaSuccessData struct {
		ID               string `json:"id"`
		MediaKey         string `json:"media_key"`
		ExpiresAfterSecs int    `json:"expires_after_secs"`
	}

	ResponseUploadMediaSuccess struct {
		Data ResponseUploadMediaSuccessData `json:"data"`
	}

	ResponseUploadMedia struct {
		Success ResponseUploadMediaSuccess `json:"success"`
	}
)

// UploadMedia calls POST https://api.x.com/2/media/upload.
func (s *media) UploadMedia(ctx context.Context, request *RequestUploadMedia) (*ResponseUploadMedia, error) {
	if s.client == nil {
		return nil, ErrMissingAccessToken
	}
	if request == nil || request.Body.Media == nil {
		return nil, ErrMissingMedia
	}

	var buffer bytes.Buffer
	writer := multipart.NewWriter(&buffer)
	if request.Body.MediaCategory != "" {
		if err := writer.WriteField("media_category", request.Body.MediaCategory); err != nil {
			return nil, err
		}
	}
	if request.Body.MediaType != "" {
		if err := writer.WriteField("media_type", request.Body.MediaType); err != nil {
			return nil, err
		}
	}
	switch media := request.Body.Media.(type) {
	case string:
		if err := writer.WriteField("media", media); err != nil {
			return nil, err
		}
	case []byte:
		part, err := writer.CreateFormFile("media", "media")
		if err != nil {
			return nil, err
		}
		if _, err := part.Write(media); err != nil {
			return nil, err
		}
	case io.Reader:
		part, err := writer.CreateFormFile("media", "media")
		if err != nil {
			return nil, err
		}
		if _, err := io.Copy(part, media); err != nil {
			return nil, err
		}
	default:
		return nil, ErrUnsupportedMedia
	}
	if err := writer.Close(); err != nil {
		return nil, err
	}

	httpRequest, err := s.client.authenticatedRequest(ctx, http.MethodPost, BaseURL+"/media/upload", &buffer)
	if err != nil {
		return nil, err
	}
	httpRequest.Header.Set("Content-Type", writer.FormDataContentType())
	var raw struct {
		Data ResponseUploadMediaSuccessData `json:"data"`
	}
	if err := s.client.Do(httpRequest, &raw); err != nil {
		return nil, err
	}
	return &ResponseUploadMedia{Success: ResponseUploadMediaSuccess{Data: raw.Data}}, nil
}
