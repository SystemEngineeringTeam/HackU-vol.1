import axios from 'axios'

export const state = () => ({
    tasks: [
        {
            id: 0,
            name: "ごはん　たべる",
            deadlineDate: "2020-08-21",
            deadlineTime: "12:00:00",
            description: "ふくだくんを焼いて食べる",
            weight: "0",
        },
        {
            id: 1,
            name: "うああああ",
            deadlineDate: "2020-8-22",
            deadlineTime: "09:12:00",
            description: "じょぼじょぼのじょぼじょぼりん",
            weight: "1"
        },
        {
            id: 2,
            name: "ウホウホミッドナイト",
            deadlineDate: "2020-8-22",
            deadlineTime: "21:30:00",
            description: "ウホウホ！！！？！？？！？！？！？！？！？ウホウホウホウホウホウホウホ！！！！ウホホホホホホホホホｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗウッキーーーーーーーーー！！！！！！！！！！！！！！！！！ｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗｗ",
            weight: "2"
        }
    ],

    post: {
        name: "",
        deadlineDate: "",
        deadlineTime: "",
        description: "",
        weight: ""
    }
})

export const mutations = {
    setTasks(state, tasks) {
        state.tasks = tasks;
    },

    addTask(state, task) {
        state.tasks.push(task);
    },

    updateTask(state, task) {
        let index = state.tasks.findIndex((element) => element.id === task.id);
        state.tasks.splice(index, 1, task);
    },

    removeTask(state, index) {
        state.tasks.splice(index, 1);
    },

    setPostName(state, name) {
        state.post.name = name;
    },

    setPostDeadlineDate(state, deadlineDate) {
        state.post.deadlineDate = deadlineDate;
    },

    setPostDeadlineTime(state, deadlineTime) {
        state.post.deadlineTime = deadlineTime;
    },

    setPostDescription(state, description) {
        state.post.description = description;
    },

    setPostWeight(state, weight) {
        state.post.weight = weight;
    }
}

export const actions = {
    async setTasks({ rootState, commit }) {
        await axios.get("localhost:8080/tasks", { params: { userToken: rootState.user.token } })
            .then((res) => {
                commit("setTasks", res.data);
            })
    },

    async postTask({ rootState, commit, dispatch }) {
        await axios.post("localhost:8080/tasks", rootState.tasks.post, { params: { userToken: rootState.user.token } })
            .then((res) => {
                if (res.status === 200) {
                    commit("addTask", res.data);
                    dispatch("postAllReset");
                }
            })
    },

    postAllReset(context) {
        context.commit("setPostName", "");
        context.commit("setPostDeadlineDate", "");
        context.commit("setPostDeadlineTime", "");
        context.commit("setPostDescription", "");
        context.commit("setPostWeight", "");
    },

    //taskIDを指定する場合
    async successTask({rootState}, taskID) {
        await axios.post("localhost:8080/tasks/success", {}, { params: { taskID: taskID, userToken: rootState.user.token} })
            .then((res) => {
                console.log(res.status);
                //何かすることがあればここに書く
            })
    }
}
