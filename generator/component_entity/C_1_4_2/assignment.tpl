<%: func ComponentEntityAssignment_C_1_4_2(cp *entitas.CP, b *bytes.Buffer) string %>
<% ms := cp.MemberList()
    for i, m := range ms {
        b.WriteString("\t\tcomponent.")
        b.WriteString(m.ID().ToLowerFirst().String())
        b.WriteString(" = ")
        b.WriteString("new")
        b.WriteString(m.ID().ToUpperFirst().String())
        b.WriteRune(';')
        if i != len(ms)-1 {
             b.WriteString("\n")
        }
    }
%><% return b.String() %>