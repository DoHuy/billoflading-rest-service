package rest

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"vtp-apis/external/domain"
	"vtp-apis/packages/ginh"
)

type VNSaleExecutionSerializer struct {
	vNSaleDomain domain.DoSomethingWithVNSale
}

func NewVNSaleSerializer(domain domain.DoSomethingWithVNSale) *VNSaleExecutionSerializer {
	return &VNSaleExecutionSerializer{
		vNSaleDomain: domain,
	}
}

// @Summary Fetch van don hanh trinh
// @Description Fecth van don hanh trinh
// @Tags [fetch chitiet]
// @Accept  json
// @Produce json
// @Success 200 {object} ginh.response{meta=ginh.ResponseMeta{code=int}}
// @failure 500 {object} ginh.response{meta=ginh.ResponseMeta{code=int}}
// @Router /api/v1/billoflading/{id} [get]
func (s *VNSaleExecutionSerializer) FetchVandonhanhtrinhByID(ctx *gin.Context) {
	//var mavandon string
	mavandon := ctx.Param("id")
	fmt.Println("mavandon", mavandon)
	detailOrder, err := s.vNSaleDomain.FetchVandonhanhtrinh(ctx, mavandon)
	if err != nil {
		ginh.LogError(ctx, "fetch failed ", err)
		ginh.BuildErrorResponse(ctx, err, nil)
	}
	ginh.LogInfo(ctx, "fetch success ")
	ginh.BuildSuccessResponse(ctx, http.StatusOK, detailOrder)
	return
}

// @Summary Update trang thai don
// @Description Update trang thai don
// @Tags [update status]
// @Accept  json
// @Produce json
// @Param Body body rest.Trangthaivandon true "Body message"
// @Success 200 {object} ginh.response{meta=ginh.ResponseMeta{code=int}}
// @failure 500 {object} ginh.response{meta=ginh.ResponseMeta{code=int}}
// @Router /api/v1/billoflading/{id} [post]
func (s *VNSaleExecutionSerializer) UpdateOrderStateByID(ctx *gin.Context) {
	jsonData, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		// Handle error
		ginh.BuildErrorResponse(ctx, err, nil)
	}
	//var mavandon string
	mavandon, _ := ctx.GetQuery("id")
	var jsonBody domain.Trangthaivandon
	if err := json.Unmarshal(jsonData, &jsonBody); err != nil {
		ginh.BuildErrorResponse(ctx, err, nil)
	}
	err = s.vNSaleDomain.UpdateOrderState(ctx, jsonBody, mavandon)
	if err != nil {
		ginh.LogError(ctx, "update failed ", err)
		ginh.BuildErrorResponse(ctx, err, nil)
	}
	ginh.LogInfo(ctx, "update success")
	ginh.BuildSuccessResponse(ctx, http.StatusOK, nil)
}
