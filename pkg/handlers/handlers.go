package handlers

import (
	"net/http"

	"github.com/codeninja/revision/pkg/config"
	"github.com/codeninja/revision/pkg/models"
	"github.com/codeninja/revision/pkg/render"
)

///TemplateData holds the data sent from handlers to templates

// Repo the repository used by the handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHanlders sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "this is the home page")
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(),"remote_ip", remoteIP)
	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello , Again"
	remoteIP := m.App.Session.GetString(r.Context(),"remote_ip")
	stringMap["remote_addr"] = remoteIP
	render.RenderTemplate(w, "about.page.tmpl",&models.TemplateData{
		StringMap: stringMap,
	})
	// sum := addValues(2, 2)
	// _, _ = fmt.Fprintf(w, fmt.Sprintf("about us page and 2+2 is %d", sum))
}
