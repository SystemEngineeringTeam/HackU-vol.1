<template>
  <v-col >
    <p>{{task_title}}</p>
    <v-img
      src="https://1.bp.blogspot.com/-DSgUUXrWoFw/XVKfz2Z_3XI/AAAAAAABUEs/a9QCrDh18-grpZCL0O_pD7r4KWC921gawCLcBGAs/s1600/fantasy_game_character_slime.png"
      class="monster"
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

.monster {
  animation: fuwafuwa 1s linear infinite alternate;
}
</style>

<script>
export default {
  name: 'Monster',

  props: ['task_title'],

  data: () => ({}),

  methods: {},

  computed: {},
}
</script>
