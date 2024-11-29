package ping_handler

import (
	"net/http"

	"github.com/umarkotak/ytkidd-api/utils/render"
	"github.com/umarkotak/ytkidd-api/worker"
)

func Ping(w http.ResponseWriter, r *http.Request) {
	worker.Enqueue(r.Context(), "PING_TEST", map[string]any{"bool": true}, nil)

	render.Response(w, r, 200, map[string]any{
		"ping": "pong",
	})
}

func ToDo(w http.ResponseWriter, r *http.Request) {
	render.Response(w, r, 200, map[string]any{
		"todo": "todo",
	})
}
