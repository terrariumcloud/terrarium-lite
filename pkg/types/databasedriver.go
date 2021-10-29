// Package types provides interfaces and structs to implement Terrarium and allow
// extensibility by 3rd parties
package types

import (
	"context"
)

// TerrariumDriver is a generic database interface to allow further database implementations for Terrarium
// if you would like to implement a different database beyond the core drivers this interface should be implemented
type TerrariumDriver interface {
	Connect(ctx context.Context) error
	Organizations() OrganizationStore
	VCS() VCSStore
	Modules() ModuleStore
}
