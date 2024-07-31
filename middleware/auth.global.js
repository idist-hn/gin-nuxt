export default defineNuxtRouteMiddleware((to, from) => {
    if (!useCookie("access_token").value && to.name && !to.name.startsWith("auth")) {
        return navigateTo({ name: "auth-login"});
    }
});
