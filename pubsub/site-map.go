package pubsub

import (
	"github.com/just-arun/micro-session/model"
	"github.com/just-arun/micro-session/session"
)

type siteMap struct{ ctx *model.GlobalCtx }

func SiteMap(con *model.GlobalCtx) *siteMap {
	return &siteMap{ctx: con}
}

func (st *siteMap) SubscribeUpdateSiteMap() *siteMap {
	st.ctx.NatsConnection.Subscribe("change-service-map", func(m *[]model.ServiceMap) {
		err := session.SiteMap().Set(st.ctx.GeneralSessionRedisDB, *m)
		if err != nil {
			return
		}
	})
	return st
}
