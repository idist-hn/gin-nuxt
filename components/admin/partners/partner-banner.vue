<template>
    <file-pond
      name="upload"
      ref="pond"
      class="filepond"
      label-idle="Chọn ảnh banner của đơn vị..."
      v-bind:allow-multiple="false"
      accepted-file-types="image/jpeg, image/png, image/webp"
      :server="server"
      v-bind:files="myFiles"
      v-on:init="handleFilePondInit"
      v-on:processfile="handleProcessFile"
    />
</template>

<script>
// Import Vue FilePond
import vueFilePond from "vue-filepond";

// Import FilePond styles
import "filepond/dist/filepond.min.css";

// Import FilePond plugins
// Please note that you need to install these plugins separately

// Import image preview plugin styles
import "filepond-plugin-image-preview/dist/filepond-plugin-image-preview.min.css";

// Import image preview and file type validation plugins
import FilePondPluginFileValidateType from "filepond-plugin-file-validate-type";
import FilePondPluginImagePreview from "filepond-plugin-image-preview";

// Create component
const FilePond = vueFilePond(FilePondPluginFileValidateType, FilePondPluginImagePreview);

export default {
  name: "PartnerBanner",
  components: { FilePond },
  props: {
    value: Object,
  },
  data() {
    return {
      entry: {},
      myFiles: "",
      server: {
        process: {
          url: `/ckfinder/core/connector/php/connector.php?command=FileUpload&type=Files&currentFolder=&hash=523cac8c2c3bbeda&responseType=json`,
          method: "POST",
          withCredentials: false,
          onload: (response) => {
            let data = JSON.parse(response);
            this.value.banner = data.currentFolder.url + data.fileName;
            return data.currentFolder.url + data.fileName;
          },
          ondata: (formData) => {
            return formData;
          },
        },
        // eslint-disable-next-line no-unused-vars
        load: (url, load, error, progress, abort, headers) => {
          fetch(url)
            .then((res) => {
              return res.blob();
            })
            .then(load);
          return {
            abort: () => {
              abort();
            },
          };
        },
      },
    };
  },
  methods: {
    handleFilePondInit: function () {
      console.log("FilePond has initialized");
    },
    async getResponseUpload(response) {
      response;
      let file = this.$refs.pond.getFile(0);

      let serverId = file.getMetadata("serverId");
    },
    handleProcessFile: function (error, file) {
      if (error == null) {
        this.value.banner = file.serverId;
      }
    },
  },
  created() {
    if (this.value.banner.length > 0) {
      this.myFiles = [
        {
          source: this.value.banner,
          options: { type: "local" },
        },
      ];
    }
  },
  mounted() {},
};
</script>

<style lang="scss">
.filepond--image-preview-overlay-success {
  display: none !important;
}

.filepond--credits {
  display: none !important;
}
</style>
