package dialect

import (
	"database/sql"
	"sync"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Dialect struct {
	Matcher   Matcher
	Config    Connector
	IsDefault bool
}

var (
	dialects       []*Dialect
	defaultDialect *Dialect
	dialectsMu     sync.Mutex
)

type Matcher interface {
	MatchName(string) bool
	Decode([]interface{}) (Connector, error)
}

const (
	QueryAppName             = "zitadel_queries"
	EventstorePusherAppName  = "zitadel_es_pusher"
	ProjectionSpoolerAppName = "zitadel_projection_spooler"
	defaultAppName           = "zitadel"
)

// DBPurpose is what the resulting connection pool is used for.
type DBPurpose int

const (
	DBPurposeQuery DBPurpose = iota
	DBPurposeEventPusher
	DBPurposeProjectionSpooler
)

func (p DBPurpose) AppName() string {
	switch p {
	case DBPurposeQuery:
		return QueryAppName
	case DBPurposeEventPusher:
		return EventstorePusherAppName
	case DBPurposeProjectionSpooler:
		return ProjectionSpoolerAppName
	default:
		return defaultAppName
	}
}

type Connector interface {
	Connect(useAdmin bool, pusherRatio, spoolerRatio float64, purpose DBPurpose) (*sql.DB, *pgxpool.Pool, error)
	Password() string
	Database
}

type Database interface {
	DatabaseName() string
	Username() string
	Type() string
	Timetravel(time.Duration) string
}

func Register(matcher Matcher, config Connector, isDefault bool) {
	dialectsMu.Lock()
	defer dialectsMu.Unlock()

	d := &Dialect{Matcher: matcher, Config: config}

	if isDefault {
		defaultDialect = d
		return
	}

	dialects = append(dialects, d)
}

func SelectByConfig(config map[string]interface{}) *Dialect {
	for key := range config {
		for _, d := range dialects {
			if d.Matcher.MatchName(key) {
				return d
			}
		}
	}

	return defaultDialect
}
