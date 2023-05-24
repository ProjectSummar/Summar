package handlers

import "summar/server/stores"

type Handlers struct {
	Store stores.Store
}

func NewHandlers(store stores.Store) *Handlers {
	return &Handlers{
		Store: store,
	}
}
