// Code generated by hero.
// source: C:\Users\alexa\go\src\github.com\SirMetathyst\go-proton\generator\component_matcher\C_1_4_2\component_matcher.hero
// DO NOT EDIT!
package generator

import (
	"bytes"

	proton "github.com/SirMetathyst/go-proton"
)

func ComponentMatcher_C_1_4_2(c *proton.Context, cp *proton.Component, isEventComponent bool, b *bytes.Buffer) string {

	ID := proton.String("")

	if isEventComponent {
		ID = eventComponentID(c, cp)
	} else {
		ID = cp.ID()
	}

	b.WriteString(`
public sealed partial class `)
	b.WriteString(c.ID().WithoutContextSuffix().ToUpperFirst().String())
	b.WriteString(`Matcher
{
    static Entitas.IMatcher<`)
	b.WriteString(c.ID().WithoutContextSuffix().ToUpperFirst().String())
	b.WriteString(`Entity> _matcher`)
	b.WriteString(ID.WithoutComponentSuffix().ToUpperFirst().String())
	b.WriteString(`;

    public static Entitas.IMatcher<`)
	b.WriteString(c.ID().WithoutContextSuffix().ToUpperFirst().String())
	b.WriteString(`Entity>`)
	b.WriteRune(' ')
	b.WriteString(ID.WithoutComponentSuffix().ToUpperFirst().String())
	b.WriteString(`
    {
        get
        {
            if (_matcher`)
	b.WriteString(ID.WithoutComponentSuffix().ToUpperFirst().String())
	b.WriteString(` == null)
            {
                var matcher = (Entitas.Matcher<`)
	b.WriteString(c.ID().WithoutContextSuffix().ToUpperFirst().String())
	b.WriteString(`Entity>)Entitas.Matcher<`)
	b.WriteString(c.ID().WithoutContextSuffix().ToUpperFirst().String())
	b.WriteString(`Entity>.AllOf(`)
	b.WriteString(c.ID().WithoutContextSuffix().ToUpperFirst().String())
	b.WriteString(`ComponentsLookup.`)
	b.WriteString(ID.WithoutComponentSuffix().ToUpperFirst().String())
	b.WriteString(`);
                matcher.componentNames = `)
	b.WriteString(c.ID().WithoutContextSuffix().ToUpperFirst().String())
	b.WriteString(`ComponentsLookup.componentNames;
                _matcher`)
	b.WriteString(ID.WithoutComponentSuffix().ToUpperFirst().String())
	b.WriteString(` = matcher;
            }

            return _matcher`)
	b.WriteString(ID.WithoutComponentSuffix().ToUpperFirst().String())
	b.WriteString(`;
        }
    }
}
`)
	return b.String()

}
