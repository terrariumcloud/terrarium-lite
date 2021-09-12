package organization

import "net/http"

type OrganizationAPIInterface interface {
	CreateOrganizationHandler() http.Handler
	GetOrganizationHandler() http.Handler
	ListOrganizationsHandler() http.Handler
	DeleteOrganizationHandler() http.Handler
}

type OrganizationAPI struct {
	Path string
}

func (o *OrganizationAPI) CreateOrganizationHandler() http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {

	})
}

func (o *OrganizationAPI) GetOrganizationHandler() http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {

	})
}

func (o *OrganizationAPI) ListOrganizationsHandler() http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {

	})
}

func (o *OrganizationAPI) DeleteOrganizationHandler() http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {

	})
}

func NewOrganizationAPI(path string) *OrganizationAPI {
	return &OrganizationAPI{
		Path: path,
	}
}
