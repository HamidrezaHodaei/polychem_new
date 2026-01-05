import { useRoute } from 'vue-router'

export const useLoadingState = () => {
  const route = useRoute()
  // اگر صفحه landing است، loading را false کن
  const initialState = route.path === '/landing' ? false : true
  return useState<boolean>('appLoading', () => initialState)
}
