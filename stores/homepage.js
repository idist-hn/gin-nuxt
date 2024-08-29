import {defineStore} from 'pinia'
import api from "~/routes/index.js";

export const useHomepageStore = defineStore('homepage', {
    state: () => ({
        collapsedMenu: false,
        theme: "light"
    }),
    actions: {
        async GetOverviewRegularPosts(filter) {
            let response = await fetch(api.GetOverviewRegularPosts,{
                method: "GET",
                params: filter,
            });
            return response.data;
        },
    },
})