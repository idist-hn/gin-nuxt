import {defineStore} from 'pinia'
import api from "~/routes/index.js";
import {toast} from "vue3-toastify";

export const useProfileStore = defineStore('profile', {
    state: () => ({
        authenticated: false,
        profile: null
    }),
    actions: {
        async PostLogin(entry) {
            let response = null
            await fetch(api.PostLogin, {
                method: "POST",
                body: JSON.stringify(entry),
            }).then((res) => res.json()).then((data) => {
                if (data.error !== undefined) {
                    toast.error(data.message);
                }
                response = data;
            });
            return response;
        },
        async PostLogout() {
            let response = null;
            await fetch(api.PostLogout, {
                method: 'POST',
            }).then((res) => res.json()).then((data) => {
                if (data.error !== undefined) {
                    toast.error(data.message);
                }
                response = data;
            });
            return response
        },
        async GetProfile() {
            let response = null;
            await fetch(api.GetProfile, {
                method: "GET",
            }).then((res) => res.json()).then((data) => {
                if (data.error !== undefined) {
                    toast.error(data.message);
                }
                response = data;
            });

            this.useState().profile = response;
            return response;
        },
    },
})