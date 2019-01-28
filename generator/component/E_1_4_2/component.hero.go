// Code generated by hero.
// source: C:\Users\alexa\go\src\github.com\SirMetathyst\go-proton\generator\component\E_1_4_2\component.hero
// DO NOT EDIT!
package generator

import (
	"bytes"

	entitas "github.com/SirMetathyst/go-entitas"
)

func Component_E_1_4_2(cp *entitas.CP, b *bytes.Buffer) string {
	b.WriteString(`
public sealed partial class `)
	b.WriteString(cp.ID().WithComponentSuffix().ToUpperFirst().String())
	b.WriteString(` : Entitas.IComponent 
{
`)

	i := 0
	for _, m := range cp.MemberSlice() {
		b.WriteRune('\t')
		b.WriteString("public ")
		b.WriteString(m.Value().String())
		b.WriteRune(' ')
		b.WriteString(m.ID().String())
		b.WriteRune(';')
		if i != len(cp.MemberSlice())-1 {
			b.WriteRune('\n')
		}
		i++
	}

	b.WriteString(`
}
`)
	return b.String()

}
