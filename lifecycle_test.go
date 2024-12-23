package bo

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHook_Start(t *testing.T) {
	t.Parallel()

	l := Hook{}

	require.NotPanics(t, func() {
		err := l.Start(context.Background())
		assert.NoError(t, err)
	})

	l = Hook{
		OnStart: func(ctx context.Context) error {
			return nil
		},
	}

	require.NotPanics(t, func() {
		err := l.Start(context.Background())
		assert.NoError(t, err)
	})
}

func TestHook_Stop(t *testing.T) {
	t.Parallel()

	l := Hook{}

	require.NotPanics(t, func() {
		err := l.Stop(context.Background())
		assert.NoError(t, err)
	})

	l = Hook{
		OnStop: func(ctx context.Context) error {
			return nil
		},
	}

	require.NotPanics(t, func() {
		err := l.Stop(context.Background())
		assert.NoError(t, err)
	})
}

func TestLifeCycle_Append(t *testing.T) {
	t.Parallel()

	l := newLifeCycle()

	l.Append(Hook{})
	l.Append(Hook{})

	assert.Len(t, l.hooks, 2)
}
