<template>
    <NuxtLoadingIndicator 
    color="#ff79c6"
    :duration="2000"
    />
    <span v-if="isMobileOrTablet == false">
        <div data-theme="dracula" class="grid grid-rows-1 grid-cols-9 justify-center bg-base-200">
            <div data-theme="dracula" class="col-span-8 col-start-3 col-end-8">
                <Header/>
                    <slot/>
                <Footer/>
            </div>
        </div>
    </span>
    <span v-else-if="isMobileOrTablet == true">
        <div data-theme="dracula" class="grid grid-rows-1 grid-cols-9 justify-center bg-base-200">
            <div data-theme="dracula" class="col-span-9 col-start-0 col-end-10">
                <Header/>
                    <slot/>
                <Footer/>
            </div>
        </div>
    </span>
</template>
<script setup>
    const { isMobileOrTablet } = useDevice();

    useHead({
        titleTemplate: (titleChunk) => {
            return titleChunk ? `${titleChunk} | Falsis` : 'Falsis'
        }
    })//Sayfa Başlığının Template'i
</script>