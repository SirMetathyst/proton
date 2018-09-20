// Code generated by hero.
// source: C:\Users\alexa\go\src\github.com\SirMetathyst\go-proton\generator\event_listener_component_entity\C_1_6_1\event_listener_component_entity.hero
// DO NOT EDIT!
package generator

import (
	"bytes"

	entitas "github.com/SirMetathyst/go-entitas"
)

func EventListenerComponentEntity_C_1_6_1(c *entitas.C, cp *entitas.CP, b *bytes.Buffer) string {
	b.WriteString(`
public partial class `)
	b.WriteString(c.ID().WithoutContextSuffix().ToUpperFirst().String())
	b.WriteString(`Entity
{
    public void Add`)
	b.WriteString(componentID(c, cp).String())
	b.WriteString(`(I`)
	b.WriteString(componentID(c, cp).String())
	b.WriteString(` value) 
    {
        var listeners = has`)
	b.WriteString(componentID(c, cp).String())
	b.WriteString(`
            ? `)
	b.WriteString(componentID(c, cp).ToLowerFirst().String())
	b.WriteString(`.value
            : new System.Collections.Generic.List<I`)
	b.WriteString(componentID(c, cp).String())
	b.WriteString(`>();
        listeners.Add(value);
        Replace`)
	b.WriteString(componentID(c, cp).String())
	b.WriteString(`(listeners);
    }
    
    public void Remove`)
	b.WriteString(componentID(c, cp).String())
	b.WriteString(`(I`)
	b.WriteString(componentID(c, cp).String())
	b.WriteString(` value, bool removeComponentWhenEmpty = true)
    {
        var listeners = `)
	b.WriteString(componentID(c, cp).ToLowerFirst().String())
	b.WriteString(`.value;
        listeners.Remove(value);
        if (removeComponentWhenEmpty && listeners.Count == 0) 
        {
            Remove`)
	b.WriteString(componentID(c, cp).String())
	b.WriteString(`();
        } 
        else 
        {
            Replace`)
	b.WriteString(componentID(c, cp).String())
	b.WriteString(`(listeners);
        }
    }
}
`)
	return b.String()

}