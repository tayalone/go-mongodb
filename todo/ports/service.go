package ports

import (
	"github.com/tayalone/go-mongodb/todo"
	TodoDTO "github.com/tayalone/go-mongodb/todo/dto"
)

/*TodoSrv is a port of services*/
type TodoSrv interface {
	GetAll() []todo.Domain
	GetByID(id string) (todo.Domain, error)
	Create(data TodoDTO.Create) (todo.Domain, error)
	Update(id string, data TodoDTO.Update) (todo.Domain, error)
	RemoveByID(id string) error
}
