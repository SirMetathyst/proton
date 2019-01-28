package generator

import (
	"bytes"

	"github.com/SirMetathyst/go-entitas"
)

// ContextObserverGenerator_C_1_4_2 ...
func ContextObserverGenerator_C_1_4_2(md *entitas.MD) ([]entitas.FI, error) {
	slice := make([]entitas.FI, 0)
	slice = append(slice, entitas.NewFileInfo("Contexts.cs", ContextObserver_C_1_4_2(md.ContextSlice(), new(bytes.Buffer)), "ContextObserverGenerator_C_1_4_2"))
	return slice, nil
}
