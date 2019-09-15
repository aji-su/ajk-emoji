<template>
  <v-container>
    <Loading :loading="loading" />
    <Errors :message.sync="errorMessage" />

    <v-layout wrap>
      <v-flex xs12>
        <v-btn @click.native="openFileDialog">
          Select image
        </v-btn>
      </v-flex>

      <v-flex xs12>
        <input
          id="file-upload"
          type="file"
          accept="image/*"
          style="display:none"
          @change="onChangeFile"
        />
      </v-flex>

      <v-flex xs12>
        <v-text-field v-model="prefix" label="Shortcode prefix" required />
      </v-flex>

      <v-flex xs12>
        <v-text-field
          v-model="xsplit"
          type="number"
          label="Number of X split"
        />
      </v-flex>

      <v-flex xs10>
        <v-btn color="primary" @click="onSubmit" :disabled="uploadReady">
          Send
        </v-btn>
      </v-flex>

      <v-flex xs8 pa-4>
        <v-img
          v-show="imageAsDataUrl"
          :src="imageAsDataUrl"
          max-height="80%"
          contain
        />
      </v-flex>
    </v-layout>
  </v-container>
</template>

<script>
import upload from "@/utils/upload.js";
import Loading from "@/components/molecules/Loading.vue";
import Errors from "@/components/molecules/Errors.vue";

export default {
  name: "FileUpload",
  props: {},
  components: {
    Loading,
    Errors
  },
  data() {
    return {
      errorMessage: "",
      imageAsDataUrl: "",
      prefix: "zzz",
      xsplit: "10",
      loading: false
    };
  },

  computed: {
    uploadReady() {
      return !this.imageAsDataUrl && !this.loading;
    }
  },

  methods: {
    openFileDialog() {
      document.getElementById("file-upload").click();
    },

    onChangeFile(e) {
      const file = e.target.files[0];
      if (file) {
        const reader = new FileReader();
        reader.onload = () => {
          this.imageAsDataUrl = reader.result;
        };
        reader.readAsDataURL(file);
      }
    },

    async onSubmit() {
      if (!this.imageAsDataUrl) {
        return;
      }
      this.errorMessage = "";
      this.loading = true;
      try {
        const requestId = await upload(
          this.prefix,
          parseInt(this.xsplit),
          this.imageAsDataUrl
        );
        this.$router.push({ name: "show", params: { requestId } });
      } catch (err) {
        this.loading = false;
        this.errorMessage = err.toString();
        throw err;
      }
    }
  }
};
</script>
