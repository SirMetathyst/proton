<%: func ContextMatcher_E_1_4_2(c *entitas.C, b *bytes.Buffer) string %>
public sealed partial class <%==s c.ID().WithoutContextSuffix().ToUpperFirst().String()%>Matcher 
{
    public static Entitas.IAllOfMatcher<<%==s c.ID().WithoutContextSuffix().ToUpperFirst().String()%>Entity> AllOf(params int[] Index) => Entitas.Matcher<<%==s c.ID().WithoutContextSuffix().ToUpperFirst().String()%>Entity>.AllOf(Index);
    public static Entitas.IAllOfMatcher<<%==s c.ID().WithoutContextSuffix().ToUpperFirst().String()%>Entity> AllOf(params Entitas.IMatcher<<%==s c.ID().WithoutContextSuffix().ToUpperFirst().String()%>Entity>[] Matcher) => Entitas.Matcher<<%==s c.ID().WithoutContextSuffix().ToUpperFirst().String()%>Entity>.AllOf(Matcher);
    public static Entitas.IAnyOfMatcher<<%==s c.ID().WithoutContextSuffix().ToUpperFirst().String()%>Entity> AnyOf(params int[] Index) => Entitas.Matcher<<%==s c.ID().WithoutContextSuffix().ToUpperFirst().String()%>Entity>.AnyOf(Index);
    public static Entitas.IAnyOfMatcher<<%==s c.ID().WithoutContextSuffix().ToUpperFirst().String()%>Entity> AnyOf(params Entitas.IMatcher<<%==s c.ID().WithoutContextSuffix().ToUpperFirst().String()%>Entity>[] Matcher) => Entitas.Matcher<<%==s c.ID().WithoutContextSuffix().ToUpperFirst().String()%>Entity>.AnyOf(Matcher);
}
<% return b.String() %>