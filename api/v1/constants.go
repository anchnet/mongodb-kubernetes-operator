package v1

type MongoDBCommunityRole string

const (
	MongoDBClusterNodeRolePrimary   MongoDBCommunityRole = "Primary"
	MongoDBClusterNodeRoleSecondary MongoDBCommunityRole = "Secondary"
	MongoDBClusterNodeRoleNone      MongoDBCommunityRole = "None"
)

type MongoDBCommunityClusterStatus string

const (
	ClusterStatusOK MongoDBCommunityClusterStatus = "Healthy"
	ClusterStatusKO MongoDBCommunityClusterStatus = "Failed"
)
const (
	DefaultMongoDBPort   = 27017
	MongoDBRolePRIMARY   = "primary"
	MongoDBRoleSECONDARY = "secondary"
)
const (
	MongoDBCommandIsMaster = "isMaster"
)
