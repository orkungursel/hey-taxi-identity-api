package api

import (
	"github.com/orkungursel/hey-taxi-identity-api/internal/server"
	"github.com/orkungursel/hey-taxi-identity-api/pkg/mongo"
)

func init() {
	server.Plug(func(s *server.Server, next server.Next) {
		config := s.Config()

		// initilize mongo
		mc, err := mongo.New(s.Context(), config)
		if err != nil {
			next(err)
			return
		}
		defer mc.Disconnect(s.Context())

		if err := Api(s, mc); err != nil {
			next(err)
			return
		}

		next(nil)

		<-s.Wait() // should wait until all http handlers are closed because of the context
	})
}
