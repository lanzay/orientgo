package orient

import (
	"github.com/dyy18/orientgo/oschema"
)

const (
	ProtoBinary = "binary"
)

type DatabaseType string
type StorageType string

const (
	DocumentDB DatabaseType = "document"
	GraphDB    DatabaseType = "graph"

	Persistent StorageType = "plocal"
	Volatile   StorageType = "memory"
)

var (
	protos = make(map[string]func(addr string) (DBConnection, error))
)

func RegisterProto(name string, dial func(addr string) (DBConnection, error)) {
	protos[name] = dial
}

type DBConnection interface {
	DatabaseExists(name string, storageType StorageType) (bool, error)
	CreateDatabase(name string, dbType DatabaseType, storageType StorageType) error
	DropDatabase(name string, storageType StorageType) error
	ListDatabases() (map[string]string, error)

	ConnectToServer(user, pass string) error
	OpenDatabase(name string, dbType DatabaseType, user, pass string) error

	Close() error
	CloseDatabase() error
	Size() (int64, error)
	ReloadSchema() error
	GetClasses() map[string]*oschema.OClass

	AddCluster(clusterName string) (clusterID int16, err error)
	DropCluster(clusterName string) (err error)

	CreateRecord(doc *oschema.ODocument) (err error)
	DeleteRecordByRID(rid string, recVersion int32) error
	DeleteRecordByRIDAsync(rid string, recVersion int32) error
	GetRecordByRID(rid oschema.ORID, fetchPlan string) (docs []*oschema.ODocument, err error)
	UpdateRecord(doc *oschema.ODocument) error
	CountRecords() (int64, error)

	CreateScriptFunc(fnc Function) error
	DeleteScriptFunc(name string) error
	UpdateScriptFunc(name string, script string) error
	CallScriptFunc(result interface{}, name string, params ...interface{}) (Records, error)

	SQLQuery(result interface{}, fetchPlan *FetchPlan, sql string, params ...interface{}) (recs Records, err error)
	SQLCommand(result interface{}, sql string, params ...interface{}) (recs Records, err error)

	ExecScript(result interface{}, lang ScriptLang, script string, params ...interface{}) (recs Records, err error)
}