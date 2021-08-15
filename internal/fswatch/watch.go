package fswatch

import (
	"log"
	"time"

	"github.com/radovskyb/watcher"

	"nodeChecker/internal/alerts"
)

type ReportsWatcher interface {
	Watch()
}

type rwatcher struct {
	folder       string
	alertmanager alerts.AlertManager
}

func (rw *rwatcher) Watch() {
	w := watcher.New()
	w.SetMaxEvents(1)
	w.FilterOps(watcher.Create)

	go func() {
		for {
			select {
			case event := <-w.Event:
				log.Println(event)
				rw.alertmanager.SendAlert(event.String(), "P1")
			case err := <-w.Error:
				log.Fatalln(err)
			case <-w.Closed:
				return
			}
		}
	}()

	if err := w.AddRecursive(rw.folder); err != nil {
		log.Fatalln(err)
	}
	if err := w.Start(time.Millisecond * 100); err != nil {
		log.Fatalln(err)
	}
}
