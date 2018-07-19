<%! import "github.com/SirMetathyst/proton/model"; %>
<%: func EventListenerArgumentPass_1_6_1(cp *model.CP, b *bytes.Buffer) string %>
<% ml := cp.MemberList()
if len(ml) > 0 && cp.EventType() == model.AddedEvent {
    b.WriteString(", ")
    for i, m := range ml {
        b.WriteString(m.Value().String())
        b.WriteRune(' ')
        b.WriteString(m.ID().ToLowerFirst().String())
        if i != len(ml)-1 {
             b.WriteString(", ")
        }
    }
}
%><% return b.String() %>