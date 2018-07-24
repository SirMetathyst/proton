package client

import (
	"encoding/json"
	"io/ioutil"

	"github.com/SirMetathyst/go-blackboard"
	"github.com/SirMetathyst/go-entitas"
	"github.com/SirMetathyst/go-entitas/builder"
)

// jsonContext ...
type jsonContext struct {
	ID string `json:"id"`
}

// jsonComponent ...
type jsonComponent struct {
	ID            string       `json:"id"`
	FlagPrefix    string       `json:"flag_prefix"`
	Unique        bool         `json:"unique"`
	EventTarget   string       `json:"event_target"`
	EventType     string       `json:"event_type"`
	EventPriority int          `json:"event_priority"`
	CleanupMode   string       `json:"cleanup_mode"`
	Context       []string     `json:"context"`
	Member        []jsonMember `json:"member"`
}

// jsonMember ...
type jsonMember struct {
	ID          string `json:"id"`
	Value       string `json:"value"`
	EntityIndex string `json:"entity_index"`
}

// jsonModel ...
type jsonModel struct {
	Namespace      string          `json:"namespace"`
	Context        []jsonContext   `json:"context"`
	DefaultContext string          `json:"default_context"`
	Component      []jsonComponent `json:"component"`
}

func getEventTarget(b string) entitas.EventTarget {
	if b == "self" {
		return entitas.SelfTarget
	} else if b == "any" {
		return entitas.AnyTarget
	}
	return entitas.NoTarget
}

func getEventType(e string) entitas.EventType {
	if e == "removed" {
		return entitas.RemovedEvent
	}
	return entitas.AddedEvent
}

func getEntityIndex(e string) entitas.EntityIndex {
	if e == "multiple" {
		return entitas.MultipleEntityIndex
	} else if e == "single" {
		return entitas.SingleEntityIndex
	}
	return entitas.NoEntityIndex
}

func getCleanupMode(m string) entitas.CleanupMode {
	if m == "destroy_entity" {
		return entitas.DestroyEntity
	} else if m == "remove_component" {
		return entitas.RemoveComponent
	}
	return entitas.NoCleanup
}

// componentID ...
func componentID(c string, cp jsonComponent) string {
	var eventTypeSuffix = ""
	if getEventType(cp.EventType) == entitas.RemovedEvent {
		eventTypeSuffix = "Removed"
	}
	var optionalContextID = ""
	if len(cp.Context) > 1 {
		optionalContextID = c
	}
	componentID := optionalContextID + entitas.String(cp.ID).WithoutComponentSuffix().ToUpperFirst().String() + eventTypeSuffix + "Listener"
	return componentID
}

func createEventComponent(mdb *builder.MDB, default_context string, cp jsonComponent) error {
	g := func(c string, cp jsonComponent) error {
		cpb := mdb.NewComponent()
		cpb.SetID(componentID(c, cp) + "Component")
		cpb.SetListener(true)
		cpb.AddContext(c)

		err := cpb.NewMember().
			SetID("value").
			SetValue("System.Collections.Generic.List<I" + componentID(c, cp) + ">").
			Build()

		if err != nil {
			return err
		}
		err = cpb.Build()
		if err != nil {
			return err
		}
		return nil
	}

	if len(cp.Context) > 0 {
		for _, c := range cp.Context {
			err := g(c, cp)
			if err != nil {
				return err
			}
		}
	} else {
		err := g(default_context, cp)
		if err != nil {
			return err
		}
	}

	return nil
}

// JSON ...
func JSON(bb *blackboard.BB) (*entitas.MD, error) {
	jm := jsonModel{}

	raw, _ := ioutil.ReadFile(suffix(*bb.StringP("File"), ".json"))

	err := json.Unmarshal(raw, &jm)
	if err != nil {
		return nil, err
	}

	mdb := builder.NewModelBuilder()
	mdb.SetNamespace(jm.Namespace)

	for _, c := range jm.Context {
		err = mdb.NewContext().
			SetID(c.ID).
			Build()

		if err != nil {
			return nil, err
		}
	}

	mdb.SetDefaultContext(jm.DefaultContext)

	for _, cp := range jm.Component {
		cpb := mdb.NewComponent().
			SetID(cp.ID).
			SetFlagPrefix(cp.FlagPrefix).
			SetUnique(cp.Unique).
			SetEventTarget(getEventTarget(cp.EventTarget)).
			SetEventType(getEventType(cp.EventType)).
			SetEventPriority(cp.EventPriority).
			SetCleanupMode(getCleanupMode(cp.CleanupMode))

		for _, c := range cp.Context {
			cpb.AddContext(c)
		}

		for _, m := range cp.Member {
			err := cpb.NewMember().
				SetID(m.ID).
				SetValue(m.Value).
				SetEntityIndex(getEntityIndex(m.EntityIndex)).
				Build()

			if err != nil {
				return nil, err
			}
		}

		err := cpb.Build()
		if err != nil {
			return nil, err
		}

		if getEventTarget(cp.EventTarget) > 0 {
			err = createEventComponent(mdb, jm.DefaultContext, cp)
			if err != nil {
				return nil, err
			}
		}
	}

	md, err := mdb.Build()
	if err != nil {
		return nil, err
	}

	return md, nil
}