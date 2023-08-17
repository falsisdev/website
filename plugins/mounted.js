export default defineNuxtPlugin((nuxtApp) => {
    nuxtApp.hook('app:mounted', () => {
        var searchParams = new URLSearchParams(window.location.search);
        if(!searchParams.toString()) {
            const OPCODES = {
                INFO: 0,
                HELLO: 1,
                INIT: 2,
                HEARTBEAT: 3,
            };
            const ws = new WebSocket("wss://api.lanyard.rest/socket"); //lanyard websocket
            /* EVENTS */
            ws.onmessage = ({ data }) => {
            const JSONdata = JSON.parse(data);
            
            if(JSONdata.op == OPCODES.HELLO) { //first start
                //authenticator identify
                ws.send(
                    JSON.stringify({
                        op: OPCODES.INIT,
                        d: {
                            subscribe_to_id: "539843855567028227",
                        }
                    })
                )//send
            
                setInterval(function() {
                    ws.send(
                        JSON.stringify({
                        op: OPCODES.HEARTBEAT,
                        })
                    )
                , JSONdata.d.heartbeat_interval}); //interval
            }else if(JSONdata.op == OPCODES.INFO){//if(JSONdata.op == OPCODES.HELLO){}
                const colors = {
                    online: "#57F287",
                    idle: "#faa61a",
                    dnd: "#ED4245",
                    offline: "#99aab5"
                };//for status
            
                if(JSONdata.t == "INIT_STATE"){ //first
                    const u = JSONdata.d; //userdata
            
                    let platform = u.active_on_discord_web ? "Internet Tarayıcısı" : u.active_on_discord_mobile ? "Mobil" : u.active_on_discord_desktop ? "Masaüstü" : ""
                    let icon = platform == "Internet Tarayıcısı" ? "fas fa-globe" : platform == "Mobil" ? "fas fa-mobile-alt" : platform == "Masaüstü" ? "fas fa-desktop" : ""
                    document.getElementById("spotify").innerHTML = u.listening_to_spotify ? `Listening to <span class="text-white">Spotify</span>` : ''
                    document.getElementById("spotifywarning").style.visibility = u.listening_to_spotify ? 'visible' : 'hidden'
                    document.getElementById("song").innerText = u.listening_to_spotify ? u.spotify.song : ''
                    document.getElementById("artist").innerText = u.listening_to_spotify ? u.spotify.artist : ''
                    document.getElementById("uname").innerText = u.discord_user.display_name
                    document.getElementById("avatar").src = `https://cdn.discordapp.com/avatars/539843855567028227/${u.discord_user.avatar}.png?size=128`
                    document.getElementById("status1").style.backgroundColor = colors[u.discord_status]
                    document.getElementById("status2").style.backgroundColor = colors[u.discord_status]
                }else if(JSONdata.t = "PRESENCE_UPDATE") { //update
                    const u = JSONdata.d;
                    let platform = u.active_on_discord_web ? "Internet Tarayıcısı" : u.active_on_discord_mobile ? "Mobil" : u.active_on_discord_desktop ? "Masaüstü" : ""
                    let icon = platform == "Internet Tarayıcısı" ? "fas fa-globe" : platform == "Mobil" ? "fas fa-mobile-alt" : platform == "Masaüstü" ? "fas fa-desktop" : ""
                    document.getElementById("spotify").innerHTML = u.listening_to_spotify ? `Listening to <span class="text-white">Spotify</span>` : ''
                    document.getElementById("spotifywarning").style.visibility = u.listening_to_spotify ? 'visible' : 'hidden'
                    document.getElementById("song").innerText = u.listening_to_spotify ? u.spotify.song : ''
                    document.getElementById("artist").innerText = u.listening_to_spotify ? u.spotify.artist : ''
                    document.getElementById("uname").innerText = u.discord_user.display_name
                    document.getElementById("avatar").src = `https://cdn.discordapp.com/avatars/539843855567028227/${u.discord_user.avatar}.png?size=128`
                    document.getElementById("status1").style.backgroundColor = colors[u.discord_status]
                    document.getElementById("status2").style.backgroundColor = colors[u.discord_status]
                }
            }
            }//onmessage
        }
    })
})