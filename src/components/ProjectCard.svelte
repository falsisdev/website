<script>
    let { projects = [] } = $props();

    const defaultProjects = [
        {
            title: "Mangile",
            category: "fullstack",
            description:
                "Modern open-source reader platform for manga, manhwa, and light novels. Features an intuitive reading UX, chapter progress caching, and Sanity CMS integration.",
            tags: ["Nuxt 4", "Vue 3", "TypeScript", "Tailwind CSS"],
            url: "https://github.com/falsisdev/mangile",
            demo: "https://mangile.vercel.app",
            featured: true,
            stars: 40,
        },
        {
            title: "Mangile Backend",
            category: "backend",
            description:
                "High-throughput API backend engineered in Go. Leverages goroutines for concurrent scraping pipelines, clean architectural boundaries, and optimized JSON payloads.",
            tags: [
                "Go (Golang)",
                "REST API",
                "Clean Architecture",
                "Concurrency",
            ],
            url: "https://github.com/falsisdev/mangile-backend",
            stars: 3,
        },
        {
            title: "Anthology",
            category: "tooling",
            description:
                "Multi-source streaming catalog and video scrapers for the Stremio and Nuvio media ecosystem. Operates 40+ modular scraping algorithms with QuickJS.",
            tags: ["JavaScript", "QuickJS", "Media Scrapers", "Stremio API"],
            url: "https://github.com/falsisdev/anthology",
            stars: 5,
        },
        {
            title: "Mangile CLI",
            category: "backend",
            description:
                "Command-line downloader and sync client written in Go for automated chapter archival, offline exports, and terminal workflows.",
            tags: ["Go", "CLI Tool", "Automation", "Developer Tools"],
            url: "https://github.com/falsisdev/mangile-cli",
            stars: 3,
        },
        {
            title: "Portfolio Website",
            category: "fullstack",
            description:
                "Ultra-modern portfolio website and digital presence powered by Astro static generation, interactive Svelte 5 islands, and Tailwind CSS v4.",
            tags: ["Astro 5", "Svelte 5", "Tailwind v4", "GitHub API"],
            url: "https://github.com/falsisdev/website",
            demo: "https://falsisdev.github.io/website",
            stars: 13,
        },
        {
            title: "Vessel",
            category: "tooling",
            description:
                "An umbrella, local-first, modular digital media consumption platform.",
            tags: ["Go", "Cloudflare", "Multilingual", "Nuvio"],
            url: "https://github.com/falsisdev/vessel",
            stars: 2,
        },
    ];

    let selectedCategory = $state("all");
    const items = $derived(
        projects && projects.length > 0 ? projects : defaultProjects,
    );

    let filteredProjects = $derived(
        selectedCategory === "all"
            ? items
            : items.filter((p) => p.category === selectedCategory),
    );
</script>

<section id="projects" class="relative px-6 py-24">
    <div class="mx-auto max-w-5xl">
        <!-- Header -->
        <div
            class="mb-12 flex flex-col md:flex-row md:items-end md:justify-between gap-6"
        >
            <div>
                <span
                    class="inline-block mb-2 text-xs font-semibold tracking-widest text-brand-light uppercase"
                >
                    Featured Work & Open Source
                </span>
                <h2 class="text-3xl font-extrabold text-white sm:text-4xl">
                    Projects & Experiments
                </h2>
                <p class="mt-2 text-sm text-zinc-400 max-w-xl">
                    Applications, APIs, and tools I have engineered and maintain
                    in open-source.
                </p>
            </div>

            <!-- Category Filter Pills -->
            <div
                class="flex flex-wrap gap-1.5 rounded-xl border border-white/10 bg-surface-alt p-1 backdrop-blur-md"
            >
                <button
                    onclick={() => (selectedCategory = "all")}
                    class="rounded-lg px-3 py-1.5 text-xs font-medium transition-all {selectedCategory ===
                    'all'
                        ? 'bg-brand text-white shadow-sm'
                        : 'text-zinc-400 hover:text-white'}"
                >
                    All
                </button>
                <button
                    onclick={() => (selectedCategory = "fullstack")}
                    class="rounded-lg px-3 py-1.5 text-xs font-medium transition-all {selectedCategory ===
                    'fullstack'
                        ? 'bg-brand text-white shadow-sm'
                        : 'text-zinc-400 hover:text-white'}"
                >
                    Web & Full-Stack
                </button>
                <button
                    onclick={() => (selectedCategory = "backend")}
                    class="rounded-lg px-3 py-1.5 text-xs font-medium transition-all {selectedCategory ===
                    'backend'
                        ? 'bg-brand text-white shadow-sm'
                        : 'text-zinc-400 hover:text-white'}"
                >
                    Backend & Systems
                </button>
                <button
                    onclick={() => (selectedCategory = "tooling")}
                    class="rounded-lg px-3 py-1.5 text-xs font-medium transition-all {selectedCategory ===
                    'tooling'
                        ? 'bg-brand text-white shadow-sm'
                        : 'text-zinc-400 hover:text-white'}"
                >
                    Scrapers & Tooling
                </button>
            </div>
        </div>

        <!-- Grid -->
        <div class="grid gap-6 sm:grid-cols-2 lg:grid-cols-3">
            {#each filteredProjects as project}
                <div
                    class="group relative flex flex-col justify-between rounded-2xl border border-white/10 bg-surface-card p-6 backdrop-blur-md transition-all duration-300 hover:border-brand/40 hover:-translate-y-1.5 hover:shadow-2xl hover:shadow-brand/5"
                >
                    <!-- Top highlight aura on hover -->
                    <div
                        class="pointer-events-none absolute -inset-px rounded-2xl bg-gradient-to-b from-brand/20 to-transparent opacity-0 transition-opacity duration-300 group-hover:opacity-100"
                    ></div>

                    <div class="relative z-10">
                        <!-- Header row: Title & Star Count / Actions -->
                        <div
                            class="mb-3 flex items-start justify-between gap-2"
                        >
                            <h3
                                class="text-lg font-bold text-white group-hover:text-brand-light transition-colors"
                            >
                                {project.title}
                            </h3>

                            <div class="flex items-center gap-2">
                                {#if project.stars}
                                    <span
                                        class="inline-flex items-center gap-1 rounded-full border border-amber-500/20 bg-amber-500/10 px-2 py-0.5 text-[11px] font-medium text-amber-300"
                                    >
                                        <svg
                                            class="h-3 w-3 fill-current"
                                            viewBox="0 0 20 20"
                                        >
                                            <path
                                                d="M9.049 2.927c.3-.921 1.603-.921 1.902 0l1.07 3.292a1 1 0 00.95.69h3.462c.969 0 1.371 1.24.588 1.81l-2.8 2.034a1 1 0 00-.364 1.118l1.07 3.292c.3.921-.755 1.688-1.54 1.118l-2.8-2.034a1 1 0 00-1.175 0l-2.8 2.034c-.784.57-1.838-.197-1.539-1.118l1.07-3.292a1 1 0 00-.364-1.118L2.98 8.72c-.783-.57-.38-1.81.588-1.81h3.461a1 1 0 00.951-.69l1.07-3.292z"
                                            />
                                        </svg>
                                        {project.stars}
                                    </span>
                                {/if}

                                {#if project.demo}
                                    <a
                                        href={project.demo}
                                        target="_blank"
                                        rel="noopener noreferrer"
                                        title="Live Demo"
                                        class="rounded-lg border border-white/10 bg-surface-alt p-1.5 text-zinc-400 hover:text-brand-light hover:border-brand/40 transition-colors"
                                    >
                                        <svg
                                            class="h-3.5 w-3.5"
                                            fill="none"
                                            stroke="currentColor"
                                            viewBox="0 0 24 24"
                                        >
                                            <path
                                                stroke-linecap="round"
                                                stroke-linejoin="round"
                                                stroke-width="2"
                                                d="M10 6H6a2 2 0 00-2 2v10a2 2 0 002 2h10a2 2 0 002-2v-4M14 4h6m0 0v6m0-6L10 14"
                                            />
                                        </svg>
                                    </a>
                                {/if}

                                {#if project.url}
                                    <a
                                        href={project.url}
                                        target="_blank"
                                        rel="noopener noreferrer"
                                        title="GitHub Repository"
                                        class="rounded-lg border border-white/10 bg-surface-alt p-1.5 text-zinc-400 hover:text-white hover:border-brand/40 transition-colors"
                                    >
                                        <svg
                                            class="h-3.5 w-3.5 fill-current"
                                            viewBox="0 0 24 24"
                                        >
                                            <path
                                                d="M12 0C5.37 0 0 5.37 0 12c0 5.31 3.435 9.795 8.205 11.385.6.105.825-.255.825-.57 0-.285-.015-1.23-.015-2.235-3.015.555-3.795-.735-4.035-1.41-.135-.345-.72-1.41-1.23-1.695-.42-.225-1.02-.78-.015-.795.945-.015 1.62.87 1.845 1.23 1.08 1.815 2.805 1.305 3.495.99.105-.78.42-1.305.765-1.605-2.67-.3-5.46-1.335-5.46-5.925 0-1.305.465-2.385 1.23-3.225-.12-.3-.54-1.53.12-3.18 0 0 1.005-.315 3.3 1.23.96-.27 1.98-.405 3-.405s2.04.135 3 .405c2.295-1.56 3.3-1.23 3.3-1.23.66 1.65.24 2.88.12 3.18.765.84 1.23 1.905 1.23 3.225 0 4.605-2.805 5.625-5.475 5.925.435.375.81 1.095.81 2.22 0 1.605-.015 2.895-.015 3.3 0 .315.225.69.825.57A12.02 12.02 0 0024 12c0-6.63-5.37-12-12-12z"
                                            />
                                        </svg>
                                    </a>
                                {/if}
                            </div>
                        </div>

                        <!-- Description -->
                        <p class="mb-5 text-sm leading-relaxed text-zinc-400">
                            {project.description}
                        </p>
                    </div>

                    <!-- Tags -->
                    <div
                        class="relative z-10 mt-auto flex flex-wrap gap-1.5 pt-3 border-t border-white/5"
                    >
                        {#each project.tags as tag}
                            <span
                                class="rounded-md border border-white/5 bg-zinc-800/70 px-2 py-0.5 text-[11px] font-medium text-zinc-300"
                            >
                                {tag}
                            </span>
                        {/each}
                    </div>
                </div>
            {/each}
        </div>
    </div>
</section>
