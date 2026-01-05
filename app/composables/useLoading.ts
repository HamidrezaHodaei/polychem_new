export const useLoadingState = () => {
  return useState<boolean>('appLoading', () => true)
}
