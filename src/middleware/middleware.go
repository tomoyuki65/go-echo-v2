package middleware

import (
	"log/slog"
	"net/http"
	"strings"

	utilContext "go-echo-v2/util/context"
	"go-echo-v2/util/logger"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
)

// リクエスト用のミドルウェア
func RequestMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		req := c.Request()

		requestId := uuid.New().String()
		req.Header.Set("X-Request-Id", requestId)

		return next(c)
	}
}

// ロガー用のミドルウェア（リクエスト単位でログ出力）
func LoggerMiddleware() echo.MiddlewareFunc {
	return echoMiddleware.RequestLoggerWithConfig(echoMiddleware.RequestLoggerConfig{
		LogRemoteIP:  true,
		LogUserAgent: true,
		LogMethod:    true,
		LogURI:       true,
		LogStatus:    true,
		LogLatency:   true,
		LogError:     true,
		HandleError:  true,
		LogValuesFunc: func(c echo.Context, v echoMiddleware.RequestLoggerValues) error {
			ctx := utilContext.CreateContext(c)

			// ログレベルの設定
			level := slog.LevelInfo
			if v.Status >= http.StatusBadRequest && v.Status < http.StatusInternalServerError {
				level = slog.LevelWarn
			} else if v.Status >= http.StatusInternalServerError {
				level = slog.LevelError
			}

			// ログ出力設定
			attrs := []slog.Attr{
				slog.String("remote_ip", v.RemoteIP),
				slog.String("user_agent", v.UserAgent),
				slog.String("method", v.Method),
				slog.String("uri", v.URI),
				slog.Int("status", v.Status),
				slog.String("latency", v.Latency.String()),
			}

			// エラーが発生している場合
			if v.Error != nil {
				if httpError, ok := v.Error.(*echo.HTTPError); ok {
					if msg, ok := httpError.Message.(string); ok {
						attrs = append(attrs, slog.String("err", msg))
					} else {
						attrs = append(attrs, slog.String("err", v.Error.Error()))
					}
				} else {
					attrs = append(attrs, slog.String("err", v.Error.Error()))
				}
			} else {
				attrs = append(attrs, slog.String("err", "-"))
			}

			// ログ出力
			logger.LogAttrs(ctx, level, "REQUEST", attrs...)

			return nil
		},
	})
}

// CORS設定用のミドルウェア
func CorsMiddleware() echo.MiddlewareFunc {
	return echoMiddleware.CORSWithConfig(echoMiddleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodPatch, http.MethodDelete, http.MethodOptions},
		AllowHeaders: []string{echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization, "X-Request-Source"},
	})
}

// 認証用のミドルウェア
func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Authorizationヘッダーからトークンを取得
		idToken := ""
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader != "" {
			idToken = strings.Replace(authHeader, "Bearer ", "", 1)
		} else {
			return echo.NewHTTPError(http.StatusUnauthorized, "Unauthorized")
		}

		// TODO: 必要な認証処理を実装する
		_ = idToken

		return next(c)
	}
}
