<template>
  <v-card>
    <v-card-title class="justify-center">
      ステータス情報
    </v-card-title>
    <v-textarea solo readonly v-model="log"></v-textarea>
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
  </v-card>
</template>

<script>
export default {
  name: 'Logbox',

  //props: [],

  data: () => ({
    log: 'log',
  }),

  methods: {},

  computed: {
    getHP() {
      return (this.$store.state.user.HP / this.$store.state.user.maxHP) * 100
    },
  },
}
</script>
