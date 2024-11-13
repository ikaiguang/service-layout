package dbv1_0_0_testing

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

// go test -v -count=1 ./app/admin-service/cmd/migration/v1.0.0/admin -test.run=TestMigrator_Upgrade
func TestMigrator_Upgrade(t *testing.T) {
	err := upHandler.Upgrade(context.Background())
	require.Nil(t, err)
}
