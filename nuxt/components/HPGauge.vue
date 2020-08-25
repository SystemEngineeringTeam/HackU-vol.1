<template>
  <div>
    <v-progress-linear
      v-if="$store.getters['game/getHP'] <= 25"
      height="25"
      color="red"
      :value="$store.getters['game/getHP']"
    ></v-progress-linear>
    <v-progress-linear
      v-else-if="$store.getters['game/getHP'] <= 50"
      height="25"
      color="yellow"
      :value="$store.getters['game/getHP']"
    ></v-progress-linear>
    <v-progress-linear
      v-else
      height="25"
      color="green"
      :value="$store.getters['game/getHP']"
    ></v-progress-linear>
    {{ this.$store.state.game.HP }}/{{ this.$store.state.game.maxHP }}
  </div>
</template>

<script>
export default {
  name: 'HPGauge',

  data: () => ({}),

  methods: {},

  // created() {
  //   this.$store.commit('game/setHP', 100)
  //   //this.$store.commit('game/setMaxHP', 100)
  // },
}
</script>
