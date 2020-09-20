package characteristics

import (
	"github.com/basilfx/go-utilities/observable"
	"github.com/go-ble/ble"

	log "github.com/sirupsen/logrus"
)

// ObservableReadHandlerFunc creates a read handler function that returns the
// value of the given observable.
func ObservableReadHandlerFunc(o *observable.Observable) ble.ReadHandlerFunc {
	return ble.ReadHandlerFunc(func(req ble.Request, rsp ble.ResponseWriter) {
		_, err := rsp.Write(o.GetValue().([]byte))

		if err != nil {
			log.Errorf("Unable to write: %v", err)
			return
		}
	})
}

// ObservableNotifyHandlerFunc creates a notification handler function that
// emites changes notifications when the observable changes.
func ObservableNotifyHandlerFunc(o *observable.Observable) ble.NotifyHandlerFunc {
	return ble.NotifyHandlerFunc(func(req ble.Request, n ble.Notifier) {
		r, c := o.Register()
		defer o.Unregister(r)

		for {
			select {
			case <-n.Context().Done():
				return
			case v := <-c:
				_, err := n.Write(v.([]byte))

				if err != nil {
					log.Errorf("Unable to notify: %v", err)
					return
				}
			}
		}
	})
}
