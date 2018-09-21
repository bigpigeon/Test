package gen_openapi_proto

type ModelDefine struct {
	Name    string
	Type    string
	Format  string
	Example string
	Items   *ModelDefine

	Require    []*ModelDefine
	Properties []*ModelDefine
}

type ApiRequest struct {
	Schema *ModelDefine
}

type ApiResponse struct {
	Description string
	Model       *ModelDefine
}

type Api struct {
	Tags   []string
	Url    string
	Method string
	Req    *ApiRequest
	Resp   map[string]ApiResponse
}

type Tags map[string][]*Api

func (t Tags) Register(api *Api, tag string) {
	t[tag] = append(t[tag], api)
	api.Tags = append(api.Tags, tag)
}

type PackageInfo struct {
	Server struct {
		Url string
	}
	Version string
	Title   string
	Contact string
	Tags    Tags
}
