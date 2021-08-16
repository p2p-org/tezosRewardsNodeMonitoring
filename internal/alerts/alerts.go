package alerts

import (
	"fmt"
	"os/exec"
)

const (
	sendAlertCommand = "lamp createAlert --description \"%s\" --message \"%s\" --apiKey \"%s\" " +
		"--users %s --priority %s"
	sendHeartbeatCommand = "lamp pingHeartbeat --name %s --apiKey %s"
)

type AlertManager interface {
	SendAlert(message, desc, priority string) (err error)
	SendHeartbeat(name string) (err error)
}

type manager struct {
	key, user string
}

func (m *manager) SendAlert(message, desc, priority string) (err error) {
	run := fmt.Sprintf(sendAlertCommand, message, desc, m.key, m.user, priority)
	_, err = exec.Command("bash", "-c", run).Output()
	return err
}

func (m *manager) SendHeartbeat(name string) (err error) {
	run := fmt.Sprintf(sendHeartbeatCommand, name, m.key)
	_, err = exec.Command("bash", "-c", run).Output()
	return err
}

func NewAlertManager(user, key string) AlertManager {
	return &manager{
		key:  key,
		user: user,
	}
}
