package wsrender

import (
	"context"
	"encoding/json"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
	"github.com/umarkotak/ytkidd-api/model"
)

type (
	WsConnWrapper struct {
		Mu   sync.Mutex
		Conn *websocket.Conn
	}
)

func Render(ctx context.Context, wsConnWrapper *WsConnWrapper, messageType string, data any) error {
	dataJson, err := json.Marshal(map[string]any{
		"message_type": messageType,
		"data":         data,
		"success":      true,
	})
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		return err
	}

	wsConnWrapper.Mu.Lock()
	defer wsConnWrapper.Mu.Unlock()

	err = wsConnWrapper.Conn.WriteMessage(websocket.TextMessage, dataJson)
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		return err
	}

	return nil
}

func Error(ctx context.Context, wsConnWrapper *WsConnWrapper, err error, customErrMsg string) error {
	jxErr, ok := model.ERR_MAP[err]
	if !ok {
		jxErr = model.ERR_MAP[model.ErrUnprocessableEntity]
	}

	errMsg := jxErr.EN

	if customErrMsg != "" {
		errMsg = customErrMsg
	}

	dataJson, err := json.Marshal(map[string]interface{}{
		"message_type":   "websocket_error",
		"error_code":     jxErr.ErrorCode,
		"error_message":  errMsg,
		"internal_error": err.Error(),
		"success":        false,
	})
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		return err
	}

	wsConnWrapper.Mu.Lock()
	defer wsConnWrapper.Mu.Unlock()

	err = wsConnWrapper.Conn.WriteMessage(websocket.TextMessage, dataJson)
	if err != nil {
		logrus.WithContext(ctx).Error(err)
		return err
	}

	return nil
}
