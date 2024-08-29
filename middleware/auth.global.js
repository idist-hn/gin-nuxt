export default defineNuxtRouteMiddleware((to, from) => {
    console.log(useCookie("access_token"), to.name)

    if (!useCookie("access_token").value && to.name && !to.name.startsWith("auth")) {

        console.log("logout, redirecting...")
        return navigateTo({ name: "auth-login"});
    }
});
