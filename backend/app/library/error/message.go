package error

import (
	"os"
	"sync"

	"github.com/jumpei00/board/backend/app/library/logger"
	"gopkg.in/yaml.v3"
)

var messageYaml *errCause
var onceYamlAttach sync.Once

func initialAttachYaml() {
	file, err := os.ReadFile("app/library/error/message.yaml")
	if err != nil {
		logger.Fatal("no open message.yaml", "error", err)
	}

	errContents := &errCause{}
	if err := yaml.Unmarshal(file, errContents); err != nil {
		logger.Fatal("unmarshal yaml file failed", "error", err)
	}

	messageYaml = errContents
}

func Message() *errCause {
	// 一度だけ実行されるようにする
	onceYamlAttach.Do(func() {
		initialAttachYaml()
	})
	return messageYaml
}

type errCause struct {
	NotSameContributor   ErrContents `yaml:"ErrNotSameContributor"`
	AlreadyUsernameExist ErrContents `yaml:"ErrAlreadyUsernameExist"`
	SignInBadRequest     ErrContents `yaml:"ErrSignInBadRequest"`
	NotThreadKey         ErrContents `yaml:"ErrNotThreadKey"`
	NotCommentKey        ErrContents `yaml:"ErrNotCommentKey"`
}

type ErrContents struct {
	message string `yaml:"message"`
}
