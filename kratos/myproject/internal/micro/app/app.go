package app

import "context"

var (
	StartC = make(chan bool)
	StopC  = make(chan bool)
)

func OnStart(ctx context.Context) {
	select {
	case <-StartC:
	default:
		close(StartC)
	}
}

func OnStop(ctx context.Context) {
	select {
	case <-StopC:
	default:
		close(StopC)
	}
}
