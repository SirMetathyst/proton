// Code generated by hero.
// source: C:\Users\alexa\go\src\github.com\SirMetathyst\go-proton\generator\component_entity_interface\C_1_4_2\component_entity_interface_flag.hero
// DO NOT EDIT!
package generator

import (
	"bytes"

	proton "github.com/SirMetathyst/go-proton"
)

func ComponentEntityInterfaceFlag_C_1_4_2(cp *proton.Component, b *bytes.Buffer) string {
	b.WriteString(`
public interface I`)
	b.WriteString(cp.ID().WithoutComponentSuffix().ToUpperFirst().String())
	b.WriteString(`Entity
{
    bool `)
	b.WriteString(cp.FlagPrefixOrDefault().ToLowerFirst().String())
	b.WriteString(cp.ID().WithoutComponentSuffix().ToUpperFirst().String())
	b.WriteString(`
    {
        get;
        set;
    }
}
`)
	return b.String()

}
