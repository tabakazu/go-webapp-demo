package input

import (
	"github.com/tabakazu/golang-webapi-demo/domain"
	"github.com/tabakazu/golang-webapi-demo/usecase/output"
)

type ItemShowUsecase interface {
	Execute(id domain.ItemID) (output.ItemShowPresenter, error)
}
