package contexted_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/uudashr/go-playground/contexted"
)

func TestBoom(t *testing.T) {
	ctx := context.Background()
	err := contexted.Boom(ctx)

	require.NoError(t, err)
	require.NoError(t, context.Cause(ctx))
}

func TestBoom_longTimeout(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := contexted.Boom(ctx)

	require.NoError(t, err)
	require.NoError(t, context.Cause(ctx))
}

func TestBoom_shortTimeout(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	err := contexted.Boom(ctx)

	require.ErrorIs(t, err, context.DeadlineExceeded)
	require.ErrorIs(t, context.Cause(ctx), context.DeadlineExceeded)
}

func TestBoom_shortTimeoutWithCause(t *testing.T) {
	ctx, cancel := context.WithTimeoutCause(context.Background(), 1*time.Second, assert.AnError)
	defer cancel()

	err := contexted.Boom(ctx)

	require.ErrorIs(t, err, context.DeadlineExceeded)
	require.ErrorIs(t, context.Cause(ctx), assert.AnError)
}

func TestBoom_canceled(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel() // Cancel the context immediately

	err := contexted.Boom(ctx)

	require.ErrorIs(t, err, context.Canceled)
	require.ErrorIs(t, context.Cause(ctx), context.Canceled)
}

func TestBoom_canceledWithCause(t *testing.T) {
	ctx, cancel := context.WithCancelCause(context.Background())
	cancel(assert.AnError)

	err := contexted.Boom(ctx)

	require.ErrorIs(t, err, context.Canceled)
	require.ErrorIs(t, context.Cause(ctx), assert.AnError)
}
