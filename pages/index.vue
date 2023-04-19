<template>
    <AboutCard />
    <div class="mx-8 flex flex-row flex-wrap">
        <GithubCard v-for="item of fetched" :key="item" :reponame="item.full_name" :repodesc="item.description" :lang="item.language" :star="item.stargazers_count" :url="item.html_url" />
    </div>
</template>
<script setup>
    const rc = useRuntimeConfig();
    const fullBaseURL = `${rc.public.apis.baseURLs.github}/${rc.public.apis.endpoints.github.userinfo}/${rc.public.apis.userinfo.githubname}/${rc.public.apis.endpoints.github.repos}`
    let fetched = await $fetch(fullBaseURL)

    useHead({
        title: 'Anasayfa'
    })
</script>