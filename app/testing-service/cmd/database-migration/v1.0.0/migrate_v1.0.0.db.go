package dbv1_0_0

import (
	"context"
	migrationpkg "github.com/ikaiguang/go-srv-kit/data/migration"
	dbv1_0_0_testing "github.com/ikaiguang/service-layout/app/testing-service/cmd/database-migration/v1.0.0/testing"
	"gorm.io/gorm"
)

const (
	Version = "v1.0.0"
)

// Upgrade update
func Upgrade(ctx context.Context, dbConn *gorm.DB, migrateRepo migrationpkg.MigrateRepo) (err error) {
	//var (
	//	serverVersion     = Version
	//	migrateIdentifier = serverVersion + ":migrate"
	//)
	//// 已进行数据库迁移
	//dataModel, _, err := migrateRepo.QueryRecord(ctx, migrateIdentifier)
	//if err != nil {
	//	return
	//}
	//if dataModel.Id > 0 {
	//	return
	//}
	//// 记录数据库迁移
	//defer func() {
	//	if err == nil {
	//		err = migrateRepo.CreateRecord(ctx, serverVersion, migrateIdentifier)
	//	}
	//}()

	// dbv1_0_0_nodeid
	err = dbv1_0_0_testing.NewMigrateHandler(dbConn, migrateRepo).Upgrade(ctx)
	if err != nil {
		return err
	}

	return err
}
