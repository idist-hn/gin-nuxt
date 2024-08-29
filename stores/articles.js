import {defineStore} from 'pinia'
import api from "~/routes/api.js";
import {toast} from "vue3-toastify";

export const useArticleStore = defineStore('articles', {
    state: () => ({}),
    actions: {
        async ListEntries(pagination) {
            let response = null;
            await fetch(api.ListArticles + "?" + new URLSearchParams(pagination), {
                method: "GET",
            }).then((res) => res.json()).then((data) => {
                if (data.error !== undefined) {
                    toast.error(data.message);
                }
                response = data.data;
            });
            return response;
        },
        async ListRelateArticles(pagination) {
            let response = null;
            await fetch(api.ListRelateArticles + "?" + new URLSearchParams(pagination), {
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
            await fetch(api.params("GetArticle", {id: entry.id}), {
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
            entry.related_articles = [];
            entry.tags = [];
            await fetch(api.CreateArticle, {
                method: "POST",
                body: JSON.stringify(entry),
            }).then((res) => res.json()).then((data) => {
                if (data.error !== undefined) {
                    toast.error(data.message);
                }
                response = data.data;
            });
            return response;
        },
        async UpdateEntry(entry) {
            let response = null;
            entry.related_articles = [];
            entry.tags = [];
            await fetch(api.params("UpdateArticle", {id: entry.id}), {
                method: "PUT",
                body: JSON.stringify(entry),
            }).then((res) => res.json()).then((data) => {
                if (data.error !== undefined) {
                    toast.error(data.message);
                }
                response = data.data;
            });
            return response;
        },
    },
})