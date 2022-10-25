package helper

import (
	"fmt"
	"os"
	"time"

	"github.com/cavaliergopher/grab/v3"
	"github.com/schollz/progressbar/v3"
)

func DownloadKubectlBinary(version string) (string, error) {
	// create client
	client := grab.NewClient()
	tempPath := os.TempDir()
	dst := tempPath + "/kubectl-" + version
	url := fmt.Sprintf("https://dl.k8s.io/release/%s/bin/linux/amd64/kubectl", version)
	req, _ := grab.NewRequest(dst, url)

	// start download
	fmt.Printf("Downloading %v...\n", req.URL())
	resp := client.Do(req)
	fmt.Printf("  %v\n", resp.HTTPResponse.Status)

	// start UI loop
	t := time.NewTicker(500 * time.Millisecond)
	defer t.Stop()
	bar := progressbar.DefaultBytes(
		resp.Size(),
		"downloading",
	)
Loop:
	for {
		select {
		case <-t.C:
			err := bar.Set(int(resp.BytesComplete()))
			if err != nil {
				return dst, err
			}
		case <-resp.Done:
			// download is complete
			break Loop
		}
	}

	// check for errors
	return dst, resp.Err()
}
