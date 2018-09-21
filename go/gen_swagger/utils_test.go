package gen_openapi_proto

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTypeNameToUrl(t *testing.T) {
	assert.Equal(t, TypeNameToUrl("UserDetail"), "user_detail")
	assert.Equal(t, TypeNameToUrl("OneToOne"), "one_to_one")
	assert.Equal(t, TypeNameToUrl("_UserDetail"), "_user_detail")
	assert.Equal(t, TypeNameToUrl("userDetail"), "user_detail")
	assert.Equal(t, TypeNameToUrl("UserDetailID"), "user_detail_id")
	assert.Equal(t, TypeNameToUrl("NameHTTPtest"), "name_http_test")
	assert.Equal(t, TypeNameToUrl("IDandValue"), "id_and_value")
	assert.Equal(t, TypeNameToUrl("toyorm.User.field"), "toyorm_user_field")
}
