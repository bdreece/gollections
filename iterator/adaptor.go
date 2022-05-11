package iterator

type AdaptorFunc[T, Ctx any] func(*Ctx) (*T, error)

type Adaptor[T, Ctx any] struct {
	ctx  Ctx
	pred AdaptorFunc[T, Ctx]
}

func NewAdaptor[T, Ctx any](ctx Ctx, pred AdaptorFunc[T, Ctx]) *Adaptor[T, Ctx] {
	return &Adaptor[T, Ctx]{ctx, pred}
}

func (a *Adaptor[T, Ctx]) Next() (*T, error) {
	return (a.pred)(&a.ctx)
}
