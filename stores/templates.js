import {defineStore} from 'pinia'
import api from "~/routes/api.js";
import {toast} from "vue3-toastify";

export const useTemplateStore = defineStore('templates', {
    state: () => ({}),
    actions: {
        async ListEntries(pagination) {
            let response = null;
            await fetch(api.ListTemplates + "?" + new URLSearchParams(pagination), {
                method: "GET",
            }).then((res) => res.json()).then((data) => {
                if (data.error !== undefined) {
                    toast.error(data.message);
                }
                response = data.data;
            });
            return response;
        },
        async CreateEntry(entry) {
            let response = null;
            await fetch(api.CreateTemplate, {
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
        async DeleteEntry(entry) {
            let response = null;
            await fetch(api.params("DeleteTemplate", {id: entry.id}), {
                method: "DELETE",
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