<template>
  <v-col>
    <p>{{ task.title }}</p>
    <v-img
      v-if="!task.weight"
      :src="mimic"
      :class="{ monster: !killing, kill_monster: killing }"
    />
    <v-img
      v-else-if="task.weight === 'ぬるい'"
      :src="slime"
      :class="{ monster: !killing, kill_monster: killing }"
    />
    <v-img
      v-else-if="task.weight === 'ふつう'"
      :src="golem"
      :class="{ monster: !killing, kill_monster: killing }"
    />
    <v-img
      v-else-if="task.weight === 'えぐい'"
      :src="dragon"
      :class="{ monster: !killing, kill_monster: killing }"
    />
  </v-col>
</template>

<style scoped>
@keyframes fuwafuwa {
  0% {
    transform: translate(0, 0) rotate(-5deg);
  }
  50% {
    transform: translate(0, -5px) rotate(0deg);
  }
  100% {
    transform: translate(0, 0) rotate(5deg);
  }
}

@keyframes flash {
  0% {
    opacity: 1;
  }

  100% {
    opacity: 0;
  }
}

.monster {
  animation: fuwafuwa 1s linear infinite alternate;
}
.kill_monster {
  animation: flash 0.15s linear infinite alternate;
}
</style>

<script>
export default {
  name: 'Monster',

  props: ['task'],

  computed: {
    store_tasks: function () {
      return this.$store.getters['tasks/tasks']
    },
  },

  watch: {
    store_tasks: {
      handler: function () {
        if (this.task.id === -1) {
          this.killing = true
          setTimeout(() => {
            this.killing = false
          }, 1000)
        }
      },
      deep: true,
    },
  },

  data: () => ({
    killing: false,
    mimic: require('@/static/monster/mimic_dot.png'),
    slime: require('@/static/monster/slime_dot.png'),
    golem: require('@/static/monster/golem_dot.png'),
    dragon: require('@/static/monster/dragon_dot.png'),
  }),

  methods: {},
}
</script>
