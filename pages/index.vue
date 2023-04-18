<template>
    <AboutCard />
    <div class="mx-8 flex flex-row flex-wrap">
        <GithubCard v-for="item of fetched.data.value" :key="item" :reponame="item.full_name" :repodesc="item.description" :lang="item.language" :star="item.stargazers_count" :url="item.html_url" />
    </div>
</template>
<script setup>
    const route = useRoute(); //request info
    const rc = useRuntimeConfig(); //nuxt.config.ts runtimeConfig info

    const fullBaseURL = `${rc.public.apis.baseURLs.github}/${rc.public.apis.endpoints.github.userinfo}/${rc.public.apis.userinfo.githubname}/${rc.public.apis.endpoints.github.repos}`
    console.log(fullBaseURL)
    const fetched = await useFetch(fullBaseURL)


    useHead({
        title: 'Anasayfa'
    })
</script>