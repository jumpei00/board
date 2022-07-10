package application_test

import (
	"os"
	"testing"

	appError "github.com/jumpei00/board/backend/app/library/error"
	"github.com/jumpei00/board/backend/app/library/logger"
	"github.com/pkg/errors"
)

func TestMain(m *testing.M) {
	if err := os.Chdir("../../"); err != nil {
		logger.Fatal("test directory transform error", "error", err)
	}

	os.Exit(m.Run())
}

func isSameError(e1, e2 error) bool {
	e1 = errors.Cause(e1)
	e2 = errors.Cause(e2)

	if e1 == nil && e2 == nil {
		return true
	} else if e1 != nil && e2 == nil {
		return false
	} else if e1 == nil && e2 != nil {
		return false
	}

	if appError.IsBadRequest(e1) && appError.IsBadRequest(e2) {
		return true
	}

	return e1 == e2
}
