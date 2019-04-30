package generator

import (
	codegeneration "github.com/SirMetathyst/go-proton/code-generation"
	_ "github.com/SirMetathyst/go-proton/code-generation/generator/event_listener_component_entity/C_1_6_1"
)

func init() {
	codegeneration.EnableGenerator("CSharpEventListenerComponentEntityGenerator_C_1_6_1", true)
}
