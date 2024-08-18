<template>
  <a-form-item prop="icon" name="icon" label="Icon">
    <file-pond
        name="upload"
        ref="pond"
        class="filepond"
        label-idle="Chọn ảnh icon..."
        v-bind:allow-multiple="false"
        accepted-file-types="image/jpeg, image/png, image/svg+xml"
        :server="server"
        v-bind:files="myFiles"
        v-on:init="handleFilePondInit"
        v-on:processfile="handleProcessFile"
    />
  </a-form-item>
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
  name: "DepartmentIcon",
  components: {FilePond},
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
            this.value.icon = data.currentFolder.url + data.fileName;
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
    },
    async getResponseUpload(response) {
      response;
      let file = this.$refs.pond.getFile(0);

      let serverId = file.getMetadata("serverId");
    },
    handleProcessFile: function (error, file) {
      if (error == null) {
        this.value.icon = file.serverId;
      }
    },
  },
  created() {
    if (this.value.icon.length > 0) {
      this.myFiles = [
        {
          source: this.value.icon,
          options: {type: "local"},
        },
      ];
    }
  },
  mounted() {
  },
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
