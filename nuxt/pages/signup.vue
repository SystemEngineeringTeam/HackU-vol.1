<template>
  <v-layout column justify-center align-center>
    <v-flex 12>
      <v-form ref="form" v-model="valid">
        <h1>サインアップ</h1>
        <v-text-field
          v-model="name"
          label="ユーザーネーム"
          required
          outlined
        ></v-text-field>

        <v-text-field
          v-model="email"
          label="E-mail"
          required
          outlined
        ></v-text-field>

        <v-text-field
          v-model="pass"
          label="パスワード"
          required
          outlined
        ></v-text-field>

        <v-btn
          :disabled="!valid"
          color="success"
          class="mr-4"
          @click="create_user"
        >
          Validate
        </v-btn>
      </v-form>
    </v-flex>
  </v-layout>
</template>

<script>
export default {
  data() {
    return {
      valid: false,
      name: '',
      email: '',
      pass: '',
    }
  },
  methods: {
    create_user() {
      console.log(this.name, this.email, this.pass)
    },
  },
}
</script>

<style></style>
