package bo

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestStartTimeout(t *testing.T) {
	t.Parallel()

	option := StartTimeout(time.Second * 100)
	applyOptions := &bootkitApplyOptions{&bootkitOptions{}}
	option.apply(applyOptions)

	assert.Equal(t, time.Second*100, applyOptions.bootkit.startTimeout)
}

func TestStopTimeout(t *testing.T) {
	t.Parallel()

	option := StopTimeout(time.Second * 100)
	applyOptions := &bootkitApplyOptions{&bootkitOptions{}}
	option.apply(applyOptions)

	assert.Equal(t, time.Second*100, applyOptions.bootkit.stopTimeout)
}
