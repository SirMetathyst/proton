<%: func Context_C_1_4_2(c *entitas.C, b *bytes.Buffer) string %>
public sealed partial class <%==s c.ID().WithContextSuffix().ToUpperFirst().String()%> : Entitas.Context<<%==s c.ID().WithoutContextSuffix().ToUpperFirst().String()%>Entity> 
{
    public <%==s c.ID().WithContextSuffix().ToUpperFirst().String()%>()
        : base(
            <%==s c.ID().WithoutContextSuffix().ToUpperFirst().String()%>ComponentsLookup.TotalComponents,
            0,
            new Entitas.ContextInfo(
                "<%==s c.ID().WithoutContextSuffix().ToUpperFirst().String()%>",
                <%==s c.ID().WithoutContextSuffix().ToUpperFirst().String()%>ComponentsLookup.componentNames,
                <%==s c.ID().WithoutContextSuffix().ToUpperFirst().String()%>ComponentsLookup.componentTypes
            ),
            (entity) =>
#if (ENTITAS_FAST_AND_UNSAFE)
            new Entitas.UnsafeAERC()
#else
            new Entitas.SafeAERC(entity)
#endif
        ) 
    {
    }
}
<% return b.String() %>