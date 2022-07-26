package initialize

import (
	"os"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/example"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"

	"github.com/flipped-aurora/gin-vue-admin/server/model/dms"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// Gorm Initialize the database and generate database global variables
func Gorm() *gorm.DB {
	switch global.GVA_CONFIG.System.DbType {
	case "mysql":
		return GormMysql()
	case "pgsql":
		return GormPgSql()
	default:
		return GormMysql()
	}
}

// RegisterTables Registration database table dedicated
func RegisterTables(db *gorm.DB) {
	err := db.AutoMigrate(
		// system module table
		system.SysApi{},
		system.SysUser{},
		system.SysBaseMenu{},
		system.JwtBlacklist{},
		system.SysAuthority{},
		system.SysDictionary{},
		system.SysOperationRecord{},
		system.SysAutoCodeHistory{},
		system.SysDictionaryDetail{},
		system.SysBaseMenuParameter{},
		system.SysBaseMenuBtn{},
		system.SysAuthorityBtn{},
		system.SysAutoCode{},

		// example module table
		example.ExaFile{},
		example.ExaCustomer{},
		example.ExaFileChunk{},
		example.ExaFileUploadAndDownload{},

		// automation module sheet
		// Code generated by github.com/flipped-aurora/gin-vue-admin/server Begin; DO NOT EDIT.

		dms.DocumentAgencies{},
		dms.DocumentCategories{},
		dms.DocumentFieldReferences{},
		dms.DocumentFields{},
		dms.DocumentFiles{},
		dms.DocumentPriorityLevels{},
		dms.DocumentRelationReferences{},
		dms.DocumentSignerReferences{},
		dms.Documents{},
		dms.DocumentUsers{},
		dms.DocumentRules{},
		// Code generated by github.com/flipped-aurora/gin-vue-admin/server End; DO NOT EDIT.
	)
	if err != nil {
		global.GVA_LOG.Error("register table failed", zap.Error(err))
		os.Exit(0)
	}
	global.GVA_LOG.Info("register table success")
}
