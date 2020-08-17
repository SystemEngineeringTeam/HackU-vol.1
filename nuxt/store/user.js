export const state = () => ({
    token: "aaa",
    email: "aaa@aaa.com",
    name: "uouo"
})

export const mutations = {
    setToken(state,token){
        state.token = token;
    },

    setEmail(state,email){
        state.email = email;
    },

    setName(state,name){
        state.name = name;
    }
}