package main

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"sync"
	"time"
)

type YAMLHandler struct {
	opts HandlerOptions
	mu  *sync.Mutex
	out io.Writer
}

type HandlerOptions struct {
	Level slog.Leveler
}

func NewYAMLHandler(out io.Writer, opts *HandlerOptions) *YAMLHandler {
	h := &YAMLHandler{out: out, mu: &sync.Mutex{}}
	if opts != nil {
		h.opts = *opts
	}
	if h.opts.Level == nil {
		h.opts.Level = slog.LevelInfo // lets info be default as is in slog
	}
	return h
}

// We have to implement the handler interface
// 1. Enabled(context.Context, Level) bool
// 2. Handle(context.Context, Record) error

func (h *YAMLHandler) Enabled(ctx context.Context, level slog.Level) bool {
	return level >= h.opts.Level.Level()
}

func (h *YAMLHandler) Handle(ctx context.Context, r slog.Record) error {
	buf := make([]byte, 0, 1024)

	// Lets deal with the core mandatory attributes first
	if !r.Time.IsZero() {
		buf = h.appendToOutput(buf, slog.Time(slog.TimeKey, r.Time))
	}
	buf = h.appendToOutput(buf, slog.Any(slog.LevelKey, r.Level))

	buf = h.appendToOutput(buf, slog.String(slog.MessageKey, r.Message))

	// Addding other custom custom attrs here
	r.Attrs(func(a slog.Attr) bool {
		buf = h.appendToOutput(buf, a)
		return true
	})

	delimiter := []byte("---\n")
	buf = append(buf, delimiter...)

	h.mu.Lock()
	defer h.mu.Unlock()
	_, err := h.out.Write(buf)
	return err
}

func (h *YAMLHandler) appendToOutput(buf []byte, a slog.Attr) []byte {
	// In case the value of the attr is a struct that implements a logvaluer interface
	a.Value.Resolve()

	if a.Equal(slog.Attr{}) { // empty attribute doesnt append
		return buf
	}

	var r string
	switch a.Value.Kind() {
	case slog.KindString:
		r = fmt.Sprintf("%s: %q\n", a.Key, a.Value.String())
	case slog.KindTime:
		r = fmt.Sprintf("%s: %s\n", a.Key, a.Value.Time().Format(time.RFC3339Nano))
	default: // handle any
		r = fmt.Sprintf("%s: %v\n", a.Key, a.Value)
	}

	return append(buf, []byte(r)...)
}

func (h *YAMLHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return h
}

func (h *YAMLHandler) WithGroup(name string) slog.Handler {
	return h
}
