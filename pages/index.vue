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
  <div class="flex pb-[25px] pt-[25px] bg-[#080808] break-word">
  <b class="pl-[35px] text-white-500 text-1x3 pr-[690px]">LISTENING TO SPOTIFY</b>
  <svg class="pr-[10px]" xmlns="http://www.w3.org/2000/svg" height="30px" width="30px" version="1.1" viewBox="0 0 168 168">
 <path fill="#1ED760" d="m83.996 0.277c-46.249 0-83.743 37.493-83.743 83.742 0 46.251 37.494 83.741 83.743 83.741 46.254 0 83.744-37.49 83.744-83.741 0-46.246-37.49-83.738-83.745-83.738l0.001-0.004zm38.404 120.78c-1.5 2.46-4.72 3.24-7.18 1.73-19.662-12.01-44.414-14.73-73.564-8.07-2.809 0.64-5.609-1.12-6.249-3.93-0.643-2.81 1.11-5.61 3.926-6.25 31.9-7.291 59.263-4.15 81.337 9.34 2.46 1.51 3.24 4.72 1.73 7.18zm10.25-22.805c-1.89 3.075-5.91 4.045-8.98 2.155-22.51-13.839-56.823-17.846-83.448-9.764-3.453 1.043-7.1-0.903-8.148-4.35-1.04-3.453 0.907-7.093 4.354-8.143 30.413-9.228 68.222-4.758 94.072 11.127 3.07 1.89 4.04 5.91 2.15 8.976v-0.001zm0.88-23.744c-26.99-16.031-71.52-17.505-97.289-9.684-4.138 1.255-8.514-1.081-9.768-5.219-1.254-4.14 1.08-8.513 5.221-9.771 29.581-8.98 78.756-7.245 109.83 11.202 3.73 2.209 4.95 7.016 2.74 10.733-2.2 3.722-7.02 4.949-10.73 2.739z"/>
</svg>
  </div>
   <div class="flex pb-[25px] bg-[#080808] break-word rounded-b-[10px]">
      <img :Src="lanyardActivities.spotify.cover" class="ml-8 rounded-md border-[3.2px] w-[8rem] h-[8rem] border-green-500">
      <div class="ml-4">
        <span class="absolute pl-[10px] text-green-500 text-1x2">{{ lanyardActivities.spotify.song }}</span>
        <span class="absolute pt-[20px] pl-[10px] text-white-500 text-1xl">by {{ lanyardActivities.spotify.artist }}</span>
        <span class="absolute pt-[40px] pl-[10px] text-white-500 text-1xl">on {{ lanyardActivities.spotify.album.large_text }}</span>
        <span class="absolute pt-[60px] pl-[10px] text-white-500 text-1xl">{{ lanyardActivities.spotify.elapsed[0] }} / {{ lanyardActivities.spotify.elapsed[1] }}</span>
      </div>
   </div>
 </div>
  <p v-if="lanyardActivities.activity.check == false"></p>
   <p v-else-if="lanyardActivities.activity.name == 'Spotify'"></p>
   <i v-else-if="lanyardActivities.activity.check == true" v-if="lanyardActivities.activity.image == null"></i>
  <div v-else class="mt-5">
  <div class="flex pb-[25px] pt-[25px] bg-[#080808] break-word">
  <b class="pl-[35px] text-white-500 text-1x3 pr-[690px]">PLAYING A GAME</b>
      <i v-if="lanyardActivities.activity.small == null"></i><img class="w-[2rem] h-[2rem] rounded-full" v-else-if="lanyardActivities.activity['small'].substring(0, 3) == 'mp:' == true" :src="`https://${lanyardActivities.activity['small'].slice(62)}`"><img v-else class="w-[2rem] h-[2rem] rounded-full" :Src="lanyardActivities.activity.small">   
        </div>
   <div class="flex pb-[25px] bg-[#080808] break-word rounded-b-[10px]">
          <i v-if="lanyardActivities.activity.image == null"></i><img v-else-if="lanyardActivities.activity['image'].charAt(0) + lanyardActivities.activity['image'].charAt(1) + lanyardActivities.activity['image'].charAt(2) == 'mp:' == true" :src="`https://${lanyardActivities.activity['image'].slice(62)}`" class="ml-8 rounded-md w-[8rem] h-[8rem]"><img v-else :src="`https://cdn.discordapp.com/app-assets/${lanyardActivities.activity.appid}/${lanyardActivities.activity['image']}.png`" class="ml-8 rounded-md w-[8rem] h-[8rem]">
      <div class="ml-4">
        <i v-if="lanyardActivities.activity.name == null"></i><span v-else class="absolute pl-[10px] text-red-500 text-1x2">{{ lanyardActivities.activity.name || "" }}</span>
        <i v-if="lanyardActivities.activity.state == null"></i><span v-else class="absolute pt-[20px] pl-[10px] text-white-500 text-1xl">{{ lanyardActivities.activity.state || "" }}</span>
        <i v-if="lanyardActivities.activity.details == null"></i><span v-else class="absolute pt-[40px] pl-[10px] text-white-500 text-1xl">{{ lanyardActivities.activity.details || "" }}</span>
        <i v-if="String(String(lanyardActivities.activity['elapsed']).split(':').join('')) == String('NaNNaN')"></i><span v-else class="absolute pt-[60px] pl-[10px] text-white-500 text-1xl">{{ lanyardActivities.activity.elapsed == null ? "" : `Elapsed: ${String(parseInt(lanyardActivities.activity.elapsed.substring(0, 2)) < 10 ? 0 + lanyardActivities.activity.elapsed : lanyardActivities.activity.elapsed)}`  }} </span>
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
        <a :href="`${item.url}`" :class="`p-2 rounded-lg pb-2 px-5 text-white bg-[${item.color}]`"><i :class="`${item.icon} mr-2`"></i> {{ item.name }}</a>
     </div>
     </div>
  </div>
</div>
<div class="mt-5">
  <div class="bg-[#080808] w-full p-4 rounded-lg">
     <div class="mt-2">
       <h1 class="font-semibold text-3xl text-white text-center"><span class="text-red-400">GitHub</span> Repositories</h1>
          <p v-if="github.repocount == 0" class="mt-3 text-center text-red-500">No results found!</p>
          <div v-else class="grid grid-cols-3 md:grid-cols-3 gap-3 mt-6">
            <div v-for="item in github.repos">
            <span v-if="item.fork == true"><a :href="`https://github.com/falsisdev/${item.name}`">
            <div class="bg-[#080707] rounded-lg ml-5 p-3">
              <h1 class="text-white text-[18px]"><i class="fa-brands fa-github text-white mr-1"></i>  {{ item.name }}</h1>
              <div class="flex gap-x-5">
                <div class="mt-3 text-white"><i class="fa-solid fa-code-branch text-white mr-1"></i> forked from {{ item.name }}</div> <!-- çaktırma ve sessizce uzaklaş -->
              </div>
             </div>
           </a></span>
            <span v-else>
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
           </span>
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
           url: "https://github.com/falsisdev"
        },{
           name: "Spotify",
           icon: "fa-brands fa-spotify",
           color: "#1db954",
           url: "https://open.spotify.com/user/315l5ib3a4fd2obidm76lipspxji?si=1ea7f02ae26d40a1"
        },{
          name: "Instagram",
          icon: "fa-brands fa-instagram",
          color: "#E1306C",
          url: "https://instagram.com/falsisdev"
        },{
          name: "Discord",
          icon: "fa-brands fa-discord",
          color: "#5865F2",
          url: "https://discord.com/users/539843855567028227"
        }, {
          name: "Osu!",
          icon: `fa-solid fa-play`,
          color: "#FF66AA",
          url: "https://osu.ppy.sh/users/25156234"
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
            elapsed: [],
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
            elapsed: [],
            state: [],
            name: [],
            details: [],
            appid: [],
            image: [],
            small: [],
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
      this.lanyardUser.avatar = `https://cdn.discordapp.com/avatars/${this.lanyard.data.discord_user.id}/${this.lanyard.data.discord_user.avatar}${this.lanyard.data.discord_user.avatar.substring(0, 2) == 'a_' ? '.gif' : '.webp'}`
      this.lanyardUser.status = this.lanyard.data.discord_status

      /* SPOTIFY */
      this.lanyardActivities.spotify.check = this.lanyard.data.listening_to_spotify //boolean
      if(this.lanyard.data.listening_to_spotify == false){
      this.lanyardActivities.spotify.trackid = null
      this.lanyardActivities.spotify.timestamps = null
      this.lanyardActivities.spotify.song = null
      this.lanyardActivities.spotify.cover = null
      }else{
        function millisToMinutesAndSeconds(millis) {
             var minutes = Math.floor(millis / 60000);
             var seconds = ((millis % 60000) / 1000).toFixed(0);
             if(minutes >= 60){
              return Math.floor(minutes / 60) + ":" +  Number(minutes - (Math.floor(minutes / 60) * 60)) + ":" + (seconds < 10 ? '0' : '') + seconds;
             }
            return minutes + ":" + (seconds < 10 ? '0' : '') + seconds;
        }
      this.lanyardActivities.spotify.elapsed = [millisToMinutesAndSeconds(Date.now() - this.lanyard.data.spotify.timestamps.start), millisToMinutesAndSeconds(this.lanyard.data.spotify.timestamps.end - this.lanyard.data.spotify.timestamps.start)]
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
      this.lanyardActivities.activity.elapsed = null
      this.lanyardActivities.activity.state = null
      this.lanyardActivities.activity.name = null
      this.lanyardActivities.activity.id = null
      this.lanyardActivities.activity.details = null
      this.lanyardActivities.activity.created_at = null
      this.lanyardActivities.activity.appid = null
      this.lanyardActivities.activity['image'] = null
      this.lanyardActivities.activity['small'] = null
      }else {
        try {
                   function millisToMinutesAndSeconds(millis) {
             var minutes = Math.floor(millis / 60000);
             var seconds = ((millis % 60000) / 1000).toFixed(0);
             if(minutes >= 60){
              return Math.floor(minutes / 60) + ":" +  Number(minutes - (Math.floor(minutes / 60) * 60)) + ":" + (seconds < 10 ? '0' : '') + seconds;
             }
            return minutes + ":" + (seconds < 10 ? '0' : '') + seconds;
      }
      this.lanyardActivities.activity.timestamps = this.lanyard.data.activities[0].timestamps.start
      this.lanyardActivities.activity.elapsed = millisToMinutesAndSeconds(Date.now() - this.lanyard.data.activities[0].timestamps.start)
            this.lanyardActivities.activity['image'] = this.lanyard.data.activities[0].assets.large_image
            this.lanyardActivities.activity['small'] = this.lanyard.data.activities[0].assets.small_image || null
                  this.lanyardActivities.activity.state = this.lanyard.data.activities[0].state
      this.lanyardActivities.activity.name = this.lanyard.data.activities[0].name
      this.lanyardActivities.activity.details = this.lanyard.data.activities[0].details
      this.lanyardActivities.activity.appid = this.lanyard.data.activities[0].application_id
        }catch(err){
          this.lanyardActivities.activity.timestamps = null
          this.lanyardActivities.activity.elapsed = null
                this.lanyardActivities.activity['image'] = null
                this.lanyardActivities.activity['small'] = null
                      this.lanyardActivities.activity.state = null
      this.lanyardActivities.activity.name = null
      this.lanyardActivities.activity.details = null
      this.lanyardActivities.activity.appid = null
        }
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
