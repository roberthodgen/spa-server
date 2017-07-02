package spa_server

import (
	"net/http"
	"os"
	"path/filepath"
)

// Serve from a public directory with specific index
type spaHandler struct {
	publicDir string // The directory from which to serve
	indexFile string // The fallback/default file to serve
}

func (rh *spaHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	serve(w, r, rh.publicDir, rh.indexFile)
}

// Returns a request handler (http.Handler) that serves a single
// page application from a given public directory (publicDir).
//
// It falls back to a supplied index (indexFile) when either condition is true:
// (1) Request (file) path is not found
// (2) Request path is a directory
func SpaHandler(publicDir string, indexFile string) http.Handler {
	return &spaHandler{publicDir, indexFile}
}

func serve(w http.ResponseWriter, r *http.Request, publicDir string, indexFile string) {
	// Clean request path; get a file path in the public directory
	rp := filepath.Join(publicDir, filepath.Clean(r.URL.Path))

	// Attempt to get info about the request path...
	if info, err := os.Stat(rp); err != nil {
		// Not found or otherwise failed to read the file
		http.ServeFile(w, r, filepath.Join(publicDir, indexFile))
		return
	} else if info.IsDir() {
		// Request path is a directory
		http.ServeFile(w, r, filepath.Join(publicDir, indexFile))
		return
	}

	// Serve the requested path
	http.ServeFile(w, r, rp)
}
