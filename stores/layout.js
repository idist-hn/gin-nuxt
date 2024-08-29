import {defineStore} from 'pinia'

export const useLayoutStore = defineStore('layout', {
    state: () => ({
        collapsedMenu: false,
        theme: "light"
    }),
    actions: {
        toggleMenu() {
            this.collapsedMenu = !this.collapsedMenu;
        },
        setTheme(theme) {
            this.theme = theme;
        },
        changeSideNavMode() {
            setTimeout(() => {
                // this.collapsedMenu = !this.collapsedMenu
                this.collapsedMenu = false
            }, 10);
        },
        changeThemeMode() {
            this.theme = this.theme === 'light' ? 'dark' : 'light'
        },
    },
    // getters: {
    //     getTheme: (state) => {
    //         return state.theme;
    //     },
    //     getMenuMode: (state) => {
    //         return state.collapsedMenu;
    //     },
    // },
})