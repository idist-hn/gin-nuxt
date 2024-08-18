// https://nuxt.com/docs/api/configuration/nuxt-config
import {theme} from "./utils/themeVariables";

// @ts-ignore
export default defineNuxtConfig({
    nitro: {
        routeRules: {
            '/api/**': { proxy: 'http://localhost:6006/api/**' },
            '/ckfinder/**': { proxy: 'http://localhost:8080/ckfinder/**' }
        }
    },
    ssr: true,
    devtools: {
        enabled: true,
        timeline: {
            enabled: true
        }
    },
    runtimeConfig: {
        public: {
            portalDomain: process.env.NUXT_ENV_PUBLIC_DOMAIN || "http://localhost:3000",
        },
    },
    devServer: {
        // @ts-ignore
        port: process.env.PORT || 3000,
        host: "0.0.0.0",
    },
    app: {
        baseURL: process.env.BASE_URL || "/",
        // pageTransition: {
        //     mode: "out-in",
        //     name: "fade",
        // },
        head: {
            title: "PTIT-CMS",
            htmlAttrs: {
                lang: "vi",
            },
            meta: [
                {charset: "utf-8"},
                {name: "viewport", content: "width=device-width, initial-scale=1"},
                {hid: "description", name: "description", content: ""},
                {name: "format-detection", content: "telephone=no"},
            ],
            link: [{rel: "icon", type: "image/x-icon", href: "/favicon.ico"}],
            script: [{src: "/ckfinder/ckfinder.js"}],
        },
        // pageTransition: {
        //     name: "fade",
        //     mode: "out-in",
        // },
    },
    modules: [
        "@formkit/auto-animate",
        "@nuxtjs/device",
        "@ant-design-vue/nuxt",
        "@nuxtjs/i18n",
        "nuxt-icon",
        "@pinia/nuxt",
    ],
    antd: {},
    alias: {
        pinia: "/node_modules/@pinia/nuxt/node_modules/pinia/dist/pinia.mjs",
    },
    pinia: {
        storesDirs: [
            "~/stores",
        ],
    },

    plugins: [
        "~/plugins/toast.client",
    ],
    css: [
        "~/assets/less/antd.less",
        "~/assets/less/custom.less",
    ],
    vite: {
        "css": {
            "preprocessorOptions": {
                "less": {
                    "javascriptEnabled": true,
                    "modifyVars": theme,
                    "math": "always"
                },
                "sass": {
                    "implementation": require("sass"),
                },
                "scss": {
                    "implementation": require("sass"),
                    "sassOptions": {},
                },
            },
        },
    },
})