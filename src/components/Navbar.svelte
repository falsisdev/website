<script>
    let { profile = { login: "falsisdev", name: "Falsis" } } = $props();
    let scrolled = $state(false);

    $effect(() => {
        const handleScroll = () => {
            scrolled = window.scrollY > 40;
        };
        window.addEventListener("scroll", handleScroll, { passive: true });
        return () => window.removeEventListener("scroll", handleScroll);
    });
</script>

<header
    class="fixed top-0 left-0 right-0 z-50 flex justify-center px-4 pt-4 transition-all duration-300 pointer-events-none"
>
    <nav
        class="pointer-events-auto flex items-center justify-between gap-4 sm:gap-8 rounded-full border border-white/10 bg-surface/80 px-5 py-2.5 backdrop-blur-xl shadow-2xl shadow-black/60 transition-all duration-300 {scrolled
            ? 'scale-95 border-brand/30 bg-surface-alt/90 shadow-brand/5'
            : ''}"
        aria-label="Main Navigation"
    >
        <a href="#about" class="flex items-center gap-2.5 group">
            <span
                class="text-sm font-semibold tracking-tight text-white group-hover:text-brand-light transition-colors"
            >
                falsis<span class="text-brand-light">.dev</span>
            </span>
        </a>

        <!-- Nav Links (hidden on mobile, visible on sm+) -->
        <div
            class="hidden sm:flex items-center gap-6 text-xs font-medium text-zinc-400"
        >
            <a
                href="#about"
                class="transition-colors hover:text-white hover:text-brand-light"
                >About</a
            >
            <a
                href="#projects"
                class="transition-colors hover:text-white hover:text-brand-light"
                >Projects</a
            >
            <a
                href="#tech-stack"
                class="transition-colors hover:text-white hover:text-brand-light"
                >Tech Stack</a
            >
        </div>

        <!-- Actions: Live status & GitHub -->
        <div class="flex items-center gap-3">
            <div
                class="hidden md:flex items-center gap-1.5 rounded-full border border-emerald-500/20 bg-emerald-500/10 px-2.5 py-1 text-[11px] font-medium text-emerald-400"
            >
                <span class="relative flex h-2 w-2">
                    <span
                        class="absolute inline-flex h-full w-full animate-ping rounded-full bg-emerald-400 opacity-75"
                    ></span>
                    <span
                        class="relative inline-flex h-2 w-2 rounded-full bg-emerald-500"
                    ></span>
                </span>
                Undergraduate
            </div>

            <a
                href="https://github.com/falsisdev"
                target="_blank"
                rel="noopener noreferrer"
                aria-label="GitHub Profile"
                class="flex items-center gap-1.5 rounded-full border border-white/10 bg-surface-card px-3 py-1.5 text-xs font-medium text-zinc-200 transition-all hover:border-brand/50 hover:bg-brand/10 hover:text-white"
            >
                <svg class="h-4 w-4 fill-current" viewBox="0 0 24 24">
                    <path
                        d="M12 0C5.37 0 0 5.37 0 12c0 5.31 3.435 9.795 8.205 11.385.6.105.825-.255.825-.57 0-.285-.015-1.23-.015-2.235-3.015.555-3.795-.735-4.035-1.41-.135-.345-.72-1.41-1.23-1.695-.42-.225-1.02-.78-.015-.795.945-.015 1.62.87 1.845 1.23 1.08 1.815 2.805 1.305 3.495.99.105-.78.42-1.305.765-1.605-2.67-.3-5.46-1.335-5.46-5.925 0-1.305.465-2.385 1.23-3.225-.12-.3-.54-1.53.12-3.18 0 0 1.005-.315 3.3 1.23.96-.27 1.98-.405 3-.405s2.04.135 3 .405c2.295-1.56 3.3-1.23 3.3-1.23.66 1.65.24 2.88.12 3.18.765.84 1.23 1.905 1.23 3.225 0 4.605-2.805 5.625-5.475 5.925.435.375.81 1.095.81 2.22 0 1.605-.015 2.895-.015 3.3 0 .315.225.69.825.57A12.02 12.02 0 0024 12c0-6.63-5.37-12-12-12z"
                    />
                </svg>
                <span>GitHub</span>
            </a>
        </div>
    </nav>
</header>
