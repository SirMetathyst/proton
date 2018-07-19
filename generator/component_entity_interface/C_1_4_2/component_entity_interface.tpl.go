// Code generated by hero.
// source: C:\Users\alexa\go\src\github.com\SirMetathyst\proton\generator\component_entity_interface\C_1_4_2\component_entity_interface.tpl
// DO NOT EDIT!
package generator

import (
	"bytes"

	"github.com/SirMetathyst/proton/model"
)

func ComponentEntityInterface_C_1_4_2(cp *model.CP, b *bytes.Buffer) string {
	b.WriteString(`
public interface I`)
	b.WriteString(cp.ID().WithoutComponentSuffix().ToUpperFirst().String())
	b.WriteString(`Entity 
{
    `)
	b.WriteString(cp.ID().WithComponentSuffix().ToUpperFirst().String())
	b.WriteRune(' ')
	b.WriteString(cp.ID().WithoutComponentSuffix().ToLowerFirst().String())
	b.WriteString(` 
    { 
        get; 
    }
    
    bool has`)
	b.WriteString(cp.ID().WithoutComponentSuffix().ToUpperFirst().String())
	b.WriteString(` 
    { 
        get; 
    }

    void Add`)
	b.WriteString(cp.ID().WithoutComponentSuffix().ToUpperFirst().String())
	b.WriteString(`(`)
	ComponentEntityInterfaceArgument_C_1_4_2(cp, b)
	b.WriteString(`);
    void Replace`)
	b.WriteString(cp.ID().WithoutComponentSuffix().ToUpperFirst().String())
	b.WriteString(`(`)
	ComponentEntityInterfaceArgument_C_1_4_2(cp, b)
	b.WriteString(`);
    void Remove`)
	b.WriteString(cp.ID().WithoutComponentSuffix().ToUpperFirst().String())
	b.WriteString(`();
}
`)
	return b.String()

}
