package admin

import (
	"github.com/enlivengo/admincore"
	"github.com/enlivengo/database"
	"github.com/enlivengo/enliven"
	"github.com/qor/qor"
)

var adminResources []interface{}

// AddResources adds models to qor/admin
func AddResources(resources ...interface{}) {
	for _, res := range resources {
		adminResources = append(adminResources, res)
	}
}

// GetAdmin returns our instance of qor/admin
func GetAdmin(ev *enliven.Enliven) *admincore.Admin {
	if a, ok := ev.GetService("admin").(*admincore.Admin); ok {
		return a
	}
	return nil
}

// NewApp generates and returns an instance of the app
func NewApp() *App {
	return &App{}
}

// App is the admin application
type App struct {
}

// Initialize sets up the qor/admin module
func (aa *App) Initialize(ev *enliven.Enliven) {
	if !ev.AppInstalled("default_database") {
		panic("The Admin app requires that the Database app is initialized with a default connection.")
	}

	db := database.GetDatabase()

	admin := admincore.New(&qor.Config{DB: db})

	for _, resource := range adminResources {
		admin.AddResource(resource)
	}

	admin.Enliven = ev
	admin.MountTo("/admin")
	admin.SetSiteName("Enliven")
	ev.Auth.AddPermission("admin-app", ev, "Administrator")
	ev.AddService("admin", admin)
}

// GetName returns the app's name
func (aa *App) GetName() string {
	return "admin"
}
