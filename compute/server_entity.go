package compute

const (
	ServerStatusRunning   = 1
	ServerStatusStopped   = 2
	ServerStatusSuspended = 3
	ServerStatusStarting  = 4
	ServerStatusStopping  = 5
	ServerStatusError     = 6
	ServerStatusUpgrading = 9

	ApplicationStatusUnhealthy = 7
	ApplicationStatusHealthy   = 8

	ClusterStatusHealthy     = 14
	ClusterStatusWorking     = 15
	ClusterStatusUnhealthy   = 16
	ClusterStatusUnavailable = 17

	TaskStatusCreating   = 18
	TaskStatusNeutral    = 10
	TaskStatusCordoned   = 11
	TaskStatusDraining   = 12
	TaskStatusRebuilding = 13
)

type ServerAction struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Command string `json:"command"`
	Sorting int    `json:"sorting"`
}

type ServerStatus struct {
	ID      int            `json:"id"`
	Name    string         `json:"name"`
	Key     string         `json:"key"`
	Actions []ServerAction `json:"actions"`
}
