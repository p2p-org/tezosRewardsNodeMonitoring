package fswatch

import (
	"fmt"
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
				rw.alertmanager.SendAlert("TRDReport", fmt.Sprintf("Report %s is created", event.FileInfo.Name()), "P4")
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

func NewReportsWatcher(folder string, alertmanager alerts.AlertManager) ReportsWatcher {
	return &rwatcher{
		folder:       folder,
		alertmanager: alertmanager,
	}
}
