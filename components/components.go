package components

type Component struct {
	Type string `json:"componentType"`
}

func NewComponent(t string) Component {
	return Component{
		Type: t,
	}
}
