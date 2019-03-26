package codegeneration

import (
	"fmt"

	proton "github.com/SirMetathyst/go-proton"
)

var (
	// ErrProtonModelIsNil is returned when the a nil model is given
	ErrProtonModelIsNil = fmt.Errorf("proton: model is nil")
	// ErrProtonFileInfoIsNil is returned when a nil FileInfo is given
	ErrProtonFileInfoIsNil = fmt.Errorf("proton: fileInfo is nil")
)

// P ...
type P struct {
	generatorInfo     []*GI
	postProcessorInfo []*PPI
}

// NewProton ...
func NewProton() *P {
	return new(P)
}

var (
	p = NewProton()
)

// Singleton ...
func Singleton() *P {
	return p
}

// GeneratorInfo ...
func (p *P) GeneratorInfo() []*GI {
	return p.generatorInfo
}

// GeneratorInfo ...
func GeneratorInfo() []*GI {
	return p.generatorInfo
}

// PostProcessorInfo ...
func (p *P) PostProcessorInfo() []*PPI {
	return p.postProcessorInfo
}

// PostProcessorInfo ...
func PostProcessorInfo() []*PPI {
	return p.postProcessorInfo
}

// AddGenerator ...
func (p *P) AddGenerator(generatorVersion string, generator G, enabled bool) {
	p.generatorInfo = append(p.generatorInfo, NewGeneratorInfo(generatorVersion, generator, enabled))
}

// AddGenerator ...
func AddGenerator(generatorVersion string, generator G, enabled bool) {
	p.AddGenerator(generatorVersion, generator, enabled)
}

// AddPostProcessor ...
func (p *P) AddPostProcessor(postProcessorVersion string, postProcessor PP, enabled bool) {
	p.postProcessorInfo = append(p.postProcessorInfo, NewPostProcessorInfo(postProcessorVersion, postProcessor, enabled))
}

// AddPostProcessor ...
func AddPostProcessor(postProcessorVersion string, postProcessor PP, enabled bool) {
	p.AddPostProcessor(postProcessorVersion, postProcessor, enabled)
}

// EnableGenerator ...
func (p *P) EnableGenerator(generatorVersion string, enabled bool) {
	for _, generatorInfo := range p.generatorInfo {
		if generatorInfo.GeneratorVersion == generatorVersion {
			generatorInfo.Enabled = enabled
		}
	}
}

// EnableGenerator ...
func EnableGenerator(generatorVersion string, enabled bool) {
	p.EnableGenerator(generatorVersion, enabled)
}

// EnablePostProcessor ...
func (p *P) EnablePostProcessor(postProcessorVersion string, enabled bool) {
	for _, postProcessorInfo := range p.postProcessorInfo {
		if postProcessorInfo.PostProcessorVersion == postProcessorVersion {
			postProcessorInfo.Enabled = enabled
		}
	}
}

// EnablePostProcessor ...
func EnablePostProcessor(postProcessorVersion string, enabled bool) {
	p.EnablePostProcessor(postProcessorVersion, enabled)
}

// RunGenerator ...
func (p *P) RunGenerator(md *proton.MD) ([]proton.FI, error) {
	if md == nil {
		return nil, ErrProtonModelIsNil
	}
	r := make([]proton.FI, 0)
	for _, generatorInfo := range p.generatorInfo {
		if generatorInfo.Enabled {
			gv, err := generatorInfo.Generator(md)
			if err != nil {
				return nil, err
			}
			r = append(r, gv...)
		}
	}
	return r, nil
}

// RunGenerator ...
func RunGenerator(md *proton.MD) ([]proton.FI, error) {
	return p.RunGenerator(md)
}

// RunPostProcessor ...
func (p *P) RunPostProcessor(fi []proton.FI) ([]proton.FI, error) {
	if fi == nil {
		return nil, ErrProtonFileInfoIsNil
	}
	r := fi
	for _, postProcessorInfo := range p.postProcessorInfo {
		if postProcessorInfo.Enabled {
			pv, err := postProcessorInfo.PostProcessor(r)
			if err != nil {
				return nil, err
			}
			r = pv
		}
	}
	return r, nil
}

// RunPostProcessor ...
func RunPostProcessor(fi []proton.FI) ([]proton.FI, error) {
	return p.RunPostProcessor(fi)
}

// Run ...
func (p *P) Run(md *proton.MD) ([]proton.FI, error) {
	if md == nil {
		return nil, ErrProtonModelIsNil
	}
	gv, err := p.RunGenerator(md)
	if err != nil {
		return nil, err
	}
	pv, err := p.RunPostProcessor(gv)
	if err != nil {
		return nil, err
	}
	return pv, nil
}

// Run ...
func Run(md *proton.MD) ([]proton.FI, error) {
	return p.Run(md)
}
