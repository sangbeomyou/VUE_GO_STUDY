<template>
  <v-container>
    <v-row justify="center" align="center" class="fill-height">
      <v-col cols="12" sm="8" md="6">
        <v-card class="elevation-12">
          <v-toolbar color="primary" dark flat>
            <v-toolbar-title>Login</v-toolbar-title>
          </v-toolbar>
          <v-card-text>
            <v-text-field
              label="Username"
              name="username"
              prepend-icon="mdi-account"
              type="text"
              v-model="username"
            ></v-text-field>
            <v-text-field
              id="password"
              label="Password"
              name="password"
              prepend-icon="mdi-lock"
              type="password"
              v-model="password"
            ></v-text-field>
          </v-card-text>
          <v-card-actions>
            <v-btn color="primary" @click="login">Login</v-btn>
          </v-card-actions>
        </v-card>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
export default {
  data: () => ({
    username: '',
    password: ''
  }),
  methods: {
    login() {
      // 로그인 로직을 여기에 작성하세요.
      console.log('Logging in with', this.username, this.password);
    }
  }
}
</script>