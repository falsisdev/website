<template data-theme="dracula">
  <main data-theme="dracula">
    <body
      class="mt-8 bg-[#2a303c] max-w-screen-lg p-5 py-1 mx-auto transition-all duration-300 scroll-auto"
      data-theme="dracula"
    >
    <div class="flex flex-wrap flex-row align-items-center justify-center">
      <div
        class="card w-72 bg-base-100 shadow-xl"
      >
      <div class="card-body align-items-center justify-center">
      <figure><div :class="`avatar pl-[10px]`">
        <div class="w-24 rounded-full">
        <img
          :src="`${lanyardUser.avatar}?size=4096`"
        />
        </div>
        </div>
        </figure>
          <h1 class="card-title align-items-center justify-center">
            {{ lanyardUser.username
            }}<span class="text-[15px] font-bold text-gray-200"
              >#{{ lanyardUser.tag }}</span
            >
            <span
              id="platform"
              :class="`text-[15px] font-bold text-[#${lanyardUser.status === 'dnd' ? 'ED4245' :  lanyardUser.status === 'idle' ? 'faa61a' : lanyardUser.status === 'online' ? '10b981' : lanyardUser.status === 'offline' ? '99aab5' : '10b981'}]`"
              ><i
                :class="device == 'Desktop' ? 'fa-solid fa-desktop' : device == 'Mobile' ? 'fa-solid fa-mobile' : device == 'Web' ? 'fa-brands fa-firefox-browser' : ''"
              ></i
            ></span>
          </h1>
          <p class="align-items-center justify-center text-center">
            An Alone Web Developer
          </p>
      </div>
      </div>
      <p v-if="lanyardActivities.spotify.check == false"></p>
      <div v-else-if="lanyardActivities.spotify.check == true" class="pl-[5px]">
        <div class="card w-72 bg-base-100 shadow-xl image-full">
          <figure>
          <img
            :src="lanyardActivities.spotify.cover"
          />
          </figure>
          <div class="card-body">
            <h1
              class="card-title"
              >{{ lanyardActivities.spotify.song }}</h1
            >
            <p>
            <span
              >by {{ lanyardActivities.spotify.artist }}</span
            ><br>
            <span
              >on {{ lanyardActivities.spotify.album.large_text }}</span
            ><br>
            <span
              >{{ lanyardActivities.spotify.elapsed[0] }} /
              {{ lanyardActivities.spotify.elapsed[1] }}</span
            >
            </p>
          </div>
        </div>
      </div>
      <span
        v-if="lanyardActivities.spotify.check == false & lanyardActivities.activity.check == true"
      >
        <i v-if="lanyardActivities.activity.image == null"></i>
        <div v-else class="pl-[5px]">
          <div class="card w-72 bg-base-100 shadow-xl image-full">
            <figure><i v-if="lanyardActivities.activity.image == null"></i
            ><img
              v-else-if="lanyardActivities.activity['image'].charAt(0) + lanyardActivities.activity['image'].charAt(1) + lanyardActivities.activity['image'].charAt(2) == 'mp:' == true"
              :src="`https://${lanyardActivities.activity['image'].slice(62)}`"
              width="400"
              height="400"
            /><img
              v-else
              :src="`https://cdn.discordapp.com/app-assets/${lanyardActivities.activity.appid}/${lanyardActivities.activity['image']}.png`"
              width="400"
              height="400"
            /></figure>
          <div class="card-body">
              <i v-if="lanyardActivities.activity.name == null"></i
              ><h1
                v-else
                class="card-title"
                >{{ lanyardActivities.activity.name || "" }}</h1
              >
              <p>
              <i v-if="lanyardActivities.activity.state == null"></i
              ><span
                v-else
                >{{ lanyardActivities.activity.state || "" }}</span
              >
              <i v-if="lanyardActivities.activity.details == null"></i
              ><span
                v-else
                >{{ lanyardActivities.activity.details || "" }}</span
              >
              <i
                v-if="String(String(lanyardActivities.activity['elapsed']).split(':').join('')) == String('NaNNaN')"
              ></i
              ><span
                v-else
                >{{ lanyardActivities.activity.elapsed == null ? "" : `Elapsed: ${lanyardActivities.activity.elapsed}`  }}
              </span>
              </p>
            </div>
          </div>
        </div></span
      ></div><br>
      <div align="center">
            <div
              v-for="item in links" v-bind:key="item"
              class="p-2 mt-3 text-center content-center inline-block"
            >
              <a
                :href="`${item.url}`"
                :class="`duration-300 p-2 rounded-lg pb-2 px-5 text-white bg-[${item.color}]`"
                ><i :class="`${item.icon} mr-2`"></i> {{ item.name }}</a
              >
            </div>
          </div>
      <div class="w-full align-items-center justify-center">
        <div class="align-items-center justify-center p-4 rounded-lg">
          <div class="mt-2">
            <h1 class="font-semibold text-3xl text-white text-center">
              <span class="text-primary">Hello</span>, I’m Falsis ✌️
            </h1>
            <p
              style="color: white"
              class="p-3 text-gray-200/90 text-[18px] mt-3"
            >
              I'm a
              <a
                href="https://en.wikipedia.org/wiki/Front-end_web_development"
                class="duration-300 hover:text-primary"
              >
                Front-end Web Developer</a
              >!<br />
              So, I'm interested in
              <a
                href="https://en.wikipedia.org/wiki/HTML"
                class="Text-primary"
                title="Hyper Text Markup Language"
              >
                <font color="orange"
                  ><i class="fa-brands fa-html5"></i></font></a
              >,
              <a
                href="https://en.wikipedia.org/wiki/CSS"
                class="Text-primary"
                title="Cascading Style Sheets"
              >
                <font color="#1589FF"
                  ><i class="fa-brands fa-css3-alt"></i></font
              ></a>
              and
              <a
                href="https://en.wikipedia.org/wiki/JavaScript"
                title="JavaScript"
                class="Text-primary"
              >
                <font color="yellow"><i class="fab fa-js-square"></i></font></a
              >!<br />
              Also, I know
              <a
                href="https://en.wikipedia.org/wiki/Go_(programming_language)"
                class="text-primary"
                title="The Go Programming Language"
                ><font color="#add8e6"
                  ><i class="fa-brands fa-golang"></i></font></a
              >,
              <a
                href="https://en.wikipedia.org/wiki/TypeScript"
                class="text-primary"
                title="TypeScript"
                >TypeScript</a
              >
              and
              <a
                href="https://tr.wikipedia.org/wiki/Dart_(programlama_dili)"
                class="text-primary"
                >Dart</a
              >...<br />
              If you wanna see them, visit my
              <a href="https://github.com/falsisdev" class="Text-primary"
                ><font color="white"><i class="fa-brands fa-github"></i> </font
              ></a>
              profile!
              <br />
              Don't forget to visit
              <NuxtLink
                to="/list"
                class="animate-pulse text-primary"
                title="My List"
              >
              My Anime List!!
              </NuxtLink>
            </p>
          </div>
        </div>
      </div>
     <div class="mt-5">
        <div class="p-4 rounded-lg">
          <div class="mt-2">
            <h1 class="font-semibold text-3xl text-white text-center">
              <span class="text-primary">GitHub</span> Repositories
            </h1>
            <p
              v-if="github.repocount == 0"
              class="mt-3 text-center text-primary"
            >
              No results found!
            </p>
            <div v-else class="grid grid-cols-3 md:grid-cols-3 gap-3 mt-6">
              <div v-for="item in github.repos" v-bind:key="item">
                <span v-if="item.fork == true"
                  ><a :href="`https://github.com/falsisdev/${item.name}`">
                    <div class="rounded-lg ml-5 p-3">
                      <h1 class="text-white text-[18px]">
                        <i class="fa-brands fa-github text-white mr-1"></i>
                        {{ item.name }}
                      </h1>
                      <div class="flex gap-x-5">
                        <div class="mt-3 text-white">
                          <i
                            class="fa-solid fa-code-branch text-white mr-1"
                          ></i>
                          forked from {{ item.name }}
                        </div>
                      </div>
                    </div>
                  </a></span
                >
                <span v-else>
                  <a :href="`https://github.com/falsisdev/${item.name}`">
                    <div class="rounded-lg ml-5 p-3">
                      <h1 class="text-white text-[18px]">
                        <i class="fa-brands fa-github text-white mr-1"></i>
                        {{ item.name }}
                      </h1>
                      <div class="flex gap-x-5">
                        <div class="mt-3 text-white">
                          <i class="fas fa-star text-white mr-1"></i
                          >{{ item.stargazers_count }}
                        </div>
                        <div class="mt-3 text-white">
                          <i class="fa-solid fa-code-branch text-white mr-1"></i
                          >{{ item.forks }}
                        </div>
                        <div class="mt-3 text-white">
                          <i class="fa-solid fa-earth text-white mr-1"></i
                          >{{ item.language }}
                        </div>
                      </div>
                    </div>
                  </a>
                </span>
              </div>
            </div>
          </div>
        </div>
      </div>
      <div class="mt-5">
        <div class="p-4 rounded-lg">
          <div class="mt-2">
            <h1 class="font-semibold text-3xl text-white text-center">
              <span class="text-primary">GitHub</span> Gists
            </h1>
            <div class="grid grid-cols-3 md:grid-cols-3 gap-3 mt-6">
              <div v-for="item in github.gists" v-bind:key="item">
                <a :href="item.html_url">
                  <div class="rounded-lg ml-5 p-3">
                    <h1 class="text-white text-[18px]">
                      <i class="fa-brands fa-github text-white mr-1"></i>
                      {{ item.description.lenght > 23 ? `${item.description.substring(0, 20)}...` : item.description }}
                    </h1>
                    <div class="flex gap-x-5">
                      <div class="mt-3 text-white">
                        <i class="fas fa-comments text-white mr-1"></i>
                        {{ item.comments }}
                      </div>
                      <div class="mt-3 text-white">
                        <i class="fas fa-file text-white mr-1"></i>
                        {{ size(item.files) }}
                      </div>
                    </div>
                  </div>
                </a>
              </div>
            </div>
          </div>
        </div>
      </div>
      <div class="py-10 text-white border-t border-zinc-600/20 mt-5">
        <p class="float-left">
          © 2021 / 2022 -
          <a
            class="duration-300 hover:text-primary font-medium"
            href="https://github.com/falsisdev/website"
            ><i class="fa-brands fa-github text-white mr-1"></i> FalsisDev</a
          >
          - All Rights
          <a
            class="duration-300 hover:text-primary font-medium"
            href="https://github.com/falsisdev/website/blob/third/LICENSE"
            >Reserved</a
          >.
        </p>
        <p class="text-white float-right">
          Designed with
          <i class="fa-regular fa-heart mx-1 text-primary"></i> by
          <a class="duration-300 hover:text-primary font-medium" href="/#top"
            >Falsis</a
          >
        </p>
      </div>
    </body>
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
         },
         {
          rel: 'stylesheet', href: 'https://cdn.jsdelivr.net/npm/daisyui@2.31.0/dist/full.css'
         },
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
   /* async mounted() {
      const OPCODES = {
        INFO: 0,
        HELLO: 1,
        INIT: 2,
        HEARTBEAT: 3,
      };
      this.$socket.$on("message", ({ data }) => {
        const JSONdata = JSON.parse(data);
        if(JSONdata.op == OPCODES.HELLO)  {
          this.$socketManager.send(
          JSON.stringify({
            op: OPCODES.INIT,
            d: {
                subscribe_to_id: "539843855567028227",
            }
          })
          )
        setInterval(function() {
        this.$socketManager.send(
            JSON.stringify({
              op: OPCODES.HEARTBEAT,
            })
        )
        , JSONdata.d.heartbeat_interval})
        }else if(JSONdata.op == OPCODES.INFO){
        const colors = {
          online: "#57F287",
          idle: "#faa61a",
          dnd: "#ED4245",
          offline: "#99aab5"
        };
        if(JSONdata.t == "INIT_STATE"){
          const u = JSONdata.d;

          document.getElementById("platform").className = `text-[15px] font-bold text-[${colors[u.discord_status]}]`
        }
        }
      })
     },*/
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
         repos: [], //aslında array içinde object ama burada object açmıyorum aşağıda fetch'leyince object oluşturulacak
         gists: []
       }
     }
   },
   methods: {
size(obj){
 var size = 0,
   key;
 for (key in obj) {
   if (obj.hasOwnProperty(key)) size++;
 }
 return size;
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
     this.github.gists = await fetch(`${githubURL}/gists`).then(res => res.json())
   }
     }
</script>
