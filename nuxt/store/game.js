import axios from 'axios'

export const state = () => ({
  HP: 750000,
  maxHP: 1000000,
  log: '',
})

export const mutations = {
  setHP(state, HP) {
    state.HP = HP
  },

  setMaxHP(state, maxHP) {
    state.maxHP = maxHP
  },

  setLog(state, log) {
    state.log = log
  },
}

export const actions = {
  async getHP({ state, commit }) {
    await axios
      .get(process.env.URL_HP, {
        params: { userToken: state.token },
      })
      .then((res) => {
        if (res.status === 200) {
          console.log(235)
          commit('setHP', res.data.HP)
          commit('setMaxHP', res.data.maxHP)
        }
      })
  },

  logout({ state, commit }) {
    commit('setHP', state.maxHP)
    commit('setLog', '')
  },
}
