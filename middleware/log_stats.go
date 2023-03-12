package middleware

import (
	"github.com/Garykom/geziyor/client"
	"github.com/Garykom/geziyor/internal"
)

// LogStats logs responses
type LogStats struct {
	LogDisabled bool
}

func (p *LogStats) ProcessResponse(r *client.Response) {
	// LogDisabled check is not necessary, but done here for performance reasons
	if !p.LogDisabled {
		internal.Logger.Printf("Crawled: (%d) <%s %s>", r.StatusCode, r.Request.Method, r.Request.URL.String())
	}
}
