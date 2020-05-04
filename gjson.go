package main

import (
	"fmt"
	"github.com/tidwall/gjson"
)

var gjdata = `
[{"sha1":"aed278f0ad641bce773ad19902c6de02c203aa54","report_size":139558,"report_uri":"files/aed278f0ad641bce773ad19902c6de02c203aa54/logs/cuckoo-win7en/c0e185ee39e38e88d62fe180c89e49289bf95136/storage","pcap_uri":"files/aed278f0ad641bce773ad19902c6de02c203aa54/logs/cuckoo-win7en/c05da0659e0ee62802e1d0455222e224e017348b/storage","time":1587521537,"pcap_size":19301}]
`

func main() {
	value := gjson.Get(gjdata, "#.sha1")
	fmt.Println(value.Array()[0].Raw)
}
