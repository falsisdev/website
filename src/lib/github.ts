export interface GitHubProfile {
  login: string;
  name: string;
  avatar_url: string;
  bio: string;
  location: string;
  public_repos: number;
  followers: number;
  following: number;
  html_url: string;
}

export interface LanguageStat {
  name: string;
  bytes: number;
  percentage: number;
  color: string;
  category: "systems" | "frontend" | "scripting" | "markup";
}

export interface GitHubTechStack {
  languages: LanguageStat[];
  totalBytes: number;
  repoCount: number;
  topTechnologies: string[];
}

export interface GitHubProject {
  title: string;
  name: string;
  category: "fullstack" | "backend" | "tooling";
  description: string;
  tags: string[];
  url: string;
  demo?: string;
  stars: number;
  featured?: boolean;
}

export const LANGUAGE_COLORS: Record<
  string,
  { color: string; category: LanguageStat["category"] }
> = {
  Go: { color: "#00ADD8", category: "systems" },
  TypeScript: { color: "#3178C6", category: "frontend" },
  Vue: { color: "#41B883", category: "frontend" },
  JavaScript: { color: "#F7DF1E", category: "frontend" },
  Svelte: { color: "#FF3E00", category: "frontend" },
  Python: { color: "#3572A5", category: "scripting" },
  Astro: { color: "#BC52EE", category: "frontend" },
  HTML: { color: "#E34F26", category: "markup" },
  CSS: { color: "#663399", category: "markup" },
};

export const FALLBACK_PROFILE: GitHubProfile = {
  login: "falsisdev",
  name: "Falsis",
  avatar_url: "https://avatars.githubusercontent.com/u/63756985?v=4",
  bio: "ECE undergrad @ YTU • Systems & Web Developer",
  location: "Izmir, Turkiye",
  public_repos: 6,
  followers: 54,
  following: 10,
  html_url: "https://github.com/falsisdev",
};

export const FALLBACK_PROJECTS: GitHubProject[] = [
  {
    title: "Mangile",
    name: "mangile",
    category: "fullstack",
    description:
      "A lightning-fast, highly intuitive web platform delivering localized scans for manga, webtoon, and novel enthusiasts.",
    tags: ["Vue", "Nuxt", "TypeScript", "Tailwind CSS"],
    url: "https://github.com/falsisdev/mangile",
    demo: "https://mangile.vercel.app",
    stars: 40,
    featured: true,
  },
  {
    title: "Website",
    name: "website",
    category: "fullstack",
    description:
      "Personal portfolio website built with modern web technologies — statically generated, zero-config deployment.",
    tags: ["Svelte", "Astro", "Tailwind CSS"],
    url: "https://github.com/falsisdev/website",
    demo: "https://falsisdev.github.io/website/",
    stars: 12,
    featured: true,
  },
  {
    title: "Anthology",
    name: "anthology",
    category: "tooling",
    description:
      "Verified multi-source streaming catalog and modular scraper suite for Stremio and Nuvio media ecosystem.",
    tags: ["JavaScript", "Addon", "Stremio API", "Scrapers"],
    url: "https://github.com/falsisdev/anthology",
    demo: "https://falsisdev.github.io/anthology/",
    stars: 5,
    featured: false,
  },
  {
    title: "Mangile Backend",
    name: "mangile-backend",
    category: "backend",
    description:
      "Low-latency, high-throughput backend engine for the Mangile platform engineered with clean architecture.",
    tags: ["Go", "REST API", "Clean Architecture", "Concurrency"],
    url: "https://github.com/falsisdev/mangile-backend",
    demo: "https://mangile.vercel.app",
    stars: 3,
    featured: false,
  },
  {
    title: "Mangile CLI",
    name: "mangile-cli",
    category: "backend",
    description:
      "Single-purpose CLI tool developed for publishing and managing Mangile content and automated archival workflows.",
    tags: ["Go", "CLI", "Sanity", "Developer Tools"],
    url: "https://github.com/falsisdev/mangile-cli",
    stars: 2,
    featured: false,
  },
  {
    title: "Vessel",
    name: "vessel",
    category: "tooling",
    description:
      "An umbrella, local-first, modular digital media consumption platform.",
    tags: ["Go", "Cloudflare", "Multilingual"],
    url: "https://github.com/falsisdev/vessel",
    stars: 1,
    featured: false,
  },
];

export const FALLBACK_TECH_STACK: GitHubTechStack = {
  totalBytes: 1198888,
  repoCount: 6,
  topTechnologies: [
    "Go",
    "JavaScript",
    "TypeScript",
    "Vue / Nuxt",
    "Svelte",
    "Tailwind CSS",
  ],
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
};

function getHeaders(): Record<string, string> {
  const headers: Record<string, string> = {
    "User-Agent": "falsisdev-portfolio-builder",
    Accept: "application/vnd.github.v3+json",
  };
  let token =
    (typeof import.meta !== "undefined" &&
      (import.meta.env.GITHUB_TOKEN || import.meta.env.GH_TOKEN)) ||
    (typeof import.meta !== "undefined" &&
      (import.meta as any).env?.GITHUB_TOKEN);

  if (
    !token &&
    typeof import.meta !== "undefined" &&
    typeof window === "undefined"
  ) {
    try {
      const childProcess = (import.meta as any).getBuiltinModule?.(
        "node:child_process",
      );
      if (childProcess?.execSync) {
        token = childProcess
          .execSync("gh auth token", {
            stdio: ["ignore", "pipe", "ignore"],
            encoding: "utf-8",
          })
          .trim();
      }
    } catch {
      // gh CLI not available or not logged in
    }
  }

  if (token) {
    headers.Authorization = `Bearer ${token}`;
  }
  return headers;
}

export async function getGitHubProfile(
  username = "falsisdev",
): Promise<GitHubProfile> {
  try {
    const controller = new AbortController();
    const timeout = setTimeout(() => controller.abort(), 4000);

    const res = await fetch(`https://api.github.com/users/${username}`, {
      headers: getHeaders(),
      signal: controller.signal,
    });
    clearTimeout(timeout);

    if (!res.ok) {
      console.warn(
        `[GitHub API] Profile fetch failed with status ${res.status}, using fallback.`,
      );
      return FALLBACK_PROFILE;
    }

    const data = await res.json();
    return {
      login: data.login || FALLBACK_PROFILE.login,
      name: data.name || FALLBACK_PROFILE.name,
      avatar_url: data.avatar_url || FALLBACK_PROFILE.avatar_url,
      bio: data.bio || FALLBACK_PROFILE.bio,
      location: data.location || FALLBACK_PROFILE.location,
      public_repos:
        typeof data.public_repos === "number"
          ? data.public_repos
          : FALLBACK_PROFILE.public_repos,
      followers:
        typeof data.followers === "number"
          ? data.followers
          : FALLBACK_PROFILE.followers,
      following:
        typeof data.following === "number"
          ? data.following
          : FALLBACK_PROFILE.following,
      html_url: data.html_url || FALLBACK_PROFILE.html_url,
    };
  } catch (err) {
    console.warn("[GitHub API] Network error during profile fetch:", err);
    return FALLBACK_PROFILE;
  }
}

export async function getGitHubTechStack(
  username = "falsisdev",
): Promise<GitHubTechStack> {
  try {
    const controller = new AbortController();
    const timeout = setTimeout(() => controller.abort(), 4500);

    const res = await fetch(
      `https://api.github.com/users/${username}/repos?per_page=100&sort=updated`,
      {
        headers: getHeaders(),
        signal: controller.signal,
      },
    );
    clearTimeout(timeout);

    if (!res.ok) {
      console.warn(
        `[GitHub API] Repos fetch failed with status ${res.status}, using fallback.`,
      );
      return FALLBACK_TECH_STACK;
    }

    const repos = await res.json();
    if (!Array.isArray(repos) || repos.length === 0) {
      return FALLBACK_TECH_STACK;
    }

    const personalRepos = repos.filter((r) => !r.fork);

    const languageTotals: Record<string, number> = {};

    const languageFetches = personalRepos.slice(0, 6).map(async (repo) => {
      try {
        if (!repo.languages_url) return;
        const lRes = await fetch(repo.languages_url, { headers: getHeaders() });
        if (lRes.ok) {
          const langs = await lRes.json();
          for (const [lang, bytes] of Object.entries(langs)) {
            if (typeof bytes === "number") {
              languageTotals[lang] = (languageTotals[lang] || 0) + bytes;
            }
          }
        }
      } catch {
        if (repo.language) {
          languageTotals[repo.language] =
            (languageTotals[repo.language] || 0) + 10000;
        }
      }
    });

    await Promise.all(languageFetches);

    const totalBytes = Object.values(languageTotals).reduce(
      (sum, b) => sum + b,
      0,
    );

    if (totalBytes === 0) {
      return FALLBACK_TECH_STACK;
    }

    const filteredLangs = Object.entries(languageTotals)
      .filter(([lang]) => !["HTML", "CSS"].includes(lang))
      .sort((a, b) => b[1] - a[1]);

    const filteredTotal =
      filteredLangs.reduce((sum, [, b]) => sum + b, 0) || totalBytes;

    const languages: LanguageStat[] = filteredLangs.map(([name, bytes]) => {
      const config = LANGUAGE_COLORS[name] || {
        color: "#6366f1",
        category: "systems" as const,
      };
      const percentage = parseFloat(((bytes / filteredTotal) * 100).toFixed(1));
      return {
        name,
        bytes,
        percentage,
        color: config.color,
        category: config.category,
      };
    });

    return {
      languages:
        languages.length > 0 ? languages : FALLBACK_TECH_STACK.languages,
      totalBytes: filteredTotal,
      repoCount: personalRepos.length,
      topTechnologies: [
        "Go",
        "TypeScript",
        "Vue / Nuxt",
        "Svelte",
        "Tailwind CSS",
        "Docker",
      ],
    };
  } catch (err) {
    console.warn("[GitHub API] Error calculating tech stack:", err);
    return FALLBACK_TECH_STACK;
  }
}

const REPO_TITLE_MAP: Record<string, string> = {
  mangile: "Mangile",
  "mangile-backend": "Mangile Backend",
  "mangile-cli": "Mangile CLI",
  anthology: "Anthology",
  website: "Portfolio",
  vessel: "Vessel",
};

function formatRepoTitle(name: string): string {
  if (REPO_TITLE_MAP[name.toLowerCase()]) {
    return REPO_TITLE_MAP[name.toLowerCase()];
  }
  return name
    .split(/[-_]/)
    .map((w) => w.charAt(0).toUpperCase() + w.slice(1))
    .join(" ");
}

function determineCategory(repo: any): "fullstack" | "backend" | "tooling" {
  const lang = (repo.language || "").toLowerCase();
  const topics: string[] = (repo.topics || []).map((t: string) =>
    t.toLowerCase(),
  );
  const name = (repo.name || "").toLowerCase();

  if (
    topics.some((t) =>
      [
        "scraper",
        "scrapers",
        "addon",
        "stremio",
        "nuvio",
        "cloudstream",
      ].includes(t),
    ) ||
    name === "anthology" ||
    name === "vessel"
  ) {
    return "tooling";
  }
  if (
    lang === "go" ||
    topics.some((t) => ["backend", "golang", "cli", "api"].includes(t))
  ) {
    return "backend";
  }
  return "fullstack";
}

function formatRepoTags(repo: any): string[] {
  const tags: string[] = [];
  if (repo.language) {
    tags.push(repo.language);
  }
  if (repo.name?.toLowerCase() === "website" && !tags.includes("Astro")) {
    tags.push("Astro", "Tailwind");
  }
  if (Array.isArray(repo.topics)) {
    for (const t of repo.topics) {
      if (t.toLowerCase() === repo.language?.toLowerCase()) continue;
      const formatted =
        t.length <= 3
          ? t.toUpperCase()
          : t.charAt(0).toUpperCase() + t.slice(1);
      if (!tags.includes(formatted)) tags.push(formatted);
      if (tags.length >= 4) break;
    }
  }
  return tags.length > 0 ? tags : [repo.language || "Code"];
}

export async function getGitHubProjects(
  username = "falsisdev",
): Promise<GitHubProject[]> {
  try {
    const controller = new AbortController();
    const timeout = setTimeout(() => controller.abort(), 4500);

    const res = await fetch(
      `https://api.github.com/users/${username}/repos?per_page=100&sort=updated`,
      {
        headers: getHeaders(),
        signal: controller.signal,
      },
    );
    clearTimeout(timeout);

    if (!res.ok) {
      console.warn(
        `[GitHub API] Repos fetch failed with status ${res.status}, using fallback projects.`,
      );
      return FALLBACK_PROJECTS;
    }

    const repos = await res.json();
    if (!Array.isArray(repos) || repos.length === 0) {
      return FALLBACK_PROJECTS;
    }

    const filtered = repos
      .filter((r) => !r.fork && r.name.toLowerCase() !== username.toLowerCase())
      .sort((a, b) => (b.stargazers_count || 0) - (a.stargazers_count || 0));

    const projects: GitHubProject[] = filtered.map((r) => ({
      title: formatRepoTitle(r.name),
      name: r.name,
      category: determineCategory(r),
      description:
        r.description || "Open source project and developer tooling.",
      tags: formatRepoTags(r),
      url: r.html_url,
      demo:
        r.homepage &&
        typeof r.homepage === "string" &&
        r.homepage.trim().length > 0
          ? r.homepage
          : undefined,
      stars: r.stargazers_count || 0,
      featured:
        (r.stargazers_count || 0) >= 10 || r.name.toLowerCase() === "mangile",
    }));

    return projects.length > 0 ? projects : FALLBACK_PROJECTS;
  } catch (err) {
    console.warn("[GitHub API] Error fetching repositories for projects:", err);
    return FALLBACK_PROJECTS;
  }
}
