package main

import (
	"go.bytecodealliance.org/cm"

	"github.com/a-skua/example-wasi/read-web/internal/gen/wasi/http/outgoing-handler"
	"github.com/a-skua/example-wasi/read-web/internal/gen/wasi/http/types"
)

func init() {
	outgoinghandler.Exports.Handle = Handle
}

func main() {}

func Handle(out types.OutgoingRequest, options cm.Option[types.RequestOptions]) (result cm.Result[outgoinghandler.ErrorCodeShape, types.FutureIncomingResponse, types.ErrorCode]) {
	panic("Handle is not supported in this example")
}
