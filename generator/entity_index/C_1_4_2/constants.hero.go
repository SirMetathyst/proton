// Code generated by hero.
// source: C:\Users\alexa\go\src\github.com\SirMetathyst\go-proton\generator\entity_index\C_1_4_2\constants.hero
// DO NOT EDIT!
package generator

import (
	"bytes"

	entitas "github.com/SirMetathyst/go-entitas"
)

func EntityIndexConstants_C_1_4_2(cp []*entitas.CP, b *bytes.Buffer) string {
	for _, ccp := range cp {
		ID := ccp.ID().WithoutComponentSuffix().ToUpperFirst().String()
		if len(ccp.MembersWithEntityIndex()) == 1 {

			b.WriteRune('\t')
			b.WriteString(`public const string`)
			b.WriteRune(' ')
			b.WriteString(ID)
			b.WriteString(` = "`)
			b.WriteString(ID)
			b.WriteString(`";`)
			b.WriteString("\n")
		} else if len(ccp.MembersWithEntityIndex()) > 1 {
			for _, m := range ccp.MemberSlice() {

				b.WriteRune('\t')
				b.WriteString(`public const string`)
				b.WriteRune(' ')
				b.WriteString(ID)
				b.WriteString(m.ID().ToUpperFirst().String())
				b.WriteString(` = "`)
				b.WriteString(ID)
				b.WriteString(m.ID().ToUpperFirst().String())
				b.WriteString(`";`)
				b.WriteString("\n")

			}
		}
	}
	return b.String()

}
