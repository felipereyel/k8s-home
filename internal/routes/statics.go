package routes

import (
	"net/http"
	"k8s-home/internal/embeded"

	"github.com/gofiber/fiber/v2/middleware/filesystem"
)

var staticsHandler = filesystem.New(filesystem.Config{
	Root:       http.FS(embeded.EmbedDirStatic),
	PathPrefix: "statics",
	MaxAge:     60 * 60,
})
