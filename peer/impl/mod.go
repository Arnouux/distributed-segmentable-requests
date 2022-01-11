package impl

// file should be on root as final

import (
	"fmt"
	"net/http"
	"time"
)

type Node struct {
}

func NewNode() *Node {
	return &Node{}
}

func (n *Node) SendPrepDwnldResp(dest string, relays map[uint]string) error {
	//TODO send PrepareDownloadReply to dest &&
	//Todo: stats on how fast get data from url server (no more than upperbound bytes)

	var client = http.Client{
		Timeout: 2 * time.Second,
	}

	time := time.Now()
	// Ping the server
	// ping artepweb-vh.akamaihd.net
	url := "arteptweb-vh.akamaihd.net"
	req, err := http.NewRequest("HEAD", url, nil)
	if err != nil {
		return err
	}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	resp.Body.Close()
	fmt.Println(resp.StatusCode)
	fmt.Println(time)

	//either m3u8 homemade or:
	// mp4 or other using:
	//https://stackoverflow.com/questions/27844307/how-to-download-only-the-beginning-of-a-large-file-with-go-http-client-and-app
	//=>implying dynamical size of chunks, to add in route calculationsé & dwnld requests (& type) -> need to decide quickly

	return nil
}
