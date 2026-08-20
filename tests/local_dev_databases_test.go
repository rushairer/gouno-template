package tests

import (
	"database/sql"
	"log"
	"os"
	"path/filepath"
	"runtime"

	_ "github.com/lib/pq"

	"{{.ModulePath}}/config"
)

func NewTestDB() *sql.DB {
	configManager := NewTestConfigManager()

	driverConfig := configManager.Config().DatabaseConfig.GetDefaultDriver()
	if driverConfig == nil {
		log.Panic("no default database driver configured")
	}

	// lib/pq 以 "postgres" 名称注册驱动;配置中的 driver 字段可能是 "pgx"/"postgres"/"sqlite3" 等。
	// 若使用其他驱动,请先在测试中导入对应的 database/sql 驱动包。
	driverName := driverConfig.Driver
	switch driverName {
	case "pgx", "postgres":
		driverName = "postgres"
	}

	db, err := sql.Open(driverName, driverConfig.DSN)
	if err != nil {
		log.Panic(err)
	}
	if err := db.Ping(); err != nil {
		log.Panic(err)
	}
	return db
}

func NewTestConfigManager() (configManager *config.ConfigManager) {
	projectRoot := projectRoot()
	configPath := filepath.Join(projectRoot, "config")
	configManager, err := config.NewConfigManager(nil, configPath, "test")
	if err != nil {
		log.Panic(err)
	}
	return
}

func projectRoot() string {
	_, b, _, _ := runtime.Caller(0)
	// This will be the directory of the current file (test_engine.go)
	currentDir := filepath.Dir(b)

	// Traverse up the directory tree to find go.mod
	for {
		goModPath := filepath.Join(currentDir, "go.mod")
		if _, err := os.Stat(goModPath); err == nil {
			return currentDir
		}
		parentDir := filepath.Dir(currentDir)
		if parentDir == currentDir { // Reached file system root
			log.Fatalf("go.mod not found in any parent directory")
		}
		currentDir = parentDir
	}
}
