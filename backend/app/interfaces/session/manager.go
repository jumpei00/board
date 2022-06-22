package session

import (
	"encoding/json"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/jumpei00/board/backend/app/config"
	"github.com/jumpei00/board/backend/app/domain"
	"github.com/pkg/errors"
	appError "github.com/jumpei00/board/backend/app/library/error"
)

const sessionkey = "session-key"

// 1day
const defaultMaxAge = 86400

type Manager interface {
	Get(c *gin.Context) (*Session, error)
	Create(c *gin.Context, user *domain.User) (*Session, error)
	Delete(c *gin.Context) error
}

type manager struct {
	setOptions    sessions.Options
	deleteOptions sessions.Options
}

func NewSessionManager() *manager {
	return &manager{
		setOptions: sessions.Options{
			Path:   "/",
			MaxAge: defaultMaxAge,
			// 本番環境の時は有効化
			Secure:   config.IsProduction(),
			HttpOnly: true,
		},
		deleteOptions: sessions.Options{
			MaxAge: -1,
		},
	}
}

func (m *manager) Get(c *gin.Context) (*Session, error) {
	session, err := m.get(c)
	if err != nil {
		return nil, err
	}

	// セッションへのアクセスを更新する
	session.updateLastAccess()

	if err := m.setAndSave(c, session); err != nil {
		return nil, err
	}

	return session, nil
}

func (m *manager) Create(c *gin.Context, user *domain.User) (*Session, error) {
	session := newSession(user.ID)

	if err := m.setAndSave(c, session); err != nil {
		return nil, errors.WithStack(err)
	}

	return session, nil
}

func (m *manager) Delete(c *gin.Context) error {
	if err := m.delete(c); err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (m *manager) get(c *gin.Context) (*Session, error) {
	sessionStore := sessions.Default(c)

	sessionJSON, ok := sessionStore.Get(sessionkey).([]byte)
	// 存在しなかった場合
	if sessionJSON == nil {
		return nil, appError.NewErrNotFound("session not found")
	}
	// キャストできなかった場合
	// 一応削除を試みる
	if !ok {
		if err := m.delete(c); err != nil {
			errors.WithStack(err)
		}
		return nil, appError.NewErrSessionCastFailed("session cast error")
	}

	var session *Session
	if err := json.Unmarshal(sessionJSON, session); err != nil {
		return nil, errors.WithStack(err)
	}

	return session, nil
}

func (m *manager) setAndSave(c *gin.Context, session *Session) error {
	sessionStore := sessions.Default(c)
	// セット用のオプション設定を付与
	sessionStore.Options(m.setOptions)

	sessionJSON, err := json.Marshal(session)
	if err != nil {
		return errors.WithStack(err)
	}

	sessionStore.Set(sessionkey, sessionJSON)

	if err := sessionStore.Save(); err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (m *manager) delete(c *gin.Context) error {
	sessionStore := sessions.Default(c)
	// 削除用のオプション設定を付与
	sessionStore.Options(m.deleteOptions)

	sessionStore.Delete(sessionkey)

	if err := sessionStore.Save(); err != nil {
		return errors.WithStack(err)
	}

	return nil
}
