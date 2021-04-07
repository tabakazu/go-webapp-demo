package output

type ItemShowOutputData struct{}

type ItemShowPresenter interface {
	Complete(data ItemShowOutputData)
}
