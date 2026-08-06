import { describe, it, expect } from 'vitest'

import { mount } from '@vue/test-utils'
import { createRouter, createMemoryHistory } from 'vue-router'
import { createPinia } from 'pinia'
import ui from '@nuxt/ui/vue-plugin'

import App from '../App.vue'
import HomePage from '../pages/HomePage.vue'

const router = createRouter({
  history: createMemoryHistory(),
  routes: [{ path: '/', component: HomePage }],
})

describe('App', () => {
  it('renders the layout and the home page', async () => {
    router.push('/')
    await router.isReady()

    const wrapper = mount(App, {
      global: { plugins: [router, createPinia(), ui] },
    })

    expect(wrapper.text()).toContain('App')

    wrapper.unmount()
  })
})
