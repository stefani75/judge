package judge

import (
	"time"

	"github.com/gearnode/judge/orn"
)

// Organization is the root aggregate of all Judge entities.
// In other words, the Organization can be represent like a
// tenant in a multi-tenant appliation.
type Organization struct {
	ORN       orn.ORN
	ID        string
	Name      string
	CreatedAt time.Time
}

// SetupNewOrganization foobar
func SetupNewOrganization() error {
	return nil
}
