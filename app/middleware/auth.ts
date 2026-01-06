export default defineNuxtRouteMiddleware((to, from) => {
  const token = localStorage.getItem('token')
  
  if (!token && to.path === '/Dashboard_125') {
    return navigateTo('/adminpannel')
  }
  
  if (token && to.path === '/adminpannel') {
    return navigateTo('/Dashboard_125')
  }
})