package http

import (
	useCase "assik3/services/contact/internal/usecase"

	"github.com/julienschmidt/httprouter"
)

type Delivery struct {
	ucContact useCase.Contact
	ucGroup   useCase.Group
	Router    *httprouter.Router
}

func New(ucContact useCase.Contact, ucGroup useCase.Group) *Delivery {
	var d = &Delivery{
		ucContact: ucContact,
		ucGroup:   ucGroup,
		Router:    httprouter.New(),
	}

	return d
}
