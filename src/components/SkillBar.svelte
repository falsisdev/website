<script>
    let {
        techStack = {
            totalBytes: 1198888,
            repoCount: 6,
            languages: [
                {
                    name: "JavaScript",
                    bytes: 689532,
                    percentage: 57.5,
                    color: "#F7DF1E",
                    category: "frontend",
                },
                {
                    name: "Vue",
                    bytes: 232160,
                    percentage: 19.4,
                    color: "#41B883",
                    category: "frontend",
                },
                {
                    name: "Go",
                    bytes: 229401,
                    percentage: 19.1,
                    color: "#00ADD8",
                    category: "systems",
                },
                {
                    name: "Python",
                    bytes: 22638,
                    percentage: 1.9,
                    color: "#3572A5",
                    category: "scripting",
                },
                {
                    name: "TypeScript",
                    bytes: 19157,
                    percentage: 1.6,
                    color: "#3178C6",
                    category: "frontend",
                },
                {
                    name: "Svelte",
                    bytes: 6385,
                    percentage: 0.5,
                    color: "#FF3E00",
                    category: "frontend",
                },
            ],
            topTechnologies: [
                "Go",
                "TypeScript",
                "Vue / Nuxt",
                "Svelte",
                "Tailwind CSS",
                "Docker",
            ],
        },
    } = $props();

    let visible = $state(false);
    let activeTab = $state("github"); // 'github' | 'ecosystem'

    const ecosystemCategories = [
        {
            title: "Core Languages",
            description:
                "Systems programming, type-safe development & scripting",
            items: [
                { name: "Go", badge: "Primary", highlight: true },
                { name: "JavaScript", badge: "Primary", highlight: true },
                { name: "TypeScript", badge: "Fluent" },
                { name: "Python", badge: "Scripting & Data" },
            ],
        },
        {
            title: "Frontend & Full-Stack",
            description:
                "High-performance reactive interfaces & static generation",
            items: [
                {
                    name: "Vue 3 & Nuxt 4",
                    badge: "Full-Stack",
                    highlight: true,
                },
                { name: "Svelte 5", badge: "UI Islands" },
                { name: "Astro", badge: "SSG Framework" },
                { name: "Tailwind CSS v4", badge: "Styling" },
            ],
        },
        {
            title: "Backend & Systems",
            description: "Scalable APIs, CLI tools & data scrapers",
            items: [
                { name: "Go REST APIs", badge: "High-perf", highlight: true },
                { name: "QuickJS & Media Scrapers", badge: "Scraping" },
                { name: "Concurrency & Goroutines", badge: "Architecture" },
                { name: "Microservices", badge: "Design" },
            ],
        },
        {
            title: "Tools & Infrastructure",
            description: "Modern developer workflow & deployment pipeline",
            items: [
                { name: "Git & GitHub", badge: "VCS", highlight: true },
                { name: "macOS", badge: "Environment", highlight: true },
                { name: "Vite", badge: "Tooling" },
            ],
        },
    ];

    function formatBytes(bytes) {
        if (!bytes) return "0 KB";
        if (bytes >= 1048576) {
            return (bytes / 1048576).toFixed(1) + " MB";
        }
        return (bytes / 1024).toFixed(0) + " KB";
    }

    $effect(() => {
        const timer = setTimeout(() => {
            visible = true;
        }, 60);

        return () => clearTimeout(timer);
    });
</script>

<section id="tech-stack" class="relative px-6 py-24">
    <!-- Subtle ambient glow -->
    <div class="pointer-events-none absolute inset-0 -z-10 overflow-hidden">
        <div
            class="absolute left-1/3 bottom-10 h-80 w-80 rounded-full bg-brand/10 blur-[130px]"
        ></div>
    </div>

    <div class="mx-auto max-w-5xl">
        <!-- Header -->
        <div
            class="mb-12 flex flex-col md:flex-row md:items-end md:justify-between gap-6"
        >
            <div>
                <span
                    class="inline-block mb-2 text-xs font-semibold tracking-widest text-brand-light uppercase"
                >
                    Live GitHub Analytics & Capabilities
                </span>
                <h2 class="text-3xl font-extrabold text-white sm:text-4xl">
                    Tech Stack & Codebase Distribution
                </h2>
                <p class="mt-2 text-sm text-zinc-400 max-w-xl">
                    Aggregated language distribution calculated directly from my
                    GitHub repositories.
                </p>
            </div>

            <!-- Tab switchers -->
            <div
                class="inline-flex rounded-xl border border-white/10 bg-surface-alt p-1 backdrop-blur-md"
            >
                <button
                    onclick={() => (activeTab = "github")}
                    class="flex items-center gap-2 rounded-lg px-4 py-2 text-xs font-medium transition-all {activeTab ===
                    'github'
                        ? 'bg-brand text-white shadow-md'
                        : 'text-zinc-400 hover:text-white'}"
                >
                    <svg class="h-3.5 w-3.5 fill-current" viewBox="0 0 24 24">
                        <path
                            d="M12 0C5.37 0 0 5.37 0 12c0 5.31 3.435 9.795 8.205 11.385.6.105.825-.255.825-.57 0-.285-.015-1.23-.015-2.235-3.015.555-3.795-.735-4.035-1.41-.135-.345-.72-1.41-1.23-1.695-.42-.225-1.02-.78-.015-.795.945-.015 1.62.87 1.845 1.23 1.08 1.815 2.805 1.305 3.495.99.105-.78.42-1.305.765-1.605-2.67-.3-5.46-1.335-5.46-5.925 0-1.305.465-2.385 1.23-3.225-.12-.3-.54-1.53.12-3.18 0 0 1.005-.315 3.3 1.23.96-.27 1.98-.405 3-.405s2.04.135 3 .405c2.295-1.56 3.3-1.23 3.3-1.23.66 1.65.24 2.88.12 3.18.765.84 1.23 1.905 1.23 3.225 0 4.605-2.805 5.625-5.475 5.925.435.375.81 1.095.81 2.22 0 1.605-.015 2.895-.015 3.3 0 .315.225.69.825.57A12.02 12.02 0 0024 12c0-6.63-5.37-12-12-12z"
                        />
                    </svg>
                    <span>GitHub Languages</span>
                </button>
                <button
                    onclick={() => (activeTab = "ecosystem")}
                    class="flex items-center gap-2 rounded-lg px-4 py-2 text-xs font-medium transition-all {activeTab ===
                    'ecosystem'
                        ? 'bg-brand text-white shadow-md'
                        : 'text-zinc-400 hover:text-white'}"
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
                            d="M19 11H5m14 0a2 2 0 012 2v6a2 2 0 01-2 2H5a2 2 0 01-2-2v-6a2 2 0 012-2m14 0V9a2 2 0 00-2-2M5 11V9a2 2 0 012-2m0 0V5a2 2 0 012-2h6a2 2 0 012 2v2M7 7h10"
                        />
                    </svg>
                    <span>Ecosystem & Tools</span>
                </button>
            </div>
        </div>

        <!-- TAB 1: GitHub Languages View -->
        {#if activeTab === "github"}
            <div class="space-y-8 transition-opacity duration-300">
                <!-- Interactive Multi-segment Bar -->
                <div
                    class="rounded-2xl border border-white/10 bg-surface-card p-6 shadow-xl backdrop-blur-md"
                >
                    <div
                        class="mb-3 flex items-center justify-between text-xs text-zinc-400"
                    >
                        <span class="font-medium text-zinc-300"
                            >Repository Language Distribution</span
                        >
                        <span>Total: {formatBytes(techStack.totalBytes)}</span>
                    </div>

                    <!-- Stacked Bar -->
                    <div
                        class="flex h-3.5 w-full overflow-hidden rounded-full bg-zinc-800/80 p-0.5 shadow-inner"
                    >
                        {#each techStack.languages as lang}
                            <div
                                class="h-full first:rounded-l-full last:rounded-r-full transition-all duration-1000 ease-out hover:opacity-90 cursor-pointer"
                                style="width: {visible
                                    ? lang.percentage
                                    : 0}%; background-color: {lang.color};"
                                title="{lang.name}: {lang.percentage}% ({formatBytes(
                                    lang.bytes,
                                )})"
                            ></div>
                        {/each}
                    </div>

                    <!-- Language Legend Pills below the bar -->
                    <div class="mt-4 flex flex-wrap gap-4 pt-2">
                        {#each techStack.languages as lang}
                            <div class="flex items-center gap-2 text-xs">
                                <span
                                    class="h-2.5 w-2.5 rounded-full shadow-sm"
                                    style="background-color: {lang.color}"
                                ></span>
                                <span class="font-medium text-zinc-200"
                                    >{lang.name}</span
                                >
                                <span class="text-zinc-500"
                                    >{lang.percentage}%</span
                                >
                            </div>
                        {/each}
                    </div>
                </div>

                <!-- Language Cards Grid -->
                <div class="grid gap-4 sm:grid-cols-2 lg:grid-cols-3">
                    {#each techStack.languages as lang, i}
                        <div
                            class="group relative overflow-hidden rounded-xl border border-white/10 bg-surface-alt/80 p-5 backdrop-blur-md transition-all duration-300 hover:border-white/20 hover:-translate-y-1 hover:shadow-lg"
                        >
                            <!-- Ambient card top highlight matching language color -->
                            <div
                                class="absolute top-0 left-0 right-0 h-1 opacity-60 group-hover:opacity-100 transition-opacity"
                                style="background-color: {lang.color}"
                            ></div>

                            <div class="flex items-center justify-between mb-3">
                                <div class="flex items-center gap-2.5">
                                    <div
                                        class="h-3 w-3 rounded-full shadow-sm"
                                        style="background-color: {lang.color}"
                                    ></div>
                                    <h3
                                        class="font-semibold text-white group-hover:text-brand-light transition-colors"
                                    >
                                        {lang.name}
                                    </h3>
                                </div>
                                <span
                                    class="rounded-full bg-zinc-800/80 px-2 py-0.5 text-[11px] font-medium text-zinc-400"
                                >
                                    {lang.category}
                                </span>
                            </div>

                            <div
                                class="flex items-baseline justify-between text-xs text-zinc-400 mb-2"
                            >
                                <span>Calculated volume</span>
                                <span
                                    class="font-mono text-zinc-300 font-semibold"
                                    >{formatBytes(lang.bytes)}</span
                                >
                            </div>

                            <!-- Individual progress bar -->
                            <div
                                class="h-1.5 w-full overflow-hidden rounded-full bg-zinc-800"
                            >
                                <div
                                    class="h-full rounded-full transition-all duration-1000 ease-out"
                                    style="width: {visible
                                        ? lang.percentage
                                        : 0}%; background-color: {lang.color}; transition-delay: {i *
                                        100}ms"
                                ></div>
                            </div>

                            <div
                                class="mt-2 text-right text-[11px] font-medium text-zinc-500"
                            >
                                {lang.percentage}% of authored code
                            </div>
                        </div>
                    {/each}
                </div>
            </div>
        {:else}
            <!-- TAB 2: Full Ecosystem Grid -->
            <div
                class="grid gap-6 sm:grid-cols-2 transition-opacity duration-300"
            >
                {#each ecosystemCategories as cat}
                    <div
                        class="rounded-2xl border border-white/10 bg-surface-card p-6 backdrop-blur-md transition-all hover:border-brand/30"
                    >
                        <h3 class="text-lg font-bold text-white mb-1">
                            {cat.title}
                        </h3>
                        <p class="text-xs text-zinc-400 mb-5">
                            {cat.description}
                        </p>

                        <div class="flex flex-wrap gap-2.5">
                            {#each cat.items as item}
                                <div
                                    class="flex items-center gap-2 rounded-lg border {item.highlight
                                        ? 'border-brand/40 bg-brand/10 text-white'
                                        : 'border-white/10 bg-surface-alt text-zinc-300'} px-3 py-2 text-xs font-medium"
                                >
                                    <span>{item.name}</span>
                                    <span
                                        class="rounded bg-black/40 px-1.5 py-0.5 text-[10px] text-zinc-400"
                                    >
                                        {item.badge}
                                    </span>
                                </div>
                            {/each}
                        </div>
                    </div>
                {/each}
            </div>
        {/if}
    </div>
</section>
