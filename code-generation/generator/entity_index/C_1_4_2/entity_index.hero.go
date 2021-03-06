// Code generated by hero.
// source: C:\Users\alexa\go\src\github.com\SirMetathyst\go-proton\generator\entity_index\C_1_4_2\entity_index.hero
// DO NOT EDIT!
package generator

import (
	"bytes"

	proton "github.com/SirMetathyst/go-proton"
)

func EntityIndex_C_1_4_2(cp []*proton.Component, ei []*proton.EntityIndex, b *bytes.Buffer) string {
	b.WriteString(`
public partial class Contexts
{
`)
	EntityIndexConstants_C_1_4_2(cp, b)
	EntityIndexCustomConstants_C_1_4_2(ei, b)
	b.WriteString(`

    [Entitas.CodeGeneration.Attributes.PostConstructor]
    public void InitializeEntityIndices()
    {
`)
	EntityIndexAddIndices_C_1_4_2(cp, b)
	EntityIndexAddCustomIndices_C_1_4_2(ei, b)
	b.WriteString(`
    }
}

public static class ContextsExtensions
{
`)
	EntityIndexGetIndices_C_1_4_2(cp, b)
	EntityIndexGetCustomIndices_C_1_4_2(ei, b)
	b.WriteString(`
}
`)
	return b.String()

}
