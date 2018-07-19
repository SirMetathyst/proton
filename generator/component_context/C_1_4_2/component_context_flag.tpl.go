// Code generated by hero.
// source: C:\Users\alexa\go\src\github.com\SirMetathyst\proton\generator\component_context\C_1_4_2\component_context_flag.tpl
// DO NOT EDIT!
package generator

import (
	"bytes"

	"github.com/SirMetathyst/proton/model"
)

func ComponentContextFlag_C_1_4_2(c *model.C, cp *model.CP, b *bytes.Buffer) string {
	b.WriteString(`
public partial class `)
	b.WriteString(c.ID().WithContextSuffix().ToUpperFirst().String())
	b.WriteString(` 
{
    public `)
	b.WriteString(c.ID().WithoutContextSuffix().ToUpperFirst().String())
	b.WriteString(`Entity `)
	b.WriteString(cp.ID().WithoutComponentSuffix().ToLowerFirst().String())
	b.WriteString(`Entity 
    { 
        get 
        { 
            return GetGroup(`)
	b.WriteString(c.ID().WithoutContextSuffix().ToUpperFirst().String())
	b.WriteString(`Matcher.`)
	b.WriteString(cp.ID().WithoutComponentSuffix().ToUpperFirst().String())
	b.WriteString(`).GetSingleEntity(); 
        } 
    }

    public bool `)
	b.WriteString(cp.FlagPrefixOrDefault().ToLowerFirst().String())
	b.WriteString(cp.ID().WithoutComponentSuffix().ToUpperFirst().String())
	b.WriteString(` 
    {
        get 
        { 
            return `)
	b.WriteString(cp.ID().WithoutComponentSuffix().ToLowerFirst().String())
	b.WriteString(`Entity != null; 
        }
        set {
            var entity = `)
	b.WriteString(cp.ID().WithoutComponentSuffix().ToLowerFirst().String())
	b.WriteString(`Entity;
            if (value != (entity != null))
            {
                if (value) 
                {
                    CreateEntity().`)
	b.WriteString(cp.FlagPrefixOrDefault().WithoutComponentSuffix().ToLowerFirst().String())
	b.WriteString(cp.ID().WithoutComponentSuffix().ToUpperFirst().String())
	b.WriteString(` = true;
                } 
                else 
                {
                    entity.Destroy();
                }
            }
        }
    }
}
`)
	return b.String()

}
