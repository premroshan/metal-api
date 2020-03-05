package machine

import (
	"github.com/emicklei/go-restful"
	"github.com/metal-stack/metal-api/cmd/metal-api/internal/metal"
)

func (r machineResource) powerResetMachine(request *restful.Request, response *restful.Response) {
	r.publishMachineCmd("powerResetMachine", metal.MachineResetCmd, request, response)
}
