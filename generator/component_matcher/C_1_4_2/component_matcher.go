package generator

import (
	"bytes"

	"github.com/SirMetathyst/go-entitas"
)

// eventComponentID ...
func eventComponentID(c *entitas.C, cp *entitas.CP) entitas.String {
	var eventTypeSuffix = ""
	if cp.EventType() == entitas.RemovedEvent {
		eventTypeSuffix = "Removed"
	}
	var optionalContextID = ""
	if len(cp.ContextList()) > 1 {
		optionalContextID = c.ID().String()
	}
	return entitas.String(optionalContextID + entitas.String(cp.ID()).WithoutComponentSuffix().ToUpperFirst().String() + eventTypeSuffix + "Listener")
}

// ComponentMatcherGenerator_C_1_4_2 ...
func ComponentMatcherGenerator_C_1_4_2(md *entitas.MD) ([]entitas.FI, error) {
	slice := make([]entitas.FI, 0)
	for _, cp := range md.ComponentList() {
		for _, c := range cp.ContextList() {
			slice = append(slice, entitas.NewFileInfo(c.ID().WithoutContextSuffix().ToUpperFirst().String()+"/Components/"+c.ID().WithoutContextSuffix().ToUpperFirst().String()+cp.ID().WithComponentSuffix().ToUpperFirst().String()+".cs", ComponentMatcher_C_1_4_2(c, cp, false, new(bytes.Buffer)), "ComponentMatcherGenerator_C_1_4_2"))
			if cp.EventTarget() != entitas.NoTarget {
				slice = append(slice, entitas.NewFileInfo(c.ID().WithoutContextSuffix().ToUpperFirst().String()+"/Components/"+c.ID().WithoutContextSuffix().ToUpperFirst().String()+eventComponentID(c, cp).WithComponentSuffix().ToUpperFirst().String()+".cs", ComponentMatcher_C_1_4_2(c, cp, true, new(bytes.Buffer)), "ComponentMatcherGenerator_C_1_4_2"))
			}
		}
	}
	return slice, nil
}
