package uptime

import (
	"fmt"

	"github.com/njfamirm/sitatus/internal/analytic"
	"github.com/njfamirm/sitatus/pkg/config"
	"github.com/njfamirm/sitatus/pkg/uptime"
)

func Uptime() {
	for _, site := range config.GetSiteList() {
		ping, err := uptime.Uptime(site.Url)

		if err != nil {
			fmt.Println(fmt.Errorf("(uptime): visit %s failed", site.Name))
			ping = 0
		} else {
			fmt.Printf("(uptime): visit %s with %dms ping\n", site.Name, ping)
		}

		// write analytic
		writeAnalyticErr := analytic.WriteAnalytic(site.Name, ping)
		if writeAnalyticErr != nil {
			fmt.Println(fmt.Errorf("(uptime): write analytic %s failed", site.Name))
			panic(writeAnalyticErr)
		}
	}
}
