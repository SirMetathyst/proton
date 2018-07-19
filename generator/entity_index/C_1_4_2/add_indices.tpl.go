// Code generated by hero.
// source: C:\Users\alexa\go\src\github.com\SirMetathyst\proton\generator\entity_index\C_1_4_2\add_indices.tpl
// DO NOT EDIT!
package generator

import (
	"bytes"

	"github.com/SirMetathyst/proton/model"
)

func EntityIndexAddIndices_C_1_4_2(cp []*model.CP, b *bytes.Buffer) string {
	for _, ccp := range cp {
		for _, c := range ccp.ContextList() {
			for _, m := range ccp.MemberList() {
				if m.EntityIndex() > 0 {

					ID := ccp.ID().WithoutComponentSuffix().ToUpperFirst().String()
					Type := m.Value().String()
					IndexType := ""
					if m.EntityIndex() == 1 {
						IndexType = "Entitas.PrimaryEntityIndex"
					} else if m.EntityIndex() > 1 {
						IndexType = "Entitas.EntityIndex"
					}

					b.WriteString("\t\t")
					b.WriteString(c.ID().WithoutContextSuffix().ToLowerFirst().String())
					b.WriteString(`.AddEntityIndex(new `)
					b.WriteString(IndexType)
					b.WriteString(`<`)
					b.WriteString(c.ID().WithoutContextSuffix().ToUpperFirst().String())
					b.WriteString(`Entity, `)
					b.WriteString(Type)
					b.WriteString(`>(
`)
					if len(ccp.MembersWithEntityIndex()) == 1 {
						b.WriteString("\t\t\t")
						b.WriteString(ID)
					} else if len(ccp.MembersWithEntityIndex()) > 1 {
						b.WriteString("\t\t\t")
						b.WriteString(ID)
						b.WriteString(m.ID().ToUpperFirst().String())

					}
					b.WriteString(`,`)
					b.WriteString(c.ID().WithoutContextSuffix().ToLowerFirst().String())
					b.WriteString(`.GetGroup(`)
					b.WriteString(c.ID().WithoutContextSuffix().ToUpperFirst().String())
					b.WriteString(`Matcher.`)
					b.WriteString(ID)
					b.WriteString(`), (e, c) => ((`)
					b.WriteString(ccp.ID().WithComponentSuffix().ToUpperFirst().String())
					b.WriteString(`)c).`)
					b.WriteString(m.ID().ToLowerFirst().String())
					b.WriteString(`));`)
					b.WriteString("\n\n")
				}
			}
		}
	}
	return b.String()

}
