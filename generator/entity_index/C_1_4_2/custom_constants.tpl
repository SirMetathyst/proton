<%: func EntityIndexCustomConstants_C_1_4_2(ei []*entitas.EI, b *bytes.Buffer) string %>
<% for _, cei := range ei { 
    ID := cei.ID().String()
    b.WriteRune('\t')%>public const string<% b.WriteRune(' ')%><%==s ID%> = "<%==s ID%>";<% b.WriteString("\n") } %>
<% return b.String() %>