<%! import "github.com/SirMetathyst/proton/model"; %>
<%: func EventSystemSelfTarget_C_1_6_1(c *model.C, cp *model.CP, b *bytes.Buffer) string %>
public sealed class <%==s ComponentID(c, cp).ToUpperFirst().String()%>EventSystem : Entitas.ReactiveSystem<<%==s c.ID().WithoutContextSuffix().ToUpperFirst().String()%>Entity> 
{
    readonly System.Collections.Generic.List<I<%==s ComponentID(c, cp).ToUpperFirst().String()%>> _listenerBuffer;
    
    public <%==s ComponentID(c, cp).ToUpperFirst().String()%>EventSystem(Contexts contexts) : base(contexts.<%==s c.ID().WithoutContextSuffix().ToLowerFirst().String()%>) 
    {
        _listenerBuffer = new System.Collections.Generic.List<I<%==s ComponentID(c, cp).ToUpperFirst().String()%>>();
    }

    protected override Entitas.ICollector<<%==s c.ID().WithoutContextSuffix().ToUpperFirst().String()%>Entity> GetTrigger(Entitas.IContext<<%==s c.ID().WithoutContextSuffix().ToUpperFirst().String()%>Entity> context) 
    {
        return Entitas.CollectorContextExtension.CreateCollector(
            context, Entitas.TriggerOnEventMatcherExtension.<%==s cp.EventType().String().String()%>(<%==s c.ID().WithoutContextSuffix().ToUpperFirst().String()%>Matcher.<%==s cp.ID().WithoutComponentSuffix().ToUpperFirst().String()%>)
        );
    }

    protected override bool Filter(<%==s c.ID().WithoutContextSuffix().ToUpperFirst().String()%>Entity entity) 
    {
        return <%==s Filter(c, cp)%>;
    }

    protected override void Execute(System.Collections.Generic.List<<%==s c.ID().WithoutContextSuffix().ToUpperFirst().String()%>Entity> entities) 
    {
        foreach (var e in entities) 
        {
            <% 
                if len(cp.MemberList()) > 0 {
                    b.WriteString("var component = e.")
                    b.WriteString(cp.ID().WithoutComponentSuffix().ToLowerFirst().String())
                    b.WriteRune(';')
                    b.WriteRune('\n')
                }
            %>
            _listenerBuffer.Clear();
            _listenerBuffer.AddRange(e.<%==s ComponentID(c, cp).ToLowerFirst().String()%>.value);
            foreach (var listener in _listenerBuffer) 
            {
                listener.On<%==s MethodID(cp).ToUpperFirst().String()%>(e<% EventSystemArgumentPass_1_6_1(cp, b)%>);
            }
        }
    }
}
<% return b.String() %>