package irisadaptor

import (
	"github.com/veerdone/gsecurity"
	"net/http"
	"time"
)

type irisAdaptor struct {
	iris.Context
}

func New(c iris.Context) *irisAdaptor {
	return &irisAdaptor{
		Context: c,
	}
}

func (i *irisAdaptor) GetFromHeader(tokenName string) string {
	return i.Context.Request().Header.Get(tokenName)
}

func (i *irisAdaptor) GetFromQuery(tokenName string) string {
	return i.URLParam(tokenName)
}

func (i *irisAdaptor) GetFromCookie(tokenName string) string {
	return i.GetCookie(tokenName)
}

func (i *irisAdaptor) SetHeader(headerName, headerValue string) {
	i.Context.Header(headerName, headerValue)
}

func (i *irisAdaptor) Get(key string) interface{} {
	return i.Context.Values().Get(key)
}

func (i *irisAdaptor) Set(key string, v interface{}) {
	i.Context.Values().Set(key, v)
}

func (i *irisAdaptor) SetCookie(conf gsecurity.Config, token string) {
	sameSite := http.SameSiteDefaultMode
	switch conf.Cookie.SameSite {
	case "None":
		sameSite = http.SameSiteNoneMode
	case "Lax":
		sameSite = http.SameSiteLaxMode
	case "strict":
		sameSite = http.SameSiteStrictMode
	}

	c := http.Cookie{
		Name:     conf.TokenName,
		Value:    token,
		Path:     conf.Cookie.Path,
		Domain:   conf.Cookie.Domain,
		Expires:  time.Now().Add(time.Second * time.Duration(conf.Timeout)),
		Secure:   conf.Cookie.Secure,
		HttpOnly: conf.Cookie.HttpOnly,
		SameSite: sameSite,
	}
	i.Context.SetCookie(&c)
}
