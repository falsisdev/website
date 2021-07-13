const app = require(`express`)();
app.listen(3000,function(){
console.log("Siteye Bağlanıldı");
});
app.set("view engine", "ejs");
app.set('views', __dirname+'/ejs');
const fetch = require("node-fetch")
app.get("/", (req, res) => {
  fetch(`https://api.lanyard.rest/v1/users/539843855567028227`).then(x => x.json()).then(z => {
    var status = z.data.discord_status
    var ac = z.data.activities
    if(ac === "[]"){
    var spotify = "";
    var fspotify = "";
    var avatar = `https://cdn.discordapp.com/avatars/${z.data.discord_user.id}/${z.data.discord_user.avatar}.webp`
    var customstatus2 = {
      state: "",
      name: "Şuanda Durumumda Birşey Bulunmuyor.",
      details: "",
      image: ""
    }
    }
    if(status === "offline"){
      var spotify = "";
    var fspotify = "";
    var avatar = `https://cdn.discordapp.com/avatars/${z.data.discord_user.id}/${z.data.discord_user.avatar}.webp`
    var customstatus2 = {
      state: "",
      name: "Şuanda Durumumda Birşey Bulunmuyor.",
      details: "",
      image: ""
    }
    }
        if(!z.data.activities){
      var spotify = "";
    var fspotify = "";
    var avatar = `https://cdn.discordapp.com/avatars/${z.data.discord_user.id}/${z.data.discord_user.avatar}.webp`
    var customstatus2 = {
      state: "",
      name: "Şuanda Durumumda Birşey Bulunmuyor.",
      details: "",
      image: ""
    }
    if(z.data.activities[0].id === "custom"){
            var spotify = "";
    var fspotify = "";
    var avatar = `https://cdn.discordapp.com/avatars/${z.data.discord_user.id}/${z.data.discord_user.avatar}.webp`
    var customstatus2 = {
      state: "",
      name: z.data.activities[0].state,
      details: "",
      image: ""
    }
    }
    }else{
      try{
    var spotify = z.data.spotify;
    var fspotify = z.data.listening_to_spotify;
    var avatar = `https://cdn.discordapp.com/avatars/${z.data.discord_user.id}/${z.data.discord_user.avatar}.webp`
    var customstatus2 = {
      state: z.data.activities[0].state,
      name: z.data.activities[0].name,
      details: z.data.activities[0].details,
      image: fspotify ? spotify.album_art_url : `https://cdn.discordapp.com/app-assets/${z.data.activities[0].application_id}/${z.data.activities[0].assets.large_image}.png`
    }
    }catch(error){
            var spotify = "";
    var fspotify = "";
    var avatar = `https://cdn.discordapp.com/avatars/${z.data.discord_user.id}/${z.data.discord_user.avatar}.webp`
    var customstatus2 = {
      state: "",
      name: "Şuanda Durumumda Birşey Bulunmuyor.",
      details: "",
      image: ""
    }
    }
    }
	res.render("index", {spotify, fspotify, status, customstatus2, avatar})
  })
})
