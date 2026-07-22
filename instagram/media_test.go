package instagram

import (
	"context"
	"errors"
	"testing"
)

func TestEndpointWithNilClientReturnsNamedError(t *testing.T) {
	_, err := NewMedia(nil).ListMedia(context.Background(), &RequestListMedia{Path: RequestListMediaPath{IgUserID: "user"}})
	if !errors.Is(err, ErrNilClient) {
		t.Fatalf("error = %v, want ErrNilClient", err)
	}
}
