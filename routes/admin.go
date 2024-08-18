package routes

import (
	"github.com/gin-gonic/gin"
	"idist-core/app/controllers/admin"
	"idist-core/app/middlewares"
)

func AdminRoutes(router *gin.RouterGroup) {
	router.Use(middlewares.AuthMiddleware().MiddlewareFunc())
	router.GET("/refresh-token", middlewares.AuthMiddleware().RefreshHandler)
	router.POST("/logout", middlewares.AuthMiddleware().LogoutHandler)
	router.GET("/profile", admin.GetProfile)
	router.POST("/profile", admin.UpdateProfile)

	// Auth
	router.POST("/profile/change-password", middlewares.Gate("", ""), admin.ChangePasswordRequest)
	router.PUT("/profile/change-password", middlewares.Gate("", ""), admin.ChangePasswordConfirm)

	// Basic
	router.GET("/provinces", middlewares.Gate("", ""), admin.ListProvinces)
	router.GET("/provinces/:id", middlewares.Gate("", ""), admin.ReadProvince)
	router.PUT("/provinces/:id", middlewares.Gate("", ""), admin.UpdateProvince)

	router.GET("/districts", middlewares.Gate("", ""), admin.ListDistricts)
	router.GET("/districts/id", middlewares.Gate("", ""), admin.ReadDistrict)
	router.PUT("/districts/id", middlewares.Gate("", ""), admin.UpdateDistrict)

	router.GET("/wards", middlewares.Gate("", ""), admin.ListWards)
	router.GET("/wards/id", middlewares.Gate("", ""), admin.ReadWard)
	router.PUT("/wards/id", middlewares.Gate("", ""), admin.UpdateWard)

	router.GET("/roles", middlewares.Gate("roles", "list"), admin.ListRoles)
	router.POST("/roles", middlewares.Gate("roles", "create"), admin.CreateRole)
	router.GET("/roles/:id", middlewares.Gate("roles", "read"), admin.ReadRole)
	router.PUT("/roles/:id", middlewares.Gate("roles", "update"), admin.UpdateRole)
	router.DELETE("/roles/:id", middlewares.Gate("roles", "delete"), admin.DeleteRole)

	router.GET("/permissions", middlewares.Gate("permissions", "list"), admin.ListPermissions)
	router.GET("/permissions/:id", middlewares.Gate("permissions", "read"), admin.ReadPermission)

	router.GET("/users", middlewares.Gate("users", "list"), admin.ListUsers)
	router.POST("/users", middlewares.Gate("users", "create"), admin.CreateUser)
	router.GET("/users/:id", middlewares.Gate("users", "read"), admin.ReadUser)
	router.PUT("/users/:id", middlewares.Gate("users", "update"), admin.UpdateUser)
	router.DELETE("/users/:id", middlewares.Gate("users", "delete"), admin.DeleteUser)

	// Categories
	router.GET("/categories", middlewares.Gate("categories", "list"), admin.ListCategories)
	router.POST("/categories", middlewares.Gate("categories", "create"), admin.CreateCategory)
	router.GET("/categories/:id", middlewares.Gate("categories", "read"), admin.GetCategory)
	router.PUT("/categories/:id", middlewares.Gate("categories", "update"), admin.UpdateCategory)
	router.DELETE("/categories/:id", middlewares.Gate("categories", "delete"), admin.DeleteCategory)

	// Tags
	router.GET("/tags", middlewares.Gate("tags", "list"), admin.ListTags)
	router.POST("/tags", middlewares.Gate("tags", "create"), admin.CreateTag)
	router.GET("/tags/:id", middlewares.Gate("tags", "read"), admin.GetTag)
	router.PUT("/tags/:id", middlewares.Gate("tags", "update"), admin.UpdateTag)
	router.DELETE("/tags/:id", middlewares.Gate("tags", "delete"), admin.DeleteTag)

	// Menus
	router.GET("/menus", middlewares.Gate("menus", "list"), admin.ListMenus)
	router.POST("/menus", middlewares.Gate("menus", "create"), admin.CreateMenu)
	router.GET("/menus/:id", middlewares.Gate("menus", "read"), admin.GetMenu)
	router.PUT("/menus/:id", middlewares.Gate("menus", "update"), admin.UpdateMenu)
	router.DELETE("/menus/:id", middlewares.Gate("menus", "delete"), admin.DeleteMenu)

	// Menu Locations
	router.GET("/menu-locations", middlewares.Gate("menu-locations", "list"), admin.ListMenuLocations)
	router.POST("/menu-locations", middlewares.Gate("menu-locations", "create"), admin.CreateMenuLocation)
	router.GET("/menu-locations/:id", middlewares.Gate("menu-locations", "read"), admin.GetMenuLocation)
	router.PUT("/menu-locations/:id", middlewares.Gate("menu-locations", "update"), admin.UpdateMenuLocation)
	router.DELETE("/menu-locations/:id", middlewares.Gate("menu-locations", "delete"), admin.DeleteMenuLocation)

	// Articles
	router.GET("/articles", middlewares.Gate("articles", "list"), admin.ListArticles)
	router.GET("/relate-articles", middlewares.Gate("articles", "list"), admin.ListRelateArticles)
	router.POST("/articles", middlewares.Gate("articles", "create"), admin.CreateArticle)
	router.GET("/articles/:id", middlewares.Gate("articles", "read"), admin.GetArticle)
	router.PUT("/articles/:id", middlewares.Gate("articles", "update"), admin.UpdateArticle)
	router.DELETE("/articles/:id", middlewares.Gate("articles", "delete"), admin.DeleteArticle)

	// Schools
	router.GET("/schools", middlewares.Gate("", ""), admin.ListSchools)
	router.POST("/schools", middlewares.Gate("", ""), admin.CreateSchool)
	router.GET("/schools/:id", middlewares.Gate("", ""), admin.ReadSchool)
	router.PUT("/schools/:id", middlewares.Gate("", ""), admin.UpdateSchool)
	router.DELETE("/schools/:id", middlewares.Gate("", ""), admin.DeleteSchool)

	// Schools
	router.GET("/admissions", middlewares.Gate("", ""), admin.ListAdmissions)
	router.POST("/admissions", middlewares.Gate("", ""), admin.CreateAdmission)
	router.GET("/admissions/:id", middlewares.Gate("", ""), admin.ReadAdmission)
	router.PUT("/admissions/:id", middlewares.Gate("", ""), admin.UpdateAdmission)
	router.DELETE("/admissions/:id", middlewares.Gate("", ""), admin.DeleteAdmission)

	// Logs
	router.GET("/logs", middlewares.Gate("logs", "list"), admin.ListLogRecords)

	// Organization Units
	router.GET("/organization-units", middlewares.Gate("", ""), admin.ListOrganizationUnits)
	router.POST("/organization-units", middlewares.Gate("", ""), admin.CreateOrganizationUnit)
	router.GET("/organization-units/:id", middlewares.Gate("", ""), admin.ReadOrganizationUnit)
	router.PUT("/organization-units/:id", middlewares.Gate("", ""), admin.UpdateOrganizationUnit)
	router.DELETE("/organization-units/:id", middlewares.Gate("", ""), admin.DeleteOrganizationUnit)

	// Dashboard Overview
	router.GET("/dashboard/overview/regular-posts", middlewares.Gate("", ""), admin.GetOverviewRegularPosts)
	router.GET("/dashboard/overview/regular-pages", middlewares.Gate("", ""), admin.GetOverviewRegularPages)
	router.GET("/dashboard/overview/traffics", middlewares.Gate("", ""), admin.GetOverviewTraffics)

	// Departments
	router.GET("/departments", middlewares.Gate("departments", "list"), admin.ListDepartments)
	router.POST("/departments", middlewares.Gate("departments", "create"), admin.CreateDepartment)
	router.GET("/departments/:id", middlewares.Gate("departments", "read"), admin.ReadDepartment)
	router.PUT("/departments/:id", middlewares.Gate("departments", "update"), admin.UpdateDepartment)
	router.DELETE("/departments/:id", middlewares.Gate("departments", "delete"), admin.DeleteDepartment)

	// Partners
	router.GET("/partners", middlewares.Gate("partners", "list"), admin.ListPartners)
	router.POST("/partners", middlewares.Gate("partners", "create"), admin.CreatePartner)
	router.GET("/partners/:id", middlewares.Gate("partners", "read"), admin.GetPartner)
	router.PUT("/partners/:id", middlewares.Gate("partners", "update"), admin.UpdatePartner)
	router.DELETE("/partners/:id", middlewares.Gate("partners", "delete"), admin.DeletePartner)

	// Descriptions
	router.GET("/descriptions", middlewares.Gate("", ""), admin.ListDescriptions)
	router.POST("/descriptions", middlewares.Gate("", ""), admin.CreateDescription)
	router.GET("/descriptions/tong-quan", middlewares.Gate("", ""), admin.GetDescriptionSummary)
	router.PUT("/descriptions/tong-quan", middlewares.Gate("", ""), admin.UpdateDescriptionSummary)
	router.GET("/descriptions/:id", middlewares.Gate("", ""), admin.GetDescription)
	router.PUT("/descriptions/:id", middlewares.Gate("", ""), admin.UpdateDescription)
	router.DELETE("/descriptions/:id", middlewares.Gate("", ""), admin.DeleteDescription)

	// Pages
	router.GET("/pages", middlewares.Gate("pages", "list"), admin.ListPages)
	router.POST("/pages", middlewares.Gate("pages", "create"), admin.CreatePage)
	router.GET("/pages/:id", middlewares.Gate("pages", "read"), admin.GetPage)
	router.PUT("/pages/:id", middlewares.Gate("pages", "update"), admin.UpdatePage)
	router.DELETE("/pages/:id", middlewares.Gate("pages", "delete"), admin.DeletePage)
	// Templates
	router.GET("/templates", middlewares.Gate("templates", "list"), admin.ListTemplates)
	router.POST("/templates", middlewares.Gate("templates", "create"), admin.CreateTemplate)
	router.GET("/templates/:id", middlewares.Gate("templates", "read"), admin.GetTemplate)
	router.PUT("/templates/:id", middlewares.Gate("templates", "update"), admin.UpdateTemplate)
	router.DELETE("/templates/:id", middlewares.Gate("templates", "delete"), admin.DeleteTemplate)
}
