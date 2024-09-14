import {defineStore} from 'pinia'
import api from "~/routes/api.js";
import {toast} from "vue3-toastify";

export const useMenuStore = defineStore('menus', {
    state: () => ({}),
    actions: {
        async ListEntries(pagination) {
            let response = null;
            await fetch(api.ListMenus + "?" + new URLSearchParams(pagination), {
                method: "GET",
            }).then((res) => res.json()).then((data) => {
                if (data.error !== undefined) {
                    toast.error(data.message);
                }
                response = data.data;
            });
            return response;
        },
        async ListMenuLocations(pagination) {
            let response = null;
            await fetch(api.ListMenuLocations + "?" + new URLSearchParams(pagination), {
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
            await fetch(api.params("GetMenu", {id: entry.id}), {
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
            await fetch(api.CreateMenu, {
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
            await fetch(api.params("UpdateMenu", {id: entry.id}), {
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
            await fetch(api.params("DeleteMenu", {id: entry.id}), {
                method: "DELETE",
            }).then((res) => res.json()).then((data) => {
                if (data.error !== undefined) {
                    toast.error(data.message);
                }
                response = data;
            });
            return response;
        },

        // Methods for Common

        async GetSubMenuEntry() {
            let response = null;
            let filter = {
                language: "vn",
                status: "active",
            }
            await fetch(api.params("GetSubMenu", {query: filter}), {
                method: "GET",

            }).then((res) => res.json()).then((data) => {
                if (data.error !== undefined) {
                    toast.error(data.message);
                }
                response = data;
                console.log(data)
            });
            return response;
        },
    },
})