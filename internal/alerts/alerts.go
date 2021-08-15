package alerts

import (
	"fmt"
	"os/exec"
)

const (
	cmd = "lamp createAlert --description \"%s\" --message \"%s\" --apiKey \"%s\" " +
		"--users %s --priority %s"
)

type AlertManager interface {
	SendAlert(message, desc, priority string) (err error)
}

type manager struct {
	key, user string
}

func (m *manager) SendAlert(message, desc, priority string) (err error) {
	run := fmt.Sprintf(cmd, message, desc, m.key, m.user, priority)
	_, err = exec.Command("bash", "-c", run).Output()
	return err
}

func NewAlertManager(user, key string) AlertManager {
	return &manager{
		key:  key,
		user: user,
	}
}
