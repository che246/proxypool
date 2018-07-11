package getter

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strings"

	"github.com/Aiicy/ProxyPool/pkg/models"
)

// IP66 get ip from 66ip.cn
func IP66() (result []*models.IP) {
	var ExprIP = regexp.MustCompile(`((25[0-5]|2[0-4]\d|((1\d{2})|([1-9]?\d)))\.){3}(25[0-5]|2[0-4]\d|((1\d{2})|([1-9]?\d)))\:([0-9]+)`)

	pollURL := "http://www.66ip.cn/mo.php?tqsl=100"
	resp, err := http.Get(pollURL)
	if err != nil {
		log.Println(err)
		return
	}

	if resp.StatusCode != 200 {
		log.Println(err)
		return
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	body_ips := string(body)
	ips := ExprIP.FindAllString(body_ips, 100)

	for index := 0; index < len(ips); index++ {
		ip := models.NewIP()
		ip.Data = strings.TrimSpace(ips[index])
		ip.Type = "http"
		fmt.Printf("ip = %s, type = %s\n", ip.Data, ip.Type)
		result = append(result, ip)
	}

	log.Println("IP66 done.")
	return
}
