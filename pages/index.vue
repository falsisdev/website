<template>
    <AboutCard />
    <div class="mx-8 flex flex-row flex-wrap">
        <GithubCard v-for="item of fetched" :key="item" :reponame="item.full_name" :repodesc="item.description" :lang="item.language" :star="item.stargazers_count" :url="item.html_url" />
    </div>
</template>
<script setup>
    const fetched = await $fetch('/api')

    useHead({
        title: 'Anasayfa'
    })
</script>