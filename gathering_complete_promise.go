package webrtc

import "context"

// GatheringCompletePromise is a Pion specific helper function that returns a channel that is closed when gathering is complete.
// This function may be helpful in cases where you are unable to trickle your ICE Candidates.
//
// It is better to not use this function, and instead trickle candidates. If you use this function you will see longer connection startup times.
// When the call is connected you will see no impact however.
func GatheringCompletePromise(pc *PeerConnection) (gatherComplete <-chan struct{}) {
	gatheringComplete, done := context.WithCancel(context.Background())

	if pc.ICEGatheringState() == ICEGatheringStateComplete {
		done()
	} else {
		pc.setGatherCompleteHdlr(func() { done() })
	}

	return gatheringComplete.Done()
}
