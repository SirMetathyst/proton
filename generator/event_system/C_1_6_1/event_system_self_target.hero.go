// Code generated by hero.
// source: C:\Users\alexa\go\src\github.com\SirMetathyst\go-proton\generator\event_system\C_1_6_1\event_system_self_target.hero
// DO NOT EDIT!
package generator

import (
	"bytes"

	entitas "github.com/SirMetathyst/go-entitas"
)

func EventSystemSelfTarget_C_1_6_1(c *entitas.C, cp *entitas.CP, b *bytes.Buffer) string {
	b.WriteString(`
public sealed class `)
	b.WriteString(componentID(c, cp).ToUpperFirst().String())
	b.WriteString(`EventSystem : Entitas.ReactiveSystem<`)
	b.WriteString(c.ID().WithoutContextSuffix().ToUpperFirst().String())
	b.WriteString(`Entity> 
{
    readonly System.Collections.Generic.List<I`)
	b.WriteString(componentID(c, cp).ToUpperFirst().String())
	b.WriteString(`> _listenerBuffer;
    
    public `)
	b.WriteString(componentID(c, cp).ToUpperFirst().String())
	b.WriteString(`EventSystem(Contexts contexts) : base(contexts.`)
	b.WriteString(c.ID().WithoutContextSuffix().ToLowerFirst().String())
	b.WriteString(`) 
    {
        _listenerBuffer = new System.Collections.Generic.List<I`)
	b.WriteString(componentID(c, cp).ToUpperFirst().String())
	b.WriteString(`>();
    }

    protected override Entitas.ICollector<`)
	b.WriteString(c.ID().WithoutContextSuffix().ToUpperFirst().String())
	b.WriteString(`Entity> GetTrigger(Entitas.IContext<`)
	b.WriteString(c.ID().WithoutContextSuffix().ToUpperFirst().String())
	b.WriteString(`Entity> context) 
    {
        return Entitas.CollectorContextExtension.CreateCollector(
            context, Entitas.TriggerOnEventMatcherExtension.`)
	b.WriteString(cp.EventType().String().String())
	b.WriteString(`(`)
	b.WriteString(c.ID().WithoutContextSuffix().ToUpperFirst().String())
	b.WriteString(`Matcher.`)
	b.WriteString(cp.ID().WithoutComponentSuffix().ToUpperFirst().String())
	b.WriteString(`)`)

	if len(cp.MemberSlice()) < 1 {
		b.WriteString(", Entitas.TriggerOnEventMatcherExtension.Removed")
		b.WriteRune('(')
		b.WriteString(c.ID().WithoutContextSuffix().ToUpperFirst().String())
		b.WriteString("Matcher.")
		b.WriteString(cp.ID().WithoutComponentSuffix().ToUpperFirst().String())
		b.WriteRune(')')
		b.WriteRune('\n')
	}
	b.WriteString(`        );
    }

    protected override bool Filter(`)
	b.WriteString(c.ID().WithoutContextSuffix().ToUpperFirst().String())
	b.WriteString(`Entity entity) 
    {
        return `)
	b.WriteString(filter(c, cp))
	b.WriteString(`;
    }

    protected override void Execute(System.Collections.Generic.List<`)
	b.WriteString(c.ID().WithoutContextSuffix().ToUpperFirst().String())
	b.WriteString(`Entity> entities) 
    {
        foreach (var e in entities) 
        {
            `)

	if len(cp.MemberSlice()) > 0 {
		b.WriteString("var component = e.")
		b.WriteString(cp.ID().WithoutComponentSuffix().ToLowerFirst().String())
		b.WriteRune(';')
		b.WriteRune('\n')
	}

	b.WriteString(`
            _listenerBuffer.Clear();
            _listenerBuffer.AddRange(e.`)
	b.WriteString(componentID(c, cp).ToLowerFirst().String())
	b.WriteString(`.value);
            foreach (var listener in _listenerBuffer) 
            {
                listener.On`)
	b.WriteString(methodID(cp).ToUpperFirst().String())
	b.WriteString(`(e`)
	EventSystemArgumentPass_1_6_1(cp, b)
	b.WriteString(`);
            }
        }
    }
}
`)
	return b.String()

}
