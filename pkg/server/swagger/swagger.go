//go:build !dev

package swagger

import (
	"github.com/orkungursel/hey-taxi-identity-api/pkg/server"
)

func Api(s *server.Server) {}
