package context

import (
	"context"

	"github.com/labstack/echo/v4"
)

type contextKey string

const (
	XRequestId     contextKey = "X-Request-Id"
	XRequestSource contextKey = "X-Request-Source"
	XUid           contextKey = "X-Uid"
)

func CreateContext(c echo.Context) context.Context {
	req := c.Request()
	requestId := req.Header.Get("X-Request-Id")
	if requestId == "" {
		requestId = "-"
	}

	requestSource := req.Header.Get("X-Request-Source")
	if requestSource == "" {
		requestSource = "-"
	}

	uid := req.Header.Get("X-Uid")
	if uid == "" {
		uid = "-"
	}

	// コンテキストの設定
	ctx := context.WithValue(req.Context(), XRequestId, requestId)
	ctx = context.WithValue(ctx, XRequestSource, requestSource)
	ctx = context.WithValue(ctx, XUid, uid)

	return ctx
}
