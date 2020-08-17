export const state = () => ({
    token: "aaa",
    name: "uouo"
})

export const mutations = {
    setToken(state,token){
        state.token = token;
    },

    setName(state,name){
        state.name = name;
    }
}