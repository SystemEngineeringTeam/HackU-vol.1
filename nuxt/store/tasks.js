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
    ]
})

export const mutations = {
    setTasks(state,tasks){
        state.tasks = tasks;
    },

    addTask(state,task){
        state.tasks.push(task);
    },

    updateTask(state, task) {
        let index = state.tasks.findIndex((element) => element.id === task.id);
        state.tasks.splice(index, 1, task);
    },

    removeTask(state, index) {
      state.tasks.splice(index, 1);
    }
}

export const actions = {
    async setTasks(context){
        await axios.get("localhost:8080/tasks").then((res) => {
            context.commit("setTasks", res.data);
        })
    },
}