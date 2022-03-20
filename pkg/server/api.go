package server

import (
	"github.com/pkg/errors"

	"github.com/labstack/echo/v4"
)

var ErrApiAlreadyExists = errors.New("api route already exists %s")

type ApiHandler interface {
	RegisterRoutes(group *echo.Group)
}

type ApiHandlerItem struct {
	prefix  string
	handler ApiHandler
	isRoot  bool
}

func (s *Server) addApiRoute(prefix string, h ApiHandler, isRoot bool) error {
	for _, ahi := range s.apis {
		if ahi.isRoot == isRoot && ahi.prefix == prefix {
			return errors.Errorf(ErrApiAlreadyExists.Error(), prefix)
		}
	}

	s.apis = append(s.apis, ApiHandlerItem{
		prefix:  prefix,
		handler: h,
		isRoot:  isRoot,
	})

	return nil
}

func (s *Server) AddApi(prefix string, h ApiHandler) error {
	return s.addApiRoute(prefix, h, false)
}

func (s *Server) AddApiToRoot(prefix string, h ApiHandler) error {
	return s.addApiRoute(prefix, h, true)
}
