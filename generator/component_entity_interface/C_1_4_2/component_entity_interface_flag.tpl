<%: func ComponentEntityInterfaceFlag_C_1_4_2(cp *entitas.CP, b *bytes.Buffer) string %>
public interface I<%==s cp.ID().WithoutComponentSuffix().ToUpperFirst().String() %>Entity 
{
    bool <%==s cp.FlagPrefixOrDefault().ToLowerFirst().String()  %><%==s cp.ID().WithoutComponentSuffix().ToUpperFirst().String() %> 
    { 
        get; 
        set; 
    }
}
<% return b.String() %>