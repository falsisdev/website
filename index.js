const app = require("express")();
const axios = require("axios");

const config = {
  web: {
    port: "3000",
  },
  account: {
    discord: "539843855567028227",
    github: "falsisdev",
  },
  api: {
    lanyardBase: "https://api.lanyard.rest",
    githubBase: "https://api.github.com",
  },
};

app.set("view engine", "ejs");
app.set("views", process.cwd() + "/src");

app.get("/redirect/:name", function (req, res) {
  const p = req.params.name;
  res.redirect(
    p
      ? p === "discord"
        ? "https://discord.com/users/" + config.account.discord
        : p === "youtube"
        ? "https://youtube.com/c/Falsis"
        : p === "github"
        ? "https://github.com" + config.account.github
        : p === "animecix"
        ? "https://animecix.net/lists/110787"
        : p === "spotify"
        ? "https://open.spotify.com/user/315l5ib3a4fd2obidm76lipspxji?si=fcf653d3bc0d4cb7"
        : "/"
      : "/"
  );
});
const t = process["env"].auth;
app.get("/", async function (req, res) {
  try {
    const githubUser = await axios.get(
      `${config.api.githubBase}/users/${config.account.github}`,
      t && {
        headers: {
          Authorization: `token ${t}`,
        },
      }
    );
    const githubRepos = await axios.get(
      `${config.api.githubBase}/users/${config.account.github}/repos`,
      t && {
        headers: {
          Authorization: `token ${t}`,
        },
      }
    );
    const lanyard = await axios.get(
      `${config.api.lanyardBase}/v1/users/${config.account.discord}`
    );

    if (req.query.q === "falsis")
      return res.json({
        repo: githubRepos.data,
        github: githubUser.data,
        discord: lanyard.data,
      });
    res.render("index", {
      githubUser: githubUser,
      githubRepos: githubRepos.data,
      user: lanyard.data,
    });
  } catch (e) {
    console.log(e);
    res.send("Sayfa oluşturulurken bir hata verdi!");
  }
});
app.get("/list", async function (req, res) {
  try {
    const githubUser = await axios.get(
      `${config.api.githubBase}/users/${config.account.github}`,
      t && {
        headers: {
          Authorization: `token ${t}`,
        },
      }
    );
    const githubRepos = await axios.get(
      `${config.api.githubBase}/users/${config.account.github}/repos`,
      t && {
        headers: {
          Authorization: `token ${t}`,
        },
      }
    );
    const lanyard = await axios.get(
      `${config.api.lanyardBase}/v1/users/${config.account.discord}`
    );

    if (req.query.q === "falsis")
      return res.json({
        repo: githubRepos.data,
        github: githubUser.data,
        discord: lanyard.data,
      });
    res.render("list", {
      githubUser: githubUser,
      githubRepos: githubRepos.data,
      user: lanyard.data,
    });
  } catch (e) {
    console.log(e);
    res.send("Sayfa oluşturulurken bir hata verdi!");
  }
});
app.get("/css", (req, res) => {
  res.sendFile(process.cwd() + "/src/assets/css/style.css");
});
app.listen(config.web.port, function () {
  console.log("📀 Successfully connected to the website");
});
