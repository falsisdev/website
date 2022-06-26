<template>
<main>
    <body class="mt-8 max-w-screen-lg p-5 py-1 mx-auto transition-all duration-300">
      <div align="center" :class="`rounded-t-[10px] w-full pb-[0.25rem] bg-[#${lanyardUser.status === 'dnd' ? 'ED4245' :  lanyardUser.status === 'idle' ? 'faa61a' : lanyardUser.status === 'online' ? '57F287' : lanyardUser.status === 'offline' ? '99aab5' : 'EB459E'}]`"> </div>
<div class="flex pb-[25px] pt-[25px] bg-[#080808] break-word rounded-b-[10px]">
    <img :class="`ml-8 rounded-full border-[3.2px] w-[8rem] h-[8rem] border-[#${lanyardUser.status === 'dnd' ? 'ED4245' :  lanyardUser.status === 'idle' ? 'faa61a' : lanyardUser.status === 'online' ? '57F287' : lanyardUser.status === 'offline' ? '99aab5' : 'EB459E'}]`" :src="lanyardUser.avatar"/>
    <div class="ml-4 mt-8">
      <h1 class="absolute pl-[10px] text-red-500 text-4xl">{{ lanyardUser.username }}<span class="text-[15px] font-bold text-gray-200">#{{ lanyardUser.tag }}</span></h1>
      
      <p class="text-[19px] text-gray-200/90 pt-10 ml-2">An Alone Web Developer</p>
      
    </div>
 </div>
 <p v-if="lanyardActivities.spotify.check == false"></p>
  <div v-else-if="lanyardActivities.spotify.check == true" class="mt-5">
   <div class="flex pb-[25px] pt-[25px] bg-[#080808] break-word rounded-b-[10px]">
      <img :Src="lanyardActivities.spotify.cover" class="ml-8 rounded-md border-[3.2px] w-[8rem] h-[8rem] border-green-500">
      <div class="ml-4">
        <span class="absolute pl-[10px] text-green-500 text-1x2">{{ lanyardActivities.spotify.song }}</span>
        <span class="absolute pt-[20px] pl-[10px] text-white-500 text-1xl">by {{ lanyardActivities.spotify.artist }}</span>
        <span class="absolute pt-[40px] pl-[10px] text-white-500 text-1xl">on {{ lanyardActivities.spotify.album }}</span>
      </div>
   </div>
 </div>
  <p v-if="lanyardActivities.activity.check == false"></p>
   <p v-else-if="lanyardActivities.activity.name == 'Spotify'"></p>
  <div v-else-if="lanyardActivities.activity.check == true" class="mt-5">
   <div class="flex pb-[25px] pt-[25px] bg-[#080808] break-word rounded-b-[10px]">
      <img :src="`https://cdn.discordapp.com/app-assets/${lanyardActivities.activity.appid}/${lanyardActivities.activity.image}.png`" class="ml-8 rounded-md w-[8rem] h-[8rem]">
      <div class="ml-4">
        <span class="absolute pl-[10px] text-red-500 text-1x2">{{ lanyardActivities.activity.name ||"" }}</span>
        <span class="absolute pt-[20px] pl-[10px] text-white-500 text-1xl">{{ lanyardActivities.activity.state || "" }}</span>
        <span class="absolute pt-[40px] pl-[10px] text-white-500 text-1xl">{{ lanyardActivities.activity.details || "" }}</span>
      </div>
   </div>
 </div>
 <div class="mt-5">
   <div class="bg-[#080808] w-full p-4 text-[18px] text-center rounded-lg">
      <h1><span class="text-red-500"><i class="fa-solid fa-triangle-exclamation"></i></span><span style="color: white;"> This website is currently in </span><span class="text-green-300">early-access</span>. <span style="color: white;">Please report bugs if any.</span></h1>
   </div>
 </div>
  <div class="mt-5">
  <div class="bg-[#080808] w-full p-4 rounded-lg">
     <div class="mt-2">
       <h1 class="font-semibold text-3xl text-white text-center"><span class="text-red-400">Hello</span>, I’m Falsis ✌️</h1>
       <p style="color: white;" class="p-3 text-gray-200/90 text-[18px] mt-3">
I'm a <a href="https://en.wikipedia.org/wiki/Front-end_web_development" class="Text-red-500"> Front-end Web Developer</a>!<br>
So, I'm interested in <a href="https://en.wikipedia.org/wiki/HTML" class="Text-red-500" title="Hyper Text Markup Language"> <font color="orange"><i class="fa-brands fa-html5"></i></font></a>, <a href="https://en.wikipedia.org/wiki/CSS" class="Text-red-500" title="Cascading Style Sheets"> <font color="#1589FF"><i class="fa-brands fa-css3-alt"></i></font></a> and <a href="https://en.wikipedia.org/wiki/JavaScript" title="JavaScript" class="Text-red-500"> <font color="yellow"><i class="fab fa-js-square"></i></font></a>!<br>
Also, I know <a href="https://en.wikipedia.org/wiki/Go_(programming_language)" class="Text-red-500" title="The Go Programming Language"><font color="#add8e6"><i class="fa-brands fa-golang"></i></font></a>, <a href="https://en.wikipedia.org/wiki/Python_(programming_language)" class="Text-red-500" title="Python"><font color="#1589FF"><i class="fa-brands fa-python"></i></font></a> and <a href="https://tr.wikipedia.org/wiki/Lua_(programlama_dili)" class="Text-red-500">Lua</a>...<br>
If you wanna see them, visit my <a href="https://github.com/falsisdev" class="Text-red-500"><font color="white"><i class="fa-brands fa-github"></i> </font></a> profile!</p>
     </div>
     <div align="center">
     <div v-for="item in links" class="p-2 mt-3 text-center content-center inline-block">
        <a :href="`/redirect/${item.url}`" :class="`p-2 rounded-lg pb-2 px-5 text-white bg-[${item.color}]`"><i :class="`${item.icon} mr-2`"></i> {{ item.name }}</a>
     </div>
     </div>
  </div>
</div>
<div class="mt-5">
  <div class="bg-[#080808] w-full p-4 rounded-lg">
     <div class="mt-2">
       <h1 class="font-semibold text-3xl text-white text-center"><span class="text-red-400">GitHub</span> Repositories</h1>
          <p v-if="github.repocount == 0" class="mt-3 text-center text-red-500">Hiç bir sonuç bulunamadı!</p>
          <div v-else class="grid grid-cols-3 md:grid-cols-3 gap-3 mt-6">
            <div v-for="item in github.repos">
          <a :href="`https://github.com/falsisdev/${item.name}`">
            <div class="bg-[#080707] rounded-lg ml-5 p-3">
              <h1 class="text-white text-[18px]"><i class="fa-brands fa-github text-white mr-1"></i>  {{ item.name }}</h1>
              <div class="flex gap-x-5">
                <div class="mt-3 text-white"><i class="fas fa-star text-white mr-1"></i>{{ item.stargazers_count }}</div>
                <div class="mt-3 text-white"><i class="fa-solid fa-code-branch text-white mr-1"></i>{{ item.forks }}</div>
                <div class="mt-3 text-white"><i class="fa-solid fa-earth text-white mr-1"></i>{{ item.language }}</div>
              </div>
             </div>
           </a>
        </div>
        </div>
     </div>
  </div>
</div>
    <div class="py-10 text-white border-t border-zinc-600/20 mt-5 ">
  <p class="float-left">© 2021 - 2022 FalsisDev - All Rights Reserved.</p>
  <p class="text-white float-right">Designed with <i class="fa fa-heart mx-1 text-pink-400"></i> by <a class="hover:text-green-400 font-medium" href="/">Falsis</a></p>
</div></body>
    </main>
</template>
  <script>
  let lanyardURL = "https://api.lanyard.rest/v1/users/539843855567028227";
  let githubURL = "https://api.github.com/users/falsisdev"
   export default {
        head() {
        return {
        title: this.title,
        link: [
          {
             rel: 'icon', type: 'image/x-icon', href: this.lanyardUser.avatar
          },
          {
            rel: 'stylesheet', href: 'https://cdnjs.cloudflare.com/ajax/libs/font-awesome/4.7.0/css/font-awesome.min.css'
          }
          ],
        script: [
          {
            src: "https://cdn.tailwindcss.com",
          },
          {
            src: "https://kit.fontawesome.com/e0330794a8.js",
            crossorigin: "anonymous"
          }
        ]
      }
      },
     data() {
      return {
        title: "Home - FalsisDev",
        links: [{
           name: "Github",
           icon: "fa-brands fa-github",
           color: "#333",
           url: "github"
        },{
           name: "Spotify",
           icon: "fa-brands fa-spotify",
           color: "#1db954",
           url: "spotify"
        }],
        lanyard: [],
        githubf: [],
        lanyardUser: {
          username: [],
          tag: [],
          id: [],
          avatar: [],
          status: []
        },
        lanyardActivities: {
          spotify: {
            check: [],
            trackid: [],
            timestamps: [],
            song: [],
            artist: [],
            album: [],
            cover: []
          },
          activity: {
            check: [],
            timestamps: [],
            state: [],
            name: [],
            details: [],
            appid: [],
            image: []
          },
        },
        device: [],
        github: {
          username: [],
          name: [],
          avatar: [],
          url: [],
          company: [],
          blog: [],
          location: [],
          bio: [],
          twitter: [],
          repocount: [],
          gistcount: [],
          followers: [],
          following: [],
          repos: [] //aslında array içinde object ama burada object açmıyorum aşağıda fetch'leyince object oluşturulacak
        }
      }
    },
    async fetch() {
      try {
      this.lanyard = await fetch(lanyardURL).then(res => res.json())
      if(this.lanyard.success == false) console.log("Lanyard Çekme Başarısız")
      }catch(err) {
        console.log(err)
      }
      try {
        this.githubf = await fetch(githubURL).then(res => res.json())
      }catch(err) {
        console.log(err)
      }

      /* DISCORD USER */
      this.lanyardUser.username = this.lanyard.data.discord_user.username
      this.lanyardUser.tag = this.lanyard.data.discord_user.discriminator
      this.lanyardUser.id = this.lanyard.data.discord_user.id
      this.lanyardUser.avatar = `https://cdn.discordapp.com/avatars/${this.lanyard.data.discord_user.id}/${this.lanyard.data.discord_user.avatar}.webp`
      this.lanyardUser.status = this.lanyard.data.discord_status

      /* SPOTIFY */
      this.lanyardActivities.spotify.check = this.lanyard.data.listening_to_spotify //boolean
      if(this.lanyard.data.listening_to_spotify == false){
      this.lanyardActivities.spotify.trackid = null
      this.lanyardActivities.spotify.timestamps = null
      this.lanyardActivities.spotify.song = null
      this.lanyardActivities.spotify.cover = null
      }else{
      this.lanyardActivities.spotify.trackid = this.lanyard.data.spotify.track_id
      this.lanyardActivities.spotify.timestamps = [this.lanyard.data.spotify.timestamps.start, this.lanyard.data.spotify.timestamps.end]
      this.lanyardActivities.spotify.song = this.lanyard.data.spotify.song
      this.lanyardActivities.spotify.artist = this.lanyard.data.spotify.artist
      this.lanyardActivities.spotify.album = this.lanyard.data.spotify.album
      this.lanyardActivities.spotify.cover = this.lanyard.data.spotify.album_art_url
      }
      this.device = this.lanyard.data.active_on_discord_web ? "Web" : this.lanyard.data.active_on_discord_desktop ? "Desktop" : this.lanyard.data.active_on_discord_mobile ? "Mobile" : "Offline"
      /* ACTIVITY */
      this.lanyardActivities.activity.check = String(this.lanyard.data.activities) == String([]) ? false : true
      if(this.lanyardActivities.activity.check == false) {
      this.lanyardActivities.activity.timestamps = null
      this.lanyardActivities.activity.state = null
      this.lanyardActivities.activity.name = null
      this.lanyardActivities.activity.id = null
      this.lanyardActivities.activity.details = null
      this.lanyardActivities.activity.created_at = null
      this.lanyardActivities.activity.appid = null
      this.lanyardActivities.activity.image = null
      }else {
        try {
      this.lanyardActivities.activity.timestamps = this.lanyard.data.activities[0].timestamps.start
        }catch(err){
          this.lanyardActivities.activity.timestamps = null
        }
      this.lanyardActivities.activity.state = this.lanyard.data.activities[0].state
      this.lanyardActivities.activity.name = this.lanyard.data.activities[0].name
      this.lanyardActivities.activity.details = this.lanyard.data.activities[0].details
      this.lanyardActivities.activity.appid = this.lanyard.data.activities[0].application_id
      this.lanyardActivities.activity.image = this.lanyard.data.activities[0].assets.large_image
      }
      /* GITHUB */
      this.github.username = this.githubf.login
      this.github.name = this.githubf.name
      this.github.avatar = this.githubf.avatar_url
      this.github.url = this.githubf.html_url
      this.github.company = this.githubf.company
      this.github.blog = this.githubf.blog
      this.github.location = this.githubf.location
      this.github.bio = this.githubf.bio
      this.github.twitter = this.githubf.twitter_username
      this.github.repocount = this.githubf.public_repos
      this.github.gistcount = this.githubf.public_gists
      this.github.followers = this.githubf.followers
      this.github.following = this.githubf.following
      this.github.repos = await fetch(`${githubURL}/repos`).then(res => res.json()) //burda array içinde objectler oluşturuldu.
    }
      }
  </script>