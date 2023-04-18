<template>
    <article class="prose mx-3 max-w-none">
        <h1 class="text-white">Anime Listesi</h1>
        <p class="text-white">
            Aşağıda <a class="text-white" href="https://animecix.net">AnimeciX</a> platformundan alınan verilerle benim anime listem yer alıyor. Bu anime listesi <b>yalnızca dizi şeklindeki animeleri içermektedir</b>. Filmler için animecix hesabımda ayrı bir listem bulunmaktadır.<br/>
            Ayrıca liste benim animeleri <i>sevme sırama</i> göre sıralanmıştır. 1. sıradaki seri en çok sevdiğim, sonuncu sıradaki seri ise en nefret ettiğim seridir.<br/>
            Not: Butonlar yeni sekmeye yönlendirir.
        </p>
    </article>
    <SerieCard v-for="item of animes" :key="item" :name="item.name" :cover="item.poster" :description="item.description" :year="item.year" :genres="item.genres" :id="item.id"/>
</template>
<script setup>
    const route = useRoute(); //request info
    const rc = useRuntimeConfig(); //nuxt.config.ts runtimeConfig info

    const listdata = await useFetch(`https://animecix.net/secure/lists/110787`)
    const animes = listdata.data.value.items.data
    useHead({
        title: 'Anime List'
    })
</script>