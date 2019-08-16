package components

type Context struct {
	Components map[string]Component `json:"contextData"`
}

func NewContext() *Context {
	return &Context{
		Components: make(map[string]Component),
	}
}

func (c *Context) AddComponent(key string, component Component) {
	c.Components[key] = component
}
