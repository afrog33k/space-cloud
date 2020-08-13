package routing

import (
	"sync"
	"text/template"

	"github.com/spaceuptech/space-cloud/gateway/config"
)

// Routing manages the routing functionality of space cloud
type Routing struct {
	lock sync.RWMutex

	routes       config.Routes
	globalConfig *config.GlobalRoutesConfig
	goTemplates  map[string]*template.Template
}

// New creates a new instance of the routing module
func New() *Routing {
	return &Routing{routes: make(config.Routes, 0), goTemplates: map[string]*template.Template{}, globalConfig: new(config.GlobalRoutesConfig)}
}
