import axios from 'axios'

export const state = () => ({
  tasks: [
    {
      id: 0,
      title: 'ごはん　たべる',
      deadlineDate: '2020-08-21',
      deadlineTime: '12:00:00',
      description: 'ふくだくんを焼いて食べる',
      weight: '0',
    },
    {
      id: 1,
      title: 'うああああ',
      deadlineDate: '2020-8-22',
      deadlineTime: '09:12:00',
      description: 'じょぼじょぼのじょぼじょぼりん',
      weight: '1',
    },
    {
      id: 2,
      title: 'ウホウホミッドナイト',
      deadlineDate: '2020-8-22',
      deadlineTime: '21:30:00',
      description:
        'ウホウホ！！！？！？？！？！？！？！？！？ウホウホウホウホウホウホウホ！！！！ウホホホホホホホホホｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗウッキーーーーーーーーー！！！！！！！！！！！！！！！！！ｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗ',
      weight: '2',
    },
    {
      id: 3,
      title: 'ウホウホミッドナイト',
      deadlineDate: '2020-8-22',
      deadlineTime: '21:30:00',
      description:
        'ウホウホ！！！？！？？！？！？！？！？！？ウホウホウホウホウホウホウホ！！！！ウホホホホホホホホホｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗウッキーーーーーーーーー！！！！！！！！！！！！！！！！！ｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗ',
      weight: '2',
    },
    {
      id: 4,
      title: 'ウホウホミッドナイト',
      deadlineDate: '2020-8-22',
      deadlineTime: '21:30:00',
      description:
        'ウホウホ！！！？！？？！？！？！？！？！？ウホウホウホウホウホウホウホ！！！！ウホホホホホホホホホｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗウッキーーーーーーーーー！！！！！！！！！！！！！！！！！ｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗ',
      weight: '2',
    },
    {
      id: 5,
      title: 'ウホウホミッドナイト',
      deadlineDate: '2020-8-22',
      deadlineTime: '21:30:00',
      description:
        'ウホウホ！！！？！？？！？！？！？！？！？ウホウホウホウホウホウホウホ！！！！ウホホホホホホホホホｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗウッキーーーーーーーーー！！！！！！！！！！！！！！！！！ｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗ',
      weight: '2',
    },
    {
      id: 6,
      title: 'ウホウホミッドナイト',
      deadlineDate: '2020-8-22',
      deadlineTime: '21:30:00',
      description:
        'ウホウホ！！！？！？？！？！？！？！？！？ウホウホウホウホウホウホウホ！！！！ウホホホホホホホホホｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗウッキーーーーーーーーー！！！！！！！！！！！！！！！！！ｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗ',
      weight: '2',
    },
  ],

  post: {
    title: '',
    deadlineDate: '',
    deadlineTime: '',
    description: '',
    weight: '',
  },

  weights: ['ぬるい', 'ふつう', 'えぐい'],
})

export const mutations = {
  setTaskIDbyDeleteAnimation(state, index) {
    let task = state.tasks[index]
    task.id = -1
    state.tasks.splice(index, 1, task)
  },

  setTasks(state, tasks) {
    state.tasks = tasks
  },

  addTask(state, task) {
    state.tasks.push(task)
  },

  updateTask(state, task) {
    let index = state.tasks.findIndex((element) => element.id === task.id)
    state.tasks.splice(index, 1, task)
  },

  removeTask(state, index) {
    state.tasks.splice(index, 1)
  },

  setPostTitle(state, title) {
    state.post.title = title
  },

  setPostDeadlineDate(state, deadlineDate) {
    state.post.deadlineDate = deadlineDate
  },

  setPostDeadlineTime(state, deadlineTime) {
    state.post.deadlineTime = deadlineTime
  },

  setPostDescription(state, description) {
    state.post.description = description
  },

  setPostWeight(state, weight) {
    state.post.weight = weight
  },

  setWeights(state, weights) {
    state.weights = weights
  },
}

export const getters = {
  tasks: (state) => {
    return state.tasks
  },
}

export const actions = {
  async setTasks({ rootState, commit }) {
    await axios
      .get(process.env.URL_TASKS, {
        params: { userToken: rootState.user.token },
      })
      .then((res) => {
        commit('setTasks', res.data)
      })
  },

  async postTask({ rootState, commit, dispatch }) {
    let post_json = JSON.parse(JSON.stringify(rootState.tasks.post))
    if (post_json.deadlineDate && post_json.deadlineTime) {
      post_json.deadlineTime = post_json.deadlineTime + ':00'
    } else if (post_json.deadlineDate && !post_json.deadlineTime) {
      post_json.deadlineTime = '23:59:59'
    } else if (!post_json.deadlineDate && post_json.deadlineTime) {
      post_json.deadlineDate = new Date().toISOString().substr(0, 10)
      post_json.deadlineTime = post_json.deadlineTime + ':00'
    } else {
      post_json.deadlineDate = ''
      post_json.deadlineTime = ''
    }
    await axios
      .post(process.env.URL_TASKS, JSON.stringify(post_json), {
        params: { userToken: rootState.user.token },
      })
      .then((res) => {
        if (res.status === 200) {
          const new_task = {
            id: res.data,
            title: post_json.title,
            deadlineDate: post_json.deadlineDate,
            deadlineTime: post_json.deadlineTime,
            description: post_json.description,
            weight: post_json.weight,
          }
          commit('addTask', new_task)
          commit('game/addTask', null, { root: true })
          dispatch('postAllReset')
        }
      })
  },

  postAllReset(context) {
    context.commit('setPostTitle', '')
    context.commit('setPostDeadlineDate', null)
    context.commit('setPostDeadlineTime', null)
    context.commit('setPostDescription', '')
    context.commit('setPostWeight', '')
  },

  async successTask({ state, rootState, commit }, index) {
    await axios
      .post(
        process.env.URL_TASKS_SUCCESS,
        {},
        {
          params: {
            taskID: state.tasks[index].id,
            userToken: rootState.user.token,
          },
        }
      )
      .then((res) => {
        if (res.status === 200) {
          commit('removeTask', index)
          commit('game/removeTask', index, { root: true })
        }
      })
  },

  async getWeights({ commit }) {
    await axios.get(process.env.URL_WEIGHTS).then((res) => {
      if (res.status === 200) {
        commit('setWeights', res.data)
      }
    })
  },

  logout({ commit }) {
    commit('setTasks', [])
  },
}
