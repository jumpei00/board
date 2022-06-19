package error

import (
	"os"

	"github.com/jumpei00/board/backend/app/library/logger"
	"gopkg.in/yaml.v3"
)

var messageYaml *errCause

func init() {
	file, err := os.ReadFile("./message.yaml")
	if err != nil {
		logger.Fatal("no open message.yaml")
	}

	errContents := &errCause{}
	yaml.Unmarshal(file, errContents)

	messageYaml = errContents
}

func Message() *errCause {
	return messageYaml
}

type errCause struct {
	NotSameContributor   ErrContents `yaml:"ErrNotSameContributor"`
	AlreadyUsernameExist ErrContents `yaml:"ErrAlreadyUsernameExist"`
	SignInBadRequest     ErrContents `yaml:"ErrSignInBadRequest"`
}

type ErrContents struct {
	message string `yaml:"message"`
}
