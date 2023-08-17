<template>
    <article class="prose mx-3 max-w-none">
        <h1 class="text-white">Anime Listesi</h1>
        <p class="text-white">
            Aşağıda <a class="text-white" href="https://myanimelist.net">MyAnimeList</a> platformundan alınan verilerle benim anime listem yer alıyor. Bu anime listesi <b>bitirdiğim, izliyor olduğum, ara verdiğim, bıraktığım veya izlemeyi planladığım bütün animeleri içermektedir</b>.<br/>
            Ayrıca liste <b>izlemeye devam ettiğim animeler - bitirdiğim animeler - beklemede olan animeler - bıraktığım animeler - izlemeyi planladığım animeler</b> şeklinde <b>alfabetik</b> sıralanmıştır. Anime isminin yanındaki rakam <b>benim animeye 10 üzerinden verdiğim puandır</b>. Puanlamadığım animeler "0" olarak görünür. Puanladığım animelerin hiçbiri "0" değildir.<br/>
            Not: Butonlar yeni sekmeye yönlendirir.
        </p>
    </article>
    <SerieCard v-for="item of animes" :key="item" :name="String(item.animeTitle)" :cover="item.animeImagePath" :description="'Bu seri bir ' + item.animeMediaTypeString + '\'dir.'" :puan="item.score" :genres="item.genres" :url="item.animeUrl"/>
</template>
<script setup>
    const route = useRoute(); //request info
    const rc = useRuntimeConfig(); //nuxt.config.ts runtimeConfig info

    const listdata = await useFetch(`${rc.public.apiurl}/animelist`)
    const animes = listdata.data.value.data
    useHead({
        title: 'Anime List'
    })
</script>