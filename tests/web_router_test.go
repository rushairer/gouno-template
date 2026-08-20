package tests

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/rushairer/gouno"
	"{{.ModulePath}}/router"
)

func TestWebRouter(t *testing.T) {
	gin.SetMode(gin.TestMode)
	engine := gin.New()
	router.RegisterWebRouter(engine)

	t.Run("GET /", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/", nil)
		engine.ServeHTTP(w, req)

		if w.Code != http.StatusOK {
			t.Fatalf("expected status 200, got %d", w.Code)
		}
		if body := w.Body.String(); body != "Hello gouno!" {
			t.Fatalf("expected body 'Hello gouno!', got %q", body)
		}
	})

	t.Run("GET /test/alive", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest(http.MethodGet, "/test/alive", nil)
		engine.ServeHTTP(w, req)

		if w.Code != http.StatusOK {
			t.Fatalf("expected status 200, got %d", w.Code)
		}

		var resp gouno.Response
		if err := json.Unmarshal(w.Body.Bytes(), &resp); err != nil {
			t.Fatalf("failed to unmarshal response: %v", err)
		}
		if resp.Code != http.StatusOK {
			t.Errorf("expected resp.Code 200, got %d", resp.Code)
		}
		if resp.Data != "pong" {
			t.Errorf("expected resp.Data 'pong', got %v", resp.Data)
		}
	})
}
