package dbv1_0_0_testing

import (
	"context"
	migrationpkg "github.com/ikaiguang/go-srv-kit/data/migration"
	errorpkg "github.com/ikaiguang/go-srv-kit/kratos/error"
	testingchemas "github.com/ikaiguang/service-layout/app/testing-service/internal/data/schema"
	"gorm.io/gorm"
)

// Migrate 数据库迁移
type Migrate struct {
	dbConn      *gorm.DB
	migrateRepo migrationpkg.MigrateRepo
}

// NewMigrateHandler 处理手柄
func NewMigrateHandler(dbConn *gorm.DB, migrateRepo migrationpkg.MigrateRepo) *Migrate {
	return &Migrate{
		dbConn:      dbConn,
		migrateRepo: migrateRepo,
	}
}

// Upgrade ...
func (s *Migrate) Upgrade(ctx context.Context) error {
	var (
		mr       migrationpkg.MigrationInterface
		migrator = s.dbConn.WithContext(ctx).Migrator()
	)

	// 创建表
	mr = testingchemas.TestingSchema.CreateTableMigrator(migrator)
	if err := s.migrateRepo.RunMigratorUp(ctx, mr); err != nil {
		e := errorpkg.ErrorInternalError("")
		return errorpkg.Wrap(e, err)
	}
	// 添加字段
	mr = testingchemas.TestingSchema.AddColumnAccessToken(migrator)
	if err := s.migrateRepo.RunMigratorUp(ctx, mr); err != nil {
		e := errorpkg.ErrorInternalError("")
		return errorpkg.Wrap(e, err)
	}
	// 创建索引
	mr = testingchemas.TestingSchema.CreateUniqueIndexForIntAndUint(migrator)
	if err := s.migrateRepo.RunMigratorUp(ctx, mr); err != nil {
		e := errorpkg.ErrorInternalError("")
		return errorpkg.Wrap(e, err)
	}
	return nil
}
