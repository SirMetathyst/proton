// Code generated by hero.
// source: C:\Users\alexa\go\src\github.com\SirMetathyst\proton\generator\component_lookup\C_1_4_2\component_lookup.tpl
// DO NOT EDIT!
package generator

import (
	"bytes"
	"strconv"

	"github.com/SirMetathyst/proton/model"
)

func ComponentLookup_C_1_4_2(c *model.C, cp []*model.CP, b *bytes.Buffer) string {
	b.WriteString(`
public static class `)
	b.WriteString(c.ID().WithoutContextSuffix().ToUpperFirst().String())
	b.WriteString(`ComponentsLookup 
{
`)
	for i, ccp := range cp {
		b.WriteString("\tpublic const int ")
		b.WriteString(ccp.ID().WithoutComponentSuffix().ToUpperFirst().String())
		b.WriteString(" = ")
		b.WriteString(strconv.Itoa(i))
		b.WriteString(";\n")
	}
	b.WriteString(`
    public const int TotalComponents = `)
	b.WriteString(strconv.Itoa(len(cp)))
	b.WriteString(`;
	
	public static readonly string[] componentNames = 
    {
`)

	for i, ccp := range cp {
		b.WriteString("\t\t\"")
		b.WriteString(ccp.ID().WithoutComponentSuffix().ToUpperFirst().String())
		b.WriteString("\"")
		if i != len(cp)-1 {
			b.WriteString(",\n")
		}
	}

	b.WriteString(`
	};
	
	public static readonly System.Type[] componentTypes = 
    {
`)

	for i, ccp := range cp {
		b.WriteString("\t\ttypeof(")
		b.WriteString(ccp.ID().WithComponentSuffix().ToUpperFirst().String())
		b.WriteRune(')')
		if i != len(cp)-1 {
			b.WriteString(",\n")
		}
	}

	b.WriteString(`
	};
}
`)
	return b.String()

}
