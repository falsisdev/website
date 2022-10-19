<template data-theme="dracula">
<main data-theme="dracula">
    <body
      class="mt-8 max-w-screen-lg py-1 mx-auto transition-all duration-300 scroll-auto"
    >
        <div 
        class="mt-5"
        >
        <h1 class="bold text-2xl text-center"> My Anime List </h1>
            <div 
            class="flex flex-wrap rounded-lg"
            >
            <span v-for="item of images" v-bind:key="item">
            <a :href="`https://animecix.net/search?query=${names[images.indexOf(item)]}`">
            <img
            :src="item"
            class="hover:animate-pulse p-[20px] w-[250px] h-[350px] rounded-lg"
            >
            </a>
            </span>
            </div>
        </div>
    </body>
</main>
</template> 
<script>
const animeNames = ['youkoso-jitsuryoku-shijou-shugi-no-kyoushitsu-e', 'codegeasshangyakunolelouch', 'monster', 'evangelion', 'mushoku-tensei-isekai-ittara-honki-dasu', 're-zero-kara-hajimeru-isekai-seikatsu', 'fullmetal-alchemist-brotherhood', 'tensei-shitara-slime-datta-ken', 'sword-art-online', 'mirai-nikki', 'kimetsu-no-yaiba', 'tokyo-ghoul', 'orange', 'charlotte', 'akamegakill', 'bokudakegainaimachi', 'owarinoseraph', 'bungoustraydogs', 'death-note', 'vinland-saga', 'platinum-end', 'domestic-na-kanojo', 'shingeki-no-kyojin', 'tokyo-revengers', 'seishun-buta-yarou-wa-bunny-girl-senpai-no-yume-wo-minai', 'ansatsu-kyoushitsu', 'shigatsu-wo-kimi-no-uso', 'lycorisrecoil', 'jujutsu-kaisen', 'kaifuku-jutsushi-no-yarinaoshi', 'kaguya-sama-wa-kokurasetai', 'spyxfamily', 'no-game-no-life', 'dororo', 'yakusoku-no-neverland', 'dr-stone', 'parasyte', 'ousama-game', 'another', 'devilman-crybaby', 'high-school-dxd', 'deadman-wonderland', 'yofukashi-no-uta', 'noragami', 'ao-no-exorcist', 'koroshi-ai', 'golden-time', 'violet-evergarden', 'darling-in-the-franxx', 'kiznaiver', 'shiki', 'boku-no-hero-academia', 'nanatsu-no-taizai', 'sono-bisque-doll-wa-koi-wo-suru', '5-toubun-no-hanayome', 'shuumatsu-no-harem', 'fairy-tail', 'shinmai-maou-no-tastement', 'wonder-egg-priority', 'tantei-wa-mou-shindeiru', 'noblesse', 'ijiranaide-nagatoro-san', 'kanojo-okarishimasu', '91days', 'goblin-slayer', 'banana-fish', 'kobayashi-san-chi-no-maid-dragon', 'kiss-x-sis', 'takt-op-destiny', 'code-geass-fukkatsu-no-lelouch', 'kimi-no-na-wa', 'tenkinoko', 'kimetsu-no-yaiba-mugen-ressha-hen', 'code-geass-boukoku-no-akito', 'tensuranikki']
export default {
    head() {
        return{
        title: `${this.title} - FalsisDev`
        }
    },
    data() {
        return {
            title: 'Anime List',
            images: [],
            names: []
        }
    },
    async fetch() {
        this.names = animeNames
        for(let item of animeNames) {
            if(item == "code-geass-boukoku-no-akito") {
                this.images.push(`https://cdn.falsis.ga/anime/images?name=code-geass-boukoku-no-akito&type=cover&order=3`)
            }else{
            this.images.push(`https://cdn.falsis.ga/anime/images?name=${item}&type=cover`)
            }
        }
    }
}
</script>
<style scoped>
.bold {
font-family: "Arial Black", "Arial Bold", Gadget, sans-serif;
}
.top-left {
  position: absolute;
  top: 8px;
  left: 16px;
}   
</style>
