package postprocessor

import (
	proton "github.com/SirMetathyst/go-proton"
	codegeneration "github.com/SirMetathyst/go-proton/code-generation"
)

func init() {
	codegeneration.AddPostProcessor("MergeContentPostProcessor_C_1_4_2", MergeContentPostProcessor_C_1_4_2, true)
}

// MergeContentPostProcessor_C_1_4_2 ...
func MergeContentPostProcessor_C_1_4_2(p *codegeneration.P, fileInfo []proton.FileInfo) ([]proton.FileInfo, error) {
	fl := make(map[string][]proton.FileInfo)
	inl := make([]proton.FileInfo, 0)

	for _, f := range fileInfo {
		fl[f.File()] = append(fl[f.File()], proton.NewFileInfo(f.File(), f.FileContent(), f.Generator()))
	}

	for k, v := range fl {

		file := k
		fileContent := ""
		generator := ""

		for i, f := range v {
			fileContent += f.FileContent()
			generator += f.Generator()
			if i != len(v)-1 {
				generator += ","
			}
		}

		inl = append(inl, proton.NewFileInfo(file, fileContent, generator))
	}

	return inl, nil
}
