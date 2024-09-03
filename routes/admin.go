package routes

import (
	"github.com/gin-gonic/gin"
	"idist-core/app/controllers/admin"
	"idist-core/middleware"
)

func AdminRoutes(router *gin.RouterGroup) {
	router.Use(middleware.AuthMiddleware().MiddlewareFunc())
	router.GET("/refresh-token", middleware.AuthMiddleware().RefreshHandler)
	router.POST("/logout", middleware.AuthMiddleware().LogoutHandler)
	router.GET("/profile", admin.GetProfile)
	router.POST("/profile", admin.UpdateProfile)

	// Auth
	router.POST("/profile/change-password", middleware.Gate("", ""), admin.ChangePasswordRequest)
	router.PUT("/profile/change-password", middleware.Gate("", ""), admin.ChangePasswordConfirm)

	// Basic
	router.GET("/provinces", middleware.Gate("", ""), admin.ListProvinces)
	router.GET("/provinces/:id", middleware.Gate("", ""), admin.ReadProvince)
	router.PUT("/provinces/:id", middleware.Gate("", ""), admin.UpdateProvince)

	router.GET("/districts", middleware.Gate("", ""), admin.ListDistricts)
	router.GET("/districts/id", middleware.Gate("", ""), admin.ReadDistrict)
	router.PUT("/districts/id", middleware.Gate("", ""), admin.UpdateDistrict)

	router.GET("/wards", middleware.Gate("", ""), admin.ListWards)
	router.GET("/wards/id", middleware.Gate("", ""), admin.ReadWard)
	router.PUT("/wards/id", middleware.Gate("", ""), admin.UpdateWard)

	router.GET("/roles", middleware.Gate("roles", "list"), admin.ListRoles)
	router.POST("/roles", middleware.Gate("roles", "create"), admin.CreateRole)
	router.GET("/roles/:id", middleware.Gate("roles", "read"), admin.ReadRole)
	router.PUT("/roles/:id", middleware.Gate("roles", "update"), admin.UpdateRole)
	router.DELETE("/roles/:id", middleware.Gate("roles", "delete"), admin.DeleteRole)

	router.GET("/permissions", middleware.Gate("permissions", "list"), admin.ListPermissions)
	router.GET("/permissions/:id", middleware.Gate("permissions", "read"), admin.ReadPermission)

	router.GET("/users", middleware.Gate("users", "list"), admin.ListUsers)
	router.POST("/users", middleware.Gate("users", "create"), admin.CreateUser)
	router.GET("/users/:id", middleware.Gate("users", "read"), admin.ReadUser)
	router.PUT("/users/:id", middleware.Gate("users", "update"), admin.UpdateUser)
	router.DELETE("/users/:id", middleware.Gate("users", "delete"), admin.DeleteUser)

	// Categories
	router.GET("/categories", middleware.Gate("categories", "list"), admin.ListCategories)
	router.POST("/categories", middleware.Gate("categories", "create"), admin.CreateCategory)
	router.GET("/categories/:id", middleware.Gate("categories", "read"), admin.GetCategory)
	router.PUT("/categories/:id", middleware.Gate("categories", "update"), admin.UpdateCategory)
	router.DELETE("/categories/:id", middleware.Gate("categories", "delete"), admin.DeleteCategory)

	// Tags
	router.GET("/tags", middleware.Gate("tags", "list"), admin.ListTags)
	router.POST("/tags", middleware.Gate("tags", "create"), admin.CreateTag)
	router.GET("/tags/:id", middleware.Gate("tags", "read"), admin.GetTag)
	router.PUT("/tags/:id", middleware.Gate("tags", "update"), admin.UpdateTag)
	router.DELETE("/tags/:id", middleware.Gate("tags", "delete"), admin.DeleteTag)

	// Menus
	router.GET("/menus", middleware.Gate("menus", "list"), admin.ListMenus)
	router.POST("/menus", middleware.Gate("menus", "create"), admin.CreateMenu)
	router.GET("/menus/:id", middleware.Gate("menus", "read"), admin.GetMenu)
	router.PUT("/menus/:id", middleware.Gate("menus", "update"), admin.UpdateMenu)
	router.DELETE("/menus/:id", middleware.Gate("menus", "delete"), admin.DeleteMenu)

	// Menu Locations
	router.GET("/menu-locations", middleware.Gate("menu-locations", "list"), admin.ListMenuLocations)
	router.POST("/menu-locations", middleware.Gate("menu-locations", "create"), admin.CreateMenuLocation)
	router.GET("/menu-locations/:id", middleware.Gate("menu-locations", "read"), admin.GetMenuLocation)
	router.PUT("/menu-locations/:id", middleware.Gate("menu-locations", "update"), admin.UpdateMenuLocation)
	router.DELETE("/menu-locations/:id", middleware.Gate("menu-locations", "delete"), admin.DeleteMenuLocation)

	// Articles
	router.GET("/articles", middleware.Gate("articles", "list"), admin.ListArticles)
	router.GET("/relate-articles", middleware.Gate("articles", "list"), admin.ListRelateArticles)
	router.POST("/articles", middleware.Gate("articles", "create"), admin.CreateArticle)
	router.GET("/articles/:id", middleware.Gate("articles", "read"), admin.GetArticle)
	router.PUT("/articles/:id", middleware.Gate("articles", "update"), admin.UpdateArticle)
	router.DELETE("/articles/:id", middleware.Gate("articles", "delete"), admin.DeleteArticle)

	// Schools
	router.GET("/schools", middleware.Gate("", ""), admin.ListSchools)
	router.POST("/schools", middleware.Gate("", ""), admin.CreateSchool)
	router.GET("/schools/:id", middleware.Gate("", ""), admin.ReadSchool)
	router.PUT("/schools/:id", middleware.Gate("", ""), admin.UpdateSchool)
	router.DELETE("/schools/:id", middleware.Gate("", ""), admin.DeleteSchool)

	// Schools
	router.GET("/admissions", middleware.Gate("", ""), admin.ListAdmissions)
	router.POST("/admissions", middleware.Gate("", ""), admin.CreateAdmission)
	router.GET("/admissions/:id", middleware.Gate("", ""), admin.ReadAdmission)
	router.PUT("/admissions/:id", middleware.Gate("", ""), admin.UpdateAdmission)
	router.DELETE("/admissions/:id", middleware.Gate("", ""), admin.DeleteAdmission)

	// Logs
	router.GET("/logs", middleware.Gate("logs", "list"), admin.ListLogRecords)

	// Organization Units
	router.GET("/organization-units", middleware.Gate("", ""), admin.ListOrganizationUnits)
	router.POST("/organization-units", middleware.Gate("", ""), admin.CreateOrganizationUnit)
	router.GET("/organization-units/:id", middleware.Gate("", ""), admin.ReadOrganizationUnit)
	router.PUT("/organization-units/:id", middleware.Gate("", ""), admin.UpdateOrganizationUnit)
	router.DELETE("/organization-units/:id", middleware.Gate("", ""), admin.DeleteOrganizationUnit)

	// Dashboard Overview
	router.GET("/dashboard/overview/regular-posts", middleware.Gate("", ""), admin.GetOverviewRegularPosts)
	router.GET("/dashboard/overview/regular-pages", middleware.Gate("", ""), admin.GetOverviewRegularPages)
	router.GET("/dashboard/overview/traffics", middleware.Gate("", ""), admin.GetOverviewTraffics)

	// Departments
	router.GET("/departments", middleware.Gate("departments", "list"), admin.ListDepartments)
	router.POST("/departments", middleware.Gate("departments", "create"), admin.CreateDepartment)
	router.GET("/departments/:id", middleware.Gate("departments", "read"), admin.ReadDepartment)
	router.PUT("/departments/:id", middleware.Gate("departments", "update"), admin.UpdateDepartment)
	router.DELETE("/departments/:id", middleware.Gate("departments", "delete"), admin.DeleteDepartment)

	// Partners
	router.GET("/partners", middleware.Gate("partners", "list"), admin.ListPartners)
	router.POST("/partners", middleware.Gate("partners", "create"), admin.CreatePartner)
	router.GET("/partners/:id", middleware.Gate("partners", "read"), admin.GetPartner)
	router.PUT("/partners/:id", middleware.Gate("partners", "update"), admin.UpdatePartner)
	router.DELETE("/partners/:id", middleware.Gate("partners", "delete"), admin.DeletePartner)

	// Descriptions
	router.GET("/descriptions", middleware.Gate("", ""), admin.ListDescriptions)
	router.POST("/descriptions", middleware.Gate("", ""), admin.CreateDescription)
	router.GET("/descriptions/tong-quan", middleware.Gate("", ""), admin.GetDescriptionSummary)
	router.PUT("/descriptions/tong-quan", middleware.Gate("", ""), admin.UpdateDescriptionSummary)
	router.GET("/descriptions/:id", middleware.Gate("", ""), admin.GetDescription)
	router.PUT("/descriptions/:id", middleware.Gate("", ""), admin.UpdateDescription)
	router.DELETE("/descriptions/:id", middleware.Gate("", ""), admin.DeleteDescription)

	// Pages
	router.GET("/pages", middleware.Gate("pages", "list"), admin.ListPages)
	router.POST("/pages", middleware.Gate("pages", "create"), admin.CreatePage)
	router.GET("/pages/:id", middleware.Gate("pages", "read"), admin.GetPage)
	router.PUT("/pages/:id", middleware.Gate("pages", "update"), admin.UpdatePage)
	router.DELETE("/pages/:id", middleware.Gate("pages", "delete"), admin.DeletePage)
	// Templates
	router.GET("/templates", middleware.Gate("templates", "list"), admin.ListTemplates)
	router.POST("/templates", middleware.Gate("templates", "create"), admin.CreateTemplate)
	router.GET("/templates/:id", middleware.Gate("templates", "read"), admin.GetTemplate)
	router.PUT("/templates/:id", middleware.Gate("templates", "update"), admin.UpdateTemplate)
	router.DELETE("/templates/:id", middleware.Gate("templates", "delete"), admin.DeleteTemplate)
}
