<%: func ComponentContextArgumentPass_1_4_2(cp *entitas.CP, b *bytes.Buffer) string %>
<% ms := cp.MemberList()
    for i, m := range ms {
        b.WriteString("new")
        b.WriteString(m.ID().ToUpperFirst().String())
        if i != len(ms)-1 {
             b.WriteString(", ")
        }
    }
%><% return b.String() %>