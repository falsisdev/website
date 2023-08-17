<template>
    <article class="prose mx-3 max-w-none">
        <h1 class="text-white">Manga Listesi</h1>
        <p class="text-white">
            Aşağıda <a class="text-white" href="https://myanimelist.net">MyAnimeList</a> platformundan alınan verilerle benim manga listem yer alıyor. Bu manga listesi <b>bitirdiğim, okuyor olduğum, ara verdiğim, bıraktığım veya okumayı planladığım bütün serileri içermektedir</b>.<br/>
            Ayrıca liste <b>okumaya devam ettiğim mangalar - bitirdiğim mangalar - beklemede olan mangalar - bıraktığım mangalar - okumayı planladığım mangalar</b> şeklinde <b>alfabetik</b> sıralanmıştır. Manga isminin yanındaki rakam <b>benim seriye 10 üzerinden verdiğim puandır</b>. Puanlamadığım seriler "0" olarak görünür. Puanladığım serilerin hiçbiri "0" değildir.<br/>
            Not: Butonlar yeni sekmeye yönlendirir.
        </p>
    </article>
    <SerieCard v-for="item of mangas" :key="item" :name="String(item.mangaTitle)" :cover="item.mangaImagePath" :description="'Bu seri bir ' + item.mangaMediaTypeString + '\'dır.'" :puan="item.score" :genres="item.genres" :url="item.mangaUrl"/>
</template>
<script setup>
    const route = useRoute(); //request info
    const rc = useRuntimeConfig(); //nuxt.config.ts runtimeConfig info

    const listdata = await useFetch(`${rc.public.apiurl}/mangalist`)
    const mangas = listdata.data.value.data
    useHead({
        title: 'Manga List'
    })
</script>