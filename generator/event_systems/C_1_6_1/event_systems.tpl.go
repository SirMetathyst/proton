// Code generated by hero.
// source: C:\Users\alexa\go\src\github.com\SirMetathyst\proton\generator\event_systems\C_1_6_1\event_systems.tpl
// DO NOT EDIT!
package generator

import (
	"bytes"
	"sort"
	"strconv"

	"github.com/SirMetathyst/proton/model"
)

func EventSystems_C_1_6_1(c *model.C, cp []*model.CP, b *bytes.Buffer) string {
	b.WriteString(`
public sealed class `)
	b.WriteString(c.ID().WithoutContextSuffix().ToUpperFirst().String())
	b.WriteString(`EventSystems : Feature
{
    public `)
	b.WriteString(c.ID().WithoutContextSuffix().ToUpperFirst().String())
	b.WriteString(`EventSystems(Contexts contexts) 
    {
`)

	sort.Sort(byPriority(cp))
	for _, ccp := range cp {
		b.WriteString("\t\tAdd(new ")
		b.WriteString(ComponentID(c, ccp).ToUpperFirst().String())
		b.WriteString("EventSystem(contexts));")
		b.WriteString(" // priority: ")
		b.WriteString(strconv.Itoa(ccp.EventPriority()))
		b.WriteRune('\n')
	}

	b.WriteString(`
    }
}
`)
	return b.String()

}
