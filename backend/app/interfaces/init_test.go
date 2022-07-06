package interfaces_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/jumpei00/board/backend/app/library/logger"
)

func TestMain(m *testing.M) {
	if err := os.Chdir("../../"); err != nil {
		logger.Fatal("test directory transform error", "error", err)
	}

	os.Exit(m.Run())
}

func executeHttpTest(r *gin.Engine, method string, path string, body io.Reader) *httptest.ResponseRecorder {
	request, _ := http.NewRequest(method, path, body)
	responseRecorder := httptest.NewRecorder()

	// execute
	r.ServeHTTP(responseRecorder, request)

	return responseRecorder
}
