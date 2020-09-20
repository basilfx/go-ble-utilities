package characteristics

import (
	"bufio"
	"io"
	"sync"

	"github.com/acomagu/bufpipe"
	"github.com/basilfx/go-utilities/observable"
	"github.com/go-ble/ble"

	log "github.com/sirupsen/logrus"
)

type commanderStream struct {
	reader io.ReadCloser
	writer io.WriteCloser
}

// BytesCommandHandlerFunc is a function that handles bytes.
type BytesCommandHandlerFunc func(command []byte) *[]byte

// CommandResponseHandlerFunc returns a writer and a notifier that processes
// commands separated by newline characters. Every line is then passed onto the
// provided BytesCommandHandlerFunc. It is agnostic to the type of data.
// Every command that is processed can return a response, or none.
func CommandResponseHandlerFunc(f BytesCommandHandlerFunc) (ble.WriteHandlerFunc, ble.NotifyHandlerFunc) {
	responder := observable.New()

	streams := make(map[string]commanderStream)
	lock := sync.Mutex{}

	writer := ble.WriteHandlerFunc(func(req ble.Request, rsp ble.ResponseWriter) {
		addr := req.Conn().RemoteAddr().String()

		lock.Lock()
		defer lock.Unlock()

		// Get or create a stream.
		stream, ok := streams[addr]

		if !ok {
			reader, writer := bufpipe.New(nil)

			stream = commanderStream{
				reader,
				writer,
			}

			streams[addr] = stream

			// Start a goroutine that reads from the pipe until it is closed.
			go func() {
				log.Debugf("New stream goroutine started for '%s'.", addr)

				scanner := bufio.NewScanner(reader)

				for scanner.Scan() {
					request := scanner.Bytes()

					response := f(request)

					if response == nil {
						continue
					}

					for _, s := range slice(append(*response, '\n'), 20) {
						responder.SetValue(s)
					}
				}

				log.Debugf("Streamer stopped.")
			}()
		}

		// Write incoming data to the stream.
		_, err := stream.writer.Write(req.Data())

		if err != nil {
			log.Errorf("Error while writing to stream: %v", err)

			delete(streams, addr)

			stream.reader.Close()
			stream.writer.Close()
		}
	})

	notifier := ObservableNotifyHandlerFunc(responder)

	return writer, notifier
}
