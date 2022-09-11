package storage

import (
	"github.com/goocarry/bootstrapper/app/pkg/client/postgresql"
	"github.com/goocarry/bootstrapper/app/pkg/logger"
)

type storage struct {
	client postgresql.Client
	logger *logger.Logger
}
