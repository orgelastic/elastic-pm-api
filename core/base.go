package core

import (
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stewie1520/elasticpmapi/config"
	"github.com/stewie1520/elasticpmapi/daos"
	"github.com/stewie1520/elasticpmapi/db"
	hook "github.com/stewie1520/elasticpmapi/hooks"
)

const (
	DefaultDataMaxOpenConns int = 120
	DefaultDataMaxIdleConns int = 20
)

var _ App = (*BaseApp)(nil)

type BaseAppConfig struct {
	*config.Config
	IsDebug          bool
	DataMaxOpenConns int
	DataMaxIdleConns int
}

type BaseApp struct {
	config BaseAppConfig
	dao    *daos.Dao

	onAfterAccountCreated *hook.Hook[*AccountCreatedEvent]
	onUnauthorizedAccess  *hook.Hook[*UnauthorizedAccessEvent]
}

func NewBaseApp(config BaseAppConfig) *BaseApp {
	app := &BaseApp{
		config:                config,
		onAfterAccountCreated: &hook.Hook[*AccountCreatedEvent]{},
		onUnauthorizedAccess:  &hook.Hook[*UnauthorizedAccessEvent]{},
	}

	return app
}

func (app *BaseApp) Bootstrap() error {
	if err := app.initDatabase(); err != nil {
		return err
	}

	if err := initSuperToken(app); err != nil {
		return err
	}

	return nil
}

func (app *BaseApp) IsDebug() bool {
	return app.config.IsDebug
}

func (app *BaseApp) Dao() *daos.Dao {
	return app.dao
}

func (app *BaseApp) DB() *pgxpool.Pool {
	if app.Dao() == nil {
		return nil
	}

	return app.Dao().DB()
}

func (app *BaseApp) Config() *config.Config {
	return app.config.Config
}

func (app *BaseApp) OnAfterAccountCreated() *hook.Hook[*AccountCreatedEvent] {
	return app.onAfterAccountCreated
}

func (app *BaseApp) OnUnauthorizedAccess() *hook.Hook[*UnauthorizedAccessEvent] {
	return app.onUnauthorizedAccess
}

func (app *BaseApp) initDatabase() error {
	maxOpenConns := DefaultDataMaxOpenConns
	maxIdleConns := DefaultDataMaxIdleConns

	if app.config.DataMaxOpenConns > 0 {
		maxOpenConns = app.config.DataMaxOpenConns
	}

	if app.config.DataMaxIdleConns > 0 {
		maxIdleConns = app.config.DataMaxIdleConns
	}

	pool, err := db.NewPostgresDBX(
		app.config.DatabaseURL,
		db.WithConnMaxIdleTime(time.Duration(maxIdleConns)),
		db.WithMaxOpenConns(maxOpenConns),
	)

	if err != nil {
		return err
	}

	app.dao = daos.New(pool)

	return nil
}
