package apitools

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/websocket"
	"github.com/lab5e/go-spanapi/v4"
)

// DataStream is a stream of data from the Span service. The data stream uses
// a websocket under the hood.
type DataStream interface {
	Recv() (spanapi.OutputDataMessage, error)
	Close() error
}

// NewCollectionDataStream creates a live data stream for devices in a
// collection. The context must contain the appropriate credentials.
func NewCollectionDataStream(ctx context.Context, config *spanapi.Configuration, collectionID string) (DataStream, error) {
	wsURL, err := serverURL(config)
	if err != nil {
		return nil, err
	}

	urlStr := fmt.Sprintf("%s/collections/%s/from", wsURL, collectionID)
	return newDataStream(ctx, urlStr)
}

// NewDeviceDataStream creates a live data stream for a single device. The
// context must contain the appropriate credentials.
func NewDeviceDataStream(ctx context.Context, config *spanapi.Configuration, collectionID, deviceID string) (DataStream, error) {
	wsURL, err := serverURL(config)
	if err != nil {
		return nil, err
	}

	urlStr := fmt.Sprintf("%s/collections/%s/devices/%s/from", wsURL, collectionID, deviceID)
	return newDataStream(ctx, urlStr)
}
func serverURL(config *spanapi.Configuration) (string, error) {
	if len(config.Servers) < 1 {
		return "", errors.New("no configured servers")
	}
	return strings.Replace(config.Servers[0].URL, "https", "wss", 1), nil
}
func newDataStream(ctx context.Context, urlStr string) (DataStream, error) {
	header := http.Header{}

	t := ctx.Value(spanapi.ContextAPIKeys)
	if t == nil {
		return nil, errors.New("no auth token in context")
	}
	key, ok := t.(map[string]spanapi.APIKey)
	if !ok {
		return nil, errors.New("not a token in context")
	}
	header.Add("X-API-Token", key["APIToken"].Key)
	dialer := websocket.Dialer{}
	ws, _, err := dialer.Dial(urlStr, header)
	if err != nil {
		return nil, fmt.Errorf("Error dialing websocket: %v", err)
	}
	return &wsDataStream{ws}, nil
}

type wsDataStream struct {
	ws *websocket.Conn
}

func (d *wsDataStream) Recv() (spanapi.OutputDataMessage, error) {
	for {
		_, msgBytes, err := d.ws.ReadMessage()
		if err != nil {
			return spanapi.OutputDataMessage{}, err
		}

		m := spanapi.OutputDataMessage{}
		err = json.Unmarshal(msgBytes, &m)
		if err != nil {
			return spanapi.OutputDataMessage{}, err
		}

		if m.Type != nil && *m.Type == spanapi.DATA {
			return m, nil
		}
	}
}

func (d *wsDataStream) Close() error {
	return d.ws.Close()
}
