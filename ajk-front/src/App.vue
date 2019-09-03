<template>
  <v-app>
    <v-app-bar app>
      <v-toolbar-title>
        <v-btn text to="/">
          <span>ajk-emoji</span>
        </v-btn>
      </v-toolbar-title>
      <v-spacer></v-spacer>
    </v-app-bar>

    <v-content>
      <router-view />
    </v-content>
  </v-app>
</template>

<script>
export default {
  name: "App"
};
</script>
