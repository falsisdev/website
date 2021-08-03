const app = require('express')();
const fetch = require('node-fetch')
const path = require('path')
////////////////////////----------------------\\\\\\\\\\\\\\\\\\\\\\\\\
const lanyard = 'https://api.lanyard.rest/v1/users/539843855567028227'
const github = 'https://api.github.com/users/falsisdev'
const repos = 'https://api.github.com/users/falsisdev/repos'
const port = {
	port: 3000
}
////////////////////////----------------------\\\\\\\\\\\\\\\\\\\\\\\\\
app.set('view engine', 'ejs');
app.set('views', __dirname + '/src');
////////////////////////----------------------\\\\\\\\\\\\\\\\\\\\\\\\\
let checkSpotify;
let spotify;
let user;
let activity;
let assets;
let appid;
let device;
let checkActivities;
////////////////////////----------------------\\\\\\\\\\\\\\\\\\\\\\\\\
////////////////////////----------------------\\\\\\\\\\\\\\\\\\\\\\\\\
app.get("/redirect/:name", function(req, res) {
	if(req.params.name === "discord"){
		res.redirect("https://discord.com/users/539843855567028227")
	}
	if(req.params.name === "youtube"){
		res.redirect("https://youtube.com/c/Falsis")
	}
	if(req.params.name === "github"){
		res.redirect("https://github.com/falsisdev")
	}
	if(req.params.name === "repo"){
		res.redirect("https://github.com/falsisdev/site")
	}
	if(!req.params.name){
		res.send("Redirect failed")
	}
	if(req.params.name === "spotify"){
	res.redirect("https://open.spotify.com/user/315l5ib3a4fd2obidm76lipspxji?si=fcf653d3bc0d4cb7")
}else{
	res.redirect(req.params.name)
}
})
app.get("/", function(req, res){
  fetch(lanyard).then(json => json.json()).then(result => {
checkActivities = JSON.stringify(result.data.activities) === "[]" //boolean
checkSpotify = result.data.listening_to_spotify//boolean
spotify = checkSpotify ? result.data.spotify  : false
user = {
	username: result.data.discord_user.username,
	public_flags: result.data.discord_user.public_flags,
	id: result.data.discord_user.id,//number
	tag: result.data.discord_user.discriminator,//number
	avatar: result.data.discord_user.avatar,
	avatarURL: `https://cdn.discordapp.com/avatars/${result.data.discord_user.id}/${result.data.discord_user.avatar}.webp`,
	status: result.data.discord_status
}//user
activity = {
	  start: checkActivities ? null : result.data.activities[0].timestamps.start,//number
    text1: checkActivities ? null : result.data.activities[0].name,
    text2: checkActivities ? null : result.data.activities[0].details,
    text3: checkActivities ? null : result.data.activities[0].state
}//activity
assets = {
	small: {
		text: checkActivities ? null : result.data.activities[0].assets.small_text,
		image: checkActivities ? null : result.data.activities[0].assets.small_image
	},//small
	large: {
		text: checkActivities ? null : result.data.activities[0].assets.large_text,
		image: checkActivities ? null : result.data.activities[0].assets.large_image
	}//large
}//assets
appid = checkActivities ? null : result.data.activities[0].application_id//number
device = {
	mobile: result.data.active_on_discord_mobile,//boolean
	pc: result.data.active_on_discord_desktop,//boolean
}//device
})//lanyard-fetch
////////////////////////----------------------\\\\\\\\\\\\\\\\\\\\\\\\\
let githubUser;
let githubRepos;
////////////////////////----------------------\\\\\\\\\\\\\\\\\\\\\\\\\
fetch(github).then(json => json.json()).then(result => {
githubUser = {
	username: result.name,
	nickname: result.login,
	avatar: result.avatar_url,
	url: result.html_url,
	website: result.blog,
	location: result.location,
	bio: result.bio,
	repos: result.public_repos,//number
	gists: result.public_gists,//number
	followers: result.followers,//number
	following: result.following//number
}//githubUser
})//github-fetch
res.render('index', {checkSpotify, checkActivities, spotify, user, activity, assets, appid, device, githubUser, githubRepos})
})//get()
app.get("/css", (req, res) => {
	res.sendFile(__dirname + '/src/assets/css/style.css');
})
app.listen(port.port, function() {
  console.log('📀 Successfully connected to the website')
});
////////////////////////----------------------\\\\\\\\\\\\\\\\\\\\\\\\\
////////////////////////----------------------\\\\\\\\\\\\\\\\\\\\\\\\\
