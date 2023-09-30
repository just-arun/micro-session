package pubsub

import (
	"fmt"

	"github.com/just-arun/micro-session/model"
	"github.com/just-arun/micro-session/session"
)

type siteMap struct{ ctx *model.GlobalCtx }

func SiteMap(con *model.GlobalCtx) *siteMap {
	return &siteMap{ctx: con}
}

func (st *siteMap) SubscribeUpdateSiteMap() *siteMap {
	st.ctx.NatsConnection.Subscribe("change-service-map", func(m *[]model.ServiceMap) {
		fmt.Println(m)
		// var serviceMap []model.ServiceMap
		// fmt.Println(string(m.Data))
		// err := json.Unmarshal(m.Data, &serviceMap)
		// if err != nil {
		// 	fmt.Println("ERR: ", err.Error())
		// 	return
		// }
		err := session.SiteMap().Set(st.ctx.GeneralSessionRedisDB, *m)
		if err != nil {
			fmt.Println("ERR: ", err.Error())
			return
		}
	})
	return st
}
