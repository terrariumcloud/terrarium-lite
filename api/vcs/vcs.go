package vcs

import "net/http"

type VCSAPIInterface interface {
	CreateVCSHandler() http.Handler
	GetVCSHandler() http.Handler
	UpdateVCSHandler() http.Handler
	ListVCSsHandler() http.Handler
	DeleteVCSHandler() http.Handler
}

type VCSAPI struct {
}
