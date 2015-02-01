package obinary

// --------

type ODatabase struct {
	Name     string
	Typ      string // DocumentDbType or GraphDbType
	Clusters []OCluster
	ClustCfg []byte
}

type OCluster struct {
	Name string
	Id   int16 // TODO: maybe change to int?
}

// --------

//
// ClientOptions should be created by the end user to configure details and
// options needed when opening a database or connecting to the OrientDB server
//
type ClientOptions struct {
	ServerHost      string
	ServerPort      string
	ClusteredConfig string // TODO: needs research - what goes here?; currently not used
}

// --------

type OServerException struct {
	Class   string
	Message string
}