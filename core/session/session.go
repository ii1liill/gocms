package session

import (
	"github.com/gorilla/sessions"

	"github.com/ii1liill/gocms/core/config"
)

var Store = sessions.NewFilesystemStore("", []byte(config.GetString("app.key")))

