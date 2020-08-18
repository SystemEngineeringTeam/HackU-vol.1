import axios from 'axios'

export const state = () => ({
  token: '',
  name: '',
})

export const mutations = {
  setToken(state, token) {
    state.token = token
  },

  setName(state, name) {
    state.name = name
  },
}

export const actions = {
  async signup({}, post_json) {
    await axios.post(process.env.URL_SIGNUP, post_json).then((res) => {
      if (res.status == 200) {
        console.log('ok!')
        this.$router.push('/login')
      }
    })
  },

  async login({ commit }, post_json) {
    await axios.post(process.env.URL_LOGIN, post_json).then((res) => {
      if (res.status == 200) {
        commit('setToken', res.data.token)
        commit('setName', res.data.name)
        this.$router.push('/')
      }
    })
  },
}
