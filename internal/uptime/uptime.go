package uptime

import (
	"fmt"

	"github.com/njfamirm/sitatus/pkg/config"
	"github.com/njfamirm/sitatus/pkg/uptime"
)

func Uptime() {
	for _, site := range config.GetSiteList() {
		ping, err := uptime.Uptime(site.Url)

		if err != nil {
			fmt.Println(fmt.Errorf("(uptime): visit %s failed", site.Name))
		} else {
			fmt.Printf("(uptime): visit %s with %d ping\n", site.Name, ping)
		}
	}
}
