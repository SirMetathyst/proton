// Code generated by hero.
// source: C:\Users\alexa\go\src\github.com\SirMetathyst\proton\generator\component_entity_interface\C_1_4_2\component_entity_interface_link.tpl
// DO NOT EDIT!
package generator

import (
	"bytes"

	"github.com/SirMetathyst/proton/model"
)

func ComponentEntityInterfaceLink_C_1_4_2(c *model.C, cp *model.CP, b *bytes.Buffer) string {
	b.WriteString(`
public partial class `)
	b.WriteString(c.ID().WithoutContextSuffix().ToUpperFirst().String())
	b.WriteString(`Entity : I`)
	b.WriteString(cp.ID().WithoutComponentSuffix().ToUpperFirst().String())
	b.WriteString(`Entity 
{ 
}
`)
	return b.String()

}
