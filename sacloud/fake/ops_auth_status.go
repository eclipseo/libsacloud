package fake

import (
	"context"

	"github.com/sacloud/libsacloud/v2/sacloud"
)

// Read is fake implementation
func (o *AuthStatusOp) Read(ctx context.Context) (*sacloud.AuthStatus, error) {
	return authStatus, nil
}
