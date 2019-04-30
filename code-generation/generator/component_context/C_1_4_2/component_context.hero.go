// Code generated by hero.
// source: C:\Users\alexa\go\src\github.com\SirMetathyst\go-proton\generator\component_context\C_1_4_2\component_context.hero
// DO NOT EDIT!
package generator

import (
	"bytes"

	proton "github.com/SirMetathyst/go-proton"
)

func ComponentContext_C_1_4_2(c *proton.Context, cp *proton.Component, b *bytes.Buffer) string {
	b.WriteString(`
public partial class `)
	b.WriteString(c.ID().WithContextSuffix().ToUpperFirst().String())
	b.WriteString(`
{
    public `)
	b.WriteString(c.ID().WithoutContextSuffix().ToUpperFirst().String())
	b.WriteString(`Entity `)
	b.WriteString(cp.ID().WithoutComponentSuffix().ToLowerFirst().String())
	b.WriteString(`Entity
    {
        get
        {
            return GetGroup(`)
	b.WriteString(c.ID().WithoutContextSuffix().ToUpperFirst().String())
	b.WriteString(`Matcher.`)
	b.WriteString(cp.ID().WithoutComponentSuffix().ToUpperFirst().String())
	b.WriteString(`).GetSingleEntity();
        }
    }

    public `)
	b.WriteString(cp.ID().WithComponentSuffix().ToUpperFirst().String())
	b.WriteRune(' ')
	b.WriteString(cp.ID().WithoutComponentSuffix().ToLowerFirst().String())
	b.WriteString(`
    {
        get
        {
            return `)
	b.WriteString(cp.ID().WithoutComponentSuffix().ToLowerFirst().String())
	b.WriteString(`Entity.`)
	b.WriteString(cp.ID().WithoutComponentSuffix().ToLowerFirst().String())
	b.WriteString(`;
        }
    }

    public bool has`)
	b.WriteString(cp.ID().WithoutComponentSuffix().ToUpperFirst().String())
	b.WriteString(`
    {
        get
        {
            return `)
	b.WriteString(cp.ID().WithoutComponentSuffix().ToLowerFirst().String())
	b.WriteString(`Entity != null;
        }
    }

    public `)
	b.WriteString(c.ID().WithoutContextSuffix().ToUpperFirst().String())
	b.WriteString(`Entity Set`)
	b.WriteString(cp.ID().WithoutComponentSuffix().ToUpperFirst().String())
	b.WriteString(`(`)
	ComponentContextArgument_1_4_2(cp, b)
	b.WriteString(`)
    {
        if (has`)
	b.WriteString(cp.ID().WithoutComponentSuffix().ToUpperFirst().String())
	b.WriteString(`)
        {
            throw new Entitas.EntitasException("Could not set `)
	b.WriteString(cp.ID().WithoutComponentSuffix().ToUpperFirst().String())
	b.WriteString(`!\n" + this + " already has an entity with `)
	b.WriteString(cp.ID().WithoutComponentSuffix().ToUpperFirst().WithComponentSuffix().String())
	b.WriteString(`!",
                "You should check if the context already has a `)
	b.WriteString(cp.ID().WithoutComponentSuffix().ToLowerFirst().String())
	b.WriteString(`Entity before setting it or use context.Replace`)
	b.WriteString(cp.ID().WithoutComponentSuffix().ToUpperFirst().String())
	b.WriteString(`().");
        }
        var entity = CreateEntity();
        entity.Add`)
	b.WriteString(cp.ID().WithoutComponentSuffix().ToUpperFirst().String())
	b.WriteString(`(`)
	ComponentContextArgumentPass_1_4_2(cp, b)
	b.WriteString(`);
        return entity;
    }

    public void Replace`)
	b.WriteString(cp.ID().WithoutComponentSuffix().ToUpperFirst().String())
	b.WriteString(`(`)
	ComponentContextArgument_1_4_2(cp, b)
	b.WriteString(`)
    {
        var entity = `)
	b.WriteString(cp.ID().WithoutComponentSuffix().ToLowerFirst().String())
	b.WriteString(`Entity;
        if (entity == null)
        {
            entity = Set`)
	b.WriteString(cp.ID().WithoutComponentSuffix().ToUpperFirst().String())
	b.WriteString(`(`)
	ComponentContextArgumentPass_1_4_2(cp, b)
	b.WriteString(`);
        }
        else
        {
            entity.Replace`)
	b.WriteString(cp.ID().WithoutComponentSuffix().ToUpperFirst().String())
	b.WriteString(`(`)
	ComponentContextArgumentPass_1_4_2(cp, b)
	b.WriteString(`);
        }
    }

    public void Remove`)
	b.WriteString(cp.ID().WithoutComponentSuffix().ToUpperFirst().String())
	b.WriteString(`()
    {
        `)
	b.WriteString(cp.ID().WithoutComponentSuffix().ToLowerFirst().String())
	b.WriteString(`Entity.Destroy();
    }
}
`)
	return b.String()

}
