package events

import "github.com/sayuthisobri/claws/internal/registry"

func init() {
	registry.Global.RegisterCustom("cloudtrail", "events", registry.Entry{
		DAOFactory:      NewEventDAO,
		RendererFactory: NewEventRenderer,
	})
}
