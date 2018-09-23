package gen_openapi_proto

import (
	"fmt"
	"git.speakin.mobi/infrastructure/api_serv.git/api_def"
	"reflect"
)

func NewModel(value reflect.Value, name string) *ModelDefine {
	typ := value.Type()
	switch value.Kind() {
	case reflect.Struct:
		model := ModelDefine{
			Type: "object",
			Name: name,
		}

		for i := 0; i < value.NumField(); i++ {
			field := value.Field(i)
			if field.Kind() == reflect.Ptr {
				field = field.Elem()
			}
			item := NewModel(field, typ.Field(i).Tag.Get("json"))
			model.Properties = append(model.Properties, item)
			if typ.Field(i).Tag.Get("binding") == "required" {
				model.Require = append(model.Require, item)
			}
		}
		return &model

	case reflect.Slice:
		model := ModelDefine{
			Type: "array",
		}
		var elem reflect.Value
		if value.Len() > 0 {
			elem = value.Index(0)
		} else {
			elem = reflect.Zero(typ.Elem())
		}
		if elem.Kind() == reflect.Ptr {
			elem = elem.Elem()
		}
		model.Items = NewModel(elem, "")
		if model.Items.Type == "object" {
			model.Name = elem.Type().Name()
		}
		return &model
	case reflect.Ptr:
		return NewModel(value.Elem(), name)
	case reflect.Int64:
		return &ModelDefine{
			Type:    "int",
			Format:  "int64",
			Example: fmt.Sprint(value.Interface()),
		}
	case reflect.Int32:
		return &ModelDefine{
			Type:    "int",
			Format:  "int32",
			Example: fmt.Sprint(value.Interface()),
		}
	case reflect.Int:
		return &ModelDefine{
			Type:    "int",
			Example: fmt.Sprint(value.Interface()),
		}
	case reflect.Float64:
		return &ModelDefine{
			Type:    "number",
			Format:  "double",
			Example: fmt.Sprint(value.Interface()),
		}
	case reflect.Float32:
		return &ModelDefine{
			Type:    "number",
			Format:  "float",
			Example: fmt.Sprint(value.Interface()),
		}
	case reflect.Bool:
		return &ModelDefine{
			Type:    "boolean",
			Example: fmt.Sprint(value.Interface()),
		}
	case reflect.String:
		return &ModelDefine{
			Type:    "string",
			Example: fmt.Sprint(value.Interface()),
		}
	default:
		panic(fmt.Sprintf("incorrect type %v", value.Kind()))
	}

}

func NewGenerate(version, title string, factories ...api_def.ApiFactory) *PackageInfo {
	info := PackageInfo{
		Version: version,
		Title:   title,
	}
	for _, factory := range factories {
		var cateName = factory.GetApiCateName()
		allApiIface := factory.GetAllApi()
		for _, apiIface := range allApiIface {
			var api Api
			req := reflect.ValueOf(apiIface.GetRequestDemo()).Elem()
			if req.Kind() == reflect.Ptr {
				req = req.Elem()
			}
			tName := req.Type().Name()
			api.Req = &ApiRequest{
				Schema: NewModel(req, tName),
			}
			api.Url = TypeNameToUrl(tName)
			//resp := apiIface.GetResponseDemo()
		}
	}
}
