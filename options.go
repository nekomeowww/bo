package bo

import "time"

type boOptions struct {
	startTimeout time.Duration
	stopTimeout  time.Duration
}

type boApplyOptions struct {
	bootkit *boOptions
}

type Option interface {
	apply(options *boApplyOptions)
}

type startTimeoutOption time.Duration

func (t startTimeoutOption) apply(m *boApplyOptions) {
	m.bootkit.startTimeout = time.Duration(t)
}

func StartTimeout(duration time.Duration) startTimeoutOption {
	return startTimeoutOption(duration)
}

type stopTimeoutOption time.Duration

func (t stopTimeoutOption) apply(m *boApplyOptions) {
	m.bootkit.stopTimeout = time.Duration(t)
}

func StopTimeout(duration time.Duration) stopTimeoutOption {
	return stopTimeoutOption(duration)
}
