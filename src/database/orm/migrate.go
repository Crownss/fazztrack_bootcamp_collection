package orm

import (
	"log"

	"github.com/crownss/fazztrack_bootchamp/week_10/src/database/orm/models"
	"github.com/crownss/fazztrack_bootchamp/week_10/src/helpers"
	"github.com/go-gormigrate/gormigrate/v2"
	"github.com/spf13/cobra"
	"gorm.io/gorm"
)

var (
	migUp   bool
	migDown bool
)

var MigrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "database migration",
	RunE:  dbMigrate,
}

func init() {
	MigrateCmd.Flags().BoolVarP(&migUp, "up", "u", false, "migration....")
	MigrateCmd.Flags().BoolVarP(&migDown, "down", "d", false, "migration rollback....")
}

func dbMigrate(cmd *cobra.Command, args []string) error {
	db, err := New()
	if err != nil {
		return err
	}

	m := gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
		{
			ID: helpers.RandomCode(5),
			Migrate: func(tx *gorm.DB) error {
				return tx.AutoMigrate(&models.History{}, &models.Users{}, &models.Vehicles{})
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Migrator().DropTable(&models.History{}, &models.Users{}, &models.Vehicles{})
			},
		},
	})
	if migUp {
		if err := m.Migrate(); err != nil {
			return err
		}
		log.Println("migration completed")
		return nil
	}
	if migDown {
		if err := m.RollbackLast(); err != nil {
			return err
		}
		log.Println("rollback migration completed")
		return nil
	}
	log.Println("init schema database done")
	return nil
}
