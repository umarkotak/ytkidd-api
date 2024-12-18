package comfy_ui_handler

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/umarkotak/ytkidd-api/config"
	"github.com/umarkotak/ytkidd-api/utils/render"
)

func Gallery(w http.ResponseWriter, r *http.Request) {
	files, err := os.ReadDir(config.Get().ComfyUIOutputDir)
	if err != nil {
		log.Fatal(err)
	}

	imageUrls := []string{}
	for _, file := range files {
		if !file.IsDir() {
			imageUrls = append(imageUrls, fmt.Sprintf("%s/%s", "ytkidd-api-m4.cloudflare-avatar-id-1.site/comfy_ui_gallery", file.Name()))
		}
	}

	render.Response(w, r, 200, map[string]any{
		"image_urls": imageUrls,
	})
}
