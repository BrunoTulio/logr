package slog

import (
	"context"
	"log/slog"
)

// MultiHandler é um slog.Handler que distribui logs para múltiplos handlers.
type MultiHandler struct {
	handlers []slog.Handler
}

// New cria um novo MultiHandler com os handlers passados.
func NewMultiHandler(handlers ...slog.Handler) slog.Handler {
	return &MultiHandler{handlers: handlers}
}

func (m *MultiHandler) Enabled(ctx context.Context, level slog.Level) bool {
	return true
}

func (m *MultiHandler) Handle(ctx context.Context, rec slog.Record) error {
	var err error
	for _, h := range m.handlers {
		if h.Enabled(ctx, rec.Level) {
			if e := h.Handle(ctx, rec); e != nil && err == nil {
				err = e
			}
		}
	}
	return err
}

func (m *MultiHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	hs := make([]slog.Handler, len(m.handlers))
	for i, h := range m.handlers {
		hs[i] = h.WithAttrs(attrs)
	}
	return &MultiHandler{handlers: hs}
}

func (m *MultiHandler) WithGroup(name string) slog.Handler {
	hs := make([]slog.Handler, len(m.handlers))
	for i, h := range m.handlers {
		hs[i] = h.WithGroup(name)
	}
	return &MultiHandler{handlers: hs}
}
