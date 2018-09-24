// Code generated by hero.
// source: C:\Users\alexa\go\src\github.com\SirMetathyst\go-proton\generator\component_entity\C_1_4_2\assignment.hero
// DO NOT EDIT!
package generator

import (
	"bytes"

	entitas "github.com/SirMetathyst/go-entitas"
)

func ComponentEntityAssignment_C_1_4_2(cp *entitas.CP, isEventComponent bool, b *bytes.Buffer) string {

	if isEventComponent {
		b.WriteString("\t\tcomponent.value")
		b.WriteString(" = ")
		b.WriteString("newValue")
		b.WriteRune(';')
	} else {
		ms := cp.MemberList()
		for i, m := range ms {
			b.WriteString("\t\tcomponent.")
			b.WriteString(m.ID().ToLowerFirst().String())
			b.WriteString(" = ")
			b.WriteString("new")
			b.WriteString(m.ID().ToUpperFirst().String())
			b.WriteRune(';')
			if i != len(ms)-1 {
				b.WriteString("\n")
			}
		}
	}

	return b.String()

}
