package handlers

import (
	"jira-for-peasants/internal/api/requests"
	"jira-for-peasants/internal/api/responses"
	"jira-for-peasants/internal/repositories"
	"jira-for-peasants/internal/services"
	errpkg "jira-for-peasants/pkg/errors"
	"jira-for-peasants/pkg/utils"

	"github.com/labstack/echo/v4"
)

type RoleHandler struct {
	roleService *services.RoleService
}

func NewRoleHandler(roleService *services.RoleService) RoleHandler {
	return RoleHandler{
		roleService: roleService,
	}
}

func (h *RoleHandler) RegisterRoutes(handler *echo.Group) {
	handler.POST("", h.handleCreateRole)
	handler.PATCH("/:id", h.handleUpdateRole)
	handler.DELETE("", h.handleDeleteRole)
	handler.GET("/organization", h.handleGetRolesForOrganization)
	handler.GET("/user", h.handleGetRolesForUser)
	handler.GET("/:id", h.handleGetRole)
}

func (h *RoleHandler) handleCreateRole(c echo.Context) error {
	u := new(requests.CreateRoleRequest)
	ctx := c.Request().Context()
	if err := c.Bind(u); err != nil {
		return errpkg.BadRequest(err.Error())
	}
	if err := c.Validate(u); err != nil {
		return err
	}

	entity, err := h.roleService.CreateRole(ctx, services.CreateRoleParams{
		Name:           u.Name,
		OrganizationId: u.OrganizationId,
		Description:    u.Description,
	})

	if err != nil {
		return err
	}

	return c.JSON(200, responses.NewRoleResponse(
		entity.ID,
		entity.Name,
		entity.Description,
		entity.Permissions,
	))
}

func (h *RoleHandler) handleUpdateRole(c echo.Context) error {
	u := new(requests.UpdateRoleRequest)
	roleId := c.QueryParam("id")
	ctx := c.Request().Context()
	if err := c.Bind(u); err != nil {
		return errpkg.BadRequest(err.Error())
	}
	if err := c.Validate(u); err != nil {
		return err
	}

	entity, err := h.roleService.UpdateRole(ctx, roleId, services.UpdateRoleParams{
		Name:           u.Name,
		OrganizationId: u.OrganizationId,
		Description:    u.Description,
	})

	if err != nil {
		return err
	}

	return c.JSON(200, responses.NewRoleResponse(
		entity.ID,
		entity.Name,
		entity.Description,
		entity.Permissions,
	))
}

func (h *RoleHandler) handleDeleteRole(c echo.Context) error {
	roleId := c.QueryParam("id")
	ctx := c.Request().Context()

	err := h.roleService.DeleteRole(ctx, roleId)

	if err != nil {
		return err
	}

	return c.JSON(200, nil)
}

func (h *RoleHandler) handleGetRolesForOrganization(c echo.Context) error {
	organizationId := c.QueryParam("organization_id")
	ctx := c.Request().Context()

	roles, err := h.roleService.GetAllRoleForOrganization(ctx, organizationId)

	if err != nil {
		return err
	}

	return c.JSON(200, utils.TransformSlice(roles, func(role repositories.RoleModel) responses.RoleResponse {
		return *responses.NewRoleResponse(
			role.ID,
			role.Name,
			role.Description,
			role.Permissions,
		)
	}))
}

func (h *RoleHandler) handleGetRole(c echo.Context) error {
	roleId := c.Param("id")
	ctx := c.Request().Context()

	role, err := h.roleService.GetRoleById(ctx, roleId)

	if err != nil {
		return err
	}

	return c.JSON(200, responses.NewRoleResponse(
		role.ID,
		role.Name,
		role.Description,
		role.Permissions,
	))
}

func (h *RoleHandler) handleGetRolesForUser(c echo.Context) error {
	ctx := c.Request().Context()
	userId := utils.GetUser(c)
	organizationId := utils.GetOrganization(c)

	roles, err := h.roleService.GetAllRoleForUser(ctx, userId, organizationId)

	if err != nil {
		return err
	}

	return c.JSON(200, utils.TransformSlice(roles, func(role repositories.RoleModel) responses.RoleResponse {
		return *responses.NewRoleResponse(
			role.ID,
			role.Name,
			role.Description,
			role.Permissions,
		)
	}))
}
