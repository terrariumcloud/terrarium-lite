// Package types provides interfaces and structs to implement Terrarium and allow
// extensibility by 3rd parties
package types

import (
	"context"
	"net/http"
)

const BadRequestPrefix string = "Bad Request"
const InternalServerErrorPrefix string = "Internal Server Error"

type APIErrorWriter interface {
	Write(rw http.ResponseWriter, err error, statusCode int)
}

type APIResponseWriter interface {
	Write(rw http.ResponseWriter, data interface{}, statusCode int)
}

type TerrariumDataResponse struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}

type TerrariumServerResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// OrganizationStore is a generic data interface for implementaing database operations relating to organizations
type OrganizationStore interface {
	Init() error
	Create(name string, email string) (*Organization, error)
	ReadAll(limit int, offset int) ([]*Organization, error)
	ReadOne(id string) (*Organization, error)
	Update(id string, name string, email string) (*Organization, error)
	Delete(id string) error
}

// TerrariumDriver is a generic database interface to allow further database implementations for Terrarium
// if you would like to implement a different database beyond the core drivers this interface should be implemented
type TerrariumDriver interface {
	Connect(ctx context.Context) error
	Organizations() OrganizationStore
}
