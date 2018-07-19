package generator

import (
	"bytes"

	"github.com/SirMetathyst/proton"

	"github.com/SirMetathyst/proton/model"
)

// ComponentGenerator_C_1_7_0 ...
func ComponentGenerator_C_1_7_0(md *model.MD) ([]proton.FileInfo, error) {
	slice := make([]proton.FileInfo, 0)
	for _, cp := range md.ComponentList() {
		if !cp.IsListener() {
			slice = append(slice, proton.NewFileInfo("../Generated_Components/"+cp.ID().WithComponentSuffix().ToUpperFirst().String()+".cs", Component_C_1_7_0(cp, new(bytes.Buffer)), "ComponentGenerator_C_1_7_0"))
		}
	}
	return slice, nil
}
