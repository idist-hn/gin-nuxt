let config = {
    PostLogout: "/api/v1/admin/logout",
    PostLogin: "/api/v1/auth/login",
    GetProfile: "/api/v1/admin/profile",

    // Users
    ListUsers: "/api/v1/admin/users",
    CreateUser: "/api/v1/admin/users",
    GetUser: "/api/v1/admin/users/{id}",
    UpdateUser: "/api/v1/admin/users/{id}",
    DeleteUser: "/api/v1/admin/users/{id}",

    // Roles
    ListRoles: "/api/v1/admin/roles",
    CreateRole: "/api/v1/admin/roles",
    GetRole: "/api/v1/admin/roles/{id}",
    UpdateRole: "/api/v1/admin/roles/{id}",
    DeleteRole: "/api/v1/admin/roles/{id}",

    // Permissions
    ListPermissions: "/api/v1/admin/permissions",
    GetPermission: "/api/v1/admin/permissions/{id}",

    // Categories
    ListCategories: "/api/v1/admin/categories",
    CreateCategory: "/api/v1/admin/categories",
    GetCategory: "/api/v1/admin/categories/{id}",
    UpdateCategory: "/api/v1/admin/categories/{id}",
    DeleteCategory: "/api/v1/admin/categories/{id}",

    // Tags
    ListTags: "/api/v1/admin/tags",
    CreateTag: "/api/v1/admin/tags",
    GetTag: "/api/v1/admin/tags/{id}",
    UpdateTag: "/api/v1/admin/tags/{id}",
    DeleteTag: "/api/v1/admin/tags/{id}",

    // Menus
    ListMenus: "/api/v1/admin/menus",
    CreateMenu: "/api/v1/admin/menus",
    GetMenu: "/api/v1/admin/menus/{id}",
    UpdateMenu: "/api/v1/admin/menus/{id}",
    DeleteMenu: "/api/v1/admin/menus/{id}",

    // Tags
    ListProvinces: "/api/v1/common/provinces",
    CreateProvince: "/api/v1/admin/provinces",
    GetProvince: "/api/v1/admin/provinces/{id}",
    UpdateProvince: "/api/v1/admin/provinces/{id}",
    DeleteProvince: "/api/v1/admin/provinces/{id}",

    // Articles
    ListArticles: "/api/v1/admin/articles",
    ListRelateArticles: "/api/v1/admin/relate-articles",
    CreateArticle: "/api/v1/admin/articles",
    GetArticle: "/api/v1/admin/articles/{id}",
    UpdateArticle: "/api/v1/admin/articles/{id}",
    DeleteArticle: "/api/v1/admin/articles/{id}",

    // Schools
    ListSchools: "/api/v1/admin/schools",
    CreateSchool: "/api/v1/admin/schools",
    GetSchool: "/api/v1/admin/schools/{id}",
    UpdateSchool: "/api/v1/admin/schools/{id}",
    DeleteSchool: "/api/v1/admin/schools/{id}",

    // Admissions
    ListAdmissions: "/api/v1/admin/admissions",
    CreateAdmission: "/api/v1/admin/admissions",
    GetAdmission: "/api/v1/admin/admissions/{id}",
    UpdateAdmission: "/api/v1/admin/admissions/{id}",
    DeleteAdmission: "/api/v1/admin/admissions/{id}",

    CreateTuyenSinh: "/api/v1/common/tuyen-sinh",

    // Loggings
    ListLogs: "/api/v1/admin/logs",
    GetLog: "/api/v1/admin/logs/{id}",

    // Organization Units
    ListOrganizationUnits: "/api/v1/admin/organization-units",
    CreateOrganizationUnit: "/api/v1/admin/organization-units",
    GetOrganizationUnit: "/api/v1/admin/organization-units/{id}",
    UpdateOrganizationUnit: "/api/v1/admin/organization-units/{id}",
    DeleteOrganizationUnit: "/api/v1/admin/organization-units/{id}",

    // menu Locations
    ListMenuLocations: "/api/v1/admin/menu-locations",
    CreateMenuLocation: "/api/v1/admin/menu-locations",
    GetMenuLocation: "/api/v1/admin/menu-locations/{id}",
    UpdateMenuLocation: "/api/v1/admin/menu-locations/{id}",
    DeleteMenuLocation: "/api/v1/admin/menu-locations/{id}",

    // Dashboard Overview
    GetOverviewRegularPosts: "/api/v1/admin/dashboard/overview/regular-posts",
    GetOverviewRegularPages: "/api/v1/admin/dashboard/overview/regular-pages",
    GetOverviewTraffics: "/api/v1/admin/dashboard/overview/traffics",

    // Departments
    ListDepartments: "/api/v1/admin/departments",
    CreateDepartment: "/api/v1/admin/departments",
    GetDepartment: "/api/v1/admin/departments/{id}",
    UpdateDepartment: "/api/v1/admin/departments/{id}",
    DeleteDepartment: "/api/v1/admin/departments/{id}",
    // Partners
    ListPartners: "/api/v1/admin/partners",
    CreatePartner: "/api/v1/admin/partners",
    GetPartner: "/api/v1/admin/partners/{id}",
    UpdatePartner: "/api/v1/admin/partners/{id}",
    DeletePartner: "/api/v1/admin/partners/{id}",

    // Descriptions
    ListDescriptions: "/api/v1/admin/descriptions",
    CreateDescription: "/api/v1/admin/descriptions",
    GetDescription: "/api/v1/admin/descriptions/{id}",
    GetDescriptionSummary: "/api/v1/admin/descriptions/tong-quan",
    UpdateDescriptionSummary: "/api/v1/admin/descriptions/tong-quan",
    DeleteDescription: "/api/v1/admin/descriptions/{id}",

    //Pages
    ListPages: "/api/v1/admin/pages",
    CreatePage: "/api/v1/admin/pages",
    GetPage: "/api/v1/admin/pages/{id}",
    UpdatePage: "/api/v1/admin/pages/{id}",
    DeletePage: "/api/v1/admin/pages/{id}",

    //Template
    ListTemplates: "/api/v1/admin/templates",
    CreateTemplate: "/api/v1/admin/templates",
    GetTemplate: "/api/v1/admin/templates/{id}",
    UpdateTemplate: "/api/v1/admin/templates/{id}",
    DeleteTemplate: "/api/v1/admin/templates/{id}",

};

let api = new Proxy(config, {
    get(target, name) {
        if (name !== "params") {
            return Reflect.get(target, name);
        } else {
            return Reflect.get(target, name);
        }
    },
});

api.params = (name, parameters) => {
    try {
        let endpoint = api[name];
        for (let value in parameters) {
            endpoint = endpoint.replace(`{${value}}`, parameters[value]);
        }
        return endpoint;
    } catch (e) {
        console.log(e);
    }
};

export default api;
