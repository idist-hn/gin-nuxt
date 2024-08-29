<template>
  <div class='message'>
    <a-popover placement='bottomLeft' trigger='click'>
      <a-badge dot :offset='[-4, 0]'>
        <a to='#' class='head-example'>
          <IdistFeatherIcons type='mail' :size='20' />
        </a>
      </a-badge>
      <template #content>
        <div class='atbd-top-dropdown'>
          <h4 class='atbd-top-dropdown__title'>
            <span class='title-text'>Chưa đọc</span>
            <a-badge class='badge-success' count='3' />
          </h4>
          <div class='atbd-top-dropdown-menu'>
            <perfect-scrollbar
              :options='{
                wheelSpeed: 1,
                swipeEasing: true,
                suppressScrollX: true,
              }'
            >
              <ul class='atbd-top-dropdown__nav'>
                <li>
                  <nuxt-link to='#'>
                    <figure class='atbd-top-dropdown__content'>
                      <img src='~/assets/images/NoPath.png' alt='' />
                      <figcaption>
                        <h4> Software <span class='color-success'>3 hrs ago </span></h4>
                        <div>
                          <span class='atbd-top-dropdownText'>Lorem ipsum dolor amet cosec...</span>
                          <span>
                            <a-badge class='badge-success' count='3' />
                          </span>
                        </div>
                      </figcaption>
                    </figure>
                  </nuxt-link>
                </li>
                <li>
                  <nuxt-link to='#'>
                    <figure class='atbd-top-dropdown__content'>
                      <img src='~/assets/images/NoPath.png' alt='' />
                      <figcaption>
                        <h4> Software <span class='color-success'>3 hrs ago</span></h4>
                        <div>
                          <span class='atbd-top-dropdownText'>Lorem ipsum dolor amet cosec...</span>
                          <span>
                            <a-badge class='badge-success' count='3' />
                          </span>
                        </div>
                      </figcaption>
                    </figure>
                  </nuxt-link>
                </li>
                <li>
                  <nuxt-link to='#'>
                    <figure class='atbd-top-dropdown__content'>
                      <img src='~/assets/images/NoPath.png' alt='' />
                      <figcaption>
                        <h4> Software <span class='color-success'>3 hrs ago</span></h4>
                        <div>
                          <span class='atbd-top-dropdownText'>Lorem ipsum dolor amet cosec...</span>
                          <span>
                            <a-badge class='badge-success' count='3' />
                          </span>
                        </div>
                      </figcaption>
                    </figure>
                  </nuxt-link>
                </li>
                <li>
                  <nuxt-link to='#'>
                    <figure class='atbd-top-dropdown__content'>
                      <img src='~/assets/images/NoPath.png' alt='' />
                      <figcaption>
                        <h4> Software <span class='color-success'>3 hrs ago</span></h4>
                        <div>
                          <span class='atbd-top-dropdownText'>Lorem ipsum dolor amet cosec...</span>
                          <span>
                            <a-badge class='badge-success' count='3' />
                          </span>
                        </div>
                      </figcaption>
                    </figure>
                  </nuxt-link>
                </li>
                <li>
                  <nuxt-link to='#'>
                    <figure class='atbd-top-dropdown__content'>
                      <img src='~/assets/images/NoPath.png' alt='' />
                      <figcaption>
                        <h4> Software <span class='color-success'>3 hrs ago</span></h4>
                        <div>
                          <span class='atbd-top-dropdownText'>Lorem ipsum dolor amet cosec...</span>
                          <span>
                            <a-badge class='badge-success' count='3' />
                          </span>
                        </div>
                      </figcaption>
                    </figure>
                  </nuxt-link>
                </li>
              </ul>
            </perfect-scrollbar>
          </div>
          <nuxt-link class='btn-seeAll' to='#'>
            Xem tất cả
          </nuxt-link>
        </div>
      </template>
    </a-popover>
  </div>
</template>

<script>
import { PerfectScrollbar } from 'vue3-perfect-scrollbar'
import IdistFeatherIcons from "~/components/commons/IdistFeatherIcons.vue";

export default {
  name: 'header-message',
  components: {IdistFeatherIcons, PerfectScrollbar },
  data: () => ({}),
  methods: {},
  created() {
  },
  mounted() {
  }
}
</script>

<style scoped>

</style>
