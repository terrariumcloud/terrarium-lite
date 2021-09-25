package terrariumstore

import (
	"context"

	"github.com/dylanrhysscott/terrarium/pkg/storage/types"
)

type TerrariunDriver interface {
	Connect(ctx context.Context) error
	Organizations() types.OrganizationStore
}
