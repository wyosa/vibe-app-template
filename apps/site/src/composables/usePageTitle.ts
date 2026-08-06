import { watchEffect, type MaybeRefOrGetter, toValue } from 'vue'

export function usePageTitle(title: MaybeRefOrGetter<string>) {
  watchEffect(() => {
    document.title = toValue(title)
  })
}
