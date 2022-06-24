package session

import "time"

type Session struct {
	UserID     string    `json:"user_id"`
	CreatedAt  time.Time `json:"created_at"`
	LastAccess time.Time `json:"last_access"`
}

func newSession(useID string) *Session {
	return &Session{
		UserID: useID,
		CreatedAt: time.Now(),
		LastAccess: time.Now(),
	}
}

func (s *Session) updateLastAccess() {
	s.LastAccess = time.Now()
}
