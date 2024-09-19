package api

import db "github.com/DreamCreatives/simplebank/db/sqlc"

type Server struct {
	store *db.Store
}
