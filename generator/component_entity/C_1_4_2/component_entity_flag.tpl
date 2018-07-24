<%: func ComponentEntityFlag_C_1_4_2(c *entitas.C, cp *entitas.CP, b *bytes.Buffer) string %>
public partial class <%==s c.ID().WithoutContextSuffix().ToUpperFirst().String()%>Entity 
{
    static readonly <%==s cp.ID().WithComponentSuffix().ToUpperFirst().String()%><% b.WriteRune(' ') %><%==s cp.ID().WithComponentSuffix().ToLowerFirst().String()%> = new <%==s cp.ID().WithComponentSuffix().ToUpperFirst().String()%>();

    public bool <%==s cp.FlagPrefixOrDefault().ToLowerFirst().String() %><%==s cp.ID().WithoutComponentSuffix().ToUpperFirst().String()%> 
    {
        get 
        { 
            return HasComponent(<%==s c.ID().WithoutContextSuffix().ToUpperFirst().String()%>ComponentsLookup.<%==s cp.ID().WithoutComponentSuffix().ToUpperFirst().String()%>); 
        }
        set 
        {
            if (value != <%==s cp.FlagPrefixOrDefault().ToLowerFirst().String() %><%==s cp.ID().WithoutComponentSuffix().ToUpperFirst().String()%>) 
            {
                var index = <%==s c.ID().WithoutContextSuffix().ToUpperFirst().String()%>ComponentsLookup.<%==s cp.ID().WithoutComponentSuffix().ToUpperFirst().String()%>;
                if (value) 
                {
                    var componentPool = GetComponentPool(index);
                    var component = componentPool.Count > 0
                            ? componentPool.Pop()
                            : <%==s cp.ID().WithComponentSuffix().ToLowerFirst().String()%>;

                    AddComponent(index, component);
                } 
                else 
                {
                    RemoveComponent(index);
                }
            }
        }
    }
}
<% return b.String() %>