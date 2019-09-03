<template>
  <v-container>
    <v-overlay :value="loading">
      <v-progress-circular indeterminate />
    </v-overlay>

    <v-snackbar top :value="errorMessage">
      {{ errorMessage }}
    </v-snackbar>

    <v-layout wrap>
      <v-flex xs12 v-if="emojis">
        <span class="emoji-row" v-for="(row, i) in emojis" :key="i">
          <img class="emoji" v-for="p in row" :key="p.shortcode" :src="p.url" />
          <br />
        </span>
      </v-flex>

      <v-flex xs12 v-if="emojis">
        <v-textarea
          label="Emoji shortcodes"
          :value="shortcodes"
          readonly
          @focus="$event.target.select()"
        />
      </v-flex>

      <v-flex xs12>
        <v-text-field
          :value="link"
          label="Share this page"
          readonly
          @focus="$event.target.select()"
        />
      </v-flex>

      <v-flex xs12>
        <v-btn color="primary" @click="onDownload" :disabled="loading">
          Download .tar.gz
        </v-btn>
      </v-flex>

      <v-flex xs8 pa-4>
        Original image:
        <v-img
          v-if="originalImageUrl"
          :src="originalImageUrl"
          max-height="80%"
          contain
        />
      </v-flex>
    </v-layout>
  </v-container>
</template>

<script>
import show from "@/utils/show.js";
import getDownloadLink from "@/utils/getDownloadLink.js";

export default {
  name: "Show",
  props: {
    requestId: {
      type: String,
      required: true
    }
  },
  data() {
    return {
      errorMessage: "",
      loading: true,
      emojis: [],
      shortcodes: "",
      originalImageUrl: ""
    };
  },
  computed: {
    link() {
      return location.href;
    }
  },
  async created() {
    try {
      const { emojis, originalImageUrl } = await show(this.requestId);
      this.loading = false;
      this.emojis = emojis;
      this.shortcodes = emojis
        .map(row => row.map(e => `:${e.shortcode}:`).join("\u200b"))
        .join("\n");
      this.originalImageUrl = originalImageUrl;
    } catch (err) {
      this.loading = false;
      this.errorMessage = err.toString();
      throw err;
    }
  },
  methods: {
    async onDownload() {
      try {
        this.loading = true;
        const url = await getDownloadLink(this.requestId);
        location.href = url;
        this.loading = false;
      } catch (err) {
        this.loading = false;
        this.errorMessage = err.toString();
        throw err;
      }
    }
  }
};
</script>

<style scoped>
.errorMessage {
  color: red;
}

.emoji-row {
  line-height: 0;
}

img.emoji {
  border: solid 1px black;
  width: 24px;
}
</style>
