export default defineEventHandler(async (event) => {
    const rc = useRuntimeConfig(); //nuxt.config.ts runtimeConfig info
    if (event.req.method === 'GET') {
      const fullBaseURL = `${rc.public.apis.baseURLs.github}/${rc.public.apis.endpoints.github.userinfo}/${rc.public.apis.userinfo.githubname}/${rc.public.apis.endpoints.github.repos}`
      let fetched = await $fetch(fullBaseURL)
      return fetched
    }
});