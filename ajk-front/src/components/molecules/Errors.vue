<template>
  <v-snackbar top :value="message">
    {{ message }}
    <v-btn color="pink" @click="onClickClose" timeout="0">
      Close
    </v-btn>
  </v-snackbar>
</template>

<script>
export default {
  props: {
    message: {
      type: String,
      required: false
    }
  },
  methods: {
    onClickClose() {
      this.$emit("update:message", "");
    }
  }
};
</script>
