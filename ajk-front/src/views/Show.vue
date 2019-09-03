<template>
  <Show :requestId="$route.params.requestId" />
</template>

<script>
import Show from "@/components/Show.vue";

export default {
  name: "show",
  components: {
    Show
  }
};
</script>
