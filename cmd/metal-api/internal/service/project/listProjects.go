package project

import (
	"context"
	"github.com/emicklei/go-restful"
	restfulspec "github.com/emicklei/go-restful-openapi"
	mdmv1 "github.com/metal-stack/masterdata-api/api/v1"
	"github.com/metal-stack/metal-api/cmd/metal-api/internal/service/helper"
	v1 "github.com/metal-stack/metal-api/cmd/metal-api/internal/service/v1"
	"github.com/metal-stack/metal-api/cmd/metal-api/internal/utils"
	"github.com/metal-stack/metal-lib/httperrors"
	"github.com/metal-stack/metal-lib/zapup"
	"go.uber.org/zap"
	"net/http"
)

func (r projectResource) addListProjectsRoute(ws *restful.WebService, tags []string) {
	ws.Route(ws.GET("/").
		To(helper.Viewer(r.listProjects)).
		Operation("listProjects").
		Doc("get all projects").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes([]v1.ProjectResponse{}).
		Returns(http.StatusOK, "OK", []v1.ProjectResponse{}).
		DefaultReturns("Error", httperrors.HTTPErrorResponse{}))
}

func (r projectResource) listProjects(request *restful.Request, response *restful.Response) {
	ps, err := r.mdc.Project().Find(context.Background(), &mdmv1.ProjectFindRequest{})
	if helper.CheckError(request, response, utils.CurrentFuncName(), err) {
		return
	}

	err = response.WriteHeaderAndEntity(http.StatusOK, ps.Projects)
	if err != nil {
		zapup.MustRootLogger().Error("Failed to send response", zap.Error(err))
		return
	}
}