package gatewayhelpers

import (
	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/proto"
)

func Validate(ctx *gin.Context, requestStruct proto.Message) error {
	if v, ok := requestStruct.(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return err
		}
	}

	return nil
}
