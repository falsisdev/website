export default defineEventHandler(async (event) => {
  const rc = useRuntimeConfig(); //nuxt.config.ts runtimeConfig info
  if (event.req.method === 'GET') {
    return "OK"
  }
});