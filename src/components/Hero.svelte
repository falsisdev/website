<script>
    let {
        profile = {
            name: "Falsis",
            avatar_url: "https://avatars.githubusercontent.com/u/63756985?v=4",
            bio: "ECE undergrad @ YTU",
            followers: 54,
            public_repos: 6,
            location: "Izmir, Turkiye",
            html_url: "https://github.com/falsisdev",
        },
    } = $props();

    let mounted = $state(false);
    let avatarLoaded = $state(false);

    $effect(() => {
        mounted = true;
    });
</script>

<section
    id="about"
    class="relative flex min-h-[92vh] flex-col items-center justify-center px-6 pt-24 pb-16 text-center overflow-hidden"
>
    <div class="pointer-events-none absolute inset-0 -z-10 overflow-hidden">
        <div
            class="absolute left-1/2 top-1/4 h-[480px] w-[480px] -translate-x-1/2 -translate-y-1/2 rounded-full bg-brand/20 blur-[140px] animate-pulse-glow"
        ></div>
        <div
            class="absolute right-1/4 top-1/2 h-[360px] w-[360px] rounded-full bg-cyan-500/10 blur-[120px]"
        ></div>
        <div
            class="absolute left-1/4 bottom-1/4 h-[300px] w-[300px] rounded-full bg-purple-600/15 blur-[120px]"
        ></div>
    </div>

    <div
        class="relative z-10 mx-auto flex max-w-3xl flex-col items-center text-center transition-all duration-1000 ease-out {mounted
            ? 'translate-y-0 opacity-100'
            : 'translate-y-10 opacity-0'}"
    >
        <div class="relative mb-5 flex justify-center">
            <div class="relative">
                <div
                    class="absolute -inset-2 rounded-full bg-gradient-to-r from-brand via-cyan-400 to-brand-light opacity-75 blur-md group-hover:opacity-100 animate-pulse transition duration-1000"
                ></div>

                <div
                    class="relative flex h-28 w-28 sm:h-32 sm:w-32 items-center justify-center overflow-hidden rounded-full border-2 border-white/20 bg-surface-card shadow-2xl"
                >
                    <img
                        src={profile.avatar_url}
                        alt={profile.name || "falsisdev profile"}
                        loading="eager"
                        decoding="async"
                        class="h-full w-full object-cover transition-transform duration-1500 hover:scale-110 {avatarLoaded
                            ? 'opacity-100'
                            : 'opacity-90'}"
                        onload={() => {
                            avatarLoaded = true;
                        }}
                    />
                </div>

                <div
                    class="absolute bottom-1 right-1 flex items-center justify-center rounded-full border-2 border-surface bg-surface-card p-1 shadow-lg"
                    title="Open to collaborate"
                >
                    <span class="relative flex h-3.5 w-3.5">
                        <span
                            class="absolute inline-flex h-full w-full animate-ping rounded-full bg-emerald-400 opacity-75"
                        ></span>
                        <span
                            class="relative inline-flex h-3.5 w-3.5 rounded-full bg-emerald-500"
                        ></span>
                    </span>
                </div>
            </div>
        </div>
        <div
            class="mb-6 inline-flex items-center gap-2 rounded-full border border-brand/30 bg-brand/10 px-4 py-1.5 text-xs font-semibold tracking-wider text-brand-light uppercase backdrop-blur-md"
        >
            <span
                class="inline-block h-1.5 w-1.5 rounded-full bg-brand-cyan animate-pulse"
            ></span>
            ECE Undergrad @ YTU &bull; Systems & Full-Stack Developer
        </div>

        <h1
            class="mb-6 text-4xl font-extrabold tracking-tight text-white sm:text-6xl md:text-7xl"
        >
            Engineering fast, <br class="hidden sm:inline" />
            <span
                class="bg-gradient-to-r from-white via-zinc-200 to-brand-light bg-clip-text text-transparent"
            >
                scalable software
            </span>
        </h1>

        <p
            class="mx-auto max-w-2xl text-base leading-relaxed text-zinc-400 sm:text-lg"
        >
            Hi, I'm <strong class="text-white font-medium">Falsis</strong>. I
            build high-performance systems and APIs in
            <span class="text-zinc-200">Go</span>, alongside reactive web
            platforms and media tooling with
            <span class="text-zinc-200">Vue, Svelte, and TypeScript</span>.
            Focused on developer tooling, clean architecture, and open source.
        </p>

        <div
            class="mt-8 flex flex-wrap items-center justify-center gap-3 sm:gap-6 text-xs text-zinc-400"
        >
            <div
                class="flex items-center gap-1.5 rounded-full border border-white/5 bg-surface-alt/70 px-3.5 py-1.5"
            >
                <span class="font-bold text-brand-light"
                    >{profile.public_repos || 6}</span
                >
                <span>Public Repos</span>
            </div>
            <div
                class="flex items-center gap-1.5 rounded-full border border-white/5 bg-surface-alt/70 px-3.5 py-1.5"
            >
                <span class="font-bold text-brand-cyan"
                    >{profile.followers || 54}</span
                >
                <span>GitHub Followers</span>
            </div>
            <div
                class="flex items-center gap-1.5 rounded-full border border-white/5 bg-surface-alt/70 px-3.5 py-1.5"
            >
                <span class="h-2 w-2 rounded-full bg-emerald-400"></span>
                <span>{profile.location || "Izmir, Turkiye"}</span>
            </div>
        </div>

        <div class="mt-10 flex flex-wrap items-center justify-center gap-4">
            <a
                href="#projects"
                class="group relative inline-flex items-center gap-2 overflow-hidden rounded-xl bg-gradient-to-r from-brand to-brand-light px-6 py-3 text-sm font-semibold text-white shadow-lg shadow-brand/25 transition-all duration-300 hover:shadow-brand/40 hover:scale-[1.02] active:scale-[0.98]"
            >
                <span>Explore Projects</span>
                <svg
                    class="h-4 w-4 transition-transform duration-200 group-hover:translate-x-1"
                    fill="none"
                    stroke="currentColor"
                    viewBox="0 0 24 24"
                >
                    <path
                        stroke-linecap="round"
                        stroke-linejoin="round"
                        stroke-width="2"
                        d="M14 5l7 7m0 0l-7 7m7-7H3"
                    />
                </svg>
            </a>

            <a
                href={profile.html_url || "https://github.com/falsisdev"}
                target="_blank"
                rel="noopener noreferrer"
                class="group inline-flex items-center gap-2 rounded-xl border border-border bg-surface-alt/80 px-6 py-3 text-sm font-semibold text-zinc-200 backdrop-blur-md transition-all duration-300 hover:border-brand/40 hover:bg-surface-card hover:text-white hover:scale-[1.02] active:scale-[0.98]"
            >
                <svg
                    class="h-4 w-4 fill-current text-zinc-400 group-hover:text-white transition-colors"
                    viewBox="0 0 24 24"
                >
                    <path
                        d="M12 0C5.37 0 0 5.37 0 12c0 5.31 3.435 9.795 8.205 11.385.6.105.825-.255.825-.57 0-.285-.015-1.23-.015-2.235-3.015.555-3.795-.735-4.035-1.41-.135-.345-.72-1.41-1.23-1.695-.42-.225-1.02-.78-.015-.795.945-.015 1.62.87 1.845 1.23 1.08 1.815 2.805 1.305 3.495.99.105-.78.42-1.305.765-1.605-2.67-.3-5.46-1.335-5.46-5.925 0-1.305.465-2.385 1.23-3.225-.12-.3-.54-1.53.12-3.18 0 0 1.005-.315 3.3 1.23.96-.27 1.98-.405 3-.405s2.04.135 3 .405c2.295-1.56 3.3-1.23 3.3-1.23.66 1.65.24 2.88.12 3.18.765.84 1.23 1.905 1.23 3.225 0 4.605-2.805 5.625-5.475 5.925.435.375.81 1.095.81 2.22 0 1.605-.015 2.895-.015 3.3 0 .315.225.69.825.57A12.02 12.02 0 0024 12c0-6.63-5.37-12-12-12z"
                    />
                </svg>
                <span>GitHub Profile</span>
            </a>

            <a
                href="#tech-stack"
                class="inline-flex items-center gap-1.5 rounded-xl border border-transparent px-4 py-3 text-sm font-medium text-zinc-400 transition-colors hover:text-zinc-200"
            >
                <span>View Tech Stack</span>
                <svg
                    class="h-4 w-4"
                    fill="none"
                    stroke="currentColor"
                    viewBox="0 0 24 24"
                >
                    <path
                        stroke-linecap="round"
                        stroke-linejoin="round"
                        stroke-width="2"
                        d="M19 9l-7 7-7-7"
                    />
                </svg>
            </a>
        </div>
    </div>
    <div
        class="absolute bottom-4 left-1/2 -translate-x-1/2 animate-bounce text-zinc-600"
    >
        <a href="#projects" aria-label="Scroll to projects">
            <svg
                class="h-5 w-5"
                fill="none"
                viewBox="0 0 24 24"
                stroke="currentColor"
            >
                <path
                    stroke-linecap="round"
                    stroke-linejoin="round"
                    stroke-width="2"
                    d="M19 14l-7 7m0 0l-7-7m7 7V3"
                />
            </svg>
        </a>
    </div>
</section>
