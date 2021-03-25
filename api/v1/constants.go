package v1

type MongoDBCommunityRole string

const (
	MongoDBClusterNodeRoleMaster MongoDBCommunityRole = "Master"
	MongoDBClusterNodeRoleSlave  MongoDBCommunityRole = "Slave"
	MongoDBClusterNodeRoleNone   MongoDBCommunityRole = "None"
)

type ClusterStatus string

const (
	ClusterStatusOK ClusterStatus = "Healthy"
	ClusterStatusKO ClusterStatus = "Failed"
)
