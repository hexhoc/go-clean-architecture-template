package loan

import (
	"github.com/hexhoc/go-mall-api/internal/entity"
)

//UseCase use case interface
type UseCase interface {
	Borrow(u *entity.User, b *entity.Book) error
	Return(b *entity.Book) error
}
