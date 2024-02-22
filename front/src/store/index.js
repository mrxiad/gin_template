import {createPinia} from 'pinia'
import { defineStore } from "pinia"

export const useUserStore = defineStore("user", {
    state: () => ({
        name: "石头",
        age: 18,
    }),
    getters: {
        getPerson: (state) => {
            return `${state.name}今年${state.age}岁了`
        }
    },
    actions: {
        changeName(name) {
            const that = this;
            setTimeout(() => {
                that.name = name;
            }, 1000)
        },
    },
})