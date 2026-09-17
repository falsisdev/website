# Architecture

## 1. Projenin Amacı

Bu repository bir kişisel portfolio / personal website projesidir. Proje şu anda **temel anasayfa aşamasındadır**. Modern, dark-theme tasarımlı bir anasayfa oluşturulmuş olup Hero, Projeler ve Yetenekler bölümleri Svelte component'leri ile geliştirilmiştir.

Bu aşamadaki amaç anasayfanın temel yapısını kurmak ve Astro + Svelte + Tailwind CSS stack'inin nasıl birlikte çalıştığını göstermektir.

---

## 2. Teknoloji Stack'i

| Teknoloji | Rol | Durum |
|---|---|---|
| **Astro** | Ana framework — routing, sayfa üretimi, static build | ✅ Kurulu |
| **Svelte** | İnteraktif UI component'leri için framework | ✅ Kurulu (3 component mevcut) |
| **Tailwind CSS** | Utility-first CSS framework — stillendirme | ✅ Kurulu |
| **Sanity** | Headless CMS — içerik yönetimi | ⏳ Planlanan (henüz entegre değil) |

---

## 3. Astro'nun Rolü

Astro bu projenin **ana framework'üdür**. Tüm sayfa yapısı, routing ve build sistemi Astro tarafından yönetilir.

### Astro'nun sorumlulukları:

- **Routing**: `src/pages/` klasöründeki dosyalar otomatik olarak URL route'larına dönüşür (file-based routing).
- **Sayfalar** (`pages`): Her `.astro` dosyası bir HTML sayfası üretir.
- **Layout'lar** (`layouts`): Sayfaların ortak HTML yapısını (head, body, meta tag'ler) tanımlar.
- **Static Generation**: Astro varsayılan olarak build zamanında statik HTML üretir. Sunucu tarafında çalışan bir runtime yoktur.
- **Build**: `astro build` komutu ile `dist/` klasörüne statik dosyalar üretilir.
- **GitHub Pages ile ilişkisi**: Üretilen statik dosyalar doğrudan GitHub Pages üzerinde sunulabilir.
- **HTML üretimi**: Her sayfa için tam HTML dosyaları üretilir.
- **JavaScript stratejisi**: Astro varsayılan olarak sayfaya **sıfır JavaScript** gönderir. JavaScript yalnızca bir component'e `client:*` direktifi verildiğinde tarayıcıya gönderilir.

### Konfigürasyon

Astro konfigürasyonu `astro.config.mjs` dosyasında tanımlanır:

```javascript
import { defineConfig } from 'astro/config';
import svelte from '@astrojs/svelte';
import tailwindcss from '@tailwindcss/vite';

export default defineConfig({
  site: 'https://falsisdev.github.io',
  base: '/website',
  integrations: [svelte()],
  vite: {
    plugins: [tailwindcss()],
  },
});
```

- `site`: Üretim URL'si (GitHub Pages adresi)
- `base`: Repository adına göre alt dizin (`/website`)
- `integrations`: Svelte entegrasyonu
- `vite.plugins`: Tailwind CSS v4 Vite plugin'i

---

## 4. Svelte'in Rolü

Svelte, Astro'nun **yerine geçen** bir framework **değildir**. Astro ile birlikte çalışan bir UI component framework'üdür.

### Astro + Svelte İlişkisi

```
Astro = sayfanın ana framework'ü (routing, layout, static HTML)
Svelte = gerektiğinde kullanılan interaktif component sistemi
```

Astro'nun "Islands Architecture" yaklaşımında, sayfanın büyük kısmı statik HTML olarak üretilir. Yalnızca interaktif olması gereken bölümler (islands) Svelte component'leri olarak tarayıcıda hydrate edilir.

### `client:*` Direktifleri

Bir Svelte component'i Astro sayfasında kullanıldığında, varsayılan olarak yalnızca HTML olarak render edilir ve tarayıcıya JavaScript gönderilmez. Component'in tarayıcıda interaktif olması için bir `client:*` direktifi gerekir:

| Direktif | Davranış |
|---|---|
| `client:load` | Sayfa yüklendiğinde hemen hydrate et |
| `client:idle` | Tarayıcı boşta kaldığında hydrate et |
| `client:visible` | Component viewport'a girdiğinde hydrate et |
| `client:media` | Belirli bir media query eşleştiğinde hydrate et |
| `client:only="svelte"` | Sunucuda render etme, yalnızca tarayıcıda render et |

### Örnek Kullanım (gelecekte)

```astro
---
import Counter from '../components/Counter.svelte';
---

<Counter client:load />
```

> **Not**: Projede şu anda 3 Svelte component'i bulunmaktadır: `Hero.svelte`, `ProjectCard.svelte`, `SkillBar.svelte`. Tümü Svelte 5 runes söz dizimi (`$props()`, `$state()`, `$effect()`) kullanmaktadır.

---

## 5. Tailwind CSS

Tailwind CSS bu projede **utility-first** yaklaşımıyla stillendirme için kullanılmaktadır.

### Prensipler

- Stillendirme için Tailwind utility class'ları kullanılacak
- Component'lerin içine gerektiğinde utility class olarak uygulanacak
- Gereksiz global CSS yazılmayacak
- İleride tasarım sistemi (design system) oluşturulabilir ancak şu an oluşturulmamıştır

### Kurulum Detayları

Tailwind CSS v4, **Vite plugin** yöntemiyle entegre edilmiştir:

1. `@tailwindcss/vite` plugin'i `astro.config.mjs` içinde Vite'a eklenir
2. `src/styles/global.css` dosyasında `@import "tailwindcss";` direktifi bulunur
3. Bu CSS dosyası `src/layouts/Layout.astro` içinde import edilir

Tailwind v4'te ayrı bir `tailwind.config.js` dosyasına gerek yoktur. Konfigürasyon gerektiğinde CSS dosyası içinde `@theme` direktifi ile yapılabilir.

---

## 6. Sanity

Sanity, ileride bu projenin **içerik yönetim sistemi** (Headless CMS) olarak kullanılması planlanmaktadır.

### Planlanan Akış

```
Sanity (içerik kaynağı)
    ↓
Astro build (build zamanında içerik çekilir)
    ↓
Static output (HTML/CSS/JS dosyaları)
    ↓
GitHub Pages (statik hosting)
```

### Sanity'nin Rolü (planlanan)

- Sanity bir **runtime backend değildir**. İçerik kaynağı olarak kullanılacaktır.
- İçerik Sanity Studio üzerinden yönetilecek.
- Astro build zamanında Sanity API'sinden içerik çekilecek ve statik HTML'e dönüştürülecek.
- Üretilen statik dosyalar GitHub Pages'ta sunulacak.
- Kullanıcı sayfayı ziyaret ettiğinde Sanity'ye herhangi bir istek yapılmayacak — her şey önceden build edilmiş olacak.

> **ÖNEMLİ**: Sanity şu anda projeye **entegre edilmemiştir**. Dependency olarak eklenmemiştir. Hiçbir Sanity kodu, schema, client veya environment variable bulunmamaktadır.

---

## 7. GitHub Pages

GitHub Pages bu projenin **hosting platformudur**.

### Deployment Akışı (planlanan)

```
GitHub repository
    ↓
Astro build (astro build)
    ↓
dist/ klasörü (statik dosyalar)
    ↓
GitHub Pages (statik hosting)
```

### GitHub Pages Kısıtlamaları

GitHub Pages **yalnızca statik dosyaları** sunabilir:

- HTML, CSS, JavaScript, görseller, fontlar vb.
- Server-side Go, Python veya Node.js uygulaması **çalıştıramaz**
- Veritabanı bağlantısı **kuramaz**
- API endpoint'i **barındıramaz**
- WebSocket veya SSE bağlantısı **sunamaz**

Bu nedenle projedeki tüm içerik build zamanında statik olarak üretilmelidir.

### Astro Konfigürasyonu

GitHub Pages uyumluluğu için `astro.config.mjs` dosyasında şu ayarlar yapılmıştır:

- `site: 'https://falsisdev.github.io'` — üretim URL'si
- `base: '/website'` — repository adına göre alt dizin

---

## 8. Backend Felsefesi

Bu proje şu anda **backend içermemektedir**. Bu bilinçli bir mimari karardır.

### Neden Backend Yok?

- GitHub Pages yalnızca statik dosya sunar
- Portfolio içeriği build zamanında üretilebilir
- Sanity entegrasyonu da build-time'da çalışacak
- Gereksiz karmaşıklık eklememek için backend katmanı oluşturulmamıştır

### Gelecekte Backend İhtiyacı

İleride backend gerektiren özellikler (örneğin gerçek zamanlı veri, form submission, authentication) eklenirse:

- Bu özellikler **GitHub Pages'ın dışında** ayrı bir servis olarak çalışmalıdır
- Serverless function'lar (Cloudflare Workers, Vercel Edge Functions vb.) kullanılabilir
- Veya ayrı bir backend repository'si oluşturulabilir

> **Not**: Eski projede bulunan Go backend bu aşamada tamamen kaldırılmıştır. Go veya herhangi bir backend dili bu repository'de bulunmamaktadır.

---

## 9. Klasör Yapısı

### Mevcut Yapı

```
website/
├── .gitignore              # Git ignore kuralları
├── ARCHITECTURE.md         # Mimari dokümantasyon (git'e açık)
├── README.md               # Proje tanıtım dosyası (İngilizce)
├── astro.config.mjs        # Astro konfigürasyonu
├── package.json            # Proje dependency'leri ve script'ler
├── package-lock.json       # Dependency lockfile
├── tsconfig.json           # TypeScript konfigürasyonu
├── svelte.config.js        # Svelte preprocessor konfigürasyonu
├── public/                 # Statik dosyalar (build'e kopyalanır)
│   └── favicon.svg         # Site favicon'u
└── src/                    # Kaynak kodlar
    ├── components/         # Svelte component'leri
    │   ├── Hero.svelte     # Anasayfa hero section (animasyonlu giriş)
    │   ├── ProjectCard.svelte  # Proje kartları grid'i
    │   └── SkillBar.svelte # Animasyonlu yetenek barları
    ├── layouts/            # Sayfa layout'ları
    │   └── Layout.astro    # Ana layout (HTML boilerplate, dark theme)
    ├── pages/              # Sayfalar (file-based routing)
    │   └── index.astro     # Ana sayfa (Hero + Projects + Skills + Footer)
    └── styles/             # Global stiller
        └── global.css      # Tailwind CSS import'u ve tema tanımları
```

### Gelecekte Eklenebilecek Klasörler (planlanan)

Aşağıdaki klasörler ihtiyaç duyulduğunda oluşturulabilir:

| Klasör | Amaç |
|---|---|
| `src/content/` | Astro Content Collections (lokal içerik) |
| `src/lib/` | Yardımcı fonksiyonlar ve modüller |

> **Kural**: Bu klasörleri gerçekten ihtiyaç duyulana kadar oluşturma. Boş klasör oluşturma.

---

## 10. Sayfa Oluşturma

Astro file-based routing kullanır. `src/pages/` klasörüne eklenen her `.astro` dosyası otomatik olarak bir route oluşturur.

### Örnekler

| Dosya | URL |
|---|---|
| `src/pages/index.astro` | `/website/` |
| `src/pages/about.astro` | `/website/about` |
| `src/pages/projects/index.astro` | `/website/projects` |
| `src/pages/blog/[slug].astro` | `/website/blog/:slug` (dinamik) |

### Yeni Sayfa Oluşturma

```astro
---
// src/pages/about.astro
import Layout from '../layouts/Layout.astro';
---

<Layout title="Hakkımda">
  <main>
    <h1>Hakkımda</h1>
    <p>İçerik buraya gelecek.</p>
  </main>
</Layout>
```

Bu dosya oluşturulduğunda `/website/about` adresinde erişilebilir olur.

---

## 11. Svelte Component Oluşturma

Gelecekte interaktif bir component gerektiğinde:

### 1. Component Dosyası Oluştur

```svelte
<!-- src/components/Counter.svelte -->
<script>
  let count = $state(0);
</script>

<button on:click={() => count++}>
  Sayaç: {count}
</button>
```

### 2. Astro Sayfasında Kullan

```astro
---
import Counter from '../components/Counter.svelte';
---

<Layout title="Örnek">
  <!-- Statik içerik -->
  <h1>Merhaba</h1>

  <!-- İnteraktif island -->
  <Counter client:load />
</Layout>
```

### Hydration Direktifi Seçimi

| Durum | Direktif |
|---|---|
| Hemen interaktif olması gerekiyor | `client:load` |
| Sayfa yüklendikten sonra yeterli | `client:idle` |
| Kullanıcı component'i görene kadar bekleyebilir | `client:visible` |
| Yalnızca belirli ekran boyutlarında | `client:media="(min-width: 768px)"` |
| Sunucuda render etmeye gerek yok | `client:only="svelte"` |

> **Not**: Mevcut component'ler (`Hero.svelte`, `ProjectCard.svelte`, `SkillBar.svelte`) bu pattern'i takip etmektedir.

---

## 12. Stil Yazma

Tailwind CSS ile stillendirme doğrudan HTML elementlerine utility class'lar eklenerek yapılır.

### Örnek

```astro
<h1 class="text-2xl font-bold text-zinc-900">Başlık</h1>
<p class="mt-2 text-zinc-600">Açıklama metni</p>
```

Svelte component'lerinde de aynı yaklaşım geçerlidir:

```svelte
<div class="rounded-lg border p-4">
  <slot />
</div>
```

### Özel Stil Gerektiğinde

Tailwind v4'te özel değerler CSS dosyasında `@theme` ile tanımlanabilir:

```css
/* src/styles/global.css */
@import "tailwindcss";

@theme {
  --color-brand: #6366f1;
}
```

> **Not**: `global.css` dosyasında dark-theme renk paleti (`brand`, `brand-light`, `surface`, `surface-alt`, `border`) ve Inter fontu tanımlanmıştır.

---

## 13. İçerik Yönetimi

İleride portfolio içeriğinin Sanity üzerinden yönetilmesi planlanmaktadır.

### Planlanan İçerik Tipleri

| İçerik | Açıklama |
|---|---|
| Projects | Portfolio projeleri |
| Posts | Blog yazıları |
| Profile | Kişisel bilgiler |
| Experiences | İş deneyimleri |

### Planlanan Akış

1. İçerik Sanity Studio üzerinden oluşturulur/düzenlenir
2. Astro build zamanında Sanity API'sinden içerik çekilir
3. İçerik statik HTML sayfalarına dönüştürülür
4. Statik dosyalar GitHub Pages'a deploy edilir

> **ÖNEMLİ**: Bu içerik tipleri ve Sanity entegrasyonu henüz **oluşturulmamıştır**. Gelecekte yapılacaktır.

---

## 14. Veri Akışı

### Mevcut Durum

```
Tarayıcı (Browser)
    ↓ HTTP isteği
GitHub Pages
    ↓ statik dosya sunumu
Statik HTML / CSS / JS
```

Şu anda herhangi bir veri kaynağı veya API çağrısı bulunmamaktadır.

### Gelecekte (Sanity entegrasyonundan sonra)

```
Sanity (CMS — içerik kaynağı)
    ↓ API çağrısı (build zamanında)
Astro build
    ↓ statik dosya üretimi
Static output (dist/)
    ↓ deploy
GitHub Pages (hosting)
    ↓ statik dosya sunumu
Tarayıcı (Browser)
```

---

## 15. Gerçek Zamanlı Özellikler

İleride gerçek zamanlı özellikler (örneğin canlı GitHub aktivitesi) eklenmek istenirse, GitHub Pages tek başına yeterli olmayacaktır.

### Gerekli Mimari (gelecekte)

```
GitHub API (veri kaynağı)
    ↓
Backend / Serverless Function (veri işleme)
    ↓ SSE veya WebSocket
Tarayıcı (Browser)
```

Bu tür bir mimari şunları gerektirir:

- Ayrı bir backend servisi veya serverless function
- GitHub Pages dışında bir hosting (backend için)
- Client-side JavaScript ile bağlantı yönetimi

> **ÖNEMLİ**: Bu sistemlerin hiçbiri şu anda projede **bulunmamaktadır**. Eski projede bulunan Go backend ve SSE sistemi kaldırılmıştır.

---

## 16. Geliştirme Kuralları

Projeye kod eklerken uyulması gereken kurallar:

1. **Gereksiz dependency ekleme.** Her yeni dependency eklenmeden önce gerçekten gerekli olup olmadığını değerlendir.

2. **Astro'nun doğal yapısını bozma.** File-based routing, layout sistemi ve static generation Astro'nun varsayılan davranışlarıdır. Bunları gereksiz yere override etme.

3. **Svelte'i yalnızca ihtiyaç olduğunda kullan.** Statik içerik (metin, görsel, layout) için Svelte component'i oluşturma. Svelte yalnızca tarayıcıda interaksiyon gerektiren bölümler için kullanılmalı.

4. **Her şeyi Svelte component'ine dönüştürme.** Astro component'leri (`.astro` dosyaları) statik içerik için yeterlidir ve sıfır JavaScript üretir.

5. **Statik içerik için gereksiz client-side JavaScript kullanma.** `client:*` direktifini yalnızca gerçekten interaktif component'ler için kullan.

6. **GitHub Pages uyumluluğunu koru.** Eklenen her özelliğin statik build ile uyumlu olduğundan emin ol.

7. **Backend gerektiren özellikleri doğrudan frontend'e zorla sokma.** Server-side işlem gerektiren özellikler için ayrı bir servis planla.

8. **Sanity ile ilgili kodu gerçekten entegrasyon yapılana kadar ekleme.** Placeholder Sanity kodu, boş client'lar veya dummy schema'lar oluşturma.

9. **Kullanılmayan klasör ve abstraction oluşturma.** Klasörler gerçekten ihtiyaç duyulduğunda oluşturulmalı.

10. **Overengineering yapma.** Basit bir özellik için karmaşık abstraction katmanları oluşturma.

---

## 17. Build ve Development

### Dependency Kurulumu

```bash
npm install
```

### Development Server

```bash
npm run dev
```

Varsayılan olarak `http://localhost:4321` adresinde çalışır.

### Production Build

```bash
npm run build
```

Statik dosyaları `dist/` klasörüne üretir.

### Build Önizleme

```bash
npm run preview
```

Üretilen statik dosyaları lokal sunucu ile önizler.

---

## 18. Git Kuralları

### Commit Mesajı Formatı

Bu proje [Conventional Commits](https://www.conventionalcommits.org/) formatını kullanmaktadır:

```
<type>: <description>
```

Kullanılan type'lar:

| Type | Kullanım |
|---|---|
| `feat` | Yeni özellik |
| `fix` | Hata düzeltme |
| `chore` | Bakım işleri (dependency güncelleme, config değişikliği) |
| `docs` | Dokümantasyon değişikliği |
| `refactor` | Davranış değiştirmeyen kod yeniden yapılandırması |
| `style` | Kod formatı değişikliği (işlevsel değişiklik yok) |

### Bu Görev

Bu scaffold kurulumu tek bir commit altında yapılmıştır:

```
feat: initialize portfolio architecture
```

Eski Go backend projesi kaldırılmış ve Astro + Svelte + Tailwind CSS tabanlı yeni proje sıfırdan oluşturulmuştur.
