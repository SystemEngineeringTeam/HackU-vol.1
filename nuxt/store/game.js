import axios from 'axios'

export const state = () => ({
  HP: 750000,
  maxHP: 1000000,
  log: '',

  logVariation: [
    'は火を吐いた！ファイアー！',
    'の毒ガス攻撃！',
    'の迫真の攻撃！',
    'は顔からビームを発射した！',
    'のグーパン！',
    'は叫んで超音波を出した！ギョエェェェエ！！',
    'は冷凍ビームを発射！',
    'はじゃんけんを強要してきた！負けた！',
    'は石を投げつけてきた！',
    'の精神的圧力！仕事をしろ！'
  ],
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
          // commit('setHP', 750000)
          commit('setHP', res.data.hp)
          commit('setMaxHP', res.data.maxHp)
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
      let logIndex = Math.random(state.logVariation.length)*state.logVariation.length
      logIndex = Math.floor(logIndex)
      log =
        element.title +
        state.logVariation[logIndex] +
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
