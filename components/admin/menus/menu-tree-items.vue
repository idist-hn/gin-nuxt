<template>
  <draggable v-bind="dragOptions" @start="drag = true"
             @end="drag = false" class="list-unstyled menu-list-items" :class="classes" tag="ul"
             :list="data" @change="sendUpdate"
             :group="{ name: 'g1' }" item-key="g1">
    <template #item="{element, index}">
      <div>
        <a-collapse accordion class="mb-10">
          <a-collapse-panel key="page">
            <template #header>
              <span>{{ element[title] }} <a-tag color="blue">{{ typeItem(element.type) }}</a-tag></span>
            </template>
            <template #extra>
              <a-button type="danger" class="mr-1" size="small" @click.prevent="remove(index)">
                <idist-feather-icons type="delete"/>
              </a-button>
            </template>
            <div class="mb-1">
              <template v-if="element.type === 'custom'">
                <menu-item-detail :value="element"/>
              </template>

              <template v-else-if="element.type === 'category'">
                <label>Tên hiển thị</label>
                <a-input type="text" v-model:value="element.title" class="mb-3"/>
              </template>

              <template v-else-if="element.type === 'sport'">
                <label>Tên hiển thị</label>
                <a-input type="text" v-model:value="element.title" class="mb-3"/>
              </template>
            </div>
          </a-collapse-panel>
        </a-collapse>
        <!--      <div>-->
        <!--        children: {{ el.children.length }} - recursive: {{ recursive }}-->
        <!--      </div>-->
        <menu-tree-items v-if="recursive" :parent_id="index" :title="title" :recursive="recursive"
                         :classes="classes + ` ml-20`"
                         :data="element.children"/>
      </div>
    </template>
  </draggable>
</template>

<script>
import draggable from 'vuedraggable';
import MenuItemDetail from "~/components/admin/menus/menu-item-detail.vue";
import IdistFeatherIcons from "~/components/commons/IdistFeatherIcons.vue";

export default {
  name: "MenuTreeItems",
  components: {IdistFeatherIcons, MenuItemDetail, draggable},

  computed: {
    dragOptions() {
      return {
        animation: 200,
        group: 'description',
        disabled: false,
        ghostClass: 'ghost',
      };
    },
  },
  data: () => ({
    new_item: '',
    entry: {
      data: [],
    },
  }),
  props: {
    parent_id: {
      type: Number,
      default: 0,
    },
    recursive: {
      type: Boolean,
      default: false,
    },
    data: {
      type: Array,
      default: () => ([]),
    },
    classes: {
      type: String,
      default: () => (''),
    },
    title: {
      type: String,
      default: () => ('title'),
    },

  },
  methods: {
    async remove(index) {
      this.data.splice(index, 1);
    },
    typeItem(type) {
      if (type == 'custom') return 'Tuỳ chỉnh';
      if (type == 'category') return 'Chuyên mục';
      return 'Khác';
    },
    sendUpdate(evt) {
    },
  },
  created() {
    this.data.map(el => {
      if (el.children == null) el.children = [];
    });
  },
};
</script>

<style scoped></style>
