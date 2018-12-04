package service

import (
	"net/http"
	"time"

	"go.uber.org/zap"

	"git.f-i-ts.de/cloud-native/metal/metal-api/cmd/metal-api/internal/datastore"
	"git.f-i-ts.de/cloud-native/metal/metal-api/metal"
	restful "github.com/emicklei/go-restful"
	restfulspec "github.com/emicklei/go-restful-openapi"
)

type SiteResource struct {
	*zap.SugaredLogger
	log *zap.Logger
	ds  *datastore.RethinkStore
}

func NewSite(log *zap.Logger, ds *datastore.RethinkStore) *restful.WebService {
	fr := SiteResource{
		SugaredLogger: log.Sugar(),
		log:           log,
		ds:            ds,
	}
	return fr.webService()
}

func (fr SiteResource) webService() *restful.WebService {
	ws := new(restful.WebService)
	ws.
		Path("/site").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)

	tags := []string{"Site"}

	ws.Route(ws.GET("/{id}").To(fr.findSite).
		Doc("get Site by id").
		Param(ws.PathParameter("id", "identifier of the Site").DataType("string")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(metal.Site{}).
		Returns(http.StatusOK, "OK", metal.Site{}).
		Returns(http.StatusNotFound, "Not Found", nil))

	ws.Route(ws.GET("/").To(fr.listSites).
		Doc("get all Sites").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes([]metal.Site{}).
		Returns(http.StatusOK, "OK", []metal.Site{}))

	ws.Route(ws.DELETE("/{id}").To(fr.deleteSite).
		Doc("deletes a Site and returns the deleted entity").
		Param(ws.PathParameter("id", "identifier of the Site").DataType("string")).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Writes(metal.Site{}).
		Returns(http.StatusOK, "OK", metal.Site{}).
		Returns(http.StatusNotFound, "Not Found", nil))

	ws.Route(ws.PUT("/").To(fr.createSite).
		Doc("create a Site. if the given ID already exists a conflict is returned").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(metal.Site{}).
		Returns(http.StatusCreated, "Created", metal.Site{}).
		Returns(http.StatusConflict, "Conflict", nil))

	ws.Route(ws.POST("/").To(fr.updateSite).
		Doc("updates a Site. if the Site was changed since this one was read, a conflict is returned").
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Reads(metal.Site{}).
		Returns(http.StatusOK, "OK", metal.Site{}).
		Returns(http.StatusNotFound, "Not Found", nil).
		Returns(http.StatusConflict, "Conflict", nil))

	return ws
}

func (fr SiteResource) findSite(request *restful.Request, response *restful.Response) {
	id := request.PathParameter("id")
	Site, err := fr.ds.FindSite(id)
	if checkError(fr.log, response, "findSite", err) {
		return
	}
	response.WriteEntity(Site)
}

func (fr SiteResource) listSites(request *restful.Request, response *restful.Response) {
	res, err := fr.ds.ListSites()
	if checkError(fr.log, response, "listSites", err) {
		return
	}
	response.WriteEntity(res)
}

func (fr SiteResource) deleteSite(request *restful.Request, response *restful.Response) {
	id := request.PathParameter("id")
	Site, err := fr.ds.DeleteSite(id)
	if checkError(fr.log, response, "deleteSite", err) {
		return
	}
	response.WriteEntity(Site)
}

func (fr SiteResource) createSite(request *restful.Request, response *restful.Response) {
	var s metal.Site
	err := request.ReadEntity(&s)
	if checkError(fr.log, response, "createSite", err) {
		return
	}
	s.Created = time.Now()
	s.Changed = s.Created
	err = fr.ds.CreateSite(&s)
	if checkError(fr.log, response, "createSite", err) {
		return
	}
	response.WriteHeaderAndEntity(http.StatusCreated, s)
}

func (fr SiteResource) updateSite(request *restful.Request, response *restful.Response) {
	var newSite metal.Site
	err := request.ReadEntity(&newSite)
	if checkError(fr.log, response, "updateSite", err) {
		return
	}

	oldSite, err := fr.ds.FindSite(newSite.ID)
	if checkError(fr.log, response, "updateSite", err) {
		return
	}

	err = fr.ds.UpdateSite(oldSite, &newSite)

	if checkError(fr.log, response, "updateSite", err) {
		return
	}
	response.WriteHeaderAndEntity(http.StatusOK, newSite)
}
