import {defineStore} from 'pinia'
import api from "~/routes/index.js";
import {toast} from "vue3-toastify";

export const usePermissionsStore = defineStore('permissions', {
    state: () => ({}),
    actions: {
        async ListEntries(pagination) {
            let response = null;
            await fetch(api.ListPermissions + "?" + new URLSearchParams(pagination), {
                method: "GET",
            }).then((res) => res.json()).then((data) => {
                if (data.error !== undefined) {
                    toast.error(data.message);
                }
                response = data.data;
            });
            return response;
        },

        async GetEntry(entry) {
            let response = null;
            await fetch(api.params("GetPermission", {id: entry.id}), {
                method: "GET",
            }).then((res) => res.json()).then((data) => {
                if (data.error !== undefined) {
                    toast.error(data.message);
                }
                response = data;
            });
            return response;
        },
    },
})