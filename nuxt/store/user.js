import axios from 'axios'

export const state = () => ({
  token: '',
  name: '',
  HP: 100,
  maxHP: 100,
})

export const mutations = {
  setToken(state, token) {
    state.token = token
  },

  setName(state, name) {
    state.name = name
  },

  setHP(state, HP) {
    state.HP = HP
  },

  setMaxHP(state, maxHP){
    state.maxHP = maxHP
  }
}

export const actions = {
  async signup({}, post_json) {
    await axios
      .post(process.env.URL_SIGNUP, JSON.stringify(post_json))
      .then((res) => {
        if (res.status == 200) {
          console.log('ok!')
          this.$router.push('/login')
        }
      })
  },

  async login({state, commit }, post_json) {
    await axios
      .post(process.env.URL_LOGIN, JSON.stringify(post_json))
      .then((res) => {
        if (res.status == 200) {
          commit('setToken', res.data.token)
          commit('setName', res.data.name)
        }else{
          return
        }
      })
    await axios
      .get(process.env.URL_HP, {
        params: { userToken: state.token },
      })
      .then((res) => {
        if (res.status === 200) {
          commit('setHP', res.data.HP)
          commit('setMaxHP', res.data.maxHP)
          this.$router.push('/')
        }
      })
  },

  logout({ commit }) {
    commit('setToken', '')
    commit('setName', '')
  },
}
