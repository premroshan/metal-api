package machine

import (
	"github.com/emicklei/go-restful"
	"github.com/metal-stack/metal-api/cmd/metal-api/internal/service/helper"
	"github.com/metal-stack/metal-api/cmd/metal-api/internal/utils"
	"github.com/metal-stack/metal-lib/zapup"
	"go.uber.org/zap"
	"net/http"
)

func (r machineResource) findMachine(request *restful.Request, response *restful.Response) {
	id := request.PathParameter("id")

	m, err := r.DS.FindMachineByID(id)
	if helper.CheckError(request, response, utils.CurrentFuncName(), err) {
		return
	}
	resp := helper.MakeMachineResponse(m, r.DS, utils.Logger(request).Sugar())
	err = response.WriteHeaderAndEntity(http.StatusOK, resp)
	if err != nil {
		zapup.MustRootLogger().Error("Failed to send response", zap.Error(err))
		return
	}
}
