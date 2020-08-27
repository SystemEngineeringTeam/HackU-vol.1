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
    mimic:
      'https://1.bp.blogspot.com/-_8wJqUxj-d4/W4PJlko8nmI/AAAAAAABOIc/Z-MzXgFr2OkbWRKja484G8tVn74a80h5QCLcBGAs/s800/character_game_mimic.png',
    slime:
      'https://1.bp.blogspot.com/-DSgUUXrWoFw/XVKfz2Z_3XI/AAAAAAABUEs/a9QCrDh18-grpZCL0O_pD7r4KWC921gawCLcBGAs/s1600/fantasy_game_character_slime.png',
    golem:
      'https://3.bp.blogspot.com/-ZWsv1eBwP-8/XDXcFKGXH2I/AAAAAAABRGs/bAVhn3sVs2wkaFSaeTzvwdAD3CuS47ZUACLcBGAs/s800/fantasy_golem.png',
    dragon:
      'https://4.bp.blogspot.com/-t0TdfnnfnH0/UT10GYML1QI/AAAAAAAAOrY/qNLEwXbzl-0/s1600/fantasy_dragon.png',
  }),

  methods: {},
}
</script>
