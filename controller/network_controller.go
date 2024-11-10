package controller

import (
	"go-rest-api/model"
	"go-rest-api/usecase"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type INetworkController interface {
	GetAllNetworks(c echo.Context) error
	GetNetworkById(c echo.Context) error
	CreateNetwork(c echo.Context) error
	UpdateNetwork(c echo.Context) error
	DeleteNetwork(c echo.Context) error
}

type networkController struct {
	tu usecase.INetworkUsecase
}

func NewNetworkController(tu usecase.INetworkUsecase) INetworkController {
	return &networkController{tu}
}

func (tc *networkController) GetAllNetworks(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]

	networksRes, err := tc.tu.GetAllNetworks(uint(userId.(float64)))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, networksRes)
}

func (tc *networkController) GetNetworkById(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]
	id := c.Param("networkId")
	networkId, _ := strconv.Atoi(id)
	networkRes, err := tc.tu.GetNetworkById(uint(userId.(float64)), uint(networkId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, networkRes)
}

func (tc *networkController) CreateNetwork(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]

	network := model.Network{}
	if err := c.Bind(&network); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	network.UserId = uint(userId.(float64))
	networkRes, err := tc.tu.CreateNetwork(network)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, networkRes)
}

func (tc *networkController) UpdateNetwork(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]
	id := c.Param("networkId")
	networkId, _ := strconv.Atoi(id)

	network := model.Network{}
	if err := c.Bind(&network); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	networkRes, err := tc.tu.UpdateNetwork(network, uint(userId.(float64)), uint(networkId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, networkRes)
}

func (tc *networkController) DeleteNetwork(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]
	id := c.Param("networkId")
	networkId, _ := strconv.Atoi(id)

	err := tc.tu.DeleteNetwork(uint(userId.(float64)), uint(networkId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusNoContent)
}
