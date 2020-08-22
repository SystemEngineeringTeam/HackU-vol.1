<template>
  <div>
    <v-progress-linear
      v-if="getHP <= 25"
      height="25"
      color="red"
      :value="getHP"
    ></v-progress-linear>
    <v-progress-linear
      v-else-if="getHP <= 50"
      height="25"
      color="yellow"
      :value="getHP"
    ></v-progress-linear>
    <v-progress-linear
      v-else
      height="25"
      color="green"
      :value="getHP"
    ></v-progress-linear>
    {{ this.$store.state.game.HP }}/{{ this.$store.state.game.maxHP }}
  </div>
</template>

<script>
export default {
  name: 'HPGauge',

  data: () => ({}),

  methods: {},

  computed: {
    getHP() {
      return (this.$store.state.game.HP / this.$store.state.game.maxHP) * 100
    },
  },

  // created() {
  //   this.$store.commit('game/setHP', 100)
  //   //this.$store.commit('game/setMaxHP', 100)
  // },
}
</script>
