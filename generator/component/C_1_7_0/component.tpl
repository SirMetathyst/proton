<%! import "github.com/SirMetathyst/proton/model"; %>
<%: func Component_C_1_7_0(cp *model.CP, b *bytes.Buffer) string %>
<%+ "attribute.tpl" %>
public sealed partial class <%==s cp.ID().WithComponentSuffix().String()%> : Entitas.IComponent 
{
<%+ "member.tpl" %>
}
<% return b.String() %>