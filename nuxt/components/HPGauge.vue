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
    {{ this.$store.state.user.HP }}/{{ this.$store.state.user.maxHP }}
  </div>
</template>

<script>
export default {
  name: 'HPGauge',

  //props: [],

  data: () => ({
    lowerHPid: null,
  }),

  methods: {
    lowerHP: function () {
      let hp = this.$store.state.user.HP
      hp -= 1
      this.$store.commit('user/setHP', hp)
    },
  },

  computed: {
    getHP() {
      return (this.$store.state.user.HP / this.$store.state.user.maxHP) * 100
    },
  },

  created() {
    this.$store.commit('user/setHP', 100)
    this.lowerHPid = setInterval(this.lowerHP, 1000)
  },

  destroyed() {
    clearInterval(this.lowerHPid)
  },
}
</script>
