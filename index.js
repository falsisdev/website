const app = require(`express`)();
////////////////////////----------------------\\\\\\\\\\\\\\\\\\\\\\\\\
app.listen(3000, function() {
  console.log("Siteye Bağlanıldı");
});
////////////////////////----------------------\\\\\\\\\\\\\\\\\\\\\\\\\
app.get("/css", function(req, res) {
  res.sendFile(__dirname + "/style.css")
})
////////////////////////----------------------\\\\\\\\\\\\\\\\\\\\\\\\\
app.set("view engine", "ejs");
app.set('views', __dirname + '/ejs');
////////////////////////----------------------\\\\\\\\\\\\\\\\\\\\\\\\\
const fetch = require("node-fetch")
////////////////////////----------------------\\\\\\\\\\\\\\\\\\\\\\\\\
app.get("/", (req, res) => {
  var foll;
  fetch("https://api.github.com/users/falsisdev").then(x => x.json()).then(z => {
    foll = z.followers
  })
  fetch(`https://api.lanyard.rest/v1/users/539843855567028227`).then(x => x.json()).then(z => {
    var followers = foll
    var status = z.data.discord_status
    var ac = z.data.activities
    if (ac === "[]") {
      var spotify = "";
      var fspotify = false;
      var avatar = `https://cdn.discordapp.com/avatars/${z.data.discord_user.id}/${z.data.discord_user.avatar}.webp`
      var customstatus2 = {
       data: z.data.activities[0],
       spalb: "",
        state: "",
        name: "I'm not playing anything",
        details: "",
        image: ""
      }
    }
    ////////////////////////----------------------\\\\\\\\\\\\\\\\\\\\\\\\\
    if (status === "offline") {
      var fspotify = false;
      var avatar = `https://cdn.discordapp.com/avatars/${z.data.discord_user.id}/${z.data.discord_user.avatar}.webp`
      var customstatus2 = {
       data: z.data.activities[0],
       spalb: "",
        state: "",
        name: "I'm not playing anything",
        details: "",
        image: ""
      }
    }
    ////////////////////////----------------------\\\\\\\\\\\\\\\\\\\\\\\\\
    if (!z.data.activities) {
      var fspotify = false;
      var avatar = `https://cdn.discordapp.com/avatars/${z.data.discord_user.id}/${z.data.discord_user.avatar}.webp`
      var customstatus2 = {
       data: z.data.activities[0],
       spalb: "",
        state: "",
        name: "I'm not playing anything",
        details: "",
        image: ""
      }
      if (z.data.activities[0].id === "custom") {
        var spotify = "";
        var fspotify = false;
        var avatar = `https://cdn.discordapp.com/avatars/${z.data.discord_user.id}/${z.data.discord_user.avatar}.webp`
        var customstatus2 = {
         data: z.data.activities[0],
         spalb: "",
          state: "",
          name: z.data.activities[0].state,
          details: "",
          image: ""
        }
      }
    } else {
      try {
        var spotify = z.data.spotify;
        var fspotify = z.data.listening_to_spotify;
        var avatar = `https://cdn.discordapp.com/avatars/${z.data.discord_user.id}/${z.data.discord_user.avatar}.webp`
        var customstatus2 = {
         data: z.data.activities[0],
         spalb: fspotify ? z.data.spotify.album_art_url : "",
          state: z.data.activities[0].state,
          name: z.data.activities[0].name,
          details: z.data.activities[0].details,
          image: fspotify ? z.data.spotify.album_art_url : `https://cdn.discordapp.com/app-assets/${z.data.activities[0].application_id}/${z.data.activities[0].assets.large_image}.png`
        }
      } catch (error) {
        var fspotify = false;
        var avatar = `https://cdn.discordapp.com/avatars/${z.data.discord_user.id}/${z.data.discord_user.avatar}.webp`
        var customstatus2 = {
          data: z.data.activities[0],
          spalb: "",
          state: "",
          name: "I'm not playing anything",
          details: "",
          image: ""
        }
      }
    }
    ////////////////////////----------------------\\\\\\\\\\\\\\\\\\\\\\\\\
    res.render("index", { spotify, fspotify, status, customstatus2, avatar, followers})//render
  })//fetch
  ////////////////////////----------------------\\\\\\\\\\\\\\\\\\\\\\\\\
})//get
