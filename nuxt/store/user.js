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

  async postTask({ rootState, commit, dispatch }) {
    await axios
      .post('localhost:8080/tasks', rootState.tasks.post, {
        params: { userToken: rootState.user.token },
      })
      .then((res) => {
        if (res.status === 200) {
          commit('addTask', res.data)
          dispatch('postAllReset')
        }
      })
  },

  postAllReset(context) {
    context.commit('setPostTitle', '')
    context.commit('setPostDeadlineDate', '')
    context.commit('setPostDeadlineTime', '')
    context.commit('setPostDescription', '')
    context.commit('setPostWeight', '')
  },

  async successTask({ state, rootState, commit }, taskID) {
    await axios
      .post(
        'localhost:8080/tasks/success',
        {},
        { params: { taskID: taskID, userToken: rootState.user.token } }
      )
      .then((res) => {
        if (res.status === 200) {
          let index = state.tasks.findIndex((element) => element.id === taskID)
          commit('removeTask', index)
        }
      })
  },
}
