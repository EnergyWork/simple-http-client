package anime_facts

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

type loggingRoundTripper struct {
	logger io.Writer
	next   http.RoundTripper
}

func (l *loggingRoundTripper) RoundTrip(r *http.Request) (*http.Response, error) {
	fmt.Fprintf(l.logger, "%s | %s | %s\n", time.Now().Format(time.RFC850), r.Method, r.URL)
	return l.next.RoundTrip(r)
}
