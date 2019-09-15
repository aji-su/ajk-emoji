<template>
  <v-overlay :value="loading">
    <v-progress-circular indeterminate />
  </v-overlay>
</template>

<script>
export default {
  props: {
    loading: {
      type: Boolean,
      required: true
    }
  }
};
</script>
