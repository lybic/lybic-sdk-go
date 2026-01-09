package lybic

type SandboxStatus string

const (
	SandboxRunning SandboxStatus = "RUNNING"
	SandboxPending SandboxStatus = "PENDING"
	SandboxStopped SandboxStatus = "STOPPED"
	SandboxError   SandboxStatus = "ERROR"
)

type SandboxStatusDto struct {
	Status SandboxStatus `json:"status"`
}
