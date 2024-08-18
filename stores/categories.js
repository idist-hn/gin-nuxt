import {defineStore} from 'pinia'
import api from "~/api/index.js";
import {toast} from "vue3-toastify";

export const useCategoryStore = defineStore('categories', {
    state: () => ({}),
    actions: {
        async ListEntries(pagination) {
            let response = null;
            await fetch(api.ListCategories + "?" + new URLSearchParams(pagination), {
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
            await fetch(api.CreateCategory, {
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
        async GetEntry(entry) {
            let response = null;
            await fetch(api.params("DeleteCategory", {id: entry.id}), {
                method: "GET",
            }).then((res) => res.json()).then((data) => {
                if (data.error !== undefined) {
                    toast.error(data.message);
                }
                response = data;
            });
            return response;
        },
        async UpdateEntry(entry) {
            let response = null;
            await fetch(api.params("UpdateCategory", {id: entry.id}), {
                method: "PUT",
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
            await fetch(api.params("DeleteCategory", {id: entry.id}), {
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