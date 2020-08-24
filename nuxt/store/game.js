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
  logCount: [],
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

  setLogCount(state, logCount) {
    state.logCount = logCount
  },

  addTask(state) {
    state.logCount.push(1)
  },

  removeTask(state, index) {
    state.logCount.splice(index, 1)
  },
}

export const actions = {
  gameInit({ commit, dispatch }) {
    dispatch('getHP')
    commit('setLog', '')
    dispatch('logCountInit')
  },

  logCountInit({ rootState, commit }) {
    let logCount = []
    rootState.tasks.tasks.forEach(() => logCount.push(1))
    commit('setLogCount', logCount)
  },

  async getHP({ state, commit }) {
    await axios
      .get(process.env.URL_HP, {
        params: { userToken: state.token },
      })
      .then((res) => {
        if (res.status === 200) {
          //commit('setHP', state.maxHP)
          commit('setHP', res.data.hp)
          commit('setMaxHP', res.data.maxHp)
        }
      })
  },

  lowerHP({ state, commit }, amountReduceHP) {
    let hp = state.HP
    const damage = amountReduceHP
    hp = Math.max(0, hp - damage)
    commit('setHP', hp)
  },

  recoveryHP({ state, rootState, commit }) {
    if (rootState.tasks.tasks.length === 1) {
      commit('setHP', state.maxHP)
    } else {
      commit('setHP', Math.min(state.HP + 200000, state.maxHP))
    }
  },

  writeDamageLog({ state, rootState, commit }) {
    let logCount = state.logCount
    rootState.tasks.tasks.forEach((element, index) => {
      let attackRnd = Math.random()
      if (attackRnd <= 0.3) {
        let logIndex = Math.random()*state.logVariation.length
        logIndex = Math.floor(logIndex)
        log =
          element.title +
          state.logVariation[logIndex] +
          rootState.user.name +
          'は' +
          1 * logCount[index] +
          'のダメージを受けた！\n' +
          log
        dispatch('lowerHP', logCount[index] * 1)
        logCount[index] = 1
      } else {
        logCount[index] += 1
      }
    })
    commit('setLogCount', logCount)
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
