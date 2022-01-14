package impl

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"
)

func GetFileSize(url string) (int64, error) {
	//shamelessy from https://www.socketloop.com/tutorials/golang-get-download-file-size

	resp, err := http.Head(url)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	// Is our request ok?

	if resp.StatusCode != http.StatusOK {
		fmt.Println(resp.Status)
		return 0, err
		// exit if not ok
	}

	// the Header "Content-Length" will let us know
	// the total file size to download
	size, _ := strconv.Atoi(resp.Header.Get("Content-Length"))
	downloadSize := int64(size)
	return downloadSize, nil

}

func DownloadPartFromExternal(url string, start uint, end uint, IOoutput io.Writer, toStorageCopy io.Writer) error {
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Range", "bytes="+fmt.Sprint(start)+"-"+fmt.Sprint(end))
	fmt.Println(req)
	var client http.Client
	resp, err := client.Do(req)

	if err != nil {
		println("couldnt do the HTTP request when trying to download/test bdwdth")
		return err
	}

	defer resp.Body.Close()

	r := resp.Body
	r2 := io.TeeReader(r, toStorageCopy)
	written, err := io.Copy(IOoutput, r2) //check io.PipeWriter to tee data into storage & to requesting node
	if written != int64(end-start) {
		return errors.New("did Not Download requested amount")
	}
	return err
}

func TestExternalBandwidth(url string, upperBound uint) float32 {
	t0 := time.Now().Unix()
	DownloadPartFromExternal(url, 0, upperBound, io.Discard, io.Discard)
	t1 := time.Now().Unix()
	
	fmt.Println(t0, t1)
	return float32(upperBound) / float32(t0-t1)
}

func Server_to_test(upperBound uint) {
	http.HandleFunc("/stats", func(w http.ResponseWriter, r *http.Request) {
		for i := 0; i < int(upperBound); i++ {
			fmt.Fprintf(w, "A")
		}
	})

	if err := http.ListenAndServe("127.0.0.1:8080", nil); err != nil {
		//we probably already launched it
		println(err, "while launching server for linkState test")
	}
}
