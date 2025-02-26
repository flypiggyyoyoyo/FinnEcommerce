package service

import (
	"context"
	shortlink "github.com/FinnTew/FinnEcommerce/src/internal/service/shortlink/kitex_gen/shortlink"
	"testing"
)

func TestCreateShortLink_Run(t *testing.T) {
	ctx := context.Background()
	s := NewCreateShortLinkService(ctx)
	// init req and assert value

	req := &shortlink.CreateShortLinkRequest{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
