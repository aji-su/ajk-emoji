<template>
  <v-snackbar top :value="message" :timeout="0">
    {{ message }}
    <v-btn color="pink" @click="onClickClose">
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
