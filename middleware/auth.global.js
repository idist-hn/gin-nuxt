export default defineNuxtRouteMiddleware((to, from) => {
    console.log({access_token: useCookie("access_token").value, "auth": to.name })

    if (!useCookie("access_token").value && to.name && !to.name.startsWith("auth")) {
        console.log({access_token: useCookie("access_token").value, "auth": to.name })
        return navigateTo({ name: "auth-login"});
    }
});
