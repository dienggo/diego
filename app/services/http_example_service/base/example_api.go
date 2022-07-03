package base

import (
	"go_base_project/app/base"
)

type ExampleApi struct{ ExampleConfig }

func (e ExampleApi) Get(endpoint string) base.NetClient {
	return base.HttpService().Get().Url(e.BaseUrl()+endpoint).AddHeader("Content-Type", "application/json")
}

func (e ExampleApi) Post(endpoint string) base.NetClient {
	return base.HttpService().Post().Url(e.BaseUrl()+endpoint).AddHeader("Content-Type", "application/json")
}

func (e ExampleApi) Put(endpoint string) base.NetClient {
	return base.HttpService().Put().Url(e.BaseUrl()+endpoint).AddHeader("Content-Type", "application/json")
}

func (e ExampleApi) Delete(endpoint string) base.NetClient {
	return base.HttpService().Delete().Url(e.BaseUrl()+endpoint).AddHeader("Content-Type", "application/json")
}