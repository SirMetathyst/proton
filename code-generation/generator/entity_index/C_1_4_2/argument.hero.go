// Code generated by hero.
// source: C:\Users\alexa\go\src\github.com\SirMetathyst\go-proton\generator\entity_index\C_1_4_2\argument.hero
// DO NOT EDIT!
package generator

import (
	"bytes"

	proton "github.com/SirMetathyst/go-proton"
)

func EntityIndexArgument_C_1_4_2(eim *proton.EntityIndexMethod, b *bytes.Buffer) string {
	ml := eim.MemberSlice()
	for i, m := range ml {
		b.WriteString(m.Value().String())
		b.WriteRune(' ')
		b.WriteString(m.ID().ToLowerFirst().String())
		if i != len(ml)-1 {
			b.WriteString(", ")
		}
	}

	return b.String()

}
