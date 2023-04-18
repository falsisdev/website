export default defineNuxtConfig({
    modules: ['@nuxtjs/tailwindcss', 'nuxt-icons'],
    css: [
        '@fortawesome/fontawesome-svg-core/styles.css'
    ],
    runtimeConfig: {
        public: {
            apis: {
                baseURLs: {
                    github: "https://api.github.com"
                },
                endpoints: {
                    github: {
                        userinfo: "users",
                        repos: "repos"
                    }
                },
                userinfo: {
                    githubname: "falsisdev"
                }
            },
            social: {
                instagram: "https://instagram.com/falsisdev",
                youtube: "https://youtube.com/c/Falsis",
                discord: "https://discord.com/users/539843855567028227",
                spotify: "https://open.spotify.com/user/315l5ib3a4fd2obidm76lipspxji?si=0dbe8cea814847f9",
                github: "https://github.com/falsisdev",
                reddit: "https://www.reddit.com/user/fnfalsiss",
                animecix: "https://animecix.net/users/65384",
                mangadex: "https://mangadex.org/user/b725b8b4-db9d-4c92-b7c6-6cdb0902da40",
                mail: "falsis@proton.me"
            },
            tools: [
                {
                    "name": "HTML",
                    "fullname": "Hyper Text Markup Language",
                    "defaultname": "HTML",
                    "color": "#f06529"
                },
                {
                    "name": "CSS",
                    "fullname": "Cascading Style Sheets",
                    "defaultname": "CSS",
                    "color": "#1572B6"
                },
                {
                    "name": "Tailwind",
                    "fullname": "TailwindCSS",
                    "defaultname": "TailwindCSS",
                    "color": "#06B6D4"
                },
                {
                    "name": "JS",
                    "fullname": "JavaScript",
                    "defaultname": "JavaScript",
                    "color": "#F7DF1E"
                },
                {
                    "name": "Node",
                    "fullname": "NodeJS",
                    "defaultname": "NodeJS",
                    "color": "#339933"
                },
                {
                    "name": "TS",
                    "fullname": "TypeScript",
                    "defaultname": "TypeScript",
                    "color": "#3178C6"
                },
                {
                    "name": "NPM",
                    "fullname": "Node Package Manager",
                    "defaultname": "NPM",
                    "color": "#CB3837"
                },
                {
                    "name": "YARN",
                    "fullname": "YARN",
                    "defaultname": "YARN",
                    "color": "#2C8EBB"
                },
                {
                    "name": "Vue",
                    "fullname": "Vuejs",
                    "defaultname": "Vue",
                    "color": "#4FC08D"
                },
                {
                    "name": "Nuxt",
                    "fullname": "Nuxtjs",
                    "defaultname": "Nuxt",
                    "color": "#00DC82"
                },
                {
                    "name": "Svelte",
                    "fullname": "Svelte",
                    "defaultname": "Svelte",
                    "color": "#ff3c00"
                },
                {
                    "name": "Express",
                    "fullname": "Expressjs",
                    "defaultname": "Express",
                    "color": "#080808"
                },
                {
                    "name": "HBS",
                    "fullname": "Handlebarsjs",
                    "defaultname": "Handlebars",
                    "color": "#CB3837"
                },
                {
                    "name": "Git",
                    "fullname": "Git",
                    "defaultname": "Git",
                    "color": "#F05032"
                },
                {
                    "name": "Vercel",
                    "fullname": "Vercel",
                    "defaultname": "Vercel",
                    "color": "#060606"
                }
            ]
        }
    }
})