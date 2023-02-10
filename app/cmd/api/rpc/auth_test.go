package rpc

import (
	"context"
	"github.com/aldlss/MiniTikTok-Social-Module/app/kitex_gen/pb/auth"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAuthRPC(t *testing.T) {
	ctx := context.Background()
	r, err := authClient.Auth(ctx, &auth.AuthRequest{Token: "1561861awg651aw6g5"})
	assert.Nil(t, err)
	log.Info(r)
}
func TestMain(m *testing.M) {
	initAuthRPC()
	m.Run()
}
