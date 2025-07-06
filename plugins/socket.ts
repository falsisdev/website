//@ts-nocheck
export default defineNuxtPlugin((nuxtApp) => {
  nuxtApp.hook("page:loading:end", () => {
    var searchParams = new URLSearchParams(window.location.search);
    if (!searchParams.toString()) {
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

        if (JSONdata.op == OPCODES.HELLO) {
          ws.send(
            JSON.stringify({
              op: OPCODES.INIT,
              d: {
                subscribe_to_id: "539843855567028227",
              },
            })
          );

          setInterval(function () {
            ws.send(
              JSON.stringify({
                op: OPCODES.HEARTBEAT,
              })
            );
          }, JSONdata.d.heartbeat_interval);
        } else if (JSONdata.op == OPCODES.INFO) {
          const getDominantColor = (imgUrl, cb) => {
            const img = new window.Image();
            img.crossOrigin = "Anonymous";
            img.src = imgUrl;
            img.onload = function () {
              const canvas = document.createElement("canvas");
              canvas.width = img.width;
              canvas.height = img.height;
              const ctx = canvas.getContext("2d");
              ctx.drawImage(img, 0, 0);
              const data = ctx.getImageData(0, 0, canvas.width, canvas.height).data;
              let r = 0, g = 0, b = 0, count = 0;
              for (let i = 0; i < data.length; i += 4) {
                r += data[i];
                g += data[i + 1];
                b += data[i + 2];
                count++;
              }
              r = Math.round(r / count);
              g = Math.round(g / count);
              b = Math.round(b / count);
              cb([r, g, b]);
            };
            img.onerror = function () {
              cb([0, 0, 0]);
            };
          };

          const getAlertClassByColor = ([r, g, b]) => {
            const max = Math.max(r, g, b);
            const min = Math.min(r, g, b);
            const isWhite = max > 230 && min > 200;
            const isBlack = max < 40 && min < 40;
            if (isWhite) return "alert-outline";
            if (isBlack) return "alert-neutral";
            // HSV'ye çevir
            const rr = r / 255, gg = g / 255, bb = b / 255;
            const mx = Math.max(rr, gg, bb), mn = Math.min(rr, gg, bb);
            const d = mx - mn;
            let h = 0;
            if (d === 0) h = 0;
            else if (mx === rr) h = ((gg - bb) / d) % 6;
            else if (mx === gg) h = (bb - rr) / d + 2;
            else h = (rr - gg) / d + 4;
            h = Math.round(h * 60);
            if (h < 0) h += 360;

           // En yakın renk grubu
            if (h >= 60 && h < 160) return "alert-success"; //Yeşil tonları
            if (h >= 160 && h < 250) return "alert-info"; //Mavi Tonları
            if ((h >= 0 && h < 20) || (h >= 340 && h <= 360)) return "alert-error"; // Kırmızı/Pembe Tonları
            if (h >= 20 && h < 60) return "alert-warning";// Sarı/Turuncu Tonları
            if (h >= 250 && h < 300) return "alert-accent"; //Mor Tonları
            if (h >= 300 && h < 340) return "alert-accent"; //Eflatun Tonları
            return "alert-neutral";
          };

          const updateSpotifyUI = (u) => {
            if (u.listening_to_spotify) {
              document
                .getElementById("spotifycheck")
                .classList.replace("hidden", "visible");

              const albumArtUrl = u.spotify.album_art_url;
              document.getElementById("albumart").src = albumArtUrl;
              document.getElementById("albumname")?.setAttribute("data-tip", u.spotify.album);
              document.getElementById("title").innerText = u.spotify.song;
              document.getElementById("artist").innerText = u.spotify.artist;

              // Şarkı ilerlemesini ilet
              if (u.spotify.timestamps?.start && u.spotify.timestamps?.end) {
                const event = new CustomEvent("spotify-progress", {
                  detail: {
                    start: u.spotify.timestamps.start,
                    end: u.spotify.timestamps.end,
                    isPlaying: true
                  }
                });
                window.dispatchEvent(event);
              } else {
                // Şarkı yoksa ilerlemeyi sıfırla
                const event = new CustomEvent("spotify-progress", {
                  detail: {
                    start: 0,
                    end: 0,
                    isPlaying: false
                  }
                });
                window.dispatchEvent(event);
              }

              // Alert class güncelle
              getDominantColor(albumArtUrl, (rgb) => {
                const alertDiv = document.getElementById("alertDiv");
                if (alertDiv) {
                  alertDiv.classList.remove(
                    "alert-success",
                    "alert-info",
                    "alert-error",
                    "alert-warning",
                    "alert-neutral",
                    "alert-outline"
                  );
                  const newClass = getAlertClassByColor(rgb);
                  if (newClass) alertDiv.classList.add(newClass);
                }
              });
            } else {
              // Spotify dinlenmiyorsa ilerlemeyi sıfırla
              const event = new CustomEvent("spotify-progress", {
                detail: {
                  start: 0,
                  end: 0,
                  isPlaying: false
                }
              });
              window.dispatchEvent(event);
            }
          };

          if (JSONdata.t == "INIT_STATE") {
            //first
            const u = JSONdata.d; //userdata
            document.getElementById(
              "avatar"
            ).src = `https://cdn.discordapp.com/avatars/539843855567028227/${u.discord_user.avatar}.png?size=4096`;
            updateSpotifyUI(u);
          } else if (JSONdata.t == "PRESENCE_UPDATE") {
            //update
            const u = JSONdata.d;
            document.getElementById(
              "avatar"
            ).src = `https://cdn.discordapp.com/avatars/539843855567028227/${u.discord_user.avatar}.png?size=4096`;
            updateSpotifyUI(u);
          }
        }
      }; //onmessage
    }
  });
});