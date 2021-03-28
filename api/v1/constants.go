package v1

type MongoDBCommunityRole string

const (
	MongoDBClusterNodeRolePrimary   MongoDBCommunityRole = "Primary"
	MongoDBClusterNodeRoleSecondary MongoDBCommunityRole = "Secondary"
	MongoDBClusterNodeRoleNone      MongoDBCommunityRole = "None"
)

type ClusterStatus string

const (
	ClusterStatusOK ClusterStatus = "Healthy"
	ClusterStatusKO ClusterStatus = "Failed"
)
const (
	DefaultMongoDBPort   = 27017
	MongoDBRolePRIMARY   = "primary"
	MongoDBRoleSECONDARY = "secondary"
)
const (
	MongoDBCommandIsMaster = "isMaster"
)
