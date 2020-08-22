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
  gameInit({ commit, dispatch }) {
    dispatch('getHP')
    commit('setLog', '')
  },

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

  lowerHP({ state, rootState, commit }) {
    let hp = state.HP
    const damage = rootState.tasks.tasks.length
    hp = Math.max(0, hp - damage)
    commit('setHP', hp)
  },

  writeDamageLog({ state, rootState, commit }) {
    let log = state.log
    rootState.tasks.tasks.forEach((element) => {
      log =
        element.title +
        'の攻撃！' +
        rootState.user.name +
        'は' +
        1 +
        'のダメージを受けた！\n' +
        log
    })
    commit('setLog', log)
  },

  writeSuccessLog({ state, commit }, title) {
    let log = title + 'を倒した！\n' + state.log
    commit('setLog', log)
  },

  logout({ state, commit }) {
    commit('setHP', state.maxHP)
    commit('setLog', '')
  },
}
