package comm


type Limiter struct {
	token chan struct{}
}

func NewLimiter(threshold int) *Limiter {
	if threshold <= 0 {
		threshold = 0
	}

	t := &Limiter{
		token: make(chan struct{}, threshold),
	}
	for range threshold {
		t.token <- struct{}{}
	}
	return t
}

func (t *Limiter) Get() {
	<-t.token
}

func (t *Limiter) Back() {
	select {
	case t.token <- struct{}{}:
	default:
	}
}

func (t *Limiter) Close() {
	close(t.token)
}

