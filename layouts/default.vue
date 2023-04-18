<template>
    <NuxtLoadingIndicator 
    color="#ff79c6"
    :duration="2000"
    />
    <div data-theme="dracula" class="grid grid-rows-1 grid-cols-9 justify-center bg-base-200">
        <div data-theme="dracula" class="col-span-8 col-start-3 col-end-8">
            <Header/>
                <slot/>
            <Footer/>
        </div>
    </div>
</template>
<script setup>
    const route = useRoute(); //request info
    const rc = useRuntimeConfig(); //nuxt.config.ts runtimeConfig info

    useHead({
        titleTemplate: (titleChunk) => {
            return titleChunk ? `${titleChunk} | Falsis` : 'Falsis'
        }
    })//Sayfa Başlığının Template'i
</script>