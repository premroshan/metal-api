package ip

import (
	"context"
	"fmt"
	"github.com/metal-stack/metal-api/cmd/metal-api/internal/service"
	"github.com/metal-stack/metal-api/cmd/metal-api/internal/service/helper"
	"net/http"

	mdmv1 "github.com/metal-stack/masterdata-api/api/v1"
	mdm "github.com/metal-stack/masterdata-api/pkg/client"

	"github.com/metal-stack/metal-api/cmd/metal-api/internal/datastore"
	"github.com/metal-stack/metal-api/cmd/metal-api/internal/ipam"
	"github.com/metal-stack/metal-api/cmd/metal-api/internal/metal"
	"github.com/metal-stack/metal-api/cmd/metal-api/internal/tags"
	"github.com/metal-stack/metal-api/cmd/metal-api/internal/utils"
	"go.uber.org/zap"

	v1 "github.com/metal-stack/metal-api/cmd/metal-api/internal/service/v1"

	restful "github.com/emicklei/go-restful"
	restfulspec "github.com/emicklei/go-restful-openapi"
	"github.com/metal-stack/metal-lib/httperrors"
	"github.com/metal-stack/metal-lib/zapup"
)

type ipResource struct {
	service.WebResource
	ipamer ipam.IPAMer
	mdc    mdm.Client
}

// NewIP returns a webservice for ip specific endpoints.
func NewIP(ds *datastore.RethinkStore, ipamer ipam.IPAMer, mdc mdm.Client) *restful.WebService {
	ir := ipResource{
		WebResource: service.WebResource{
			DS: ds,
		},
		ipamer: ipamer,
		mdc:    mdc,
	}
	return ir.webService()
}

func (ir ipResource) webService() *restful.WebService {
	ws := new(restful.WebService)
	ws.
		Path(service.BasePath + "v1/ip").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)

	tags := []string{"ip"}

	ws.Route(ws.GET("/{id}").
		To(helper.Viewer(ir.findIP)).
		Operation("findIP").
		Doc("get ip by id").
		Param(ws.PathParameter("id", "identifier of the ip").DataType("string")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(v1.IPResponse{}).
		Returns(http.StatusOK, "OK", v1.IPResponse{}).
		DefaultReturns("Error", httperrors.HTTPErrorResponse{}))

	ws.Route(ws.GET("/").
		To(helper.Viewer(ir.listIPs)).
		Operation("listIPs").
		Doc("get all ips").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes([]v1.IPResponse{}).
		Returns(http.StatusOK, "OK", []v1.IPResponse{}).
		DefaultReturns("Error", httperrors.HTTPErrorResponse{}))

	ws.Route(ws.POST("/find").
		To(helper.Viewer(ir.findIPs)).
		Operation("findIPs").
		Doc("get all ips that match given properties").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(v1.IPFindRequest{}).
		Writes([]v1.IPResponse{}).
		Returns(http.StatusOK, "OK", []v1.IPResponse{}).
		DefaultReturns("Error", httperrors.HTTPErrorResponse{}))

	ws.Route(ws.POST("/free/{id}").
		To(helper.Editor(ir.freeIP)).
		Operation("freeIP").
		Doc("frees an ip").
		Param(ws.PathParameter("id", "identifier of the ip").DataType("string")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(v1.IPResponse{}).
		Returns(http.StatusOK, "OK", v1.IPResponse{}).
		DefaultReturns("Error", httperrors.HTTPErrorResponse{}))

	ws.Route(ws.POST("/").
		To(helper.Editor(ir.updateIP)).
		Operation("updateIP").
		Doc("updates an ip. if the ip was changed since this one was read, a conflict is returned").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(v1.IPUpdateRequest{}).
		Writes(v1.IPResponse{}).
		Returns(http.StatusOK, "OK", v1.IPResponse{}).
		Returns(http.StatusConflict, "Conflict", httperrors.HTTPErrorResponse{}).
		DefaultReturns("Error", httperrors.HTTPErrorResponse{}))

	ws.Route(ws.POST("/allocate").
		To(helper.Editor(ir.allocateIP)).
		Operation("allocateIP").
		Doc("allocate an ip in the given network.").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(v1.IPAllocateRequest{}).
		Writes(v1.IPResponse{}).
		Returns(http.StatusCreated, "Created", v1.IPResponse{}).
		DefaultReturns("Error", httperrors.HTTPErrorResponse{}))

	ws.Route(ws.POST("/allocate/{ip}").
		To(helper.Editor(ir.allocateIP)).
		Operation("allocateSpecificIP").
		Param(ws.PathParameter("ip", "ip to try to allocate").DataType("string")).
		Doc("allocate a specific ip in the given network.").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(v1.IPAllocateRequest{}).
		Writes(v1.IPResponse{}).
		Returns(http.StatusCreated, "Created", v1.IPResponse{}).
		DefaultReturns("Error", httperrors.HTTPErrorResponse{}))

	return ws
}

func (ir ipResource) findIP(request *restful.Request, response *restful.Response) {
	id := request.PathParameter("id")

	ip, err := ir.DS.FindIPByID(id)
	if helper.CheckError(request, response, utils.CurrentFuncName(), err) {
		return
	}
	err = response.WriteHeaderAndEntity(http.StatusOK, v1.NewIPResponse(ip))
	if err != nil {
		zapup.MustRootLogger().Error("Failed to send response", zap.Error(err))
		return
	}
}

func (ir ipResource) listIPs(request *restful.Request, response *restful.Response) {
	ips, err := ir.DS.ListIPs()
	if helper.CheckError(request, response, utils.CurrentFuncName(), err) {
		return
	}

	var result []*v1.IPResponse
	for i := range ips {
		result = append(result, v1.NewIPResponse(&ips[i]))
	}
	err = response.WriteHeaderAndEntity(http.StatusOK, result)
	if err != nil {
		zapup.MustRootLogger().Error("Failed to send response", zap.Error(err))
		return
	}
}

func (ir ipResource) findIPs(request *restful.Request, response *restful.Response) {
	var requestPayload datastore.IPSearchQuery
	err := request.ReadEntity(&requestPayload)
	if helper.CheckError(request, response, utils.CurrentFuncName(), err) {
		return
	}

	var ips metal.IPs
	err = ir.DS.SearchIPs(&requestPayload, &ips)
	if helper.CheckError(request, response, utils.CurrentFuncName(), err) {
		return
	}

	var result []*v1.IPResponse
	for i := range ips {
		result = append(result, v1.NewIPResponse(&ips[i]))
	}
	err = response.WriteHeaderAndEntity(http.StatusOK, result)
	if err != nil {
		zapup.MustRootLogger().Error("Failed to send response", zap.Error(err))
		return
	}
}

func (ir ipResource) freeIP(request *restful.Request, response *restful.Response) {
	id := request.PathParameter("id")

	ip, err := ir.DS.FindIPByID(id)
	if helper.CheckError(request, response, utils.CurrentFuncName(), err) {
		return
	}

	err = validateIPDelete(ip)
	if helper.CheckError(request, response, utils.CurrentFuncName(), err) {
		return
	}

	err = ir.ipamer.ReleaseIP(*ip)
	if helper.CheckError(request, response, utils.CurrentFuncName(), err) {
		return
	}

	err = ir.DS.DeleteIP(ip)
	if helper.CheckError(request, response, utils.CurrentFuncName(), err) {
		return
	}
	err = response.WriteHeaderAndEntity(http.StatusOK, v1.NewIPResponse(ip))
	if err != nil {
		zapup.MustRootLogger().Error("Failed to send response", zap.Error(err))
		return
	}
}

func validateIPDelete(ip *metal.IP) error {
	s := ip.GetScope()
	if s != metal.ScopeProject && ip.Type == metal.Static {
		return fmt.Errorf("ip with scope %s can not be deleted", ip.GetScope())
	}
	return nil
}

// Checks whether an ip update is allowed:
// (1) allow update of ephemeral to static
// (2) allow update within a scope
// (3) allow update from and to scope project
// (4) deny all other updates
func validateIPUpdate(old *metal.IP, new *metal.IP) error {
	// constraint 1
	if old.Type == metal.Static && new.Type == metal.Ephemeral {
		return fmt.Errorf("cannot change type of ip address from static to ephemeral")
	}
	os := old.GetScope()
	ns := new.GetScope()
	// constraint 2
	if os == ns {
		return nil
	}
	// constraint 3
	if os == metal.ScopeProject || ns == metal.ScopeProject {
		return nil
	}
	return fmt.Errorf("can not use ip of scope %v with scope %v", os, ns)
}

func processTags(ts []string) ([]string, error) {
	t := tags.New(ts)
	return t.Unique(), nil
}

func (ir ipResource) allocateIP(request *restful.Request, response *restful.Response) {
	specificIP := request.PathParameter("ip")
	var requestPayload v1.IPAllocateRequest
	err := request.ReadEntity(&requestPayload)
	if helper.CheckError(request, response, utils.CurrentFuncName(), err) {
		return
	}

	if requestPayload.NetworkID == "" {
		if helper.CheckError(request, response, utils.CurrentFuncName(), fmt.Errorf("networkid should not be empty")) {
			return
		}
	}
	if requestPayload.ProjectID == "" {
		if helper.CheckError(request, response, utils.CurrentFuncName(), fmt.Errorf("projectid should not be empty")) {
			return
		}
	}

	var name string
	if requestPayload.Name != nil {
		name = *requestPayload.Name
	}
	var description string
	if requestPayload.Description != nil {
		description = *requestPayload.Description
	}

	nw, err := ir.DS.FindNetworkByID(requestPayload.NetworkID)
	if helper.CheckError(request, response, utils.CurrentFuncName(), err) {
		return
	}

	p, err := ir.mdc.Project().Get(context.Background(), &mdmv1.ProjectGetRequest{Id: requestPayload.ProjectID})
	if helper.CheckError(request, response, utils.CurrentFuncName(), err) {
		return
	}

	tags := requestPayload.Tags
	if requestPayload.MachineID != nil {
		tags = append(tags, metal.IpTag(metal.TagIPMachineID, *requestPayload.MachineID))
	}

	tags, err = processTags(tags)
	if helper.CheckError(request, response, utils.CurrentFuncName(), err) {
		return
	}

	// TODO: Following operations should span a database transaction if possible

	ipAddress, ipParentCidr, err := helper.AllocateIP(nw, specificIP, ir.ipamer)
	if helper.CheckError(request, response, utils.CurrentFuncName(), err) {
		return
	}
	utils.Logger(request).Sugar().Debugw("found an ip to allocate", "ip", ipAddress, "network", nw.ID)

	ipType := metal.Ephemeral
	if requestPayload.Type == metal.Static {
		ipType = metal.Static
	}

	ip := &metal.IP{
		IPAddress:        ipAddress,
		ParentPrefixCidr: ipParentCidr,
		Name:             name,
		Description:      description,
		NetworkID:        nw.ID,
		ProjectID:        p.GetProject().GetMeta().GetId(),
		Type:             ipType,
		Tags:             tags,
	}

	err = ir.DS.CreateIP(ip)
	if helper.CheckError(request, response, utils.CurrentFuncName(), err) {
		return
	}
	err = response.WriteHeaderAndEntity(http.StatusCreated, v1.NewIPResponse(ip))
	if err != nil {
		zapup.MustRootLogger().Error("Failed to send response", zap.Error(err))
		return
	}
}

func (ir ipResource) updateIP(request *restful.Request, response *restful.Response) {
	var requestPayload v1.IPUpdateRequest
	err := request.ReadEntity(&requestPayload)
	if helper.CheckError(request, response, utils.CurrentFuncName(), err) {
		return
	}

	oldIP, err := ir.DS.FindIPByID(requestPayload.IPAddress)
	if helper.CheckError(request, response, utils.CurrentFuncName(), err) {
		return
	}

	newIP := *oldIP
	if requestPayload.Name != nil {
		newIP.Name = *requestPayload.Name
	}
	if requestPayload.Description != nil {
		newIP.Description = *requestPayload.Description
	}
	if requestPayload.Tags != nil {
		newIP.Tags = requestPayload.Tags
	}
	if requestPayload.Type == metal.Static || requestPayload.Type == metal.Ephemeral {
		newIP.Type = requestPayload.Type
	}

	err = ir.validateAndUpateIP(oldIP, &newIP)
	if helper.CheckError(request, response, utils.CurrentFuncName(), err) {
		return
	}
	err = response.WriteHeaderAndEntity(http.StatusOK, v1.NewIPResponse(&newIP))
	if err != nil {
		zapup.MustRootLogger().Error("Failed to send response", zap.Error(err))
		return
	}
}

func (ir ipResource) validateAndUpateIP(oldIP, newIP *metal.IP) error {
	err := validateIPUpdate(oldIP, newIP)
	if err != nil {
		return err
	}
	tags, err := processTags(newIP.Tags)
	if err != nil {
		return err
	}
	newIP.Tags = tags

	err = ir.DS.UpdateIP(oldIP, newIP)
	if err != nil {
		return err
	}
	return nil
}