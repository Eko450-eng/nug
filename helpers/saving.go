package helpers

import (
	"fmt"
	"net/http"
	"os"

	gap "github.com/muesli/go-app-paths"
)

func SyncToWebDav() {
	url := os.Getenv("NUG_URL")
	username := os.Getenv("NUG_USERNAME")
	password := os.Getenv("NUG_PASSWORD")

	if url != "" || username != "" || password != "" {
		scope := gap.NewScope(gap.User, "nug")
		dirs, err := scope.DataDirs()
		appPath := dirs[0] + "/nug.db"

		file, err := os.Open(appPath)
		req, err := http.NewRequest("PUT", url, file)

		if username != "" && password != "" {
			req.SetBasicAuth(username, password)
		}

		req.Header.Set("Overwrite", "T")

		client := &http.Client{}
		resp, err := client.Do(req)
		CheckErr(err)
		defer resp.Body.Close()
		CheckErr(err)

		if resp.StatusCode != http.StatusCreated && resp.StatusCode != http.StatusOK {
			LogToFile(fmt.Sprintf("failed to upload, status: %s", resp.Status))
		}
	}
}
