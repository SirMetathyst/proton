// Code generated by hero.
// source: C:\Users\alexa\go\src\github.com\SirMetathyst\proton\generator\entity_index\C_1_4_2\get_indices.tpl
// DO NOT EDIT!
package generator

import (
	"bytes"

	"github.com/SirMetathyst/proton/model"
)

func EntityIndexGetIndices_C_1_4_2(cp []*model.CP, b *bytes.Buffer) string {
	for _, ccp := range cp {
		for _, c := range ccp.ContextList() {
			for _, m := range ccp.MemberList() {
				if m.EntityIndex() > 0 {

					ID := ccp.ID().WithoutComponentSuffix().ToUpperFirst().String()
					Type := m.Value().String()
					MemberName := m.ID().ToLowerFirst().String()
					IndexName := ""
					IndexType := ""
					if m.EntityIndex() == 1 {
						IndexType = "Entitas.PrimaryEntityIndex"
					} else if m.EntityIndex() > 1 {
						IndexType = "Entitas.EntityIndex"
					}
					entityIndexCount := len(ccp.MembersWithEntityIndex())
					if entityIndexCount == 1 {
						IndexName = ID
					} else if entityIndexCount > 1 {
						IndexName = ID + m.ID().ToUpperFirst().String()
					}
					if m.EntityIndex() == 1 {
						b.WriteRune('\t')
						b.WriteString(`public static `)
						b.WriteString(c.ID().WithoutContextSuffix().ToUpperFirst().String())
						b.WriteString(`Entity GetEntityWith`)
						b.WriteString(IndexName)
						b.WriteString(`(this `)
						b.WriteString(c.ID().WithContextSuffix().ToUpperFirst().String())
						b.WriteString(` context, `)
						b.WriteString(Type)
						b.WriteRune(' ')
						b.WriteString(MemberName)
						b.WriteString(`)
    {
        return ((`)
						b.WriteString(IndexType)
						b.WriteString(`<`)
						b.WriteString(c.ID().WithoutContextSuffix().ToUpperFirst().String())
						b.WriteString(`Entity, `)
						b.WriteString(Type)
						b.WriteString(`>)context.GetEntityIndex(Contexts.`)
						b.WriteString(IndexName)
						b.WriteString(`)).GetEntity(`)
						b.WriteString(MemberName)
						b.WriteString(`); 
    }`)
						b.WriteRune('\n')

					} else if m.EntityIndex() > 1 {
						b.WriteRune('\t')
						b.WriteString(`public static System.Collections.Generic.HashSet<`)
						b.WriteString(c.ID().WithoutContextSuffix().ToUpperFirst().String())
						b.WriteString(`Entity> GetEntitiesWith`)
						b.WriteString(IndexName)
						b.WriteString(`(this `)
						b.WriteString(c.ID().WithContextSuffix().ToUpperFirst().String())
						b.WriteString(` context, `)
						b.WriteString(Type)
						b.WriteRune(' ')
						b.WriteString(MemberName)
						b.WriteString(`)
    {
        return ((`)
						b.WriteString(IndexType)
						b.WriteString(`<`)
						b.WriteString(c.ID().WithoutContextSuffix().ToUpperFirst().String())
						b.WriteString(`Entity, `)
						b.WriteString(Type)
						b.WriteString(`>)context.GetEntityIndex(Contexts.`)
						b.WriteString(IndexName)
						b.WriteString(`)).GetEntities(`)
						b.WriteString(MemberName)
						b.WriteString(`); 
    }`)
					}
				}
			}
		}
	}
	return b.String()

}
