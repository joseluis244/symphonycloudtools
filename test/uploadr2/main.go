package main

import (
	"fmt"
	"joseluis244/symphonycloudtools"
	"sync"
	"time"
)

func main() {
	license := "f13507ac823268eee7fc0262a845c9d8e324f9442249e9df067a5af8825b3039331fcd0c4463eb1c03967529089ac0d27aa7e45af5a55736a89f7373cea529e6d21c1e62788c82bcc74f940b669aedb15b046ea8dc99bcf8e932eec718f013c27cd981ca192fa6a5b4945bf073e125ece3c30aa345513b712874d80558924136c25a925154a61ce71cee9fca097db0afa626e8ad5b104ab5f97593ebef7f4c8012e6124555b7b5b663d1f3cd0d61ead55db02091cbe46fe977de9567c9705aaee92b59ac8b95be4bad9f37bab837c61bc0803e249bcdc45f9903a608374fe445a5513c038949ffc1494c16d3ace1682b04242200004d3fcd31380f073b5338eef84a7d27f3661528c0ce681d95a23a257f4a71140fdcdfaa04a5bbceffc1d43c0882e9a8c88bd90fe4cbca17efbb11fb544565b420134326be88102217a8b19e852a389c70f7c8a3ab44762222091af7f382957cac25311d524de336006675eb194b133d182d1582d61e1b7c72514ed70efd8c3947bd77db50d8e8c939218e4168cd95001677caaa2b645ff89cc5b0368d464186b5bc9c5970a6dd2fe9"

	start := time.Now()

	symphonycloudtools.Init(license)
	var dcmurl *string
	var imgurl *string
	var zipurl *string
	WG := sync.WaitGroup{}
	WG.Add(3)
	go func() {
		defer WG.Done()
		url, err := symphonycloudtools.R2.UploadDCM("/Users/josecamacho/Pictures/mediglobe2.svg", "111-222-333")
		if err != nil {
			fmt.Println(err)
		}
		dcmurl = &url
	}()
	go func() {
		defer WG.Done()
		url, err := symphonycloudtools.R2.UploadIMG("/Users/josecamacho/Pictures/mediglobe2.svg", "111-222-333")
		if err != nil {
			fmt.Println(err)
		}
		imgurl = &url
	}()
	go func() {
		defer WG.Done()
		url, err := symphonycloudtools.R2.UploadZIP("/Users/josecamacho/Pictures/mediglobe2.svg", "111-222-333")
		if err != nil {
			fmt.Println(err)
		}
		zipurl = &url
	}()
	WG.Wait()
	fmt.Println("DCM", *dcmurl)
	fmt.Println("IMG", *imgurl)
	fmt.Println("ZIP", *zipurl)
	elapsed := time.Since(start)
	fmt.Println("Execution time:", elapsed)
}
