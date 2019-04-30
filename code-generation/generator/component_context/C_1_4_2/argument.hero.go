// Code generated by hero.
// source: C:\Users\alexa\go\src\github.com\SirMetathyst\go-proton\generator\component_context\C_1_4_2\argument.hero
// DO NOT EDIT!
package generator

import (
	"bytes"

	proton "github.com/SirMetathyst/go-proton"
)

func ComponentContextArgument_1_4_2(cp *proton.Component, b *bytes.Buffer) string {
	ms := cp.MemberSlice()
	for i, m := range ms {
		b.WriteString(m.Value().String())
		b.WriteRune(' ')
		b.WriteString("new")
		b.WriteString(m.ID().ToUpperFirst().String())
		if i != len(ms)-1 {
			b.WriteString(", ")
		}
	}

	return b.String()

}
