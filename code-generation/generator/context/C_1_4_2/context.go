package generator

import (
	"bytes"

	"github.com/SirMetathyst/go-entitas"
	proton "github.com/SirMetathyst/go-proton/pkg"
)

func init() {
	proton.AddGenerator("CSharpContextGenerator_C_1_4_2", ContextGenerator_C_1_4_2, false)
}

// ContextGenerator_C_1_4_2 ...
func ContextGenerator_C_1_4_2(md *entitas.MD) ([]entitas.FI, error) {
	slice := make([]entitas.FI, 0)
	for _, c := range md.ContextSlice() {
		slice = append(slice, entitas.NewFileInfo(c.ID().WithoutContextSuffix().ToUpperFirst().String()+"/"+c.ID().WithContextSuffix().ToUpperFirst().String()+".cs", Context_C_1_4_2(c, new(bytes.Buffer)), "ContextGenerator_C_1_4_2"))
	}
	return slice, nil
}