import {defineStore} from 'pinia'
import api from "~/api/index.js";
import {toast} from "vue3-toastify";

export const useDepartmentStore = defineStore('departments', {
    state: () => ({}),
    actions: {
        async ListEntries(pagination) {
            let response = null;
            await fetch(api.ListDepartments + "?" + new URLSearchParams(pagination), {
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
            await fetch(api.params("GetDepartment", {id: entry.id}), {
                method: "GET",
            }).then((res) => res.json()).then((data) => {
                if (data.error !== undefined) {
                    toast.error(data.message);
                }
                response = data;
            });
            return response;
        },
        async CreateEntry(entry) {
            let response = null;
            await fetch(api.CreateDepartment, {
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
        async UpdateEntry(entry) {
            let response = null;
            await fetch(api.params("UpdateDepartment", {id: entry.id}), {
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
            await fetch(api.params("DeleteDepartment", {id: entry.id}), {
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